package model

import "fmt"

type PopcornPopper struct {
}

func (p *PopcornPopper) On() {
	fmt.Println("Popcorn Popper On")
}

func (p *PopcornPopper) Off() {
	fmt.Println("Popcorn Popper Off")
}

func (p *PopcornPopper) Pop() {
	fmt.Println("Popcorn Popper popping popcorn!")
}
