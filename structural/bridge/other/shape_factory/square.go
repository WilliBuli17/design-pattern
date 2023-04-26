package shape_factory

import (
	"design-patterns/structural/bridge/other"
	"fmt"
	"math"
)

type Square struct {
	other.Color
	Length float64
}

func (s Square) Draw() {
	if s.Color != nil {
		fmt.Printf("drawing a circle with square %s\n", s.Color.Hex())
	} else {
		fmt.Println("drawing a square")
	}
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

func (s Square) Clone() other.Shape {
	return Square{
		Color:  s.Color,
		Length: s.Length,
	}
}
