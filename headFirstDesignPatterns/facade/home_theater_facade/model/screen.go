package model

import "fmt"

type Screen struct{}

func (s *Screen) Down() {
  fmt.Println("Theater Screen going down")
}

func (s *Screen) Up() {
  fmt.Println("Theater Screen going up")
}
