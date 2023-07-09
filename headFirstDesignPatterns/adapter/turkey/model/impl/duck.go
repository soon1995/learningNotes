package impl

import "fmt"

type MallardDuck struct {}

func (*MallardDuck) Quack() {
  fmt.Println("Quack")
}

func (*MallardDuck) Fly() {
  fmt.Println("I'm flying")
}
