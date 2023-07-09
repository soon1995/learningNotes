package main

import "example.com/model/store"

func main() {
  nystore := store.NewNYStylePizzaStore()
  nystore.OrderPizza("cheese")
  nystore.OrderPizza("clam")
}
