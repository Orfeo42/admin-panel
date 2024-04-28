package data

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceModel struct {
	gorm.Model
	Customer string
	Amount   float64
	Date     time.Time
	IsPaid   bool
}
