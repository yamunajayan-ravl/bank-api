package main

import (
	"encoding/json"
	"errors"
	"os"
)

const dataFile = "accounts.json"

func loadAccounts() error {
	file, err := os.Open(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()

	var accountList []Account
	err = json.NewDecoder(file).Decode(&accountList)
	if err != nil {
		return err
	}
	mu.Lock()
	defer mu.Unlock()

	for _, acc := range accountList {
		accounts[acc.ID] = acc
	}
	return nil
}

func saveAccounts() error {
	mu.Lock()
	defer mu.Unlock()
	return saveAccountsLocked()
}

func saveAccountsLocked() error {
	accountList := make([]Account, 0, len(accounts))
	for _, acc := range accounts {
		accountList = append(accountList, acc)
	}

	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(accountList)
}
