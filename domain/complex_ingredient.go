package domain

type ComplexIngredient struct {
	Name              string
	SimpleIngredients []SimpleIngredient
	Recipe            string
	FrequencyUse      int
}
