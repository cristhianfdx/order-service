package ports

import "github.com/cristhianfdx/order-service/internal/domain"

type OrderPublisher interface {
	PublishOrderCreated(order *domain.Order) error
}
