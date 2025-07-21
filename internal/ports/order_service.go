package ports

import "github.com/cristhianfdx/order-service/internal/domain"

type OrderService interface {
	CreateOrder(item string, amount float64) (*domain.Order, error)
	GetOrder(id string) (*domain.Order, error)
	MarkOrderComplete(id string) error
}
