package adapter

import (
	"design-patterns/structural/adapter/refactoring_guru"
	"design-patterns/structural/adapter/refactoring_guru/Factory"
	"fmt"
	"testing"
)

func TestWindowsAdapter(t *testing.T) {
	client := &refactoring_guru.Client{}

	mac := &Factory.Mac{}
	fmt.Println(client.InsertLightningConnectorIntoComputer(mac))

	windows := &Factory.Windows{}
	windowsMachineAdapter := &Factory.WindowsAdapter{
		WindownsMachine: windows,
	}
	fmt.Println(client.InsertLightningConnectorIntoComputer(windowsMachineAdapter))
}
