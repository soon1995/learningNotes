package duckbehavior

import "fmt"

type FlyNoWay struct{}

func (b *FlyNoWay) Fly() {
	fmt.Println("I cant' fly")
}
