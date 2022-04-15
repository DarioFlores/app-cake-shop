package usecases

import (
	"github.com/DarioFlores/app-cake-shop/domain"
	"log"
)

var instance SimpleIngredientRepository

func RegisterSimpleIngredientRepository(repository SimpleIngredientRepository) {
	instance = repository
}

func GetInstanceSimpleIngredientRepository() SimpleIngredientRepository {
	if instance == nil {
		log.Fatal("instance simpleIngredientRepository not found")
	}
	return instance
}

type SimpleIngredientRepository interface {
	Save(simpleIngredient *domain.SimpleIngredient) error
	FindAll() ([]domain.SimpleIngredient, error)
	FindByID(id string) (*domain.SimpleIngredient, error)
	Update(id string, simpleIngredient *domain.SimpleIngredient) error
}
