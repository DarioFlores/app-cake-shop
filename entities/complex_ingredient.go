package entities

type ComplexIngredient struct {
	Name              string
	SimpleIngredients []SimpleIngredient
	Recipe            string
	FrequencyUse      int
}
