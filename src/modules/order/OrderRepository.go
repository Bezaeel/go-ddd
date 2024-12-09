package order

import (
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderById(id int64) *OrderEntity
	CreateOrder(order OrderEntity) (*OrderEntity, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (repo *OrderRepositoryImpl) CreateOrder(order OrderEntity) (*OrderEntity, error) {
	result := repo.db.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepositoryImpl) GetOrderById(id int64) *OrderEntity {
	var order OrderEntity
	order.Id = 1
	order.ShipmentNumber = 3

	return &order
}
