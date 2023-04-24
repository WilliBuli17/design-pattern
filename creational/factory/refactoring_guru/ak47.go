package refactoring_guru

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun{
			Name:  "AK47",
			Power: 4,
		},
	}
}
