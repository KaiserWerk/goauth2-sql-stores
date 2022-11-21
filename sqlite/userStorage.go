package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/KaiserWerk/goauth2/storage"
	_ "modernc.org/sqlite"
)

type UserStorage struct {
	conn *sql.DB
}

func (s *UserStorage) Get(id uint) (storage.OAuth2User, error) {
	u := storage.User{}

	rows, err := s.conn.Query(userSelectQuery, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("no entry found")
	}

	if err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Disabled); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserStorage) GetByUsername(username string) (storage.OAuth2User, error) {
	u := storage.User{}

	rows, err := s.conn.Query(userSelectByUsernameQuery, username)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("no entry found")
	}

	if err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Disabled); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserStorage) Add(user storage.OAuth2User) error {
	password := user.GetPassword() // TODO
	res, err := s.conn.Exec(userInsertQuery, user.GetUsername(), user.GetEmail(), password, user.IsDisabled())
	if err != nil {
		return err
	}

	newID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.SetID(uint(newID))
	return nil
}

func (s *UserStorage) Edit(user storage.OAuth2User) error {
	password := user.GetPassword() // TODO
	_, err := s.conn.Exec(userUpdateQuery, user.GetUsername(), user.GetEmail(), password, user.IsDisabled())
	return err
}

func (s *UserStorage) Remove(id uint) error {
	_, err := s.conn.Exec(userDeleteQuery, id)
	return err
}

func (s *UserStorage) Close() error {
	return s.conn.Close()
}