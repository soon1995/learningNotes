package main

import (
	duckbehavior "example.com/behavior"
	"example.com/duck"
)

func main() {
	// duck1 := duck.NewMallardDuck()
	// duck1.PerformFly()
	// duck1.PerformQuack()

	model := duck.NewModelDuck()
	model.PerformFly()
	model.SetFlyBehavior(&duckbehavior.FlyRocketPowered{})
	model.PerformFly()
}
