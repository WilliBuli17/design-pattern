package other

import (
	"strings"
	"time"
)

type CreatorEdu struct {
}

func (c *CreatorEdu) Produce(time time.Time) Content {
	if strings.ToLower(time.Weekday().String()) == "sunday" {
		return &CloudContent{}
	} else if strings.ToLower(time.Weekday().String()) == "wednesday" {
		return &CodingContent{}
	}

	return nil
}

func (c *CreatorEdu) Type() ContentType {
	return Edukasi
}
