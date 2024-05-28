package services

import (
	"fmt"

	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/labstack/gommon/log"
)

func GetInvoiceFromFilter(filter repositories.InvoiceFilter) (*[]repositories.Invoice, error) {
	items, err := repositories.GetAllInvoice(filter)
	if err != nil {
		log.Errorf("error searching invoice by filter: %+v", err)
		return nil, err
	}
	return &items, nil
}

func CreateNewInvoice(invoiceIn repositories.Invoice) (*repositories.Invoice, error) {

	customer, err := repositories.GetCustomerByID(invoiceIn.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("customer not valid: %+v", err)
	}

	invoice, err := repositories.CreateInvoice(repositories.Invoice{
		CustomerID:          customer.ID,
		Customer:            *customer,
		Number:              invoiceIn.Number,
		PaymentMethod:       invoiceIn.PaymentMethod,
		Date:                invoiceIn.Date,
		PaymentDate:         invoiceIn.PaymentDate,
		Amount:              invoiceIn.Amount,
		PaidAmount:          invoiceIn.PaidAmount,
		ExpectedPaymentDate: invoiceIn.ExpectedPaymentDate,
		Note:                invoiceIn.Note,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating the invoice: %+v", err)
	}

	return invoice, nil
}

func UpdateInvoice(id uint, invoiceIn repositories.Invoice) (*repositories.Invoice, error) {

	inv, err := repositories.GetInvoiceByID(id)
	if err != nil {
		return nil, fmt.Errorf("invoice not fount: %+v", err)
	}

	customer, err := repositories.GetCustomerByID(invoiceIn.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("customer not valid: %+v", err)
	}

	inv.CustomerID = customer.ID

	inv.Customer = *customer

	inv.Number = invoiceIn.Number

	inv.PaymentMethod = invoiceIn.PaymentMethod

	inv.Date = invoiceIn.Date

	inv.PaymentDate = invoiceIn.PaymentDate

	inv.Amount = invoiceIn.Amount

	inv.PaidAmount = invoiceIn.PaidAmount

	inv.ExpectedPaymentDate = invoiceIn.ExpectedPaymentDate

	inv.Note = invoiceIn.Note

	updateInvoice, err := repositories.UpdateInvoice(*inv)
	if err != nil {
		return nil, fmt.Errorf("error updating the invoice: %+v", err)
	}

	return updateInvoice, nil
}
