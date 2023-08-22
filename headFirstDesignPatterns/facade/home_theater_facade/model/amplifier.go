package model

import "fmt"

type Amplifier struct{
  player StreamingPlayer
  vol int
}

func (a *Amplifier) On() {
  fmt.Println("Amplifier on")
}

func (a *Amplifier) Off() {
  fmt.Println("Amplifier off")
}

func (a *Amplifier) SetStreamingPlayer(player StreamingPlayer) {
  fmt.Println("Amplifier setting Streaming player to Streaming Player")
  a.player = player
}
func (a *Amplifier) SetSurroundSound() {
  fmt.Println("Amplifier surround sound on (5 speakers, 1 subwoofer)")
}
func (a *Amplifier) SetVolume(vol int) {
  fmt.Printf("Amplifier setting volume to %d\n", vol)
  a.vol = vol
}
