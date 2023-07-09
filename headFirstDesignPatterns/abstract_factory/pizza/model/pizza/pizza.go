package pizza

import (
	"fmt"

	ingredientfactory "example.com/ingredient_factory"
	"example.com/model"
)

type CheesePizza struct {
	fac ingredientfactory.PizzaIngredientFactory
	*model.AbstractPizza
}

func NewCheesePizza(fac ingredientfactory.PizzaIngredientFactory) *CheesePizza {
	base := &model.AbstractPizza{}
	pizza := &CheesePizza{fac: fac, AbstractPizza: base}
	pizza.AbstractPizza.Prepare = pizza.Prepare
	return pizza
}

func (p *CheesePizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.Name)
	p.Dough = p.fac.CreateDough()
	p.Sauce = p.fac.CreateSauce()
	p.Cheese = p.fac.CreateCheese()
}

type ClamPizza struct {
	fac ingredientfactory.PizzaIngredientFactory
	*model.AbstractPizza
}

func NewClamPizza(fac ingredientfactory.PizzaIngredientFactory) *ClamPizza {
	base := &model.AbstractPizza{}
	pizza := &ClamPizza{fac: fac, AbstractPizza: base}
	pizza.AbstractPizza.Prepare = pizza.Prepare
	return pizza
}

func (p *ClamPizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.Name)
	p.Dough = p.fac.CreateDough()
	p.Sauce = p.fac.CreateSauce()
	p.Cheese = p.fac.CreateCheese()
	p.Clam = p.fac.CreateCheese()
}
