package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/KaiserWerk/goauth2/storage"
	_ "modernc.org/sqlite"
)

type SessionStorage struct {
	conn *sql.DB
}

// NewSessionStorage returns a new instance of the SessionStorage using an SQLite3 DSN.
func NewSessionStorage(dsn string) (*SessionStorage, error) {
	var (
		err error
		cs  = &SessionStorage{}
	)
	cs.conn, err = sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	_, err = cs.conn.Exec(sessionCreateTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create client table: %w", err)
	}

	return cs, nil
}

func (s *SessionStorage) Get(id string) (storage.OAuth2Session, error) {
	sess := storage.Session{}
	rows, err := s.conn.Query(sessionSelectQuery, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("no entry found")
	}

	if err = rows.Scan(&sess.ID, &sess.UserID, &sess.Expires); err != nil {
		return nil, err
	}

	return sess, nil
}

func (s *SessionStorage) Add(session storage.OAuth2Session) error {
	_, err := s.conn.Exec(sessionInsertQuery, session.GetID(), session.GetUserID(), session.GetExpires())
	return err
}

func (s *SessionStorage) Remove(id string) error {
	_, err := s.conn.Exec(sessionDeleteQuery, id)
	return err
}

func (s *SessionStorage) Close() error {
	return s.conn.Close()
}
