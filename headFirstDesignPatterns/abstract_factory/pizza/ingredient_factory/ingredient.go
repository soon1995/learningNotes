package ingredientfactory

import (
	"example.com/model"
)

type PizzaIngredientFactory interface {
	CreateDough() model.Dough
	CreateSauce() model.Sauce
	CreateCheese() model.Cheese
	CreateVeggies() []model.Veggies
	CreatePepperoni() model.Pepperoni
	CreateClam() model.Clams
}
