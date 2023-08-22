package model

import (
	"bytes"
	"fmt"
)

const (
	SoldOut = iota
	NoQuater
	HasQuater
	Sold
)

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	dispense()
	Refill()
}

type GumballMachine struct {
	soldOutState    State
	noQuarterState  State
	hasQuarterState State
	soldState       State
	winnerState     State

	state State
	count int
}

func NewGumballMachine(count int) *GumballMachine {
	machine := &GumballMachine{
		count: count,
	}
	machine.soldOutState = NewSoldOutState(machine)
	machine.soldState = NewSoldState(machine)
	machine.noQuarterState = NewNoQuaterState(machine)
	machine.hasQuarterState = NewHasQuaterState(machine)
	machine.winnerState = NewWinnerState(machine)
	if count > 0 {
		machine.state = machine.noQuarterState
	} else {
		machine.state = machine.soldOutState
	}
	return machine
}

func (g *GumballMachine) InsertQuarter() {
	g.state.InsertQuarter()
}

func (g *GumballMachine) EjectQuarter() {
	g.state.EjectQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.state.TurnCrank()
	g.state.dispense()
}

func (g *GumballMachine) ReleaseBall() {
	fmt.Println("A gumball comes rolling out the slot...")
	if g.count > 0 {
		g.count--
	}
}

func (g *GumballMachine) Refill(count int) {
	g.count += count
	fmt.Printf("The gumball machine was just refilled; its new count is : %d\n", g.count)
	g.state.Refill()
}

var states []string = []string{
	NoQuater:  "NO_QUATER",
	HasQuater: "HAS_QUARTER",
	Sold:      "SOLD",
	SoldOut:   "SOLD_OUT",
}

func (g *GumballMachine) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "\nInventory: %d gumballs\n", g.count)
	fmt.Fprintf(buf, "State: %T\n", g.state)
	return buf.String()
}
