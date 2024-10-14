package models

type Order struct {
	Id         string  `json:"id"`
	TotalPrice float64 `json:"total_price"`
	CustomerId int     `json:"customer_id"`
}
