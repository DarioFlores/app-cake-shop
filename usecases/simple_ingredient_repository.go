package usecases

import (
	"github.com/DarioFlores/app-cake-shop/domain"
	"log"
)

var simpleIngredientRepository SimpleIngredientRepository

func RegisterSimpleIngredientRepository(repository SimpleIngredientRepository) {
	simpleIngredientRepository = repository
}

func NewSimpleIngredientRepository() SimpleIngredientRepository {
	if simpleIngredientRepository == nil {
		log.Fatal("simpleIngredientRepository not found")
	}
	return simpleIngredientRepository
}

type SimpleIngredientRepository interface {
	Save(simpleIngredient *domain.SimpleIngredient) error
	FindAll() ([]domain.SimpleIngredient, error)
	FindByID(id string) (*domain.SimpleIngredient, error)
	Update(id string, simpleIngredient *domain.SimpleIngredient) error
}
