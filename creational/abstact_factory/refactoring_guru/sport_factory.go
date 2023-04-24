package refactoring_guru

import (
	"design-patterns/creational/abstact_factory/refactoring_guru/factory"
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/adidas"
	productAdidas "design-patterns/creational/abstact_factory/refactoring_guru/factory/adidas/product"
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"
	"design-patterns/creational/abstact_factory/refactoring_guru/factory/nike"
	productNike "design-patterns/creational/abstact_factory/refactoring_guru/factory/nike/product"
	"fmt"
	"strings"
)

type SportFactory interface {
	MakeShoe() factory.Shoe
	MakeShirt() factory.Shirt
}

func GetSportFactory(brand string) (SportFactory, error) {
	if strings.ToLower(brand) == "adidas" {
		return &adidas.AdidasFactory{
			AdidasShoe: productAdidas.AdidasShoe{
				ShoePrice:    1_199_000,
				ShoeMaterial: constant.Nylon,
				ShoeType:     constant.Basketball,
			},
			AdidasShirt: productAdidas.AdidasShirt{
				ShirtPrice:    399_000,
				ShirtMaterial: constant.Spandex,
				ShirtType:     constant.Undershirts,
			},
		}, nil
	} else if strings.ToLower(brand) == "nike" {
		return &nike.NikeFactory{
			NikeShoe: productNike.NikeShoe{
				ShoePrice:    1_799_000,
				ShoeMaterial: constant.Lycra,
				ShoeType:     constant.Golf,
			},
			NikeShirt: productNike.NikeShirt{
				ShirtPrice:    399_000,
				ShirtMaterial: constant.Microfiber,
				ShirtType:     constant.Oversize,
			},
		}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
