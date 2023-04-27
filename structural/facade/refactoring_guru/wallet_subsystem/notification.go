package wallet_subsystem

import "fmt"

type Notification struct {
}

func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
	return
}

func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
	return
}
