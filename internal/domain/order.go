package domain

type OrderStatus string

const (
	Pending   OrderStatus = "PENDING"
	Completed OrderStatus = "COMPLETED"
)

type Order struct {
	ID     string      `json:"id"`
	Item   string      `json:"item"`
	Amount float64     `json:"amount"`
	Status OrderStatus `json:"status"`
}
