package other

type Dimension struct {
	Length, Width, Height int
}

type Table interface {
	IPrice
	Size() Dimension
}
