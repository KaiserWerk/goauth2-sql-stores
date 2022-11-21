package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/KaiserWerk/goauth2/storage"
	_ "modernc.org/sqlite"
)

type TokenStorage struct {
	conn *sql.DB
}

func (s *TokenStorage) FindByCodeChallenge(cc string) (storage.OAuth2Token, error) {
	t := storage.Token{}
	rows, err := s.conn.Query(tokenSelectByCodeChallengeQuery, cc)
	if err != nil {
		return nil, err
	}

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
