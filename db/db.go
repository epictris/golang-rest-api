package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	dropTableSQL := `
	DROP TABLE IF EXISTS users;
	`
	_, err = database.Exec(dropTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = database.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	return database
}
