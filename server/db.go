package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func setupDB() *sql.DB {
	dbPath := "data.db"
        var db *sql.DB
	// Check if the database file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil
		}
		if _, err := db.ExecContext(context.Background(), ddl); err != nil {
			return nil
		}
	} else {
		db, err = sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil
		}
	}
        return db
}
