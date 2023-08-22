package model

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Box()
	Cut()
	GetName() string
	SetName(string)
}

type AbstractPizza struct {
	Name string

	Dough     Dough
	Sauce     Sauce
	Veggies   []Veggies
	Cheese    Cheese
	Pepperoni Pepperoni
	Clam      Clams

	Prepare func()
}

func (p *AbstractPizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *AbstractPizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *AbstractPizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *AbstractPizza) GetName() string {
	return p.Name
}

func (p *AbstractPizza) SetName(name string) {
	p.Name = name
}
