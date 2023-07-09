package model

import "fmt"

type HomeTheaterFacade struct {
	amp       *Amplifier
	tuner     *Tuner
	player    *StreamingPlayer
	projector *Projector
	lights    *TheaterLights
	screen    *Screen
	popper    *PopcornPopper
}

func NewHomeTheaterFacade(
	amp *Amplifier,
	tuner *Tuner,
	player *StreamingPlayer,
	projector *Projector,
	lights *TheaterLights,
	screen *Screen,
	popper *PopcornPopper,
) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp,
		tuner,
		player,
		projector,
		lights,
		screen,
		popper,
	}
}

func (f *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	f.popper.On()
	f.popper.Pop()
	f.lights.Dim(10)
	f.screen.Down()
	f.projector.On()
	f.projector.WideScreenMode()
	f.amp.On()
	f.amp.SetStreamingPlayer(*f.player)
	f.amp.SetSurroundSound()
	f.amp.SetVolume(5)
	f.player.On()
	f.player.Play(movie)
}

func (f *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	f.popper.Off()
	f.lights.On()
	f.screen.Up()
	f.projector.Off()
	f.amp.Off()
	f.player.Stop()
	f.player.Off()
}
