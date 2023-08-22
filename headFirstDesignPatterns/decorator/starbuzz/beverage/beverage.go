package beverage

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type AbstractBeverage struct {
	description string
	Cost        func() float64
}

func NewBeverage() *AbstractBeverage {
	return &AbstractBeverage{"Unknown Beverage", nil}
}

func (b *AbstractBeverage) GetDescription() string {
	return b.description
}

type Expresso struct {
	*AbstractBeverage
}

func NewExpresso() *Expresso {
	base := NewBeverage()
	expresso := &Expresso{base}
	expresso.AbstractBeverage.Cost = expresso.Cost
	expresso.description = "Espresso"
	return expresso
}

func (b *Expresso) Cost() float64 {
	return 1.99
}

type HouseBlend struct {
	*AbstractBeverage
}

func NewHouseBlend() *HouseBlend {
	base := NewBeverage()
	houseBlend := &HouseBlend{base}
	houseBlend.AbstractBeverage.Cost = houseBlend.Cost
	houseBlend.description = "House Blend Coffee"
	return houseBlend
}

func (b *HouseBlend) Cost() float64 {
	return .89
}

type DarkRoast struct {
	*AbstractBeverage
}

func NewDarkRoast() *DarkRoast {
	base := NewBeverage()
	darkroast := &DarkRoast{base}
	darkroast.AbstractBeverage.Cost = darkroast.Cost
	darkroast.description = "Dark Roast"
	return darkroast
}

func (b *DarkRoast) Cost() float64 {
	return .99
}
