package db_update

import (
	"github.com/Orfeo42/admin-panel/cmd/preload/validation"
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/db"
	"github.com/labstack/gommon/log"
)

func LoadData(customerList *[]data.Customer, invoiceList *[]data.Invoice) {
	log.Info("Start creating customers")
	customerList, err := initializeCustomersData(customerList)
	if err != nil {
		log.Error("Error Creating customers", err)
		return
	}
	log.Info("All customers are created")
	log.Info("Start creating Invoices")
	log.Info("Setting correct customer to all invoices")

	for number, invoice := range *invoiceList {
		customer, err := validation.FindCustomerFromName(customerList, invoice.Customer.Name)
		if err != nil {
			log.Errorf("Error Finding Customer from Name with name %s: %+v", invoice.Customer.Name, err)
			continue
		}
		invoice.Customer = customer
		(*invoiceList)[number] = invoice
	}
	log.Info("End of customer assignation")
	_, err = initializeInvoiceData(invoiceList)
	if err != nil {
		log.Error("Error creating Invoices", err)
		return
	}
	log.Info("All Invoices are created")
}

func SchemaUpdate() error {
	log.Info("Updating Schema...")
	dbInstance, errConnection := db.GetInstance()

	if errConnection != nil {
		return errConnection
	}
	err := dbInstance.AutoMigrate(
		&data.Customer{},
		&data.Invoice{},
		&data.InvoiceRow{},
		&data.Order{},
		&data.OrderRow{},
	)
	if err != nil {
		log.Fatalf("Error in Updating Schema")
		return err
	}
	log.Info("Schema Updated!")
	return nil
}

func initializeCustomersData(customerList *[]data.Customer) (*[]data.Customer, error) {
	return data.CreateCustomerList(customerList)
}

func initializeInvoiceData(invoiceList *[]data.Invoice) (*[]data.Invoice, error) {
	return data.CreateInvoiceListNoTransaction(invoiceList)

}
