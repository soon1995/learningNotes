package main

import (
	"fmt"

	"example.com/model"
)

func main() {
	pancakeHouseMenu := model.NewPancakeHouseMenu()
	dinerMenu := model.NewDinerMenu()
	cafeMenu := model.NewCafeMenu()
	waitress := NewWaitress(pancakeHouseMenu, dinerMenu, cafeMenu)
	waitress.PrintMenu()

}

type Waitress struct {
	pancakeHouseMenu *model.PancakeHouseMenu
	dinerMenu        *model.DinerMenu
	cafeMenu         *model.CafeMenu
}

func NewWaitress(
	pancakeHouseMenu *model.PancakeHouseMenu,
	dinerMenu *model.DinerMenu,
	cafeMenu *model.CafeMenu,
) *Waitress {
	return &Waitress{
		pancakeHouseMenu: pancakeHouseMenu,
		dinerMenu:        dinerMenu,
		cafeMenu:         cafeMenu,
	}
}

func (w *Waitress) PrintMenu() {
	pancakeIterator := w.pancakeHouseMenu.CreateIterator()
	dinerIterator := w.dinerMenu.CreateIterator()
	cafeIterator := w.cafeMenu.CreateIterator()

	fmt.Printf("MENU\n---\nBREAKFAST\n")
	w.PrintMenuItems(pancakeIterator)
	fmt.Printf("\nLUNCH\n")
	w.PrintMenuItems(dinerIterator)
	fmt.Printf("\nDINNER\n")
	w.PrintMenuItems(cafeIterator)
}

func (w *Waitress) PrintMenuItems(iterator model.Iterator) {
	for iterator.HasNext() {
		item := iterator.Next().(*model.MenuItem)
		fmt.Printf("%s, %.2f -- %s\n", item.Name, item.Price, item.Description)
	}
}
