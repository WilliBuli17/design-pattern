package refactoring_guru

import (
	"design-patterns/creational/builder/refactoring_guru/constant"
	"strings"
)

type CarBuilder interface {
	SetSeat(numOfSeat int)
	SetEngine(engine constant.CarEngine)
	SetTripComputer(tripComputer bool)
	SetGPS(gps bool)
	GetCar() Car
}

func GetCarBuilder(builderType string) CarBuilder {
	if strings.ToLower(builderType) == "manual" {
		return NewManualCarBuilder()
	} else if strings.ToLower(builderType) == "matic" {
		return NewMeticCarBuilder()
	}
	return nil
}
