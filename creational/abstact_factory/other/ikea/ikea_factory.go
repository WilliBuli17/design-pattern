package ikea

import (
	"design-patterns/creational/abstact_factory/other"
	"design-patterns/creational/abstact_factory/other/ikea/product"
)

type IkeaFactory struct {
}

func (i *IkeaFactory) CreateChair() other.Chair {
	return &product.Leifarne{}
}

func (i *IkeaFactory) CreateTable() other.Table {
	return &product.Vittsjo{}
}

func (i *IkeaFactory) CreateSofa() other.Sofa {
	return &product.Hemlingby{}
}
