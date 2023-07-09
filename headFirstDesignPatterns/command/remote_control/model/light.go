package model

import "fmt"

type Light struct {
	name string
}

func NewLight(name string) *Light {
	return &Light{
		name,
	}
}

func (m *Light) On() {
	fmt.Printf("Light %s is On\n", m.name)
}

func (m *Light) Off() {
	fmt.Printf("Light %s is Off\n", m.name)
}
