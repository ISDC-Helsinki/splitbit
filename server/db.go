package main

import (
	"database/sql"
	"log"
        _ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func setupDB() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(" Db Error " + err.Error())
	}
	boil.SetDB(db)
}
