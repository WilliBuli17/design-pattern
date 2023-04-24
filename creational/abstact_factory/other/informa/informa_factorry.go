package informa

import (
	"design-patterns/creational/abstact_factory/other"
	"design-patterns/creational/abstact_factory/other/informa/product"
)

type InformaFactory struct {
}

func (i *InformaFactory) CreateChair() other.Chair {
	return &product.BeanBag{
		ProductPrice:  1_199_000,
		SoftnessLevel: 15,
	}
}

func (i *InformaFactory) CreateTable() other.Table {
	return nil
}

func (i *InformaFactory) CreateSofa() other.Sofa {
	return nil
}
