package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/KaiserWerk/goauth2/storage"
)

const urlSeparator = `|`

type SQLiteClientStorage struct {
	conn *sql.DB
}

// New returns a new instance of the SQLiteClientStorage using an SQLite3 DSN.
func New(dsn string) (*SQLiteClientStorage, error) {
	var (
		err error
		cs  = &SQLiteClientStorage{}
	)
	cs.conn, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	_, err = cs.conn.Query(createQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create client table: %w", err)
	}

	return cs, nil
}

func (s *SQLiteClientStorage) Get(id string) (storage.OAuth2Client, error) {
	c := storage.Client{}
	rows, err := s.conn.Query(selectQuery, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("no entry found")
	}

	var (
		confidential int
		urlsRaw      string
	)
	if err = rows.Scan(&c.ID, &c.Secret, &confidential, &c.AppName, &urlsRaw); err != nil {
		return nil, err
	}

	if confidential == 1 {
		c.Confidential = true
	} else {
		c.Confidential = false
	}

	c.RedirectURLs = strings.Split(urlsRaw, urlSeparator)

	return c, nil
}

func (s *SQLiteClientStorage) Add(client storage.OAuth2Client) error {
	// check if client exists?
	var (
		confi int
		urls  string
	)

	if client.IsConfidential() {
		confi = 1
	} else {
		confi = 0
	}

	urls = strings.Join(client.GetRedirectURLs(), urlSeparator)

	_, err := s.conn.Query(insertQuery, client.GetID(), client.GetSecret(), confi, client.GetApplicationName(), urls)
	return err
}

func (s *SQLiteClientStorage) Edit(client storage.OAuth2Client) error {
	var (
		confi int
		urls  string
	)

	if client.IsConfidential() {
		confi = 1
	} else {
		confi = 0
	}

	urls = strings.Join(client.GetRedirectURLs(), urlSeparator)

	_, err := s.conn.Query(updateQuery, client.GetSecret(), confi, client.GetApplicationName(), urls, client.GetID())
	return err
}

func (s *SQLiteClientStorage) Remove(client storage.OAuth2Client) error {
	_, err := s.conn.Query(deleteQuery, client.GetID())
	return err
}

func (s *SQLiteClientStorage) Close() error {
	return s.conn.Close()
}
