package repository

import (
	"fmt"
	"github.com/DarioFlores/app-cake-shop/domain"
	"github.com/DarioFlores/app-cake-shop/usecases"
)

func NewSimpleIngredientSliceRepository() usecases.SimpleIngredientRepository {
	//this variable simulate the real kvs client
	var simpleIngredientsSliceClient []domain.SimpleIngredient
	return &simpleIngredientSliceRepository{
		simpleIngredients: simpleIngredientsSliceClient,
	}
}

type simpleIngredientSliceRepository struct {
	simpleIngredients []domain.SimpleIngredient
}

func (repo *simpleIngredientSliceRepository) FindAll() ([]domain.SimpleIngredient, error) {
	return repo.simpleIngredients, nil
}

func (repo *simpleIngredientSliceRepository) FindByID(id string) (*domain.SimpleIngredient, error) {
	var res = new(domain.SimpleIngredient)
	for _, ingredient := range repo.simpleIngredients {
		if ingredient.ID == id {
			*res = ingredient
			return res, nil
		}
	}
	return nil, nil
}

func (repo *simpleIngredientSliceRepository) Update(id string, simpleIngredient *domain.SimpleIngredient) error {
	if id != simpleIngredient.ID {
		return fmt.Errorf("id different from the ingredient id")
	}
	var index = -1
	for i, ingredient := range repo.simpleIngredients {
		if ingredient.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("not found ingredient with ID: %s", id)
	}
	repo.simpleIngredients[index] = *simpleIngredient
	return nil
}

func (repo *simpleIngredientSliceRepository) Save(simpleIngredient *domain.SimpleIngredient) error {
	repo.simpleIngredients = append(repo.simpleIngredients, *simpleIngredient)
	return nil
}
