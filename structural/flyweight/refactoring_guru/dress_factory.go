package refactoring_guru

import (
	"design-patterns/structural/flyweight/refactoring_guru/dress_concrete"
	"fmt"
)

const (
	TerroristDressType = "Terrorist Dress"
	ArmyDressType      = "Army Dress"
)

var dressFactorySingleInstance = &DressFactory{
	DressMap: make(map[string]Dress),
}

type DressFactory struct {
	DressMap map[string]Dress
}

func (d *DressFactory) GetDressByType(dressType string) (Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.DressMap[dressType] = dress_concrete.NewTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == ArmyDressType {
		d.DressMap[dressType] = dress_concrete.NewArmyDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
