package wallet_subsystem

import "fmt"

type Ledger struct {
}

func (l *Ledger) MakeEntry(accountID string, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}
