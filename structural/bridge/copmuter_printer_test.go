package bridge

import (
	"design-patterns/structural/bridge/refactoring_guru/computer_factory"
	"design-patterns/structural/bridge/refactoring_guru/printer_factory"
	"testing"
)

var (
	hpPrinter    = &printer_factory.Hp{}
	epsonPrinter = &printer_factory.Epson{}
)

func TestComputerPrinterForMac(t *testing.T) {
	macComputer := &computer_factory.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
}

func TestComputerPrinterForWindows(t *testing.T) {
	windowsComputer := &computer_factory.Windows{}

	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
}
