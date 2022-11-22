package mysql

import (
	"database/sql"
	"fmt"

	"github.com/KaiserWerk/goauth2/storage"
	_ "github.com/go-sql-driver/mysql"
)

type TokenStorage struct {
	conn *sql.DB
}

func NewTokenStorage(dsn string) (*TokenStorage, error) {
	var (
		err error
		cs  = &TokenStorage{}
	)
	cs.conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	_, err = cs.conn.Exec(tokenCreateTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create token table: %w", err)
	}

	return cs, nil
}

func (s *TokenStorage) FindByCodeChallenge(cc string) (storage.OAuth2Token, error) {
	t := storage.Token{}
	rows, err := s.conn.Query(tokenSelectByCodeChallengeQuery, cc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, fmt.Errorf("no entry found")
	}

	var scopeRaw string
	if err = rows.Scan(&t.ClientID, &t.AccessToken, &t.TokenType, &t.ExpiresIn, &t.RefreshToken, &scopeRaw, &t.State,
		&t.CodeChallenge, &t.AuthorizationCode); err != nil {
		return nil, err
	}
	t.Scope = storage.NewScope(scopeRaw)

	return t, nil
}

func (s *TokenStorage) FindByAccessToken(at string) (storage.OAuth2Token, error) {
	t := storage.Token{}
	rows, err := s.conn.Query(tokenSelectByAccessTokenQuery, at)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, fmt.Errorf("no entry found")
	}

	var scopeRaw string
	if err = rows.Scan(&t.ClientID, &t.AccessToken, &t.TokenType, &t.ExpiresIn, &t.RefreshToken, &scopeRaw, &t.State,
		&t.CodeChallenge, &t.AuthorizationCode); err != nil {
		return nil, err
	}
	t.Scope = storage.NewScope(scopeRaw)

	return t, nil
}

func (s *TokenStorage) Add(token storage.OAuth2Token) error {
	_, err := s.conn.Exec(tokenInsertQuery, token.GetClientID(), token.GetAccessToken(), token.GetTokenType(),
		token.GetExpiresIn(), token.GetRefreshToken(), token.GetScope().String(), token.GetState(), token.GetCodeChallenge(),
		token.GetAuthorizationCode())
	return err
}

func (s *TokenStorage) Remove(token storage.OAuth2Token) error {
	_, err := s.conn.Exec(tokenDeleteQuery, token.GetAccessToken())
	return err
}

func (s *TokenStorage) Close() error {
	return s.conn.Close()
}
