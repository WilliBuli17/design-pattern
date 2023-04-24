package product

import "design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"

type AdidasShirt struct {
	ShirtPrice int64
	constant.ShirtType
	constant.ShirtMaterial
}

func (a *AdidasShirt) Price() int64 {
	return a.ShirtPrice
}

func (a *AdidasShirt) IsOversize() bool {
	if a.ShirtType == constant.Oversize {
		return true
	}

	return false
}

func (a *AdidasShirt) Material() constant.ShirtMaterial {
	return a.ShirtMaterial
}
