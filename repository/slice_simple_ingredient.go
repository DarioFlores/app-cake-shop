package repository

import (
	"github.com/DarioFlores/app-cake-shop/entities"
	"github.com/DarioFlores/app-cake-shop/gateway"
)

func NewSimpleIngredientRepository() gateway.SimpleIngredientRepository {
	//this variable simulate the real kvs client
	var simpleIngredientsSliceClient []entities.SimpleIngredient
	return &simpleIngredientRepository{
		simpleIngredients: simpleIngredientsSliceClient,
	}
}

type simpleIngredientRepository struct {
	simpleIngredients []entities.SimpleIngredient
}

func (repository *simpleIngredientRepository) Save(simpleIngredient *entities.SimpleIngredient) error {
	repository.simpleIngredients = append(repository.simpleIngredients, *simpleIngredient)
	return nil
}

func (repository *simpleIngredientRepository) GetAll() ([]entities.SimpleIngredient, error) {
	return repository.simpleIngredients, nil
}
