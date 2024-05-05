package data

import (
	"github.com/labstack/gommon/log"
	"time"

	"github.com/Orfeo42/admin-panel/db"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	CustomerID          uint
	Customer            Customer
	Year                int    `gorm:"uniqueIndex:year_number"`
	Number              string `gorm:"uniqueIndex:year_number"`
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

func GetAllInvoice(filter InvoiceFilter) ([]Invoice, error) {
	dbInstance, err := db.GetInstance()
	if err != nil {
		return []Invoice{}, err
	}

	var items []Invoice
	queryDB := dbInstance.Where("")

	if filter.CustomerID != nil {
		queryDB.Where("customer_id = ?", *filter.CustomerID)
	}
	if filter.Number != nil {
		queryDB.Where("number = ?", *filter.Number)
	}
	if filter.DateFrom != nil {
		queryDB.Where("date >= ?", *filter.DateFrom)
	}
	if filter.DateTo != nil {
		queryDB.Where("date <= ?", *filter.DateTo)
	}
	if filter.PaymentDateFrom != nil {
		queryDB.Where("payment_date >= ?", *filter.PaymentDateFrom)
	}
	if filter.PaymentDateTo != nil {
		queryDB.Where("payment_date <= ?", *filter.PaymentDateTo)
	}
	if filter.AmountFrom != nil {
		queryDB.Where("amount >= ?", *filter.AmountFrom)
	}
	if filter.AmountTo != nil {
		queryDB.Where("amount <= ?", *filter.AmountTo)
	}
	if filter.PaidAmountFrom != nil {
		queryDB.Where("paid_amount >= ?", *filter.PaidAmountFrom)
	}
	if filter.PaidAmountTo != nil {
		queryDB.Where("paid_amount <= ?", *filter.PaidAmountTo)
	}
	if filter.IsPaid != nil {
		if *filter.IsPaid {
			queryDB.Where("amount = paid_amount")
		}
		if !*filter.IsPaid {
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
	var result []Invoice
	dbInstance, err := db.GetInstance()
	if err != nil {
		return result, err
	}
	err = dbInstance.Transaction(func(tx *gorm.DB) error {
		for _, invoice := range invoiceList {
			if err := tx.Create(&invoice).Error; err != nil {
				log.Errorf("Error in creating invoice %+v: %+v", invoice, err)
				return err
			}
			result = append(result, invoice)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
