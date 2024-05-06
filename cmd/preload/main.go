package main

import (
	"github.com/Orfeo42/admin-panel/cmd/preload/db_update"
	"github.com/Orfeo42/admin-panel/cmd/preload/validation"
	"github.com/labstack/gommon/log"
)

func main() {
	err := db_update.SchemaUpdate()
	if err != nil {
		log.Error("Error Updating Schema: %v", err)
		return
	}
	customerList, invoiceList, err := validation.ValidateCsv()
	if err != nil {
		return
	}
	db_update.LoadData(customerList, invoiceList)
}
