package duck

import (
	"fmt"

	duckbehavior "example.com/behavior"
)

type MallardDuck struct {
	*Duck
}

func NewMallardDuck() *MallardDuck {
	flyWithWing := &duckbehavior.FlyWithWings{}
	quack := &duckbehavior.Quack{}
	duck := &Duck{flyWithWing, quack}

	return &MallardDuck{duck}
}

func (d *MallardDuck) Display() {
	fmt.Println("I'm a model duck")
}
