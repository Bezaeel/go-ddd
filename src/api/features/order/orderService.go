package order

import result "api/common"

type OrderService interface {
	CreateOrder(command CreateOrderCommand) result.Result[*CreateOrderResponse]
	GetOrderById(id int64) result.Result[*OrderEntity]
}

type OrderServiceImpl struct {
	orderRepository OrderRepository
}

func (service OrderServiceImpl) GetOrderById(id int64) result.Result[*OrderEntity] {
	order := service.orderRepository.GetOrderById(id)
	return result.Result[*OrderEntity]{
		Value: order,
	}
}

func (service OrderServiceImpl) CreateOrder(command CreateOrderCommand) result.Result[*CreateOrderResponse] {
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
	return &OrderServiceImpl{orderRepository: orderRepository}
}
