package other_2

import "fmt"

type Ubuntu struct{}

func (u *Ubuntu) OpenLink() {
	fmt.Println("opening a link in ubuntu")
}
