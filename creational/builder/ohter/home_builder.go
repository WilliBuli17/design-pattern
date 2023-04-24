package ohter

import (
	"strings"
)

type HomeInterface interface {
	BuildWindow() *HomeBuilder
	BuildDoor() *HomeBuilder
	SetGarage(numOfCar int) *HomeBuilder
	SetSwimmingPool(types string) *HomeBuilder
	BuilderHome() *Home
}

type HomeBuilder struct {
	Home
}

func (h *HomeBuilder) BuildWindow() *HomeBuilder {
	if h.NumOfWindows < 5 {
		h.NumOfWindows++
	}

	return h
}

func (h *HomeBuilder) BuildDoor() *HomeBuilder {
	h.NumOfDoors++
	return h
}

func (h *HomeBuilder) SetGarage(numOfCar int) *HomeBuilder {
	h.HasGarage = false

	if numOfCar > 1 {
		h.HasGarage = true
	}

	return h
}

func (h *HomeBuilder) SetSwimmingPool(types string) *HomeBuilder {
	h.HasSwimmingPool = false
	if strings.ToLower(types) == "yes" {
		h.HasSwimmingPool = true
	}

	return h
}

func (h *HomeBuilder) BuilderHome() *Home {
	return &h.Home
}
