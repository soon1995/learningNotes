package model

import "fmt"

type MenuItem struct {
	Name        string
	Description string
	Vegetarian  bool
	Price       float64
	*MenuComponent
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	menuComponent := &MenuComponent{}
	menuItem := &MenuItem{
		Name:        name,
		Description: description,
		Vegetarian:  vegetarian,
		Price:       price,
	}
	menuComponent.GetDescFn = menuItem.GetDescription
	menuComponent.GetPriceFn = menuItem.GetPrice
	menuComponent.IsVegeFn = menuItem.IsVegetarian
	menuComponent.PrintFn = menuItem.Print
	menuItem.MenuComponent = menuComponent
	return menuItem
}

func (mi *MenuItem) GetDescription() string {
	return mi.Description
}

func (mi *MenuItem) GetPrice() float64 {
	return mi.Price
}

func (mi *MenuItem) IsVegetarian() bool {
	return mi.Vegetarian
}

func (mi *MenuItem) Print() {
	fmt.Printf("  %s", mi.Name)
	if mi.IsVegetarian() {
		fmt.Printf("(v)")
	}
	fmt.Printf(", %.2f\n   --%s\n", mi.GetPrice(), mi.GetDescription())
}
