package duck

import (
	"fmt"

	duckbehavior "example.com/behavior"
)

type ModelDuck struct {
	*Duck
}

func NewModelDuck() *ModelDuck {
	flyWithWing := &duckbehavior.FlyNoWay{}
	quack := &duckbehavior.Quack{}
	duck := &Duck{flyWithWing, quack}

	return &ModelDuck{duck}
}

func (d *ModelDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}
