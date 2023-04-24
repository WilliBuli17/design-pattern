package other

import "time"

type CreatorEnt struct {
}

func (c *CreatorEnt) Produce(time time.Time) Content {
	return &MukbangContent{}
}

func (c *CreatorEnt) Type() ContentType {
	return Hiburan
}
