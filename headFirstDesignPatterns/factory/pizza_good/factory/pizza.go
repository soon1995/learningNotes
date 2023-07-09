package factory

import (
	"fmt"

	"example.com/model"
)

type SimplePizzaFactory struct{}

func (f *SimplePizzaFactory) CreatePizza(typ string) model.Pizza {
	switch typ {
	case "cheese":
		return new(model.CheesePizza)
	case "greek":
		return new(model.GreekPizza)
	case "pepperoni":
		return new(model.PepperoniPizza)
	}
	panic(fmt.Sprintf("%s not implemented", typ))
}
