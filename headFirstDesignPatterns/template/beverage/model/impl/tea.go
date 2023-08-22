package impl

import (
	"fmt"

	"example.com/model"
)

type tea struct {
  *model.CaffeineBeverage
}

func NewTea() *tea {
  b := &model.CaffeineBeverage{}
  b.AddCondimentsFn = func() {
    fmt.Println("Adding Lemon")
  }
  b.BrewFn = func() {
    fmt.Println("Steeping the tea")
  }
  t := &tea{b}
  return t
}
