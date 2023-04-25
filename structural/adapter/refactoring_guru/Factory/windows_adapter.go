package Factory

import "fmt"

type WindowsAdapter struct {
	WindownsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindownsMachine.insertIntoUSBPort()
}
