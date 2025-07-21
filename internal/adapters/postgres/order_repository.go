package postgres

import (
	"github.com/cristhianfdx/order-service/internal/domain"
	"github.com/cristhianfdx/order-service/internal/ports"
	"gorm.io/gorm"
)

type OrderModel struct {
	ID     string `gorm:"primaryKey"`
	Item   string
	Amount float64
	Status string
}

func toOrderModel(order *domain.Order) *OrderModel {
	return &OrderModel{
		ID:     order.ID,
		Item:   order.Item,
		Amount: order.Amount,
		Status: string(order.Status),
	}
}

func toDomain(model *OrderModel) *domain.Order {
	return &domain.Order{
		ID:     model.ID,
		Item:   model.Item,
		Amount: model.Amount,
		Status: domain.OrderStatus(model.Status),
	}
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.OrderRepository {
	db.AutoMigrate(&OrderModel{})
	return &GormOrderRepository{db}
}

func (r *GormOrderRepository) Save(o *domain.Order) error {
	return r.db.Create(toOrderModel(o)).Error
}

func (r *GormOrderRepository) FindByID(id string) (*domain.Order, error) {
	var model OrderModel
	if err := r.db.First(&model, "id= ?", id).Error; err != nil {
		return nil, err
	}
	return toDomain(&model), nil
}

func (r *GormOrderRepository) UpdateStatus(id string, status domain.OrderStatus) error {
	return r.db.Model(&OrderModel{}).Where("id = ?", id).Update("status", status).Error
}
