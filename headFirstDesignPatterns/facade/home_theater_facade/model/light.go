package model

import "fmt"

type TheaterLights struct {
}

func (l *TheaterLights) On() {
  fmt.Println("Theater Ceiling Lights on")
}

func (l *TheaterLights) Off() {}

func (l *TheaterLights) Dim(v int) {
	fmt.Printf("Theater Ceiling Lights dimming to %d%%\n", v)
}
