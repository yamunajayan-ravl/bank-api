package main

func addAccount(acc Account) error {
	mu.Lock()
	defer mu.Unlock()

	accounts[acc.ID] = acc

	if err := saveAccountsLocked(); err != nil {
		delete(accounts, acc.ID)
		return err
	}

	return nil
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
