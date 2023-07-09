package ingredientfactory

import "example.com/model"

type ChicagoPizzaIngredientFactory struct{}

func (i *ChicagoPizzaIngredientFactory) CreateDough() model.Dough {
	return model.NewThickCrustDough()
}

func (i *ChicagoPizzaIngredientFactory) CreateSauce() model.Sauce {
	return model.NewPlumTomatoSauce()
}

func (i *ChicagoPizzaIngredientFactory) CreateCheese() model.Cheese {
	return model.NewMozzarellaCheese()

}

func (i *ChicagoPizzaIngredientFactory) CreateVeggies() []model.Veggies {
	var veggies []model.Veggies
	veggies = append(veggies, model.NewSpinach(), model.NewEggPlant(), model.NewMushroom(), model.NewBlackOlives())
	return veggies
}

func (i *ChicagoPizzaIngredientFactory) CreatePepperoni() model.Pepperoni {
	return model.NewSlicedPepperoni()

}

func (i *ChicagoPizzaIngredientFactory) CreateClam() model.Clams {
	return model.NewFrozenClams()
}
