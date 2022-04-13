package simple_ingredient

import (
	"fmt"
	"github.com/DarioFlores/app-cake-shop/domain"
	"time"
)

type SimpleIngredientInterface interface {
	Save(simpleIngredientParams *SaveSimpleIngredientRequest) error
}

type simpleIngredientInteractor struct{}

func NewSimpleIngredientUseCase() SimpleIngredientInterface {
	return &simpleIngredientInteractor{}
}

func (s simpleIngredientInteractor) Save(request *SaveSimpleIngredientRequest) error {
	userRepository := NewSimpleIngredientRepository()
	var simpleIngredient = new(domain.SimpleIngredient)
	simpleIngredient.Name = request.Name
	simpleIngredient.Price = request.Price
	simpleIngredient.Amount = request.Amount
	simpleIngredient.Unit = request.Unit
	simpleIngredient.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	simpleIngredient.FrequencyUse = 0
	now := time.Now()
	simpleIngredient.LastModified = &now
	if err := userRepository.Save(simpleIngredient); err != nil {
		return err
	}
	return nil
}
