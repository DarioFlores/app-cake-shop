package usecases

import (
	"fmt"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// Dado cuando entonces
// Given When Should

func TestCreateOK(t *testing.T) {
	repository := NewMockSimpleIngredientRepository()
	RegisterSimpleIngredientRepository(repository)
	usecase := NewSimpleIngredientUseCase()

	t.Log("Given the need to test creating simple ingredient.")
	{
		t.Logf("\tWhen checking that the creation returns no error")
		{
			params := CreateSimpleIngredient{
				Name:   "Harina",
				Price:  210,
				Amount: 1,
				Unit:   "KG",
			}
			err := usecase.Create(&params)
			if err != nil {
				t.Errorf("\t\tShould receive an error equal to nil. %v %v", err, ballotX)
			}
			t.Logf("\t\tShould receive an error equal to nil. %v", checkMark)
		}
	}
}

func TestCreateFailed(t *testing.T) {
	repository := NewMockSimpleIngredientRepository()
	RegisterSimpleIngredientRepository(repository)
	usecase := NewSimpleIngredientUseCase()

	t.Log("Given the need to test creating simple ingredient.")
	{
		t.Logf("\tWhen checking that the creation returns an error")
		{
			params := CreateSimpleIngredient{
				Name:   "Harina",
				Price:  210,
				Amount: 1,
				Unit:   "KG",
			}
			SetErrorSimpleIngredientRepository("save", params.Name, fmt.Errorf("error repository"))
			err := usecase.Create(&params)
			if err == nil {
				t.Errorf("\t\tShould receive an error different from nil. %v", ballotX)
			}
			t.Logf("\t\tShould receive an error different from nil. %v %v", err, checkMark)
		}
	}
}

func TestUseOK(t *testing.T) {
	repository := NewMockSimpleIngredientRepository()
	RegisterSimpleIngredientRepository(repository)
	usecase := NewSimpleIngredientUseCase()

	t.Log("Given the need to test use simple ingredient.")
	{
		t.Logf("\tWhen checking that the using returns no error")
		{
			params := CreateSimpleIngredient{
				Name:   "Harina",
				Price:  210,
				Amount: 1,
				Unit:   "KG",
			}
			simpleIngredient := params.toModel()
			SetResponseSimpleIngredientRepositoryFindByID(simpleIngredient.ID, simpleIngredient)
			err := usecase.Use(simpleIngredient.ID)
			if err != nil {
				t.Errorf("\t\tShould receive an error equal to nil. %v %v", err, ballotX)
			}
			t.Logf("\t\tShould receive an error equal to nil. %v", checkMark)
		}
	}
}

func TestUseNotFount(t *testing.T) {
	repository := NewMockSimpleIngredientRepository()
	RegisterSimpleIngredientRepository(repository)
	usecase := NewSimpleIngredientUseCase()
	id := "id_not_fount"
	message := "not found ingrediente with id:"

	t.Log("Given the need to test use simple ingredient.")
	{
		t.Logf("\tWhen checking that the using returns no error")
		{
			err := usecase.Use(id)
			if err == nil {
				t.Fatalf("\t\tShould receive an error different from nil. %v", ballotX)
			} else {
				t.Logf("\t\tShould receive an error different from nil. %v", checkMark)
			}
			if err.Error() == (message + " " + id) {
				t.Logf("\t\tShould receive an error message '%s'. %v", message+" "+id, checkMark)
			} else {
				t.Errorf("\t\tShould receive an error message '%s'. %v", message+" "+id, ballotX)
				t.Logf("\t\tBut receive an error message '%s'", err.Error())
			}
		}
	}
}
