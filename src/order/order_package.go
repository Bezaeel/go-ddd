package order

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DI struct {
	OrderRepository OrderRepository
	OrderService    OrderService
}

// wire dependencies
func RegisterPackage(app *fiber.App, db *gorm.DB) {
	orderRepository := NewOrderRepository(db)
	orderService := NewOrderService(orderRepository)
	NewOrderController(app, orderService)
}
