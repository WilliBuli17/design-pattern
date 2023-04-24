package factory

import (
	"design-patterns/creational/factory/other"
	"fmt"
	"testing"
	"time"
)

func TestContentFactory(t *testing.T) {
	contentCreator := &other.CreatorEdu{}
	//contentCreator := &CreatorEnt{}

	contentType := contentCreator.Type()
	contentDesc := contentCreator.Produce(time.Now()).Play()

	fmt.Println(contentType)
	fmt.Println(contentDesc)
}
