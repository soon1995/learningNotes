package ingredientfactory

import "example.com/model"

type NYPizzaIngredientFactory struct{}

func (i *NYPizzaIngredientFactory) CreateDough() model.Dough {
	return model.NewThinCrustDough()
}

func (i *NYPizzaIngredientFactory) CreateSauce() model.Sauce {
	return model.NewMarinaraSauce()
}

func (i *NYPizzaIngredientFactory) CreateCheese() model.Cheese {
	return model.NewReggianoCheese()

}

func (i *NYPizzaIngredientFactory) CreateVeggies() []model.Veggies {
	var veggies []model.Veggies
	veggies = append(veggies, model.NewGarlic(), model.NewOnion(), model.NewMushroom(), model.NewRedPepper())
	return veggies
}

func (i *NYPizzaIngredientFactory) CreatePepperoni() model.Pepperoni {
	return model.NewSlicedPepperoni()

}

func (i *NYPizzaIngredientFactory) CreateClam() model.Clams {
	return model.NewFreshClams()
}
