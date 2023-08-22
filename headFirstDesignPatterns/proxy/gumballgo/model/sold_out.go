package model

import (
	"fmt"
)

type SoldOutState struct {
	gumballMachine *GumballMachine
}

func NewSoldOutState(gumballMachine *GumballMachine) *SoldOutState {
	return &SoldOutState{
		gumballMachine: gumballMachine,
	}
}

func (s *SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (s *SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

func (s *SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (s *SoldOutState) dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *SoldOutState) Refill() {
  s.gumballMachine.state = s.gumballMachine.noQuarterState
}
