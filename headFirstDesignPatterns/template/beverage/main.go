package main

import (
	"fmt"

	"example.com/model/impl"
)

func main() {
  myTea := impl.NewTea()
  myTea.PrepareRecipe()

  fmt.Println("")

  myCoffee := impl.NewCoffee()
  myCoffee.PrepareRecipe()
}
