package models

import "time"

type RoleModel struct {
	RoleID    uint      `json:"role_id" gorm:"primaryKey"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
