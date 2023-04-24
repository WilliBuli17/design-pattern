package refactoring_guru

import "design-patterns/creational/builder/refactoring_guru/constant"

type Car struct {
	NumOfSeat int
	constant.CarEngine
	IsTripComputer, IsGPS bool
}
