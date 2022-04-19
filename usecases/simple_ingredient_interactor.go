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

// Create a simple ingredient
func (s simpleIngredientInteractor) Create(params *CreateSimpleIngredient) error {
	repository := GetInstanceSimpleIngredientRepository()
	simpleIngredient := params.toModel()
	if err := repository.Save(&simpleIngredient); err != nil {
		return err
	}
	return nil
}

// UpdatePrice of simple ingredient
func (s simpleIngredientInteractor) UpdatePrice(id string, price float64) error {
	ingredient, err := s.FindByID(id)
	if err != nil {
		return err
	}
	if ingredient == nil {
		return fmt.Errorf("not found ingrediente with id: %s", id)
	}
	ingredient.Price = price
	ingredient.LastModifiedPrice = time.Now()
	err = s.Update(ingredient)
	if err != nil {
		return err
	}
	return nil
}

// Use a simple ingredient
func (s simpleIngredientInteractor) Use(id string) error {
	var ingredient = new(domain.SimpleIngredient)
	var err error
	if ingredient, err = s.FindByID(id); err != nil {
		return err
	}
	if ingredient == nil {
		return fmt.Errorf("not found ingrediente with id: %s", id)
	}
	ingredient.FrequencyUse++
	if err = s.Update(ingredient); err != nil {
		return err
	}
	return nil
}

// StopUse a simple ingredient
func (s simpleIngredientInteractor) StopUse(id string) error {
	ingredient, err := s.FindByID(id)
	if err != nil {
		return err
	}
	if ingredient == nil {
		return fmt.Errorf("not found ingrediente with id: %s", id)
	}
	ingredient.FrequencyUse--
	err = s.Update(ingredient)
	if err != nil {
		return err
	}
	return nil
}

// FindByID get a simple ingredient by ID
func (s simpleIngredientInteractor) FindByID(id string) (*domain.SimpleIngredient, error) {
	repository := GetInstanceSimpleIngredientRepository()
	ingredient, err := repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error finding an ingredient by id: %s - ERROR: %s", id, err.Error())
	}
	return ingredient, nil
}

// Update a simple ingredient
func (s simpleIngredientInteractor) Update(params *domain.SimpleIngredient) error {
	repository := GetInstanceSimpleIngredientRepository()
	if err := repository.Update(params.ID, params); err != nil {
		return fmt.Errorf("error updating an ingredient with id: %s - ERROR: %s", params.ID, err.Error())
	}
	return nil
}

// NewSimpleIngredientUseCase generates a new instance of the use case
func NewSimpleIngredientUseCase() SimpleIngredientInterface {
	return &simpleIngredientInteractor{}
}

// CreateSimpleIngredient struct for creation of a simple ingredient
type CreateSimpleIngredient struct {
	Name   string
	Price  float64
	Amount float64
	Unit   string
}

// toModel returns the model of a simple ingredient
func (c *CreateSimpleIngredient) toModel() domain.SimpleIngredient {
	simpleIngredient := domain.SimpleIngredient{
		ID:                fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:              c.Name,
		Price:             c.Price,
		Amount:            c.Amount,
		Unit:              c.Unit,
		FrequencyUse:      0,
		LastModifiedPrice: time.Now(),
	}
	return simpleIngredient
}
