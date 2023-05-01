package flyweight

import (
	"design-patterns/structural/flyweight/refactoring_guru"
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	game := refactoring_guru.NewGame()
	var err error

	//Add Terrorist
	err = game.AddPlayer("Terrorist")
	err = game.AddPlayer("Terrorist")
	err = game.AddPlayer("Terrorist")
	err = game.AddPlayer("Terrorist")
	err = game.AddPlayer("Terrorist")
	err = game.AddPlayer("Terrorist")

	//Add Army
	err = game.AddPlayer("Army")
	err = game.AddPlayer("Army")
	err = game.AddPlayer("Army")
	err = game.AddPlayer("Army")

	if err != nil {
		panic(err)
	}

	dressFactoryInstance := refactoring_guru.GetDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
