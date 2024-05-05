package update

import (
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/db"
	"github.com/labstack/gommon/log"
)

func LoadData(customerList *[]data.Customer, invoiceList *[]data.Invoice) {
	log.Info("Start creating customers")
	_, err := initializeCustomersData(customerList)
	if err != nil {
		log.Fatalf("Error Creating customers")
		return
	}
	log.Info("All customers are created")
	log.Info("Start creating Invoices")
	_, err = initializeInvoiceData(invoiceList)
	if err != nil {
		log.Fatalf("Error creating Invoices")
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

func initializeCustomersData(customerList *[]data.Customer) ([]data.Customer, error) {
	return data.CreateCustomerList(customerList)
}

func initializeInvoiceData(invoiceList *[]data.Invoice) ([]data.Invoice, error) {
	return data.CreateInvoiceList(invoiceList)

}
