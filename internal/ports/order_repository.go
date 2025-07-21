package ports

import (
	"github.com/cristhianfdx/order-service/internal/domain"
)

type OrderRepository interface {
	Save(order *domain.Order) error
	FindByID(id string) (*domain.Order, error)
	UpdateStatus(id string, status domain.OrderStatus) error
}
