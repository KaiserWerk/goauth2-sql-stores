package mysql

const (
	clientCreateTableQuery = `CREATE TABLE IF NOT EXISTS client (
	  id int(10) UNSIGNED NOT NULL,
	  client_id varchar(255) NOT NULL,
	  client_secret varchar(255) NOT NULL,
	  confidential tinyint(1) NOT NULL,
	  app_name varchar(255) NOT NULL,
	  urls text NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	clientSelectQuery = `SELECT client_id, client_secret, confidential, app_name, urls FROM client WHERE client_id = ?`
	clientInsertQuery = `INSERT INTO client (client_id, client_secret, confidential, app_name, urls) VALUES (?, ?, ?, ?, ?)`
	clientUpdateQuery = `UPDATE client SET client_secret = ?, confidential = ?, app_name = ?, urls = ? WHERE client_id = ?`
	clientDeleteQuery = `DELETE FROM client WHERE client_id = ?`

	sessionCreateTableQuery = `CREATE TABLE IF NOT EXISTS session (
	  id bigint(20) UNSIGNED NOT NULL,
	  session_id varchar(255) NOT NULL,
	  user_id bigint(20) UNSIGNED NOT NULL,
	  expires datetime NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	sessionSelectQuery = `SELECT session_id, user_id, expires FROM session WHERE session_id = ?`
	sessionInsertQuery = `INSERT INTO session (session_id, user_id, expires) VALUES(?, ?, ?)`
	sessionDeleteQuery = `DELETE FROM session WHERE session_id = ?`

	tokenCreateTableQuery = `CREATE TABLE IF NOT EXISTS token (
	  id bigint(20) UNSIGNED NOT NULL,
	  client_id bigint(20) UNSIGNED NOT NULL,
	  access_token varchar(255) NOT NULL,
	  token_type varchar(25) NOT NULL,
	  expires_in bigint(20) UNSIGNED NOT NULL,
	  refresh_token varchar(255) NOT NULL,
	  scope varchar(255) NOT NULL,
	  state varchar(255) NOT NULL,
	  code_challenge varchar(255) NOT NULL,
	  authorization_code varchar(255) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	tokenSelectByAccessTokenQuery   = `SELECT client_id, access_token, token_type, expires_in, refresh_token, scope, state, code_challenge, authorization_code FROM token WHERE access_token = ?`
	tokenSelectByCodeChallengeQuery = `SELECT client_id, access_token, token_type, expires_in, refresh_token, scope, state, code_challenge, authorization_code FROM token WHERE code_challenge = ?`
	tokenInsertQuery                = `INSERT INTO token (client_id, access_token, token_type, expires_in, refresh_token, scope, state, code_challenge, authorization_code) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	tokenDeleteQuery                = `DELETE FROM token WHERE access_token = ?`

	userCreateTableQuery = `CREATE TABLE IF NOT EXISTS user (
	  id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	  username varchar(255) NOT NULL,
	  email varchar(255) NOT NULL,
	  password varchar(255) NOT NULL,
	  disabled tinyint(1) NOT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	userSelectQuery           = `SELECT id, username, email, password, disabled FROM user WHERE id = ?`
	userSelectByUsernameQuery = `SELECT id, username, email, password, disabled FROM user WHERE username = ?`
	userInsertQuery           = `INSERT INTO user (username, email, password, disabled) VALUES(?, ?, ?, ?)`
	userUpdateQuery           = `UPDATE user SET username = ?, email = ?, password = ?, disabled = ? WHERE id = ?`
	userDeleteQuery           = `DELETE FROM user WHERE id = ?`
)
