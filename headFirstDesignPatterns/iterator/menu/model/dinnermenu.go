package model

import (
	"fmt"
	"os"
)

type DinerMenu struct {
	numberOfItems int
	max           int
	MenuItems     [6]*MenuItem
}

func NewDinerMenu() *DinerMenu {
	menu := &DinerMenu{max: 6}
	menu.AddItem("Vegetarian BLT",
		"(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	menu.AddItem("BLT",
		"Bacon with lettuce & tomato on whole wheat", false, 2.99)
	menu.AddItem("Soup of the day",
		"Soup of the day, with a side of potato salad", false, 3.29)
	menu.AddItem("Hotdog",
		"A hot dog, with sauerkraut, relish, onions, topped with cheese",
		false, 3.05)
	return menu
}

func (m *DinerMenu) AddItem(name, description string, vegetarian bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarian, price)
	if m.numberOfItems >= m.max {
		fmt.Fprintln(os.Stderr, "Sorry, menu is full! Can't add item to menu")
	} else {
		m.MenuItems[m.numberOfItems] = menuItem
		m.numberOfItems++
	}
}

func (m *DinerMenu) CreateIterator() Iterator {
	return &DinerMenuIterator{
		pos:    0,
		menu: m.MenuItems,
	}
}

type DinerMenuIterator struct {
	pos    int
	menu [6]*MenuItem
}

func (p *DinerMenuIterator) HasNext() bool {
	return p.pos < len(p.menu) && p.menu[p.pos] != nil
}

func (p *DinerMenuIterator) Next() interface{} {
  menu := p.menu[p.pos]
	p.pos++
	return menu
}
