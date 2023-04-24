package product

import "design-patterns/creational/abstact_factory/other"

type Hemlingby struct {
}

func (h *Hemlingby) Price() int64 {
	return 1_795_000
}

func (h *Hemlingby) Style() other.SofaStyle {
	return other.EuropeanStyle
}
