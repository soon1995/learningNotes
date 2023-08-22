package beverage

type CondimentDecorator struct {
	Beverage       Beverage
	GetDescription func() string
}

func NewCondimentDecorator() *CondimentDecorator {
	return &CondimentDecorator{}
}

type Mocha struct {
	*CondimentDecorator
}

func NewMocha(beverage Beverage) *Mocha {
	base := NewCondimentDecorator()
	mocha := &Mocha{base}
	mocha.Beverage = beverage
	mocha.CondimentDecorator.GetDescription = mocha.GetDescription
	return mocha
}

func (c *Mocha) GetDescription() string {
	return c.Beverage.GetDescription() + ", Mocha"
}

func (c *Mocha) Cost() float64 {
	return c.Beverage.Cost() + .20
}

type Soy struct {
	*CondimentDecorator
}

func NewSoy(beverage Beverage) *Soy {
	base := NewCondimentDecorator()
	soy := &Soy{base}
	soy.Beverage = beverage
	soy.CondimentDecorator.GetDescription = soy.GetDescription
	return soy
}

func (c *Soy) GetDescription() string {
	return c.Beverage.GetDescription() + ", Soy"
}

func (c *Soy) Cost() float64 {
	return c.Beverage.Cost() + .15
}

type Whip struct {
	*CondimentDecorator
}

func NewWhip(beverage Beverage) *Whip {
	base := NewCondimentDecorator()
	whip := &Whip{base}
	whip.Beverage = beverage
	whip.CondimentDecorator.GetDescription = whip.GetDescription
	return whip
}

func (c *Whip) GetDescription() string {
	return c.Beverage.GetDescription() + ", Whip"
}

func (c *Whip) Cost() float64 {
	return c.Beverage.Cost() + .10
}
