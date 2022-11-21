package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/KaiserWerk/goauth2/storage"
	_ "github.com/go-sql-driver/mysql"
)

const urlSeparator = `|`

type ClientStorage struct {
	conn *sql.DB
}

func NewClientStorage(dsn string) (*ClientStorage, error) {
	cs := &ClientStorage{}
	var err error
	cs.conn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	cs.conn.SetMaxOpenConns(10)
	cs.conn.SetMaxIdleConns(10)

	_, err = cs.conn.Query(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create client table: %w", err)
	}

	return cs, nil
}

func (s *ClientStorage) Get(id string) (storage.OAuth2Client, error) {
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

	_, err := s.conn.Exec(insertQuery, client.GetID(), client.GetSecret(), confi, client.GetApplicationName(), urls)
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

	_, err := s.conn.Exec(updateQuery, client.GetSecret(), confi, client.GetApplicationName(), urls, client.GetID())
	return err
}

func (s *ClientStorage) Remove(client storage.OAuth2Client) error {
	_, err := s.conn.Exec(deleteQuery, client.GetID())
	return err
}

func (s *ClientStorage) Close() error {
	return s.conn.Close()
}
