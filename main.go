package main

import (
	"fmt"
	"github.com/DarioFlores/app-cake-shop/entities"
	"github.com/DarioFlores/app-cake-shop/gateway"
	"github.com/DarioFlores/app-cake-shop/repository"
	"github.com/DarioFlores/app-cake-shop/use_case"
)

func main() {
	fmt.Println("AppCakeShop")
	simpleIngredientRepository := repository.NewSimpleIngredientRepository()
	gateway.RegisterSimpleIngredientRepository(simpleIngredientRepository)
	saveSimpleIngredient := use_case.NewSaveSimpleIngredientUseCase()
	simpleIngredient := entities.SimpleIngredient{
		Name:         "Harina",
		Price:        120,
		Amount:       1,
		Unit:         "KG",
		FrequencyUse: 0,
		LastModified: nil,
	}
	err := saveSimpleIngredient.Save(&simpleIngredient)

	if err != nil {
		fmt.Println("Error al guardar ingrediente:", err.Error())
	}
	ingredients, _ := simpleIngredientRepository.GetAll()
	fmt.Println("Salio todo bien")
	fmt.Println(fmt.Sprintf("Ingredientes guardados: %+v", ingredients))
}
