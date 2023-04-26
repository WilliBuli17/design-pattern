package computer_factory

import (
	"design-patterns/structural/bridge/refactoring_guru"
	"fmt"
)

type Windows struct {
	printer refactoring_guru.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer refactoring_guru.Printer) {
	w.printer = printer
}
