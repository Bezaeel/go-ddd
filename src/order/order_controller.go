package order

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	app          fiber.Router
	orderService OrderService
}

func NewOrderController(_app *fiber.App, _orderService OrderService) {

	v1 := _app.Group("/api")
	o := &orderController{
		app:          v1,
		orderService: _orderService,
	}
	o.GetOrderById()
	o.CreateOrder()

	v2 := _app.Group("/api/v2")
	p := &orderController{
		app:          v2,
		orderService: _orderService,
	}

	p.GetOrderById()

}

func (o *orderController) CreateOrder() fiber.Router {
	return o.app.Post("/orders", func(ctx *fiber.Ctx) error {
		p := new(CreateOrderCommand)
		_ = ctx.BodyParser(p)

		order := o.orderService.CreateOrder(*p)

		return ctx.Status(fiber.StatusOK).JSON(order)
	})
}

func (o *orderController) GetOrderById() fiber.Router {
	return o.app.Get("/orders/:id", func(ctx *fiber.Ctx) error {

		orderId := ctx.Params("id")
		id, err := strconv.ParseInt(orderId, 10, 64)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON("Order id is not valid")
		}

		order := o.orderService.GetOrderById(id)

		if order.Value == nil {
			return ctx.Status(fiber.StatusNotFound).JSON("Order Not Found, sorry! :(")
		}

		return ctx.Status(fiber.StatusOK).JSON(order)
	})
}
