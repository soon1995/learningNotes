package model

type PancakeHouseMenu struct {
	MenuItems []*MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menu := &PancakeHouseMenu{}
	menu.AddItem("K&B's Pancake Breakfast",
		"Pancakes with scrambled eggs and toast",
		true,
		2.99)
	menu.AddItem("Regular Pancake Breakfast",
		"Pancakes with fried eggs, sausage",
		false,
		2.99)
	menu.AddItem("Blueberry Pancakes",
		"Pancakes made with fresh blueberries",
		true,
		3.49)
	menu.AddItem("Waffles",
		"Waffles with your choice of blueberries or strawberries",
		true,
		3.59)

	return menu
}

func (m *PancakeHouseMenu) CreateIterator() Iterator {
	return &PancakeHouseMenuIterator{
		pos:    0,
		menu: m.MenuItems,
	}
}

func (m *PancakeHouseMenu) AddItem(name, description string, vegetarian bool, price float64) {
	m.MenuItems = append(m.MenuItems, NewMenuItem(name, description, vegetarian, price))
}

type PancakeHouseMenuIterator struct {
	pos    int
	menu []*MenuItem
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	return p.pos < len(p.menu)
}

func (p *PancakeHouseMenuIterator) Next() interface{} {
  menu := p.menu[p.pos]
	p.pos++
	return menu
}
