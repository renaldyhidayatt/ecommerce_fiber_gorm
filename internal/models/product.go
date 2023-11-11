package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string   `json:"name"`
	CategoryID   uint     `json:"category_id"`
	Category     Category `json:"category"`
	Description  string   `json:"description"`
	Price        int      `json:"price"`
	CountInStock int      `json:"count_in_stock"`
	Brand        string   `json:"brand"`
	Weight       int      `json:"weight"`
	Rating       float64  `json:"rating"`
	SlugProduct  string   `json:"slug_product"`
	ImageProduct string   `json:"image_product"`
}
