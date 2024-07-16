package order

import (
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

type OrderRepository interface {
	GetOrderById(id int64) *Order
	CreateOrder(order Order) (*Order, error)
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (repo *orderRepository) CreateOrder(order Order) (*Order, error) {
	result := repo.db.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *orderRepository) GetOrderById(id int64) *Order {
	var order Order
	order.Id = 1
	order.ShipmentNumber = 3

	return &order
}
