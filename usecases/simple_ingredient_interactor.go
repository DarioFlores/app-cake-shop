package usecases

import (
	"fmt"
	"github.com/DarioFlores/app-cake-shop/domain"
	"time"
)

type SimpleIngredientInterface interface {
	Create(params *CreateSimpleIngredient) error
	UpdatePrice(id string, price float64) error
	Use(id string) error
	StopUse(id string) error
	FindByID(id string) (*domain.SimpleIngredient, error)
}

type simpleIngredientInteractor struct{}

func (s simpleIngredientInteractor) UpdatePrice(id string, price float64) error {
	repository := NewSimpleIngredientRepository()
	ingredient, err := s.FindByID(id)
	if err != nil {
		return fmt.Errorf("error finding an ingredient by id: %s - ERROR: %s", id, err.Error())
	}
	if ingredient == nil {
		return fmt.Errorf("not found ingrediente with id: %s", id)
	}
	ingredient.Price = price
	now := time.Now()
	ingredient.LastModifiedPrice = &now
	err = repository.Update(id, ingredient)
	if err != nil {
		return fmt.Errorf("error updating an ingredient with id: %s - ERROR: %s", id, err.Error())
	}
	return nil
}

func (s simpleIngredientInteractor) Use(id string) error {
	ingredient, err := s.FindByID(id)
	if err != nil {
		return fmt.Errorf("error finding an ingredient by id: %s - ERROR: %s", id, err.Error())
	}
	if ingredient == nil {
		return fmt.Errorf("not found ingrediente with id: %s", id)
	}
	ingredient.FrequencyUse++
	err = s.Update(ingredient)
	if err != nil {
		return fmt.Errorf("error updating an ingredient with id: %s - ERROR: %s", id, err.Error())
	}
	return nil
}

func (s simpleIngredientInteractor) StopUse(id string) error {
	ingredient, err := s.FindByID(id)
	if err != nil {
		return fmt.Errorf("error finding an ingredient by id: %s - ERROR: %s", id, err.Error())
	}
	if ingredient == nil {
		return fmt.Errorf("not found ingrediente with id: %s", id)
	}
	ingredient.FrequencyUse--
	err = s.Update(ingredient)
	if err != nil {
		return fmt.Errorf("error updating an ingredient with id: %s - ERROR: %s", id, err.Error())
	}
	return nil
}

func (s simpleIngredientInteractor) FindByID(id string) (*domain.SimpleIngredient, error) {
	//TODO implement me
	panic("implement me")
}

func (s simpleIngredientInteractor) Create(params *CreateSimpleIngredient) error {
	repository := NewSimpleIngredientRepository()
	simpleIngredient := params.toModel()
	if err := repository.Save(&simpleIngredient); err != nil {
		return err
	}
	return nil
}

func (s simpleIngredientInteractor) Update(params *domain.SimpleIngredient) error {
	repository := NewSimpleIngredientRepository()
	if err := repository.Update(params.ID, params); err != nil {
		return err
	}
	return nil
}

func NewSimpleIngredientUseCase() SimpleIngredientInterface {
	return &simpleIngredientInteractor{}
}

type CreateSimpleIngredient struct {
	Name   string
	Price  float64
	Amount float64
	Unit   string
}

func (c *CreateSimpleIngredient) toModel() domain.SimpleIngredient {
	now := time.Now()
	simpleIngredient := domain.SimpleIngredient{
		ID:                fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:              c.Name,
		Price:             c.Price,
		Amount:            c.Amount,
		Unit:              c.Unit,
		FrequencyUse:      0,
		LastModifiedPrice: &now,
	}
	return simpleIngredient
}
