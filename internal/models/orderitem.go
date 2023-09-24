package models

import "gorm.io/gorm"

type OrderItemsModel struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name"`
	Quantity int    `json:"quantity" gorm:"column:quantity"`
	Price    int    `json:"price" gorm:"column:price"`
	OrderID  uint   `json:"order_id" gorm:"column:order_id"`

	Order OrderModel `json:"order" gorm:"foreignKey:OrderID"`
}
