package store

import (
	"fmt"

	"example.com/model"
)

type PizzaStore interface {
	OrderPizza(typ string) model.Pizza
	createPizza(typ string) model.Pizza
}

type abstractPizzaStore struct {
	createPizza func(typ string) model.Pizza
}

func (p *abstractPizzaStore) OrderPizza(typ string) model.Pizza {
	pizza := p.createPizza(typ)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type NYStylePizzaStore struct {
	*abstractPizzaStore
}

func NewNYStylePizzaStore() *NYStylePizzaStore {
	basePizzaStore := &abstractPizzaStore{}
	nyPizzaStore := &NYStylePizzaStore{basePizzaStore}
	nyPizzaStore.abstractPizzaStore.createPizza = nyPizzaStore.createPizza
	return nyPizzaStore
}

func (p *NYStylePizzaStore) createPizza(typ string) model.Pizza {
	switch typ {
	case "cheese":
		return model.NewNYStyleCheesePizza()
	case "clam":
		return new(model.NYStyleClamPizza)
	case "pepperoni":
		return new(model.NYStylePepperoniPizza)
	}
	panic(fmt.Sprintf("%s not implemented", typ))
}

type ChicagoStylePizzaStore struct {
	*abstractPizzaStore
}

func NewChicagoStylePizzaStore() *ChicagoStylePizzaStore {
	basePizzaStore := &abstractPizzaStore{}
	chicagoPizzaStore := &ChicagoStylePizzaStore{basePizzaStore}
	chicagoPizzaStore.abstractPizzaStore.createPizza = chicagoPizzaStore.createPizza
	return chicagoPizzaStore
}

func (p *ChicagoStylePizzaStore) createPizza(typ string) model.Pizza {
	switch typ {
	case "cheese":
		return model.NewChicagoStyleCheesePizza()
	case "clam":
		return new(model.ChicagoStyleClamPizza)
	case "pepperoni":
		return new(model.ChicagoStylePepperoniPizza)
	}
	panic(fmt.Sprintf("%s not implemented", typ))
}
