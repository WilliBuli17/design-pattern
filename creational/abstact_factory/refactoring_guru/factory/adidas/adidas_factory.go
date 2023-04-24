package adidas

import (
	"design-patterns/creational/abstact_factory/refactoring_guru/factory"
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/adidas/product"
)

type AdidasFactory struct {
	product.AdidasShoe
	product.AdidasShirt
}

func (a *AdidasFactory) MakeShoe() factory.Shoe {
	return &a.AdidasShoe

}

func (a *AdidasFactory) MakeShirt() factory.Shirt {
	return &a.AdidasShirt
}
