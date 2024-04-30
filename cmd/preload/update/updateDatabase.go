package update

import (
	"github.com/Orfeo42/admin-panel/cmd/preload/validation"
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/db"
	"github.com/labstack/gommon/log"
)

func UpdateSchemaAndLoadData() {
	updateSchema()
	log.Info("Start creating customers")
	customers, err := inizializeCustomersData()
	if err != nil {
		log.Fatalf("Error Creating customers")
		return
	}
	log.Info("All customers are created")
	log.Info("Start creating Invoices")
	inizializeInvoiceData(customers)
	log.Info("All Invoices are created")
}

func updateSchema() error {
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

func inizializeCustomersData() ([]data.Customer, error) {
	return data.CreateCustomerList(validation.ValidateCustomersCsv())
}

func inizializeInvoiceData(customers []data.Customer) ([]data.Invoice, error) {
	invoices, err := validation.ValidateInvoiceCsv(customers)
	if err != nil {
		log.Fatalf("Error validation invoices")
		return nil, err
	}
	return data.CreateInvoiceList(invoices)

}
