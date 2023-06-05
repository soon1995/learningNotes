package anothepkg

import "fmt"
import "example.go/anotherpkg1"

func embed() {
	w := anotherpkg1.Wheel{}
  w.X = 1
	fmt.Printf("%#v\n", w)
}
