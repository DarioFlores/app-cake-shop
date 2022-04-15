package usecases

import (
	"github.com/DarioFlores/app-cake-shop/domain"
)

var operationErrors = map[string]map[string]error{"findByID": {}, "findAll": {}, "save": {}, "update": {}}
var responseFindByID = map[string]domain.SimpleIngredient{}
var responseFindAll []domain.SimpleIngredient

type mockSimpleIngredientRepository struct{}

func NewMockSimpleIngredientRepository() SimpleIngredientRepository {
	operationErrors = map[string]map[string]error{"findByID": {}, "findAll": {}, "save": {}, "update": {}}
	responseFindByID = map[string]domain.SimpleIngredient{}
	responseFindAll = []domain.SimpleIngredient{}
	return new(mockSimpleIngredientRepository)
}

func (m mockSimpleIngredientRepository) Save(params *domain.SimpleIngredient) error {
	if err := operationErrors["save"][params.Name]; err != nil {
		return err
	}
	return nil
}

func (m mockSimpleIngredientRepository) FindAll() ([]domain.SimpleIngredient, error) {
	if err := operationErrors["findAll"]["findAll"]; err != nil {
		return nil, err
	}
	return responseFindAll, nil
}

func (m mockSimpleIngredientRepository) FindByID(id string) (*domain.SimpleIngredient, error) {
	if err := operationErrors["findByID"][id]; err != nil {
		return nil, err
	}
	res, ok := responseFindByID[id]
	if !ok {
		return nil, nil
	}
	return &res, nil
}

func (m mockSimpleIngredientRepository) Update(id string, _ *domain.SimpleIngredient) error {
	if err := operationErrors["update"][id]; err != nil {
		return err
	}
	return nil
}

func SetErrorSimpleIngredientRepository(operation, key string, err error) {
	if operationErrors[operation] == nil {
		panic("operationErrors should be initialized")
	}
	if operation == "findAll" {
		key = "findAll"
	}
	operationErrors[operation][key] = err
}

func SetResponseSimpleIngredientRepositoryFindByID(key string, res domain.SimpleIngredient) {
	responseFindByID[key] = res
}

func SetResponseSimpleIngredientRepositoryFindAll(res []domain.SimpleIngredient) {
	responseFindAll = res
}
