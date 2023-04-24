package abstact_factory_test

import (
	"design-patterns/creational/abstact_factory/other"
	"design-patterns/creational/abstact_factory/other/ikea"
	"design-patterns/creational/abstact_factory/other/informa"
	"fmt"
	"testing"
)

func TestFurnitureFactory(t *testing.T) {
	ikeaFurniture := new(ikea.IkeaFactory)
	informaFurniture := new(informa.InformaFactory)

	createFurniture(ikeaFurniture)
	createFurniture(informaFurniture)
}

func createFurniture(factory other.FurnitureFactory) {
	chair := factory.CreateChair()
	table := factory.CreateTable()
	sofa := factory.CreateSofa()

	if chair != nil {
		fmt.Printf("Chair Price %d, is ergonomic : %t \n", chair.Price(), chair.IsErgonomic())
	}
	if table != nil {
		fmt.Printf("Table Price %d, size : %dcm x %dcm x %dcm \n", table.Price(), table.Size().Length, table.Size().Width, table.Size().Height)
	}
	if sofa != nil {
		fmt.Printf("Sofa Price %d, style : %s \n\n", sofa.Price(), sofa.Style())
	}

}
