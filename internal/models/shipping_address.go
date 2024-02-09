package models

import (
	"gorm.io/gorm"
)

type ShippingAddress struct {
	gorm.Model
	Alamat   string `json:"alamat"`
	Provinsi string `json:"provinsi"`
	Negara   string `json:"negara"`
	Kota     string `json:"kota"`
	OrderID  uint   `json:"order_id"`
}
