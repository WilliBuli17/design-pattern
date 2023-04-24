package product

type Leifarne struct {
}

func (l *Leifarne) Price() int64 {
	return 1_095_000
}

func (l *Leifarne) IsErgonomic() bool {
	return false
}
