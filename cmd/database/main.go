package main

import (
	"fmt"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/db"
)

func main() {
	updateSchema()
}

func updateSchema() error {
	fmt.Println("Updating Schema...")
	dbInstance, errConnection := db.GetInstance()

	if errConnection != nil {
		return errConnection
	}
	err := dbInstance.AutoMigrate(
		&data.CustomerModel{},
		&data.InvoiceModel{},
		&data.OrderModel{},
	)
	if err != nil {
		fmt.Println("Error in Updating Schema")
		return err
	}
	fmt.Println("Schema Updated!")
	return nil
}
