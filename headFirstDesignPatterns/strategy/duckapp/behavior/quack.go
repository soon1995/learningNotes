package duckbehavior

import "fmt"

type Quack struct{}

func (m *Quack) Quack() {
	fmt.Println("Quack")
}
