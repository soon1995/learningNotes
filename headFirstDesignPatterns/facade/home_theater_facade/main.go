package main

import "example.com/model"

func main() {
	homeTheater := model.NewHomeTheaterFacade(
		&model.Amplifier{},
		&model.Tuner{},
		&model.StreamingPlayer{},
		&model.Projector{},
		&model.TheaterLights{},
		&model.Screen{},
		&model.PopcornPopper{},
	)

  homeTheater.WatchMovie("Raiders of the Lost Ark")
  homeTheater.EndMovie()
}
