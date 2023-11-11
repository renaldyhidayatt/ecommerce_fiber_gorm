package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Name      string  `json:"name"`
	Comment   string  `json:"comment" gorm:"type:text"`
	Rating    int     `json:"rating"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	Sentiment string  `json:"sentiment"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
}
