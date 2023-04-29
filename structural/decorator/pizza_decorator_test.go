package decorator

import (
	"design-patterns/structural/decorator/refactoring_guru"
	"design-patterns/structural/decorator/refactoring_guru/concrete_decorator"
	"fmt"
	"testing"
)

func TestVeganPizzaDecorator(t *testing.T) {
	veganPizza := &refactoring_guru.VegePizza{}
	fmt.Printf("Price of Vegan Pizza is %d\n", veganPizza.GetPrice())

	veganPizzaWithTomatoTopping := concrete_decorator.NewTomatoToping(veganPizza)
	fmt.Printf("Price of Vegan Pizza with Tomato Topping is %d\n", veganPizzaWithTomatoTopping.GetPrice())

	veganPizzaWithCheeseTopping := concrete_decorator.NewCheeseToping(veganPizza)
	fmt.Printf("Price of Vegan Pizza with Cheese Topping is %d\n", veganPizzaWithCheeseTopping.GetPrice())

	veganPizzaWithCheeseAndTomatoTopping := concrete_decorator.NewTomatoToping(veganPizzaWithCheeseTopping)
	fmt.Printf("Price of Vegan Pizza with Tomato and Cheese Topping is %d\n", veganPizzaWithCheeseAndTomatoTopping.GetPrice())
}

func TestRegularPizzaDecorator(t *testing.T) {
	regularPizza := &refactoring_guru.RegularPizza{}
	fmt.Printf("Price of Regular Pizza is %d\n", regularPizza.GetPrice())

	regularPizzaWithTomatoTopping := concrete_decorator.NewTomatoToping(regularPizza)
	fmt.Printf("Price of Regular Pizza with Tomato Topping is %d\n", regularPizzaWithTomatoTopping.GetPrice())

	regularPizzaWithCheeseTopping := concrete_decorator.NewCheeseToping(regularPizza)
	fmt.Printf("Price of Regular Pizza with Cheese Topping is %d\n", regularPizzaWithCheeseTopping.GetPrice())

	regularPizzaWithBeefTopping := concrete_decorator.NewBeefToping(regularPizza)
	fmt.Printf("Price of Regular Pizza with Beef Topping is %d\n", regularPizzaWithBeefTopping.GetPrice())

	regularPizzaWithCheeseAndTomatoTopping := concrete_decorator.NewTomatoToping(regularPizzaWithCheeseTopping)
	fmt.Printf("Price of Regular Pizza with Tomato and Cheese Topping is %d\n", regularPizzaWithCheeseAndTomatoTopping.GetPrice())

	regularPizzaWithCheeseAndTomatoAndBeefTopping := concrete_decorator.NewBeefToping(regularPizzaWithCheeseAndTomatoTopping)
	fmt.Printf("Price of Regular Pizza with Tomato, Cheese, and Beef Topping is %d\n", regularPizzaWithCheeseAndTomatoAndBeefTopping.GetPrice())
}
