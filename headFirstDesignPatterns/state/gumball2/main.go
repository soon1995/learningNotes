package main

import (
	"fmt"

	"example.com/model"
)

func main() {
  gumballMachine := model.NewGumballMachine(5)

  fmt.Println(gumballMachine)

  gumballMachine.InsertQuarter()
  gumballMachine.TurnCrank()

  fmt.Println(gumballMachine)

  gumballMachine.InsertQuarter()
  gumballMachine.EjectQuarter()
  gumballMachine.TurnCrank()

  fmt.Println(gumballMachine)

  gumballMachine.InsertQuarter()
  gumballMachine.TurnCrank()
  gumballMachine.InsertQuarter()
  gumballMachine.TurnCrank()
  gumballMachine.EjectQuarter()

  fmt.Println(gumballMachine)

  gumballMachine.InsertQuarter()
  gumballMachine.InsertQuarter()
  gumballMachine.TurnCrank()
  gumballMachine.InsertQuarter()
  gumballMachine.TurnCrank()
  gumballMachine.InsertQuarter()
  gumballMachine.TurnCrank()

  fmt.Println(gumballMachine)

  gumballMachine.Refill(10)
  fmt.Println(gumballMachine)
}
