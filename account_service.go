package main

func addAccount(acc *Account) error {

	query := `
	INSERT INTO accounts (name, balance, cards)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	return db.QueryRow(query, acc.Name, acc.Balance, acc.Cards).Scan(&acc.ID)
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