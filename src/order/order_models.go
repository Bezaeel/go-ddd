package order

// CreateOrderResponse
// @Description Response from creating Order
type CreateOrderResponse struct {
	OrderId int64
}

func MapFromEntity(order Order) *CreateOrderResponse {
	return &CreateOrderResponse{
		OrderId: order.Id,
	}
}

// CreateOrderCommand
// @Description Request about creating Order
type CreateOrderCommand struct {
	// shipment no of Order
	ShipmentNumber int64 `json:"shipmentNumber" validate:"required"`
	// cargo id of Order
	CargoId int `json:"cargoId" validate:"required"`
	// cargo id of Order
	OrderLineItems []CreateOrderLineItemCommand `json:"lineItems" validate:"required"`
}

func MapToOrder(command CreateOrderCommand) Order {
	var items []OrderLineItem

	if command.OrderLineItems != nil {
		for _, lineItemCommand := range command.OrderLineItems {
			items = append(items, OrderLineItem{
				SellerId:  lineItemCommand.SellerId,
				ProductId: lineItemCommand.ProductId,
			})
		}
	}

	return Order{
		IsShipped:      false,
		CargoId:        command.CargoId,
		ShipmentNumber: command.ShipmentNumber,
		OrderLineItems: items,
	}
}

type CreateOrderLineItemCommand struct {
	// product id of Order line items
	ProductId int64 `json:"productId" validate:"required"`
	// product id of Order line items
	SellerId int64 `json:"sellerId" validate:"required"`
}
