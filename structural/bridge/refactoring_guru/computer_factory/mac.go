package computer_factory

import (
	"design-patterns/structural/bridge/refactoring_guru"
	"fmt"
)

type Mac struct {
	printer refactoring_guru.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer refactoring_guru.Printer) {
	m.printer = printer
}
