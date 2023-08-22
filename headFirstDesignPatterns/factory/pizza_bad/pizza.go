package pizza

type Pizza interface {
	Prepare()
	Bake()
	Box()
}

type CheesePizza struct{}
type GreekPizza struct{}
type PepperoniPizza struct{}

func OrderPizza(name string) Pizza {
	var pizza Pizza
	switch name {
	case "cheese":
		pizza = new(CheesePizza)
	case "greek":
		pizza = new(GreekPizza)
	case "pepperoni":
		pizza = new(PepperoniPizza)
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Box()
	return pizza
}

func (p *CheesePizza) Prepare() {}
func (p *CheesePizza) Bake()    {}
func (p *CheesePizza) Box()     {}

func (p *GreekPizza) Prepare() {}
func (p *GreekPizza) Bake()    {}
func (p *GreekPizza) Box()     {}

func (p *PepperoniPizza) Prepare() {}
func (p *PepperoniPizza) Bake()    {}
func (p *PepperoniPizza) Box()     {}
