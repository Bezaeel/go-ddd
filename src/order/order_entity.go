package order

import "time"

type Order struct {
	Id             int64           `gorm:"column:id;primaryKey"`
	ShipmentNumber int64           `gorm:"column:shipment_number"`
	CargoId        int             `gorm:"column:cargo_id"`
	IsShipped      bool            `gorm:"column:is_shipped"`
	CreatedAt      time.Time       `gorm:"column:created_at"`
	OrderLineItems []OrderLineItem `gorm:"referenceKey:OrderId"`
}


type OrderLineItem struct {
	Id        int64 `gorm:"column:id;primaryKey"`
	ProductId int64 `gorm:"column:product_id"`
	SellerId  int64 `gorm:"column:seller_id"`
	OrderId   int64 `gorm:"column:order_id"`
}