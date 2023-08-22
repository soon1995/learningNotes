package duckbehavior

import "fmt"

type FlyWithWings struct{}

func (b *FlyWithWings) Fly() {
	fmt.Println("I'm flying!!")
}
