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
	PaymentMethod       string
	Amount              int
	PaidAmount          int
	Date                time.Time
	PaymentDate         time.Time
	ExpectedPaymentDate time.Time
	Rows                []InvoiceRow
	Note                string
}

type InvoiceRow struct {
	gorm.Model
	InvoiceID uint
	Invoice   Invoice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Number    int
	Amount    int
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
