package main

import (
	"database/sql"
	"log"
	"os"
	"errors"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error

	connStr := os.Getenv("DB_CONN")
	if connStr == "" {
		log.Fatal("DB_CONN environment variable is not set")
	}

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

// create user from email and password and insert into users table and return the user id and error if any
func createUser(email, password, firstName, lastName string) (int, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return 0, err
	}

	var userID int
	query := `
	INSERT INTO users (email, password, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	err = db.QueryRow(query, email, hashedPassword, firstName, lastName).Scan(&userID)
	return userID, err
}

// create a login function that accepts email and password, checks if the user exists and the password is correct, and returns the user id, firstname and error if any
func login(email, password string) (int, string, error) {
	var hashedPassword, firstName string
	var userID int
	query := `SELECT id, password, first_name FROM users WHERE email = $1`

	err := db.QueryRow(query, email).Scan(&userID, &hashedPassword, &firstName)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", errors.New("invalid credentials")
		}
		return 0, "", err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return 0, "", errors.New("invalid credentials")
	}

	return userID, firstName, nil
}