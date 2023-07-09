package model

import "fmt"

type GarageDoor struct {
	name string
}

func NewGarageDoor(name string) *GarageDoor {
	return &GarageDoor{
		name,
	}
}

func (m *GarageDoor) On() {
	fmt.Printf("GarageDoor %s is On\n", m.name)
}

func (m *GarageDoor) Off() {
	fmt.Printf("GarageDoor %s is Off\n", m.name)
}
