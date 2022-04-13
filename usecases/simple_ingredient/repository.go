package simple_ingredient

import "github.com/DarioFlores/app-cake-shop/domain"

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
	Save(simpleIngredient *domain.SimpleIngredient) error
	GetAll() ([]domain.SimpleIngredient, error)
}
