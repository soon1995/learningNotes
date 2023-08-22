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

type GumballMachine struct {
	state int
	count int
}

func NewGumballMachine(count int) *GumballMachine {
	return &GumballMachine{
		count: count,
		state: NoQuater,
	}
}

func (g *GumballMachine) InsertQuarter() {
	if g.state == HasQuater {
		fmt.Println("You can't insert another quarter")
	} else if g.state == NoQuater {
		fmt.Println("You inserted a quarter")
		g.state = HasQuater
	} else if g.state == SoldOut {
		fmt.Println("You can't insert a quarter, the machine is sold out")
	} else if g.state == Sold {
		fmt.Println("Please wait, we're already giving you a gumball")
	}
}

func (g *GumballMachine) EjectQuarter() {
	if g.state == HasQuater {
		fmt.Println("Quarter returned")
		g.state = NoQuater
	} else if g.state == NoQuater {
		fmt.Println("You haven't inserted a quarter")
	} else if g.state == SoldOut {
		fmt.Println("You can't eject, you haven't inserted a quarter yet")
	} else if g.state == Sold {
		fmt.Println("Sorry, you already turned the crank")
	}
}

func (g *GumballMachine) TurnCrank() {
	if g.state == HasQuater {
		fmt.Println("You turned...")
		g.state = Sold
		g.dispense()
	} else if g.state == NoQuater {
		fmt.Println("You turned but there's no quarter")
	} else if g.state == SoldOut {
		fmt.Println("You turned, but there are no gumballs")
	} else if g.state == Sold {
		fmt.Println("Turning twice doesn't get you another gumball!")
	}
}

func (g *GumballMachine) dispense() {
	if g.state == HasQuater {
		fmt.Println("You need to turn the crank")
	} else if g.state == NoQuater {
		fmt.Println("You need to pay first")
	} else if g.state == SoldOut {
		fmt.Println("No gumball dispensed")
	} else if g.state == Sold {
		fmt.Println("A gumball comes rolling out the slot")
		g.count--
		if g.count == 0 {
			fmt.Println("Oops, out of gumballs!")
			g.state = SoldOut
		} else {
			g.state = NoQuater
		}
	}
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
	fmt.Fprintf(buf, "State: %s\n", states[g.state])
	return buf.String()
}
