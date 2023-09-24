package models

import (
	"time"

	"gorm.io/gorm"
)

type CartModel struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	Price     string `json:"price" gorm:"column:price"`
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	Image     string `json:"image" gorm:"column:image"`
	Quantity  int    `json:"quantity" gorm:"column:quantity"`
	Weight    int    `json:"weight" gorm:"column:weight"`
	ProductID uint   `json:"product_id" gorm:"column:product_id"`

	User    UserModel    `json:"user" gorm:"foreignKey:UserID"`
	Product ProductModel `json:"product" gorm:"foreignKey:ProductID"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
