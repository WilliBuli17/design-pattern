package product

import (
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"
)

type NikeShirt struct {
	ShirtPrice int64
	constant.ShirtType
	constant.ShirtMaterial
}

func (n *NikeShirt) Price() int64 {
	return n.ShirtPrice
}

func (n *NikeShirt) IsOversize() bool {
	if n.ShirtType == constant.Oversize {
		return true
	}

	return false
}

func (n *NikeShirt) Material() constant.ShirtMaterial {
	return n.ShirtMaterial
}
