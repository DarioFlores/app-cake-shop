package repository

import (
	"github.com/DarioFlores/app-cake-shop/domain"
	"github.com/DarioFlores/app-cake-shop/usecases/simple_ingredient"
)

func NewSimpleIngredientRepository() simple_ingredient.SimpleIngredientRepository {
	//this variable simulate the real kvs client
	var simpleIngredientsSliceClient []domain.SimpleIngredient
	return &simpleIngredientRepository{
		simpleIngredients: simpleIngredientsSliceClient,
	}
}

type simpleIngredientRepository struct {
	simpleIngredients []domain.SimpleIngredient
}

func (repository *simpleIngredientRepository) Save(simpleIngredient *domain.SimpleIngredient) error {
	repository.simpleIngredients = append(repository.simpleIngredients, *simpleIngredient)
	return nil
}

func (repository *simpleIngredientRepository) GetAll() ([]domain.SimpleIngredient, error) {
	return repository.simpleIngredients, nil
}
