package other

import "fmt"

type Door struct {
}

func (d Door) Open() {
	fmt.Println("Door is Open")
}

func (d Door) TrigerAlaram() {
	fmt.Println("Failed to Open the Door")
}
