package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID          string     `json:"id" gorm:"primaryKey,column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	Slug        string     `json:"slug" gorm:"column:slug"`
	SKU         string     `json:"sku" gorm:"column:sku"`
	Description string     `json:"description" gorm:"column:description"`
	Price       int64      `json:"price" gorm:"column:price"`
	Qty         int        `json:"qty" gorm:"column:qty"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedBy   string     `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   string     `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
