package model

type CafeMenu struct {
	menuItems map[string]*MenuItem
}

func NewCafeMenu() *CafeMenu {
	menu := &CafeMenu{
		menuItems: make(map[string]*MenuItem),
	}
	menu.AddItem("Veggie Burger and Air Fries",
		"Veggie burger on a whole wheat bun, lettuce, tomato, and fries",
		true, 3.99)
	menu.AddItem("Soup of the day",
		"A cup of the soup of the day, with a side salad",
		false, 3.69)
	menu.AddItem("Burrito",
		"A large burrito, with whole pinto beans, salsa, guacamole",
		true, 4.29)

	return menu
}

func (m *CafeMenu) AddItem(name, description string, vegetarian bool, price float64) {
	menuItem := NewMenuItem(name, description, vegetarian, price)
	m.menuItems[name] = menuItem

}

func (m *CafeMenu) CreateIterator() Iterator {
  menu := make([]*MenuItem,0, len(m.menuItems))
	for _, v := range m.menuItems {
		menu = append(menu, v)
	}
	return &CafeMenuIterator{
		pos:  0,
		menu: menu,
	}
}

type CafeMenuIterator struct {
	pos  int
	menu []*MenuItem
}

func (p *CafeMenuIterator) HasNext() bool {
	return p.pos < len(p.menu)
}

func (p *CafeMenuIterator) Next() interface{} {
	menu := p.menu[p.pos]
	p.pos++
	return menu
}
