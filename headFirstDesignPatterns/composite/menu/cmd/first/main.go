package main

import "example.com/model"

func main() {
	var pancakeHouseMenu model.IMenuComponent = model.NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	var dinerMenu model.IMenuComponent = model.NewMenu("DINER MENU", "Lunch")
	var cafeMenu model.IMenuComponent = model.NewMenu("CAFE MENU", "Dinner")
	var dessertMenu model.IMenuComponent = model.NewMenu("DESSERT MENU", "Dessert of course!")
	var allMenus model.IMenuComponent = model.NewMenu("ALL MENUS", "All menus combined")
	allMenus.Add(pancakeHouseMenu)
	allMenus.Add(dinerMenu)
	allMenus.Add(cafeMenu)
	// add menu items here
	dinerMenu.Add(model.NewMenuItem(
		"Pasta",
		"Spaghetti with Marinara Sauce, and a slice of sourdough bread",
		true,
		3.89))
	dinerMenu.Add(dessertMenu)
	dessertMenu.Add(model.NewMenuItem(
		"Apple Pie",
		"Apple pie with a flakey crust, topped with vanilla ice cream",
		true,
		1.59))
	// add more menu items here
	waitress := NewWaitress(allMenus)
	waitress.printMenu()

}

type Waitress struct {
	allMenus model.IMenuComponent
}

func NewWaitress(allMenus model.IMenuComponent) *Waitress {
	return &Waitress{
		allMenus: allMenus,
	}
}

func (w *Waitress) printMenu() {
	w.allMenus.Print()
}
