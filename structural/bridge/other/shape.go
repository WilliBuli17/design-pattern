package other

type Shape interface {
	Clone() Shape
	Draw()
	Area() float64
}
