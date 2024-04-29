package data

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint
	Customer   Customer
	Amount     int
	Date       time.Time
}

type OrderRow struct {
	gorm.Model
	OrderID uint
	Order   Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Number  int
	Amount  int
}
