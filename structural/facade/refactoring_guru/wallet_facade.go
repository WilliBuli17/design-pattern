package refactoring_guru

import (
	"design-patterns/structural/facade/refactoring_guru/wallet_subsystem"
	"fmt"
)

type WalletFacade struct {
	account      *wallet_subsystem.Account
	wallet       *wallet_subsystem.Wallet
	securityCode *wallet_subsystem.SecurityCode
	notification *wallet_subsystem.Notification
	ledger       *wallet_subsystem.Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Sending wallet debit notification")
	walletFacade := &WalletFacade{
		account:      wallet_subsystem.NewAccount(accountID),
		securityCode: wallet_subsystem.NewSecurityCode(code),
		wallet:       wallet_subsystem.NewWallet(),
		notification: &wallet_subsystem.Notification{},
		ledger:       &wallet_subsystem.Ledger{},
	}

	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}

	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
