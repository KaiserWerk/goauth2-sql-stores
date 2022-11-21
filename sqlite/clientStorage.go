package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/KaiserWerk/goauth2/storage"
	_ "modernc.org/sqlite"
)

const urlSeparator = `|`

type ClientStorage struct {
	conn *sql.DB
}

// NewClientStorage returns a new instance of the ClientStorage using an SQLite3 DSN.
func NewClientStorage(dsn string) (*ClientStorage, error) {
	var (
		err error
		cs  = &ClientStorage{}
	)
	cs.conn, err = sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	_, err = cs.conn.Exec(clientCreateTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create client table: %w", err)
	}

	return cs, nil
}

func (s *ClientStorage) Get(id string) (storage.OAuth2Client, error) {
	c := storage.Client{}
	rows, err := s.conn.Query(clientSelectQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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

func (s *ClientStorage) Add(client storage.OAuth2Client) error {
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

	_, err := s.conn.Exec(clientInsertQuery, client.GetID(), client.GetSecret(), confi, client.GetApplicationName(), urls)
	return err
}

func (s *ClientStorage) Edit(client storage.OAuth2Client) error {
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

	_, err := s.conn.Exec(clientUpdateQuery, client.GetSecret(), confi, client.GetApplicationName(), urls, client.GetID())
	return err
}

func (s *ClientStorage) Remove(client storage.OAuth2Client) error {
	_, err := s.conn.Exec(clientDeleteQuery, client.GetID())
	return err
}

func (s *ClientStorage) Close() error {
	return s.conn.Close()
}
