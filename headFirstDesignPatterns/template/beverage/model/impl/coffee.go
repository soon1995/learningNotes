package impl

import (
	"fmt"
	"strings"

	"example.com/model"
	"log"
)

type coffee struct {
	*model.CaffeineBeverage
}

func NewCoffee() *coffee {
	b := &model.CaffeineBeverage{}
	b.AddCondimentsFn = func() {
		fmt.Println("Adding Sugar and Milk")
	}
	b.BrewFn = func() {
		fmt.Println("Dripping Coffee through filter")
	}
	t := &coffee{b}
	t.WantCondimentsFn = func() bool {
		answer := t.getUserInput()
		if strings.ToLower(answer) == "y" {
			return true
		} else {
			return false
		}
	}
	return t
}

func (c *coffee) getUserInput() string {
	var answer string
	fmt.Printf("Would you like milk and sugar with your coffee (y/n)?")
	_, err := fmt.Scan(&answer)
	if err != nil {
		log.Printf("cannot scan user input\n")
		answer = "n"
	}
	return answer
}
