package main

import (
	"net/http"
)

func accountsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		createAccountHandler(w, r)

	case http.MethodGet:
		getAccountHandler(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {

	var newAccount Account

	err := readJSON(r, &newAccount)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = addAccount(&newAccount)
	if err != nil {
		http.Error(w, "Failed to save account", http.StatusInternalServerError)
		return
	}
	writeJSON(w, newAccount)
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {

	accounts, err := getAllAccounts()
	if err != nil {
		http.Error(w, "Failed to retrieve accounts", http.StatusInternalServerError)
		return
	}

	writeJSON(w, accounts)
}

// add a register end point that accepts a JSON payload with username and password, then calls createUser function
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
	}

	err := readJSON(r, &newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := createUser(newUser.Email, newUser.Password, newUser.FirstName, newUser.LastName)
	if err != nil {
        if err.Error() == "user already exists" {
            http.Error(w, err.Error(), http.StatusConflict)
            return
        }
        http.Error(w, "Server error", http.StatusInternalServerError)
        return
	}

	writeJSON(w, map[string]interface{}{
		"message": "User created successfully",
		"userID":  userID,
	})
}

// add a login end point that accepts a JSON payload with email and password, then calls login function and returns the user id and firstname if successful
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := readJSON(r, &credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, firstName, err := login(credentials.Email, credentials.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	
	writeJSON(w, map[string]interface{}{
		"message": "Welcome " + firstName,
		"userID":  userID,
	})
}
