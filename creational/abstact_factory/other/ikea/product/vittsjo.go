package product

import "design-patterns/creational/abstact_factory/other"

type Vittsjo struct {
}

func (v *Vittsjo) Price() int64 {
	return 899_000
}

func (v *Vittsjo) Size() other.Dimension {
	return other.Dimension{
		Length: 80,
		Width:  78,
		Height: 40,
	}
}
