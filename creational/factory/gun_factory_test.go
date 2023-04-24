package factory

import (
	"design-patterns/creational/factory/refactoring_guru"
	"fmt"
	"testing"
)

func TestGunFactory(t *testing.T) {
	ak47, err := refactoring_guru.GetGun("ak47")
	musket, err := refactoring_guru.GetGun("musket")

	if err != nil {
		panic(err)
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun refactoring_guru.IGun) {
	fmt.Printf("Gun %s, Power %d \n", gun.GetName(), gun.GetPower())
}
