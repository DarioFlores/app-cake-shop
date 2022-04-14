package domain

import "time"

type SimpleIngredient struct {
	ID                string
	Name              string
	Price             float64
	Amount            float64
	Unit              string
	FrequencyUse      int
	LastModifiedPrice *time.Time
}
