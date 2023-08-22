package model

import (
	"fmt"
)

type WinnerState struct {
	gumballMachine *GumballMachine
}

func NewWinnerState(gumballMachine *GumballMachine) *WinnerState {
	return &WinnerState{
		gumballMachine: gumballMachine,
	}
}

func (s *WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *WinnerState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s *WinnerState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (s *WinnerState) dispense() {
	s.gumballMachine.ReleaseBall()
	if s.gumballMachine.count == 0 {
		s.gumballMachine.state = s.gumballMachine.soldOutState
	} else {
		s.gumballMachine.ReleaseBall()
		fmt.Println("YOU'RE A WINNER! You got two gumballs for your quarter")
		if s.gumballMachine.count > 0 {
			s.gumballMachine.state = s.gumballMachine.noQuarterState
		} else {
			fmt.Println("Oops, out of gumballs!")
			s.gumballMachine.state = s.gumballMachine.soldOutState
		}
	}
}

func (s *WinnerState) Refill() {}
