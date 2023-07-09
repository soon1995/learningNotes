package main

import (
	"fmt"

	"example.com/model"
	"example.com/model/impl"
)

func main() {
  var duck model.Duck
  duck = &impl.MallardDuck{}

  var turkey model.Turkey
  turkey = &impl.WildTurkey{}
  turkeyAdapter := model.NewTurkeyAdapter(turkey)

  fmt.Println("The Turkey says...")
  turkey.Gobble()
  turkey.Fly()

  fmt.Println("\nThe Duck says...")
  testDuck(duck)

  fmt.Println("\nThe TurkeyAdapter says...")
  testDuck(turkeyAdapter)
  
}

func testDuck(duck model.Duck) {
  duck.Quack()
  duck.Fly()
}
