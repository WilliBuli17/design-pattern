package refactoring_guru

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun{
			Name:  "Musket",
			Power: 1,
		},
	}
}
