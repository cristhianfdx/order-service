package app

import (
	"log"

	domain "github.com/cristhianfdx/order-service/internal/domain"
	servicePort "github.com/cristhianfdx/order-service/internal/ports"
	uuid "github.com/google/uuid"
)

type Service struct {
	repo      servicePort.OrderRepository
	publisher servicePort.OrderPublisher
}

func NewOrderService(repo servicePort.OrderRepository, pub servicePort.OrderPublisher) servicePort.OrderService {
	return &Service{repo: repo, publisher: pub}
}

func (s *Service) CreateOrder(item string, amount float64) (*domain.Order, error) {
	order := &domain.Order{
		ID:     uuid.New().String(),
		Item:   item,
		Amount: amount,
		Status: domain.Pending,
	}

	err := s.repo.Save(order)

	if err != nil {
		return nil, err
	}

	s.publisher.PublishOrderCreated(order)
	return order, nil
}

func (s *Service) GetOrder(id string) (*domain.Order, error) {
	return s.repo.FindByID(id)
}

func (s *Service) MarkOrderComplete(id string) error {
	log.Println("Making order as complete", id)
	return s.repo.UpdateStatus(id, domain.Completed)
}
