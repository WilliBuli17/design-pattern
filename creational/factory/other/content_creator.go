package other

import "time"

type ContentType string

const (
	Hiburan ContentType = "Hiburan"
	Edukasi ContentType = "Edukasi"
)

type ContentCreator interface {
	Produce(time time.Time) Content
	Type() ContentType
}
