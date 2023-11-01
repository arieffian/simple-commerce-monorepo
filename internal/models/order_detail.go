package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model

	OrderID   string    `json:"order_id" gorm:"primaryKey,column:order_id"`
	ProductID string    `json:"product_id" gorm:"primaryKey,column:product_id"`
	SubTotal  uint64    `json:"sub_total" gorm:"column:sub_total"`
	Qty       uint      `json:"qty" gorm:"column:qty"`
	Price     uint32    `json:"price" gorm:"column:price"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`

	Order Order `gorm:"foreignKey:OrderID;references:ID"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}
