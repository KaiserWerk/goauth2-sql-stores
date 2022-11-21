package mysql

const (
	createTableQuery = `CREATE TABLE client (
	  id int(10) UNSIGNED NOT NULL,
	  client_id varchar(255) NOT NULL,
	  client_secret varchar(255) NOT NULL,
	  confidential tinyint(1) NOT NULL,
	  app_name varchar(255) NOT NULL,
	  urls text NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	selectQuery = `SELECT client_id, client_secret, confidential, app_name, urls FROM client WHERE client_id = ?`
	insertQuery = `INSERT INTO client (client_id, client_secret, confidential, app_name, urls) VALUES (?, ?, ?, ?, ?)`
	updateQuery = `UPDATE client SET client_secret = ?, confidential = ?, app_name = ?, urls = ? WHERE client_id = ?`
	deleteQuery = `DELETE FROM client WHERE client_id = ?`
)
