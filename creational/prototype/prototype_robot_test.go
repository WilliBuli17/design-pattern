package prototype

import (
	"design-patterns/creational/prototype/other"
	"fmt"
	"testing"
)

func TestRobotPrototype(t *testing.T) {
	robotWilli := other.NewRobot()

	cloneRobot := robotWilli.Clone()
	robotNeo := cloneRobot.(*other.Robot)
	robotNeo.Id = "Neo-X"
	robotNeo.Name = "Neo"

	fmt.Println(robotWilli.Name)
	fmt.Println(robotNeo.Name)
}
