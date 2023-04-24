package builder

import (
	"design-patterns/creational/builder/refactoring_guru"
	"design-patterns/creational/builder/refactoring_guru/constant"
	"fmt"
	"testing"
)

func TestCarBuilder(t *testing.T) {
	maticCarBuilder := refactoring_guru.GetCarBuilder("matic")
	maticCarBuilder.SetSeat(4)
	maticCarBuilder.SetEngine(constant.Electric)
	maticCarBuilder.SetGPS(true)
	maticCarBuilder.SetTripComputer(true)
	maticCar := maticCarBuilder.GetCar()

	manualCarBuilder := refactoring_guru.GetCarBuilder("manual")
	manualCarBuilder.SetSeat(6)
	manualCarBuilder.SetEngine(constant.Diesel)
	manualCarBuilder.SetGPS(false)
	manualCarBuilder.SetTripComputer(true)
	manualCar := manualCarBuilder.GetCar()

	descCar(maticCar)
	descCar(manualCar)
}

func descCar(car refactoring_guru.Car) {
	fmt.Printf("I Have a Car with \n\tNum Of Seat : %d\n\tEngine Type : %s\n\tHave Trip Computer? : %t\n\tHave GPS? : %t\n\n", car.NumOfSeat, car.CarEngine, car.IsTripComputer, car.IsGPS)
}
