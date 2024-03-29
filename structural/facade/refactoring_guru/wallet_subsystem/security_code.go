package wallet_subsystem

import "fmt"

type SecurityCode struct {
	Code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		Code: code,
	}
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.Code != incomingCode {
		return fmt.Errorf("security code is incorect")
	}

	fmt.Println("Security Code Verified")
	return nil
}
