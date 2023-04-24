package nike

import (
	"design-patterns/creational/abstact_factory/refactoring_guru/factory"
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/nike/product"
)

type NikeFactory struct {
	product.NikeShoe
	product.NikeShirt
}

func (n *NikeFactory) MakeShoe() factory.Shoe {
	return &n.NikeShoe
}

func (n *NikeFactory) MakeShirt() factory.Shirt {
	return &n.NikeShirt
}
