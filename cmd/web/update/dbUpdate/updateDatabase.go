package dbUpdate

import (
	"admin-panel/cmd/web/update/validation"
	"admin-panel/internal/database"

	"github.com/labstack/gommon/log"
)

func LoadData(customerList []database.Customer, invoiceList []database.Invoice) error {
	log.Info("Start creating customers")
	customerList, err := initializeCustomersData(customerList)
	if err != nil {
		log.Error("Error Creating customers: ", err)
		return err
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
		return err
	}
	log.Info("All Invoices are created")
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

func EmptyData() error {
	db := database.DBInstance()
	result := db.Exec("TRUNCATE TABLE invoices RESTART IDENTITY CASCADE")
	if result.Error != nil {
		log.Errorf("Error removing all invoices")
		return result.Error
	}
	result = db.Exec("TRUNCATE TABLE customers RESTART IDENTITY CASCADE")
	if result.Error != nil {
		log.Errorf("Error removing all customers")
		return result.Error
	}
	return nil
}
