package duckbehavior

import "fmt"

type Squeak struct{}

func (m *Squeak) Quack() {
	fmt.Println("Squeak")
}
