package dress_concrete

type ArmyDress struct {
	Color string
}

func (a *ArmyDress) GetColor() string {
	return a.Color
}

func NewArmyDress() *ArmyDress {
	return &ArmyDress{
		Color: "Blue",
	}
}
