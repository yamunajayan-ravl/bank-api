package main

import (
	"net/http"
	"encoding/json"
)


func createAccountHandler(w http.ResponseWriter, r *http.Request) {

	var newAccount Account

	err := json.NewDecoder(r.Body).Decode(&newAccount)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	mu.Lock()
	accounts[newAccount.ID] = newAccount
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAccount)
}