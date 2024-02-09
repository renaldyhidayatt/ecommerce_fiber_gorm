package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint            `json:"user_id"`
	User            User            `json:"user"`
	Name            string          `json:"name"`
	Phone           string          `json:"phone"`
	Email           string          `json:"email"`
	Courier         string          `json:"courier"`
	ShippingMethod  string          `json:"shipping_method"`
	ShippingCost    int             `json:"shipping_cost"`
	TotalProduct    string          `json:"total_product"`
	TotalPrice      int             `json:"total_price"`
	TransactionID   string          `json:"transaction_id"`
	OrderItems      []OrderItems    `json:"order_items" gorm:"foreignKey:OrderID"`
	ShippingAddress ShippingAddress `json:"shipping_address" gorm:"foreignKey:OrderID"`
}
