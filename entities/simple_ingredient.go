package entities

import "time"

type SimpleIngredient struct {
	Name         string
	Price        float64
	Amount       float64
	Unit         string
	FrequencyUse int
	LastModified *time.Time
}
