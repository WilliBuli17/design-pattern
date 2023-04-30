package company_composite

type BabuKroco struct {
}

func (b BabuKroco) GetSalary() int {
	return 3
}

func (b BabuKroco) TotalDivisionSalary() int {
	return b.GetSalary()
}
