package bridge

import (
	"design-patterns/structural/bridge/other/color_factory"
	"design-patterns/structural/bridge/other/shape_factory"
	"fmt"
	"testing"
)

func TestCircleColorShape(t *testing.T) {
	circle := shape_factory.Circle{
		Radius: 10,
	}
	circle.Draw()
	fmt.Println(circle.Area())

	redCircle := shape_factory.Circle{
		Radius: 12,
		Color:  color_factory.Red{},
	}
	redCircle.Draw()
	fmt.Println(redCircle.Area())

	// ini digabungkan dengan prototype pattern
	circleClone := redCircle.Clone()
	greenCircleClone := circleClone.(shape_factory.Circle)
	greenCircleClone.Color = color_factory.Green{}
	greenCircleClone.Draw()
	fmt.Println(greenCircleClone.Area())
}

func TestSquareColorShape(t *testing.T) {
	square := shape_factory.Square{
		Length: 10,
	}
	square.Draw()
	fmt.Println(square.Area())

	greenSquare := shape_factory.Square{
		Length: 12,
		Color:  color_factory.Green{},
	}
	greenSquare.Draw()
	fmt.Println(greenSquare.Area())

	// ini digabungkan dengan prototype pattern
	cloneSquare := square.Clone()
	redSQuareClone := cloneSquare.(shape_factory.Square)
	redSQuareClone.Color = color_factory.Red{}
	redSQuareClone.Draw()
	fmt.Println(redSQuareClone.Area())
}
