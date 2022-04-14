package main

import (
	"encoding/json"
	"fmt"
	"github.com/DarioFlores/app-cake-shop/interfaces/repository"
	"github.com/DarioFlores/app-cake-shop/usecases"
)

func main() {
	fmt.Println("AppCakeShop")
	simpleIngredientRepository := repository.NewSimpleIngredientRepository()
	usecases.RegisterSimpleIngredientRepository(simpleIngredientRepository)
	saveSimpleIngredient := usecases.NewSimpleIngredientUseCase()
	simpleIngredient := usecases.CreateSimpleIngredient{
		Name:   "Harina",
		Price:  120,
		Amount: 1,
		Unit:   "KG",
	}
	err := saveSimpleIngredient.Create(&simpleIngredient)

	if err != nil {
		fmt.Println("Error al guardar ingrediente:", err.Error())
	}
	ingredients, _ := simpleIngredientRepository.GetAll()
	fmt.Println("Salio todo bien")
	jsonIngredient, _ := json.MarshalIndent(ingredients, "", "    ")
	fmt.Println(fmt.Sprintf("Ingredientes guardados: %s", jsonIngredient))
}
