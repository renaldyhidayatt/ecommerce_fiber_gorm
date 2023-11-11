package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Name      string  `json:"name"`
	Price     string  `json:"price"`
	Image     string  `json:"image"`
	Quantity  int     `json:"quantity"`
	Weight    int     `json:"weight"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
}
