package concrete_decorator

import (
	"design-patterns/structural/decorator/other"
	"fmt"
	"math/rand"
)

type MagicKey struct {
	opener other.Opener
}

func NewMagicKeyDoor(opener other.Opener) *MagicKey {
	return &MagicKey{
		opener: opener,
	}
}

func (m *MagicKey) Open() {
	fmt.Println("This is Magic Key. Magic Key is Dangerous")
	ok := m.DoMagic()
	if ok {
		fmt.Println("Magic is Working")
		m.opener.Open()
	} else {
		m.TrigerAlaram()
	}
}

func (m *MagicKey) DoMagic() bool {
	fmt.Println("Ostium magicum patefacio")
	return rand.Int() > 0
}

func (m *MagicKey) TrigerAlaram() {
	fmt.Println("Your Magic is Not Working, Use Your Hand")
	m.opener.TrigerAlaram()
}
