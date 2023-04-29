package concrete_decorator

import "design-patterns/structural/decorator/refactoring_guru"

type TomatoToping struct {
	pizza refactoring_guru.Pizza
}

func NewTomatoToping(pizza refactoring_guru.Pizza) *TomatoToping {
	return &TomatoToping{
		pizza: pizza,
	}
}

func (t *TomatoToping) GetPrice() int {
	pizzaPrice := t.pizza.GetPrice()
	return pizzaPrice + 7
}
