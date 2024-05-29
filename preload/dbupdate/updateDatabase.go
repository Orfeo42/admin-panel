package dbupdate

import (
	db "admin-panel/internal/database"
	"admin-panel/mvc/repositories"
	"admin-panel/preload/validation"

	"github.com/labstack/gommon/log"
)

func LoadData(customerList *[]repositories.Customer, invoiceList *[]repositories.Invoice) {
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
	dbInstance, errConnection := db.DBInstance()

	if errConnection != nil {
		return errConnection
	}
	err := dbInstance.AutoMigrate(
		&repositories.Customer{},
		&repositories.Invoice{},
		&repositories.InvoiceRow{},
	)
	if err != nil {
		log.Fatalf("Error in Updating Schema")
		return err
	}
	log.Info("Schema Updated!")
	return nil
}

func initializeCustomersData(customerList *[]repositories.Customer) (*[]repositories.Customer, error) {
	return repositories.CreateCustomerList(customerList)
}

func initializeInvoiceData(invoiceList *[]repositories.Invoice) (*[]repositories.Invoice, error) {
	return repositories.CreateInvoiceList(invoiceList)

}
