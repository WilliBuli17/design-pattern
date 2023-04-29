package decorator

import (
	"design-patterns/structural/decorator/other"
	"design-patterns/structural/decorator/other/concrete_decorator"
	"fmt"
	"testing"
)

func TestDoor(t *testing.T) {
	d := other.Door{}
	d.Open()

	fmt.Println("=== === === === === === === === ===")

	ed := concrete_decorator.NewElectronicKeyDoor(d)
	ed.Open()

	fmt.Println("=== === === === === === === === ===")

	md := concrete_decorator.NewMagicKeyDoor(d)
	md.Open()

	fmt.Println("=== === === === === === === === ===")

	mdxed := concrete_decorator.NewMagicKeyDoor(ed)
	mdxed.Open()
}
