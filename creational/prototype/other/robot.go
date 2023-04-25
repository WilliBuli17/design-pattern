package other

type Robot struct {
	Id, Name, processor string
}

func (r Robot) Clone() Prototype {
	return &Robot{
		Id:        r.Name,
		Name:      r.Name,
		processor: r.processor,
	}
}

func NewRobot() Robot {
	return Robot{
		Id:        "Willi-01",
		Name:      "Willi Buli",
		processor: "NVIDIA Tegra",
	}
}
