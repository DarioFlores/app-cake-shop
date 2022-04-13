package gateway

import "github.com/DarioFlores/app-cake-shop/entities"

var simpleIngredientRepository SimpleIngredientRepository

func RegisterSimpleIngredientRepository(repository SimpleIngredientRepository) {
	simpleIngredientRepository = repository
}

func NewSimpleIngredientRepository() SimpleIngredientRepository {
	if simpleIngredientRepository == nil {
		panic("simpleIngredientRepository not found")
	}
	return simpleIngredientRepository
}

type SimpleIngredientRepository interface {
	Save(simpleIngredient *entities.SimpleIngredient) error
	GetAll() ([]entities.SimpleIngredient, error)
}
