package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	Name         string  `json:"name" gorm:"column:name"`
	Image        string  `json:"image" gorm:"column:image"`
	CategoryID   uint    `json:"category_id" gorm:"column:category_id"`
	Description  string  `json:"description" gorm:"column:description"`
	Price        int     `json:"price" gorm:"column:price"`
	CountInStock int     `json:"count_in_stock" gorm:"column:count_in_stock"`
	Brand        string  `json:"brand" gorm:"column:brand"`
	Weight       int     `json:"weight" gorm:"column:weight"`
	Rating       float64 `json:"rating" gorm:"column:rating"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	Category    CategoryModel `json:"category" gorm:"foreignKey:CategoryID"`
	ReviewsUser []ReviewModel `json:"reviews_user" gorm:"foreignKey:ProductID"`
	Carts       []CartModel   `json:"carts" gorm:"foreignKey:ProductID"`
}
