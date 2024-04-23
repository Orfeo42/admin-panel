package data

import "time"

type OrderModel struct {
	Customer string
	Amount   float64
	Date     time.Time
}
