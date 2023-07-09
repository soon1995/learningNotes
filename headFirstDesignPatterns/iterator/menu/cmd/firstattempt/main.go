package main

import (
	"fmt"

	"example.com/model"
)

func main() {
	pancakeHouseMenu := model.NewPancakeHouseMenu()
	dinerMenu := model.NewDinerMenu()
	for _, v := range pancakeHouseMenu.MenuItems {
		fmt.Println(v.Name)
		fmt.Println(v.Price)
		fmt.Println(v.Description)
	}
  fmt.Println("")
	for _, v := range dinerMenu.MenuItems {
		fmt.Println(v.Name)
		fmt.Println(v.Price)
		fmt.Println(v.Description)
    // here panic because of empty menuitems
	}
}
