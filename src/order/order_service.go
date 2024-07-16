package order

import result "go-ddd/src/common"

type orderService struct {
	orderRepository OrderRepository
}

type OrderService interface {
	CreateOrder(command CreateOrderCommand) result.Result[*CreateOrderResponse]
	GetOrderById(id int64) result.Result[*Order]
}

func (service orderService) GetOrderById(id int64) result.Result[*Order] {
	order := service.orderRepository.GetOrderById(id)
	return result.Result[*Order]{
		Value: order,
	}
}

func (service orderService) CreateOrder(command CreateOrderCommand) result.Result[*CreateOrderResponse] {
	order := MapToOrder(command)
	_order, err := service.orderRepository.CreateOrder(order)
	if err != nil {
		return result.Result[*CreateOrderResponse]{Err: err}
	}
	return result.Result[*CreateOrderResponse]{
		Value: MapFromEntity(*_order),
	}
}

func NewOrderService(orderRepository OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}
