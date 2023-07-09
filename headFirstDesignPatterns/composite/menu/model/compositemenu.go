package model

import "fmt"

type Menu struct {
	menuComponents []IMenuComponent
	Name           string
	description    string
	*MenuComponent
}

func NewMenu(name, description string) *Menu {
	menuComponent := &MenuComponent{}
	menu := &Menu{
		Name:        name,
		description: description,
	}
	menuComponent.Addfn = menu.Add
	menuComponent.Removefn = menu.Remove
	menuComponent.GetChildFn = menu.GetChild
	menuComponent.PrintFn = menu.Print
	menu.MenuComponent = menuComponent
	return menu

}

func (m *Menu) Add(menuComponent IMenuComponent) {
	m.menuComponents = append(m.menuComponents, menuComponent)
}

func (m *Menu) Remove(menuComponent IMenuComponent) {
	idx := 0
	for i, mc := range m.menuComponents {
		if mc == menuComponent {
			idx = i
			break
		}
	}
	copy(m.menuComponents[idx:], m.menuComponents[idx+1:])
	m.menuComponents = m.menuComponents[:len(m.menuComponents)-1]
}

func (m *Menu) GetChild(i int) IMenuComponent {
	return m.menuComponents[i]
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) Print() {
	fmt.Printf("\n%s, %s\n", m.Name, m.GetDescription())
	fmt.Println("-------------")

	for _, mc := range m.menuComponents {
		mc.Print()
	}
}
