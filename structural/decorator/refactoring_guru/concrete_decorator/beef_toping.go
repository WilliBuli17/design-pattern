package concrete_decorator

import "design-patterns/structural/decorator/refactoring_guru"

type BeefToping struct {
	pizza refactoring_guru.Pizza
}

func NewBeefToping(pizza refactoring_guru.Pizza) *BeefToping {
	return &BeefToping{
		pizza: pizza,
	}
}

func (b *BeefToping) GetPrice() int {
	pizzaPrice := b.pizza.GetPrice()
	return pizzaPrice + 15
}
