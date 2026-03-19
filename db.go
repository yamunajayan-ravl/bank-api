package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error

	connStr := os.Getenv("DB_CONN")
	fmt.Println("DB_CONN:", connStr)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL")
}