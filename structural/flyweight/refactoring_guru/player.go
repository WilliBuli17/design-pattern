package refactoring_guru

type Player struct {
	Dress
	PlayerType string
	Lat, Long  int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, err := GetDressFactorySingleInstance().GetDressByType(dressType)
	if err != nil {
		panic(err)
	}

	return &Player{
		PlayerType: playerType,
		Dress:      dress,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.Lat = lat
	p.Long = long
}
