package store

import (
	"fmt"

	ingredientfactory "example.com/ingredient_factory"
	"example.com/model"
	"example.com/model/pizza"
)

type NYStylePizzaStore struct {
	*model.AbstractPizzaStore
}

func NewNYStylePizzaStore() *NYStylePizzaStore {
	basePizzaStore := &model.AbstractPizzaStore{}
	nyPizzaStore := &NYStylePizzaStore{basePizzaStore}
	nyPizzaStore.AbstractPizzaStore.CreatePizza = nyPizzaStore.CreatePizza
	return nyPizzaStore
}

func (p *NYStylePizzaStore) CreatePizza(typ string) model.Pizza {
	ingredientFac := &ingredientfactory.NYPizzaIngredientFactory{}
	var piz model.Pizza
	switch typ {
	case "cheese":
		piz = pizza.NewCheesePizza(ingredientFac)
		piz.SetName("New York Style Cheese Pizza")
	case "clam":
		piz = pizza.NewClamPizza(ingredientFac)
		piz.SetName("New York Style Clam Pizza")
	default:
		panic(fmt.Sprintf("%s not implemented", typ))
	}
	return piz
}
