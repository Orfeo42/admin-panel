package validation

import (
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/labstack/gommon/log"
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
)

const customerSheetName = "Totale per Cliente"
const excelFilePath = "resources/Prima Nota.xlsx"

/*
func ValidateExcel() (*[]data.Customer, *[]data.Invoice, error) {
	log.Info("Start validating Excel")
	customerList, err := ValidateCustomersCsv()
	if err != nil {
		log.Error("Error validating csv: ", err)
		return nil, nil, err
	}
	invoiceList, err := ValidateInvoiceCsv()
	if err != nil {
		log.Error("Error validating csv: ", err)
		return nil, nil, err
	}
	return customerList, invoiceList, nil
}*/

func ValidateExcel() (*[]data.Customer, *[]data.Invoice, error) {
	xlFile, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		log.Errorf("Error reading excel %s, %+v", excelFilePath, err)
		return nil, nil, err
	}
	log.Infof("Reading customer sheet")
	var customerList []data.Customer
	for _, sheet := range xlFile.Sheets {
		if sheet.Name == customerSheetName {
			for _, row := range sheet.Rows {
				text := row.Cells[0].String()
				if text == "" {
					continue
				}
				customerList = append(customerList, data.Customer{Name: row.Cells[0].String()})
			}
		}
	}
	log.Infof("Found %d customers", len(customerList))

	log.Info("Reading Invoices sheets")
	var invoiceList []data.Invoice
	for _, sheet := range xlFile.Sheets {
		if _, err := strconv.Atoi(sheet.Name); err == nil {
			log.Infof("Reading sheet %s", sheet.Name)
			for _, row := range sheet.Rows {
				customerName := row.Cells[0].String()
				if customerName == "" {
					continue
				}
				invoice := excelRowToInvoice(row.Cells)
				if invoice == nil {
					continue
				}
				invoiceList = append(invoiceList, *invoice)
			}
		}
	}
	log.Infof("Found %d invoices", len(invoiceList))
	return &customerList, &invoiceList, nil
}

func excelRowToInvoice(row []*xlsx.Cell) *data.Invoice {
	date := getDateFromExcel(row[1])
	if date == nil {
		return nil
	}
	year := date.Year()
	paymentDate := getDateFromExcel(row[5])
	expectedPaymentDate := getDateFromExcel(row[7])

	return &data.Invoice{
		Customer:            data.Customer{Name: row[0].String()},
		Date:                date,
		Year:                year,
		Number:              row[2].String(),
		Amount:              utils.ParseAmount(row[3].Value),
		PaidAmount:          utils.ParseAmount(row[4].Value),
		PaymentDate:         paymentDate,
		PaymentMethod:       utils.ParseString(row[6].String()),
		ExpectedPaymentDate: expectedPaymentDate,
		Note:                getNote(row),
	}
}

func getNote(cells []*xlsx.Cell) *string {
	rowSize := len(cells)
	if rowSize < 10 {
		return nil
	}
	return utils.ParseString(cells[9].String())
}

func getDateFromExcel(cell *xlsx.Cell) *time.Time {
	dateValue, err := cell.GetTime(false)
	if err != nil {
		//log.Errorf("Errore conversione data fattura: %+v", err)
		return nil
	}
	return &dateValue
}
