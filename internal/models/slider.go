package models

import "gorm.io/gorm"

type Slider struct {
	gorm.Model
	Name  string `json:"name"`
	Image string `json:"image"`
}
