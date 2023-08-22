package model

import "fmt"

type ICoffeineBeverate interface {
	PrepareRecipe()
	Brew()
	AddCondiments()
	BoilWater()
	PourInCup()
	CustomerWantsCondiments() bool
}

type CaffeineBeverage struct {
	BrewFn           func()
	AddCondimentsFn  func()
	WantCondimentsFn func() bool
}

func (c *CaffeineBeverage) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	if c.CustomerWantsCondiments() {
		c.AddCondiments()
	}
}

func (c *CaffeineBeverage) Brew() {
	c.BrewFn()
}

func (c *CaffeineBeverage) AddCondiments() {
	c.AddCondimentsFn()
}

func (c *CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling Water")
}

func (c *CaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (c *CaffeineBeverage) CustomerWantsCondiments() bool {
	if c.WantCondimentsFn != nil {
		return c.WantCondimentsFn()
	}
	return true
}
