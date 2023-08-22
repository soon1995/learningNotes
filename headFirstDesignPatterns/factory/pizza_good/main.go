package main

import (
	"fmt"

	"example.com/store"
)

func main() {
	nyPizzaStore := store.NewNYStylePizzaStore()
	chicagoPizzaStore := store.NewChicagoStylePizzaStore()

	pizza := nyPizzaStore.OrderPizza("cheese")
	fmt.Printf("Ethan ordered a %s \n\n", pizza.GetName())

	pizza = chicagoPizzaStore.OrderPizza("cheese")
	fmt.Printf("Joel ordered a %s \n\n", pizza.GetName())
}
