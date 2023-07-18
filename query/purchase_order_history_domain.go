package query

import (
	_const "com.ai.orch-purchase-order-inquiry/const"
)

type PurchaseOrderHistory struct {
	ItemId      int                `gorm:"column:item_id;primaryKey"`
	Price       int                `gorm:"column:price"`
	Quantity    int                `gorm:"column:quantity"`
	CustomerId  int                `gorm:"column:customer_id"`
	TotalAmount int                `gorm:"column:total_amount"`
	Status      _const.OrderStatus `gorm:"column:status"`
}
