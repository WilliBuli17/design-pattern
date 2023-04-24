package refactoring_guru

import (
	"design-patterns/creational/builder/refactoring_guru/constant"
)

type ManualCarBuilder struct {
	Car
}

func NewManualCarBuilder() *ManualCarBuilder {
	return &ManualCarBuilder{}
}

func (m *ManualCarBuilder) SetSeat(numOfSeat int) {
	m.NumOfSeat = numOfSeat
}

func (m *ManualCarBuilder) SetEngine(engine constant.CarEngine) {
	m.CarEngine = engine
}

func (m *ManualCarBuilder) SetTripComputer(tripComputer bool) {
	m.IsTripComputer = tripComputer
}

func (m *ManualCarBuilder) SetGPS(gps bool) {
	m.IsGPS = gps
}

func (m *ManualCarBuilder) GetCar() Car {
	return m.Car
}
