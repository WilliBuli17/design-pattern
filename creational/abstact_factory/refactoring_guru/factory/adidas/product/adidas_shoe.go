package product

import "design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"

type AdidasShoe struct {
	ShoePrice int64
	constant.ShoeMaterial
	constant.ShoeType
}

func (a *AdidasShoe) Price() int64 {
	return a.ShoePrice
}

func (a *AdidasShoe) Type() constant.ShoeType {
	return a.ShoeType
}

func (a *AdidasShoe) IsNylon() bool {
	if a.ShoeMaterial == constant.Nylon {
		return true
	}

	return false
}
