package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	UserID         uint   `json:"user_id" gorm:"column:user_id"`
	Name           string `json:"name" gorm:"column:name"`
	Phone          string `json:"phone" gorm:"column:phone"`
	Email          string `json:"email" gorm:"column:email"`
	Courier        string `json:"courier" gorm:"column:courier"`
	ShippingMethod string `json:"shipping_method" gorm:"column:shipping_method"`
	ShippingCost   int    `json:"shipping_cost" gorm:"column:shipping_cost"`
	TotalProduct   string `json:"total_product" gorm:"column:total_product"`
	TotalPrice     string `json:"total_price" gorm:"column:total_price"`
	TransactionID  string `json:"transaction_id" gorm:"column:transaction_id"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	OrderItems []OrderItemsModel `json:"order_items" gorm:"foreignKey:OrderID"`
}
