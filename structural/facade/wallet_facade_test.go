package facade

import (
	"design-patterns/structural/facade/refactoring_guru"
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	var (
		accountID    = "willi"
		securityCode = 123
		creditAmount = 10
		debitAmount  = 5
	)

	fmt.Println()
	walletFacade := refactoring_guru.NewWalletFacade(accountID, securityCode)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet(accountID, securityCode, creditAmount)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet(accountID, securityCode, debitAmount)
	if err != nil {
		panic(err.Error())
	}
}
