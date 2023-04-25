package other

type Kaset struct {
	PitaMusik string
}

type Walkman struct {
}

func (w Walkman) Play(kaset Kaset) string {
	return "Playing kaset with data " + kaset.PitaMusik
}
