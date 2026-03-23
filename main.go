package main

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/joho/godotenv"
)

var accounts = make(map[int]Account)
var mu sync.RWMutex

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bank API is healthy!") 
	
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Bank API!")
}


func main() {

	err :=godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	initDB()

	http.HandleFunc("/", welcomeHandler)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/accounts", accountsHandler)

	fmt.Println("Starting Bank API server on port 8080...")	
	http.ListenAndServe(":8080", nil)
}

