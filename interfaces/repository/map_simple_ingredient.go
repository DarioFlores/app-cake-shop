package repository

import (
	"fmt"
	"github.com/DarioFlores/app-cake-shop/domain"
	"github.com/DarioFlores/app-cake-shop/usecases"
)

func NewSimpleIngredientMapRepository() usecases.SimpleIngredientRepository {
	//this variable simulate the real kvs client
	var simpleIngredientsMapClient map[string]domain.SimpleIngredient
	return &simpleIngredientMapRepository{
		simpleIngredients: simpleIngredientsMapClient,
	}
}

type simpleIngredientMapRepository struct {
	simpleIngredients map[string]domain.SimpleIngredient
}

func (repo *simpleIngredientMapRepository) FindAll() ([]domain.SimpleIngredient, error) {
	var list []domain.SimpleIngredient
	for _, ingredient := range repo.simpleIngredients {
		list = append(list, ingredient)
	}
	return list, nil
}

func (repo *simpleIngredientMapRepository) FindByID(id string) (*domain.SimpleIngredient, error) {
	ingredient := repo.simpleIngredients[id]
	return &ingredient, nil
}

func (repo *simpleIngredientMapRepository) Update(id string, simpleIngredient *domain.SimpleIngredient) error {
	if id != simpleIngredient.ID {
		return fmt.Errorf("id different from the ingredient id")
	}
	_, ok := repo.simpleIngredients[id]
	if !ok {
		return fmt.Errorf("not found ingredient with ID: %s", id)

	}
	repo.simpleIngredients[id] = *simpleIngredient
	return nil
}

func (repo *simpleIngredientMapRepository) Save(simpleIngredient *domain.SimpleIngredient) error {
	repo.simpleIngredients[simpleIngredient.ID] = *simpleIngredient
	return nil
}
