package models

import "time"

type Role struct {
	RoleID    uint `gorm:"primaryKey"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
