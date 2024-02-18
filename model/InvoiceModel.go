package model

import "time"

type InvoiceModel struct {
	Customer string
	Amount   float64
	Date     time.Time
	IsPaid   bool
}
