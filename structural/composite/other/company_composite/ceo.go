package company_composite

import "design-patterns/structural/composite/other"

type CEO struct {
	Subordinates []other.Employee
}

func (c CEO) GetSalary() int {
	return 10
}

func (c CEO) TotalDivisionSalary() int {
	var sum int
	for _, subordinate := range c.Subordinates {
		sum += subordinate.TotalDivisionSalary()
	}

	return c.GetSalary() + sum
}
