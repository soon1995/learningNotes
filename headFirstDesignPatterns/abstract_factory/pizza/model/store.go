package model

type PizzaStore interface {
	OrderPizza(typ string) Pizza
	CreatePizza(typ string) Pizza
}

type AbstractPizzaStore struct {
	CreatePizza func(typ string) Pizza
}

func (p *AbstractPizzaStore) OrderPizza(typ string) Pizza {
	pizza := p.CreatePizza(typ)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
