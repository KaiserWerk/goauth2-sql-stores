package sqlite

const (
	clientCreateTableQuery = `CREATE TABLE IF NOT EXISTS "client" (
		"id"	INTEGER NOT NULL UNIQUE,
		"client_id"	TEXT NOT NULL UNIQUE,
		"client_secret"	TEXT NOT NULL,
		"confidential"	NUMERIC NOT NULL DEFAULT 0,
		"app_name"	TEXT NOT NULL,
		"redirect_urls"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`
	clientSelectQuery = `SELECT client_id, client_secret, confidential, app_name, redirect_urls FROM client WHERE client_id = ?`
	clientInsertQuery = `INSERT INTO client (client_id, client_secret, confidential, app_name, redirect_urls) VALUES(?, ?, ?, ?, ?)`
	clientUpdateQuery = `UPDATE client SET client_secret = ?, confidential = ?, app_name = ?, redirect_urls = ? WHERE client_id = ?`
	clientDeleteQuery = `DELETE FROM client WHERE client_id = ?`

	sessionCreateTableQuery = `CREATE TABLE IF NOT EXISTS "session" (
		"id"	INTEGER NOT NULL UNIQUE,
		"session_id"	TEXT NOT NULL UNIQUE,
		"user_id"	NUMERIC NOT NULL,
		"expires"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`
	sessionSelectQuery = `SELECT session_id, user_id, expires FROM session WHERE session_id = ?`
	sessionInsertQuery = `INSERT INTO session (session_id, user_id, expires) VALUES(?, ?, ?)`
	sessionDeleteQuery = `DELETE FROM session WHERE session_id = ?`

	tokenCreateTableQuery = `CREATE TABLE IF NOT EXISTS "token" (
		"id"	INTEGER NOT NULL UNIQUE,
		"client_id"	TEXT NOT NULL UNIQUE,
		"access_token"	TEXT NOT NULL,
		"token_type"	TEXT NOT NULL,
		"expires_in"	numeric NOT NULL,
		"refresh_token"	TEXT NOT NULL,
		"scope"	TEXT NOT NULL,
		"state"	TEXT NOT NULL,
		"code_challenge"	TEXT NOT NULL,
		"authorization_code"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`
	tokenSelectByAccessTokenQuery   = `SELECT client_id, access_token, token_type, expires_in, refresh_token, scope, state, code_challenge, authorization_code FROM token WHERE access_token = ?`
	tokenSelectByCodeChallengeQuery = `SELECT client_id, access_token, token_type, expires_in, refresh_token, scope, state, code_challenge, authorization_code FROM token WHERE code_challenge = ?`
	tokenInsertQuery                = `INSERT INTO token (client_id, access_token, token_type, expires_in, refresh_token, scope, state, code_challenge, authorization_code) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	tokenDeleteQuery                = `DELETE FROM token WHERE access_token = ?`

	userCreateTableQuery = `CREATE TABLE IF NOT EXISTS user (
	  id int(10) UNSIGNED NOT NULL,
	  username int(255) NOT NULL,
	  email int(255) NOT NULL,
	  password varchar(255) NOT NULL,
	  disabled tinyint(1) NOT NULL DEFAULT '0'
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	userSelectQuery           = `SELECT id, username, email, password, disabled FROM user WHERE id = ?`
	userSelectByUsernameQuery = `SELECT id, username, email, password, disabled FROM user WHERE username = ?`
	userInsertQuery           = `INSERT INTO user (username, email, password, disabled) VALUES(?, ?, ?, ?)`
	userUpdateQuery           = `UPDATE user SET username = ?, email = ?, password = ?, disabled = ? WHERE id = ?`
	userDeleteQuery           = `DELETE FROM user WHERE id = ?`
)
