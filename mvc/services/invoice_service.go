package services

import (
	"fmt"
	"time"

	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/labstack/gommon/log"
)

type InvoiceDTO struct {
	CustomerID          uint
	Year                int
	Number              string
	PaymentMethod       *string
	Amount              int
	PaidAmount          int
	Date                *time.Time
	PaymentDate         *time.Time
	ExpectedPaymentDate *time.Time
	Note                *string
}

/*func (dto InvoiceDTO) validateDTO() error {

}*/

func GetInvoiceFromFilter(filter repositories.InvoiceFilter) (*[]repositories.Invoice, error) {
	items, err := repositories.GetAllInvoice(filter)
	if err != nil {
		log.Errorf("error searching invoice by filter: %+v", err)
		return nil, err
	}
	return &items, nil
}

func CreateNewInvoice(dto InvoiceDTO) (*repositories.Invoice, error) {

	customer, err := repositories.GetCustomerByID(dto.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("customer not valid: %+v", err)
	}

	invoice, err := repositories.CreateInvoice(repositories.Invoice{
		CustomerID:  customer.ID,
		Customer:    *customer,
		Number:      dto.Number,
		Date:        dto.Date,
		PaymentDate: dto.PaymentDate,
		Amount:      dto.Amount,
		PaidAmount:  dto.PaidAmount,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating the invoice: %+v", err)
	}

	return invoice, nil
}

func UpdateInvoice(id uint, dto InvoiceDTO) (*repositories.Invoice, error) {

	inv, err := repositories.GetInvoiceByID(id)
	if err != nil {
		return nil, fmt.Errorf("invoice not fount: %+v", err)
	}

	customer, err := repositories.GetCustomerByID(dto.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("customer not valid: %+v", err)
	}

	inv.CustomerID = customer.ID

	inv.Customer = *customer

	inv.Number = dto.Number

	inv.Date = dto.Date

	inv.PaymentDate = dto.PaymentDate

	inv.Amount = dto.Amount

	inv.PaidAmount = dto.PaidAmount

	updateInvoice, err := repositories.UpdateInvoice(*inv)
	if err != nil {
		return nil, fmt.Errorf("error updating the invoice: %+v", err)
	}

	return updateInvoice, nil
}
