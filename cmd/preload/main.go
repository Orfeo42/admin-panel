package main

import (
	"fmt"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/db"
)

func main() {
	updateSchema()
	inizializeCustomersData()
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

func inizializeCustomersData() ([]data.CustomerModel, error) {
	customerList := []data.CustomerModel{
		{
			Name:    "Pippo",
			Surname: "Pippos",
			Address: "Via di pippo 1",
			Phone:   "0575-617368",
			Email:   "pippo@aaa.com",
		},
		{
			Name:    "Pluto",
			Surname: "Plutos",
			Address: "Via di Pluto 2",
			Phone:   "0575-666666",
			Email:   "pluto@bbb.com",
		},
	}
	return data.CreateCustomerList(customerList)
}
