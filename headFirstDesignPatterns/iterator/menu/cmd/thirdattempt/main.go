package main

import (
	"fmt"

	"example.com/model"
)

func main() {
	pancakeHouseMenu := model.NewPancakeHouseMenu()
	dinerMenu := model.NewDinerMenu()
	cafeMenu := model.NewCafeMenu()
	var menus []model.Menu
	menus = append(menus, pancakeHouseMenu, dinerMenu, cafeMenu)
	waitress := NewWaitress(menus)
	waitress.PrintMenu()
}

type Waitress struct {
	menus []model.Menu
}

func NewWaitress(
	menus []model.Menu,
) *Waitress {
	return &Waitress{
		menus: menus,
	}
}

func (w *Waitress) PrintMenu() {
	for _, m := range w.menus {
		w.PrintMenuItems(m.CreateIterator())
	}
}

func (w *Waitress) PrintMenuItems(iterator model.Iterator) {
	for iterator.HasNext() {
		item := iterator.Next().(*model.MenuItem)
		fmt.Printf("%s, %.2f -- %s\n", item.Name, item.Price, item.Description)
	}
}
