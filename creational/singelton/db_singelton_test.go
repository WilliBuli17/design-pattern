package singelton

import (
	"design-patterns/creational/singelton/other"
	"fmt"
	"testing"
)

func TestDBSingelton(t *testing.T) {
	db := other.NewDB()
	fmt.Println(db.SayHello("Willi"))
}
