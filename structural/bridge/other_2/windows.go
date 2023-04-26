package other_2

import "fmt"

type Windows struct{}

func (w *Windows) OpenLink() {
	fmt.Println("opening a link in windows")
}
