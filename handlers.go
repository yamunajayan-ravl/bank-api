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

	err = addAccount(newAccount)
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
