package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	OrderID  uint   `json:"order_id"`
	Order    Order  `json:"order"`
}
