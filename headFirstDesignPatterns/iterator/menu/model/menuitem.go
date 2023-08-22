package model

type MenuItem struct {
	Name        string
	Description string
	Vegetarian  bool
	Price       float64
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		name,
		description,
		vegetarian,
		price,
	}
}
