package main

import (
	"fmt"

	"example.com/beverage"
)

func main() {
	bvr := beverage.NewExpresso()
	fmt.Printf("%s $ %.2f\n", bvr.GetDescription(), bvr.Cost())

	var bvr2 beverage.Beverage
	bvr2 = beverage.NewDarkRoast()
	bvr2 = beverage.NewMocha(bvr2)
	bvr2 = beverage.NewMocha(bvr2)
	bvr2 = beverage.NewWhip(bvr2)
	fmt.Printf("%s $ %.2f\n", bvr2.GetDescription(), bvr2.Cost())

	var bvr3 beverage.Beverage
	bvr3 = beverage.NewHouseBlend()
	bvr3 = beverage.NewSoy(bvr3)
	bvr3 = beverage.NewMocha(bvr3)
	bvr3 = beverage.NewWhip(bvr3)
	fmt.Printf("%s $ %.2f\n", bvr3.GetDescription(), bvr3.Cost())
}
