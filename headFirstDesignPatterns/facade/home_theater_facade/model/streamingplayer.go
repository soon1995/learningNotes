package model

import "fmt"

type StreamingPlayer struct {
  movie string
}

func (s *StreamingPlayer) On() {
	fmt.Println("Streaming Player on")
}

func (s *StreamingPlayer) Off()              {
	fmt.Println("Streaming Player off")
}

func (s *StreamingPlayer) Play(movie string) {
	fmt.Printf("Streaming Player playing \"%s\"\n", movie)
  s.movie = movie
}

func (s *StreamingPlayer) Stop()             {
	fmt.Printf("Streaming Player stopped \"%s\"\n", s.movie)
  s.movie = ""
}
