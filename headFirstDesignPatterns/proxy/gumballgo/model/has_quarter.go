package model

import (
	"fmt"
	"math/rand"
	"time"
)

type HasQuarterState struct {
	random         *rand.Rand
	gumballMachine *GumballMachine
}

func NewHasQuaterState(gumballMachine *GumballMachine) *HasQuarterState {
	source := rand.NewSource(time.Now().UnixNano())
	return &HasQuarterState{
		gumballMachine: gumballMachine,
		random:         rand.New(source),
	}
}

func (s *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (s *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	s.gumballMachine.state = s.gumballMachine.noQuarterState
}

func (s *HasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	winner := s.random.Intn(10)
	if winner == 0 && s.gumballMachine.count > 1 {
		s.gumballMachine.state = s.gumballMachine.winnerState
	} else {
		s.gumballMachine.state = s.gumballMachine.soldState
	}
}

func (s *HasQuarterState) dispense() {
	fmt.Println("No gumball dispensed")
}

func (s *HasQuarterState) Refill() {}
