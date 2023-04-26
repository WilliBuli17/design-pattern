package shape_factory

import (
	"design-patterns/structural/bridge/other"
	"fmt"
	"math"
)

type Circle struct {
	other.Color
	Radius float64
}

func (c Circle) Draw() {
	if c.Color != nil {
		fmt.Printf("drawing a circle with color %s\n", c.Color.Hex())
	} else {
		fmt.Println("drawing a circle")
	}
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Clone() other.Shape {
	return Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}
