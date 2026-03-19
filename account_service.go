package main

// func addAccount(acc Account) error {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	accounts[acc.ID] = acc

// 	if err := saveAccountsLocked(); err != nil {
// 		delete(accounts, acc.ID)
// 		return err
// 	}

// 	return nil
// }

// func getAllAccounts() []Account {

// 	mu.RLock()
// 	defer mu.RUnlock()

// 	accountList := []Account{}

// 	for _, account := range accounts {
// 		accountList = append(accountList, account)
// 	}

// 	return accountList
// }

func addAccount(acc Account) error {

	query := `
	INSERT INTO accounts (id, name, balance, cards)
	VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(query, acc.ID, acc.Name, acc.Balance, acc.Cards)
	return err
}


func getAllAccounts() ([]Account, error) {

	rows, err := db.Query("SELECT id, name, balance, cards FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var acc Account

		err := rows.Scan(&acc.ID, &acc.Name, &acc.Balance, &acc.Cards)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	return accounts, nil
}