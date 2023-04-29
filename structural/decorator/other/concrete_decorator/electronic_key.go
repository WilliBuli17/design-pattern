package concrete_decorator

import (
	"design-patterns/structural/decorator/other"
	"fmt"
)

type ElectronicKey struct {
	opener other.Opener
}

func NewElectronicKeyDoor(opener other.Opener) *ElectronicKey {
	return &ElectronicKey{
		opener: opener,
	}
}

func (e *ElectronicKey) Open() {
	fmt.Println("This is Electronic Key")
	ok := e.ConnectToWifi()
	if ok {
		fmt.Println("Success Connecting to Wifi")
		e.opener.Open()
	} else {
		e.TrigerAlaram()
	}
}

func (e *ElectronicKey) ConnectToWifi() bool {
	fmt.Println("Connecting to Wifi")
	return true
}

func (e *ElectronicKey) TrigerAlaram() {
	fmt.Println("Failed Connecting to Wifi")
	e.opener.TrigerAlaram()
}
