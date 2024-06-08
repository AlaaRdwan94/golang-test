package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}
