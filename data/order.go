package data

import (
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	Customer string
	Amount   float64
	Date     time.Time
}
