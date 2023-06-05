package main

import (
	"fmt"
	"math"
)

type A struct{}

func main() {

	// a := A{}
	// fmt.Println(unsafe.Sizeof(a))
	// fmt.Printf("%T", unsafe.Sizeof(float64(0)))
	fmt.Println(math.Float64bits(1.6))
}
