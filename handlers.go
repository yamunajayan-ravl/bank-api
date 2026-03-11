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

func getAccountHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.RLock()

	accountList := []Account{}
	for _, account := range accounts {
		accountList = append(accountList, account)
	}
	mu.RUnlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accountList)
}
	