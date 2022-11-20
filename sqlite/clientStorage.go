package sqlite

import (
	"database/sql"
	"github.com/KaiserWerk/goauth2/storage"
)

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

	// TODO: set up table if not exist

	return cs, nil
}

// TODO: close() methode

func (s *SQLiteClientStorage) Get(id string) (storage.OAuth2Client, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SQLiteClientStorage) Set(client storage.OAuth2Client) error {
	//TODO implement me
	panic("implement me")
}
