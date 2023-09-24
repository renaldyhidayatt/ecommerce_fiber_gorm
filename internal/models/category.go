package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryModel struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	Products []ProductModel `json:"products" gorm:"foreignKey:CategoryID"`
}
