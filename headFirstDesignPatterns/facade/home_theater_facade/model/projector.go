package model

import "fmt"

type Projector struct {
}

func (p *Projector) On() {
	fmt.Println("Projector on")
}

func (p *Projector) Off() {
	fmt.Println("Projector off")
}

func (p *Projector) WideScreenMode() {
	fmt.Println("Projector in widescreen mode (16x9 aspect ratio)")
}
