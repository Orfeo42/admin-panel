package data

import (
	"log"
	"time"

	"github.com/Orfeo42/admin-panel/db"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	CustomerID          uint
	Customer            Customer
	Number              string
	PaymentMethod       *string
	Amount              int
	PaidAmount          int
	Date                *time.Time
	PaymentDate         *time.Time
	ExpectedPaymentDate *time.Time
	Rows                *[]InvoiceRow
	Note                *string
}

type InvoiceRow struct {
	gorm.Model
	InvoiceID uint
	Invoice   Invoice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Number    int
	Amount    int
}

type InvoiceFilter struct {
	CustomerID      *uint
	Number          *string
	DateFrom        *time.Time
	DateTo          *time.Time
	PaymentDateFrom *time.Time
	PaymentDateTo   *time.Time
	AmountFrom      *int
	AmountTo        *int
	PaidAmountFrom  *int
	PaidAmountTo    *int
	IsPaid          *bool
	PageSize        int
	Page            int
}

func NewBaseFilter() InvoiceFilter {

	dateTo := time.Now()
	dateFrom := dateTo.AddDate(0, -1, 0)

	return InvoiceFilter{
		DateFrom: &dateFrom,
		DateTo:   &dateTo,
		Page:     1,
		PageSize: 10,
	}
}

type InvoiceResult struct {
	Invoices []Invoice
	Page     int
	PageSize int
}

func GetAllInvoice(fileter InvoiceFilter) ([]Invoice, error) {
	dbInstance, err := db.GetInstance()
	if err != nil {
		return []Invoice{}, err
	}

	var items []Invoice
	queryDB := dbInstance.Where("")

	if fileter.CustomerID != nil {
		queryDB.Where("customer_id = ?", *fileter.CustomerID)
	}
	if fileter.Number != nil {
		queryDB.Where("number = ?", *fileter.Number)
	}
	if fileter.DateFrom != nil {
		queryDB.Where("date >= ?", *fileter.DateFrom)
	}
	if fileter.DateTo != nil {
		queryDB.Where("date <= ?", *fileter.DateTo)
	}
	if fileter.PaymentDateFrom != nil {
		queryDB.Where("payment_date >= ?", *fileter.PaymentDateFrom)
	}
	if fileter.PaymentDateTo != nil {
		queryDB.Where("payment_date <= ?", *fileter.PaymentDateTo)
	}
	if fileter.AmountFrom != nil {
		queryDB.Where("amount >= ?", *fileter.AmountFrom)
	}
	if fileter.AmountTo != nil {
		queryDB.Where("amount <= ?", *fileter.AmountTo)
	}
	if fileter.PaidAmountFrom != nil {
		queryDB.Where("paid_amount >= ?", *fileter.PaidAmountFrom)
	}
	if fileter.PaidAmountTo != nil {
		queryDB.Where("paid_amount <= ?", *fileter.PaidAmountTo)
	}
	if fileter.IsPaid != nil {
		if *fileter.IsPaid {
			queryDB.Where("amount = paid_amount")
		}
		if !*fileter.IsPaid {
			queryDB.Where("amount <> paid_amount")
		}
	}
	result := queryDB.Find(&items)

	return items, result.Error
}

func CreateInvoice(invoice Invoice) (Invoice, error) {

	dbInstance, err := db.GetInstance()
	if err != nil {
		return Invoice{}, err
	}

	result := dbInstance.Create(&invoice)

	return invoice, result.Error
}

func CreateInvoiceList(invoiceList []Invoice) ([]Invoice, error) {
	result := []Invoice{}
	dbInstance, err := db.GetInstance()
	if err != nil {
		return result, err
	}
	dbInstance.Transaction(func(tx *gorm.DB) error {
		for _, invoice := range invoiceList {
			if err := tx.Create(&invoice).Error; err != nil {
				log.Fatalln("Error in creating invoice!", invoice, err)
				return err
			}
			result = append(result, invoice)
		}
		return nil
	})

	return result, nil
}
