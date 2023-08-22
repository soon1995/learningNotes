package duck

import (
	"fmt"

	"example.com/behavior"
)

type Duck struct {
	flyBehavior   duckbehavior.FlyBehavior
	quackBehavior duckbehavior.QuackBehavior
}

func (d *Duck) SetFlyBehavior(b duckbehavior.FlyBehavior) {
	d.flyBehavior = b
}

func (d *Duck) SetQuackBehavior(b duckbehavior.QuackBehavior) {
	d.quackBehavior = b
}

func (d *Duck) PerformFly() {
	d.flyBehavior.Fly()
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d *Duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *Duck) Display() {}
