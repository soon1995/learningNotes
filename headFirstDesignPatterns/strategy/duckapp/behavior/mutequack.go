package duckbehavior

import "fmt"

type MuteQuack struct{}

func (m *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}
