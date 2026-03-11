package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", welcomeHandler)

	http.HandleFunc("/health", healthHandler)

	fmt.Println("Starting Bank API server on port 8080...")	
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bank API is healthy!")
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Bank API!")
}