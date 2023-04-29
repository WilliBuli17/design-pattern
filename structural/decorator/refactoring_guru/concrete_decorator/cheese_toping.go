package concrete_decorator

import "design-patterns/structural/decorator/refactoring_guru"

type CheeseToping struct {
	pizza refactoring_guru.Pizza
}

func NewCheeseToping(pizza refactoring_guru.Pizza) *CheeseToping {
	return &CheeseToping{
		pizza: pizza,
	}
}

func (c *CheeseToping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 10
}
