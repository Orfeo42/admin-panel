package data

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	CustomerID  uint
	Customer    Customer
	Amount      int
	PaidAmount  int
	Date        time.Time
	PaymentDate time.Time
	Rows        []InvoiceRow
}

type InvoiceRow struct {
	gorm.Model
	InvoiceID uint
	Invoice   Invoice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Number    int
	Amount    int
}
