package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	ID         string    `json:"id" gorm:"primaryKey,column:id"`
	UserID     string    `json:"user_id" gorm:"column:user_id"`
	GrandTotal uint64    `json:"grand_total" gorm:"column:grand_total"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	Status     string    `json:"status" gorm:"column:status"`

	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID;references:ID"`
}

func (Order) TableName() string {
	return "orders"
}
