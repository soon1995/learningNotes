package impl

import "fmt"

type WildTurkey struct{}

func (*WildTurkey) Gobble() {
  fmt.Println("Gobble gobble")
}

func (*WildTurkey) Fly() {
  fmt.Println("I'm flying a short distance")
}
