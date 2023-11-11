package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name          string `json:"name"`
	Description   string `json:"description"`
	SlugCategory  string `json:"slug_category"`
	ImageCategory string `json:"image_category"`
}
