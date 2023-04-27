package wallet_subsystem

import "fmt"

type Account struct {
	Name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func (a *Account) CheckAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("account Name is incorect")
	}

	fmt.Println("Account Verified")
	return nil
}
