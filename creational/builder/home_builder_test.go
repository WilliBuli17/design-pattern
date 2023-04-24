package builder

import (
	"design-patterns/creational/builder/ohter"
	"fmt"
	"testing"
)

func TestHomeBuilder(t *testing.T) {
	homeBuilder := ohter.HomeBuilder{}
	home := homeBuilder.
		BuildWindow().
		BuildWindow().
		BuildWindow().
		BuildWindow().
		BuildWindow().
		BuildWindow().
		BuildWindow().
		BuildWindow().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		BuildDoor().
		SetGarage(2).
		BuilderHome()

	fmt.Printf("I have a Home with \nNum Of Windows : %d \nNum Of Doors : %d \nHas Garage : %t \nHas Swimming Pool : %t \n", home.NumOfWindows, home.NumOfDoors, home.HasGarage, home.HasSwimmingPool)
}
