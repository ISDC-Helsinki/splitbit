package main

import (
	"database/sql"
	"log"
        _ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(" Db Error " + err.Error())
	}
	boil.SetDB(db)
        return db
}
