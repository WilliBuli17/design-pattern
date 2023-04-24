package refactoring_guru

import (
	"fmt"
	"strings"
)

func GetGun(gunType string) (IGun, error) {
	if strings.ToLower(gunType) == "ak47" {
		return NewAk47(), nil
	} else if strings.ToLower(gunType) == "musket" {
		return NewMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
