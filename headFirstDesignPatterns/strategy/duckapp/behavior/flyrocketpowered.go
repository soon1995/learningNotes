package duckbehavior

import "fmt"

type FlyRocketPowered struct{}

func (b *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
