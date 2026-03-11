package main

func addAccount(acc Account) {
	mu.Lock()
	defer mu.Unlock()

	accounts[acc.ID] = acc
}

func getAllAccounts() []Account {

	mu.RLock()
	defer mu.RUnlock()

	accountList := []Account{}

	for _, account := range accounts {
		accountList = append(accountList, account)
	}

	return accountList
}