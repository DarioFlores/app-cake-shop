package entities

import "time"

type Invoice struct {
	Items        []Item
	ClientName   string
	PaymentForm  string
	OrderDate    *time.Time
	DeliveryDate *time.Time
}
