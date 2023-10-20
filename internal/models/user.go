package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        string     `json:"id" gorm:"primaryKey,column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	Email     string     `json:"email" gorm:"column:email"`
	Type      string     `json:"type" gorm:"column:type"`
	Status    string     `json:"status" gorm:"column:status"`
	CreatedBy string     `json:"created_by" gorm:"column:created_by"`
	UpdatedBy string     `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`

	Creator *User `gorm:"foreignKey:CreatedBy;references:ID"`
	Updater *User `gorm:"foreignKey:UpdatedBy;references:ID"`
}

func (User) TableName() string {
	return "users"
}
