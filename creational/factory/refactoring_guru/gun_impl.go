package refactoring_guru

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) GetPower() int {
	return g.Power
}
