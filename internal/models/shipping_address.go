package models

import (
	"time"

	"gorm.io/gorm"
)

type ShippingAddressModel struct {
	gorm.Model
	Alamat   string `json:"alamat" gorm:"column:alamat"`
	Provinsi string `json:"provinsi" gorm:"column:provinsi"`
	Negara   string `json:"negara" gorm:"column:negara"`
	Kota     string `json:"kota" gorm:"column:kota"`
	OrderID  uint   `json:"order_id" gorm:"column:order_id"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
