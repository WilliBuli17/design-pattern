package refactoring_guru

import (
	"design-patterns/creational/builder/refactoring_guru/constant"
)

type MeticCarBuilder struct {
	Car
}

func NewMeticCarBuilder() *MeticCarBuilder {
	return &MeticCarBuilder{}
}

func (m *MeticCarBuilder) SetSeat(numOfSeat int) {
	m.NumOfSeat = numOfSeat
}

func (m *MeticCarBuilder) SetEngine(engine constant.CarEngine) {
	m.CarEngine = engine
}

func (m *MeticCarBuilder) SetTripComputer(tripComputer bool) {
	m.IsTripComputer = tripComputer
}

func (m *MeticCarBuilder) SetGPS(gps bool) {
	m.IsGPS = gps
}

func (m *MeticCarBuilder) GetCar() Car {
	return m.Car
}
