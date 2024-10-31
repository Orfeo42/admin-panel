package dbupdate

import (
	"admin-panel/internal/database"
	"admin-panel/preload/validation"

	"github.com/labstack/gommon/log"
)

func LoadData(customerList []database.Customer, invoiceList []database.Invoice) {
	log.Info("Start creating customers")
	customerList, err := initializeCustomersData(customerList)
	if err != nil {
		log.Error("Error Creating customers: ", err)
		return
	}
	log.Info("All customers are bean created")
	log.Info("Start creating Invoices")
	log.Info("Setting correct customer to all invoices")

	for number, invoice := range invoiceList {
		customer, err := validation.FindCustomerFromName(customerList, invoice.Customer.Name)
		if err != nil {
			log.Errorf("Error Finding Customer from Name with name %s: %+v", invoice.Customer.Name, err)
			continue
		}
		invoice.Customer = customer
		(invoiceList)[number] = invoice
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
	dbInstance := database.DBInstance()

	err := dbInstance.AutoMigrate(
		&database.Customer{},
		&database.Invoice{},
		&database.InvoiceRow{},
	)
	if err != nil {
		log.Fatalf("Error in Updating Schema")
		return err
	}
	log.Info("Schema Updated!")
	return nil
}

func initializeCustomersData(customerList []database.Customer) ([]database.Customer, error) {
	custRepo := database.CustomerRepositoryInstance()
	return custRepo.CreateListInTransaction(customerList)
}

func initializeInvoiceData(invoiceList []database.Invoice) ([]database.Invoice, error) {
	invRepo := database.InvoiceRepositoryInstance()

	return invRepo.CreateListInTransaction(invoiceList)
}
