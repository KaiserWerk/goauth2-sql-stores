package sqlite

const (
	createQuery = `CREATE TABLE IF NOT EXISTS "client" (
	"id"	INTEGER NOT NULL UNIQUE,
	"client_id"	TEXT NOT NULL UNIQUE,
	"client_secret"	TEXT NOT NULL,
	"confidential"	NUMERIC NOT NULL DEFAULT 0,
	"app_name"	TEXT NOT NULL,
	"redirect_urls"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
)`
	selectQuery = `SELECT client_id, client_secret, confidential, app_name, redirect_urls FROM client WHERE client_id = ?`
	insertQuery = `INSERT INTO client (client_id, client_secret, confidential, app_name, redirect_urls) VALUES(?, ?, ?, ?, ?)`
	updateQuery = `UPDATE client SET client_secret = ?, confidential = ?, app_name = ?, redirect_urls = ? WHERE client_id = ?`
	deleteQuery = `DELETE FROM client WHERE client_id = ?`
)
