package model

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Box()
	Cut()
	GetName() string
}

type AbstractPizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p *AbstractPizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding source...")
	fmt.Println("Adding toppings: ")
	for _, v := range p.toppings {
		fmt.Printf("\t%s\n", v)
	}
}

func (p *AbstractPizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *AbstractPizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}
func (p *AbstractPizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *AbstractPizza) GetName() string {
	return p.name
}

type CheesePizza struct{ *AbstractPizza }
type GreekPizza struct{ *AbstractPizza }
type PepperoniPizza struct{ *AbstractPizza }

// func (p *CheesePizza) Prepare() {}
// func (p *CheesePizza) Bake()    {}
// func (p *CheesePizza) Box()     {}
// func (p *CheesePizza) Cut()     {}

// func (p *GreekPizza) Prepare() {}
// func (p *GreekPizza) Bake()    {}
// func (p *GreekPizza) Box()     {}
// func (p *GreekPizza) Cut()     {}

// func (p *PepperoniPizza) Prepare() {}
// func (p *PepperoniPizza) Bake()    {}
// func (p *PepperoniPizza) Box()     {}
// func (p *PepperoniPizza) Cut()     {}

type NYStyleCheesePizza struct{ *AbstractPizza }
type NYStyleClamPizza struct{ *AbstractPizza }
type NYStylePepperoniPizza struct{ *AbstractPizza }

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	basePizza := &AbstractPizza{}
	pizza := &NYStyleCheesePizza{basePizza}
	pizza.name = "NY Style Sauce and Cheese Pizza"
	pizza.dough = "Thin Crust Dough"
	pizza.sauce = "Marinara Sauce"
	pizza.toppings = append(pizza.toppings, "Grated Reggiano Cheese")
	return pizza
}

// func (p *NYStyleCheesePizza) Prepare() {}
// func (p *NYStyleCheesePizza) Bake()    {}
// func (p *NYStyleCheesePizza) Box()     {}
// func (p *NYStyleCheesePizza) Cut()     {}

// func (p *NYStyleClamPizza) Prepare() {}
// func (p *NYStyleClamPizza) Bake()    {}
// func (p *NYStyleClamPizza) Box()     {}
// func (p *NYStyleClamPizza) Cut()     {}

// func (p *NYStylePepperoniPizza) Prepare() {}
// func (p *NYStylePepperoniPizza) Bake()    {}
// func (p *NYStylePepperoniPizza) Box()     {}
// func (p *NYStylePepperoniPizza) Cut()     {}

type ChicagoStyleCheesePizza struct{ *AbstractPizza }
type ChicagoStyleClamPizza struct{ *AbstractPizza }
type ChicagoStylePepperoniPizza struct{ *AbstractPizza }

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	basePizza := &AbstractPizza{}
	pizza := &ChicagoStyleCheesePizza{basePizza}
	pizza.name = "Chicago Style Deep Dish Cheese Pizza"
	pizza.dough = "Extra Thick Crust Doug"
	pizza.sauce = "Marinara Sauce"
	pizza.toppings = append(pizza.toppings, "Grated Reggiano Cheese")
	return pizza
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}

// func (p *ChicagoStyleCheesePizza) Prepare() {}
// func (p *ChicagoStyleCheesePizza) Bake()    {}
// func (p *ChicagoStyleCheesePizza) Box()     {}
// func (p *ChicagoStyleCheesePizza) Cut()     {}

// func (p *ChicagoStyleClamPizza) Prepare() {}
// func (p *ChicagoStyleClamPizza) Bake()    {}
// func (p *ChicagoStyleClamPizza) Box()     {}
// func (p *ChicagoStyleClamPizza) Cut()     {}

// func (p *ChicagoStylePepperoniPizza) Prepare() {}
// func (p *ChicagoStylePepperoniPizza) Bake()    {}
// func (p *ChicagoStylePepperoniPizza) Box()     {}
// func (p *ChicagoStylePepperoniPizza) Cut()     {}
