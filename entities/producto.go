package entities

type Product struct {
	Nombre             string
	SimpleIngredients  []SimpleIngredient
	ComplexIngredients []ComplexIngredient
	PercentageProfit   float64
	Recipe             string
}
