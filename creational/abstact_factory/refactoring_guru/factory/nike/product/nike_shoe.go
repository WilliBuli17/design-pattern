package product

import (
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"
)

type NikeShoe struct {
	ShoePrice int64
	constant.ShoeMaterial
	constant.ShoeType
}

func (n *NikeShoe) Price() int64 {
	return n.ShoePrice
}

func (n *NikeShoe) Type() constant.ShoeType {
	return n.ShoeType
}

func (n *NikeShoe) IsNylon() bool {
	if n.ShoeMaterial == constant.Nylon {
		return true
	}

	return false
}
