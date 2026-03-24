package models

// Order represents a payment order
type Order struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
