package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email;uniqueIndex" json:"email"`
	Password string `gorm:"column:password" json:"-"`
	IsActive bool   `gorm:"column:is_active;default:false" json:"is_active"`

	Reviews []ReviewModel `gorm:"foreignKey:UserID" json:"reviews"`
	Carts   []CartModel   `gorm:"foreignKey:UserID" json:"carts"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
