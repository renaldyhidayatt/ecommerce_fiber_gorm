package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name"`
	Email    string   `json:"email" gorm:"unique"`
	Password string   `json:"password" gorm:"type:varchar; not null"`
	IsStaff  bool     `json:"is_staff"`
	Reviews  []Review `json:"reviews" gorm:"many2many:user_reviews;"`
	Carts    []Cart   `json:"carts" gorm:"many2many:user_carts;"`
}
