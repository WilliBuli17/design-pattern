package other_2

type Chrome interface {
	OpenURL()
}

type Beta struct {
	OS OS
}

func (b *Beta) OpenURL() {
	b.OS.OpenLink()
}

func NewChromeBeta(os OS) Chrome {
	return &Beta{
		OS: os,
	}
}
