package main

import (
	"fmt"
	"github.com/DarioFlores/app-cake-shop/interfaces/repository"
	"github.com/DarioFlores/app-cake-shop/usecases/simple_ingredient"
)

func main() {
	fmt.Println("AppCakeShop")
	simpleIngredientRepository := repository.NewSimpleIngredientRepository()
	simple_ingredient.RegisterSimpleIngredientRepository(simpleIngredientRepository)
	saveSimpleIngredient := simple_ingredient.NewSimpleIngredientUseCase()
	simpleIngredient := simple_ingredient.SaveSimpleIngredientRequest{
		Name:   "Harina",
		Price:  120,
		Amount: 1,
		Unit:   "KG",
	}
	err := saveSimpleIngredient.Save(&simpleIngredient)

	if err != nil {
		fmt.Println("Error al guardar ingrediente:", err.Error())
	}
	ingredients, _ := simpleIngredientRepository.GetAll()
	fmt.Println("Salio todo bien")
	fmt.Println(fmt.Sprintf("Ingredientes guardados: %+v", ingredients))
}
