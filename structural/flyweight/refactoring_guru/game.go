package refactoring_guru

import (
	"fmt"
	"strings"
)

type Game struct {
	Terrorist, Army []*Player
}

func NewGame() *Game {
	return &Game{
		Terrorist: make([]*Player, 1),
		Army:      make([]*Player, 1),
	}
}

func (g Game) AddPlayer(playerType string) error {
	if strings.ToLower(playerType) == "terrorist" {
		player := NewPlayer(strings.ToUpper(playerType), TerroristDressType)
		g.Terrorist = append(g.Terrorist, player)
		return nil
	}
	if strings.ToLower(playerType) == "army" {
		player := NewPlayer(strings.ToUpper(playerType), ArmyDressType)
		g.Army = append(g.Army, player)
		return nil
	}

	return fmt.Errorf("wrong player type passed")
}
