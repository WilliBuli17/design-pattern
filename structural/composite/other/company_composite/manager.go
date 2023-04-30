package company_composite

import "design-patterns/structural/composite/other"

type Manager struct {
	Subordinates []other.Employee
}

func (m Manager) GetSalary() int {
	return 7
}

func (m Manager) TotalDivisionSalary() int {
	var sum int
	for _, subordinate := range m.Subordinates {
		sum += subordinate.TotalDivisionSalary()
	}

	return m.GetSalary() + sum
}
