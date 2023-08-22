package model

import (
	"fmt"
)

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuaterState(gumballMachine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{
		gumballMachine: gumballMachine,
	}
}

func (s *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	s.gumballMachine.state = s.gumballMachine.hasQuarterState
}

func (s *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (s *NoQuarterState) TurnCrank() {
	fmt.Println("You turned but there's no quarter")
}

func (s *NoQuarterState) dispense() {
	fmt.Println("You need to pay first")
}

func (s *NoQuarterState) Refill() {}
