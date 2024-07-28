package main

import (
	"admin-panel/preload/dbupdate"
	"admin-panel/preload/validation"

	"github.com/labstack/gommon/log"
)

func main() {
	err := dbupdate.SchemaUpdate()
	if err != nil {
		log.Error("error Updating Schema: ", err)
		return
	}
	customerList, invoiceList, err := validation.ValidateExcel()
	if err != nil {
		return
	}
	dbupdate.LoadData(customerList, invoiceList)
}
