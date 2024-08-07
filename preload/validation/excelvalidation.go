package validation

import (
	"strconv"
	"time"

	"admin-panel/internal/database"
	"admin-panel/utils"

	"github.com/labstack/gommon/log"
	"github.com/tealeg/xlsx"
)

const customerSheetName = "Totale per Cliente"
const excelFilePath = "resources/Prima Nota.xlsx"

func ValidateExcel() ([]database.Customer, []database.Invoice, error) {
	xlFile, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		log.Errorf("Error reading excel %s, %+v", excelFilePath, err)
		return nil, nil, err
	}
	log.Infof("Reading customer sheet")
	var customerList []database.Customer
	for _, sheet := range xlFile.Sheets {
		if sheet.Name == customerSheetName {
			for index, row := range sheet.Rows {
				if index == 0 {
					continue
				}
				text := row.Cells[0].String()
				if text == "" {
					continue
				}
				customerList = append(customerList, database.Customer{Name: row.Cells[0].String()})
			}
		}
	}
	log.Infof("Found %d customers", len(customerList))

	log.Info("Reading Invoices sheets")
	var invoiceList []database.Invoice
	for _, sheet := range xlFile.Sheets {
		if _, err := strconv.Atoi(sheet.Name); err == nil {
			log.Infof("Reading sheet %s", sheet.Name)
			for index, row := range sheet.Rows {
				if index == 0 {
					continue
				}
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
	return customerList, invoiceList, nil
}

func excelRowToInvoice(row []*xlsx.Cell) *database.Invoice {
	date := getDateFromExcel(row[1])
	if date == nil {
		return nil
	}
	invoiceNumber := row[2].String()
	if invoiceNumber == "" {
		invoiceNumber = "Cassa del: " + utils.FormatTimePtrToTable(date)
	}
	year := date.Year()
	paymentDate := getDateFromExcel(row[5])
	expectedPaymentDate := getDateFromExcel(row[7])

	return &database.Invoice{
		Customer:            database.Customer{Name: row[0].String()},
		Date:                date,
		Year:                year,
		Number:              invoiceNumber,
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
		return nil
	}
	return &dateValue
}
