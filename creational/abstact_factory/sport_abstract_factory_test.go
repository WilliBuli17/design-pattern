package abstact_factory

import (
	"design-patterns/creational/abstact_factory/refactoring_guru"
	"fmt"
	"testing"
)

func TestSportFactory(t *testing.T) {
	adidas, _ := refactoring_guru.GetSportFactory("adidas")
	nike, _ := refactoring_guru.GetSportFactory("nike")

	getProduct(adidas)
	getProduct(nike)
}

func getProduct(factory refactoring_guru.SportFactory) {
	shoe := factory.MakeShoe()
	shirt := factory.MakeShirt()

	fmt.Printf("Shoe Price : %d, Shoe Type : %s, is Nylon Material : %t \n", shoe.Price(), shoe.Type(), shoe.IsNylon())
	fmt.Printf("Shirt Price : %d, Shirt Material : %s, is Oversize Type : %t \n\n", shirt.Price(), shirt.Material(), shirt.IsOversize())
}
