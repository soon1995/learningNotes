package model

type IMenuComponent interface {
	Add(menu IMenuComponent)
	Remove(menu IMenuComponent)
	GetChild(int) IMenuComponent
	GetDescription() string
	GetPrice() float64
	IsVegetarian() bool
	Print()
}

type MenuComponent struct {
	Addfn      func(IMenuComponent)
	Removefn   func(IMenuComponent)
	GetChildFn func(int) IMenuComponent
	GetDescFn  func() string
	GetPriceFn func() float64
	IsVegeFn   func() bool
	PrintFn    func()
}

func (c *MenuComponent) Add(menu IMenuComponent) {
	if c.Addfn != nil {
		c.Addfn(menu)
		return
	}
	panic("Unsupported Operation")
}

func (c *MenuComponent) Remove(menu IMenuComponent) {
	if c.Removefn != nil {
		c.Removefn(menu)
		return
	}
	panic("Unsupported Operation")
}

func (c *MenuComponent) GetChild(i int) IMenuComponent {
	if c.GetChildFn != nil {
		return c.GetChildFn(i)
	}
	panic("Unsupported Operation")
}

func (c *MenuComponent) GetDescription() string {
	if c.GetDescFn != nil {
		return c.GetDescFn()
	}
	panic("Unsupported Operation")
}

func (c *MenuComponent) GetPrice() float64 {
	if c.GetPriceFn != nil {
		return c.GetPriceFn()
	}
	panic("Unsupported Operation")
}

func (c *MenuComponent) IsVegetarian() bool {
	if c.IsVegeFn != nil {
		return c.IsVegeFn()
	}
	panic("Unsupported Operation")
}

func (c *MenuComponent) Print() {
	if c.PrintFn != nil {
		c.PrintFn()
		return
	}
	panic("Unsupported Operation")
}
