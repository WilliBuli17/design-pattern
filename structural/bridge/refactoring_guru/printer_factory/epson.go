package printer_factory

import "fmt"

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
