package model

import "fmt"

type GumballMonitor struct {
	machine *GumballMachine
}

func NewGumballMonitor(machine *GumballMachine) *GumballMonitor {
	return &GumballMonitor{
		machine: machine,
	}
}

func (m *GumballMonitor) Report() {
	fmt.Printf("Gumball Machine: %s\n", m.machine.Location)
	fmt.Println(m.machine.String())
}
