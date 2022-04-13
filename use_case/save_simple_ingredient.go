package use_case

import (
	"github.com/DarioFlores/app-cake-shop/entities"
	"github.com/DarioFlores/app-cake-shop/gateway"
)

type (
	SaveSimpleIngredientInterface interface {
		Save(simpleIngredientParams *entities.SimpleIngredient) error
	}

	saveSimpleIngredientInterface struct{}
)

func NewSaveSimpleIngredientUseCase() SaveSimpleIngredientInterface {
	return &saveSimpleIngredientInterface{}
}

func (s saveSimpleIngredientInterface) Save(simpleIngredientParams *entities.SimpleIngredient) error {
	userRepository := gateway.NewSimpleIngredientRepository()
	if err := userRepository.Save(simpleIngredientParams); err != nil {
		return err
	}
	return nil
}
