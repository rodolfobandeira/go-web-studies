package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// SqliteConnection - Connects to Sqlite3
func SqliteConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/store.sq3")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
