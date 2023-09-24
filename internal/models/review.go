package models

import (
	"time"

	"gorm.io/gorm"
)

type ReviewModel struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	Comment   string `json:"comment" gorm:"column:comment"`
	Rating    int    `json:"rating" gorm:"column:rating"`
	UserID    uint   `json:"user_id" gorm:"column:user_id"`
	ProductID uint   `json:"product_id" gorm:"column:product_id"`

	User    UserModel    `json:"user" gorm:"foreignKey:UserID"`
	Product ProductModel `json:"product" gorm:"foreignKey:ProductID"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
