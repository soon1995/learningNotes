package model

import (
	"fmt"
)

type SoldState struct {
	gumballMachine *GumballMachine
}

func NewSoldState(gumballMachine *GumballMachine) *SoldState {
	return &SoldState{
		gumballMachine: gumballMachine,
	}
}

func (s *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s *SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (s *SoldState) dispense() {
  s.gumballMachine.ReleaseBall()
  if s.gumballMachine.count > 0 {
    s.gumballMachine.state = s.gumballMachine.noQuarterState
  } else {
    fmt.Println("Oops, out of gumballs!")
    s.gumballMachine.state = s.gumballMachine.soldOutState
  }
}

func (s *SoldState) Refill() {}
