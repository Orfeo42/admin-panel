package validation

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/labstack/gommon/log"
)

func ValidateCsv() (*[]data.Customer, *[]data.Invoice, error) {
	log.Info("Start validating csv")
	customerList, err := ValidateCustomersCsv()
	if err != nil {
		log.Error("Error validating csv: ", err)
		return nil, nil, err
	}
	invoiceList, err := ValidateInvoiceCsv(customerList)
	if err != nil {
		log.Error("Error validating csv: ", err)
		return nil, nil, err
	}
	return customerList, invoiceList, nil
}

func ValidateCustomersCsv() (*[]data.Customer, error) {
	var customerList []data.Customer

	csvData, err := utils.ReadCsvFile("resources/customers.csv")
	if err != nil {
		return nil, err
	}
	for number, row := range csvData {
		if number == 0 {
			continue
		}
		customerList = append(customerList, csvRowToCustomer(row))
	}

	return &customerList, nil
}

func csvRowToCustomer(row []string) data.Customer {
	return data.Customer{
		Name: row[0],
	}
}

func ValidateInvoiceCsv(customers *[]data.Customer) (*[]data.Invoice, error) {

	var invList []data.Invoice

	log.Info("Starting validating invoice 2022")
	invList2022, err := validateInvoiceSingleCsv("resources/invoices - 2022.csv", customers)
	if err != nil {
		log.Info("Error in adding 2022 invoices")
		return nil, err
	}
	invList = append(invList, invList2022...)

	log.Info("Starting validating invoice 2023")
	invList2023, err := validateInvoiceSingleCsv("resources/invoices - 2023.csv", customers)
	if err != nil {
		log.Infof("Error in adding 2023 invoices")
		return nil, err
	}
	invList = append(invList, invList2023...)

	log.Info("Starting validating invoice 2024")
	invList2024, err := validateInvoiceSingleCsv("resources/invoices - 2024.csv", customers)
	if err != nil {
		log.Infof("Error in adding 2024 invoices")
		return nil, err
	}
	invList = append(invList, invList2024...)

	return &invList, nil
}

func validateInvoiceSingleCsv(csvPath string, customerList *[]data.Customer) ([]data.Invoice, error) {

	var invList []data.Invoice

	csvData, err := utils.ReadCsvFile(csvPath)
	if err != nil {
		return invList, err
	}
	for number, row := range csvData {
		if number == 0 {
			continue
		}
		invoice, err := csvRowToInvoice(row, customerList)
		if err != nil {
			log.Errorf("Error in converting csv to invoice in row %d: %+v", number, err)
		}
		invList = append(invList, invoice)
	}

	return invList, nil
}

func csvRowToInvoice(row []string, customerList *[]data.Customer) (data.Invoice, error) {
	customer, err := FindCustomerFromName(customerList, row[0])
	if err != nil {
		return data.Invoice{}, err
	}
	date := parseDate(row[1])
	year := date.Year()
	return data.Invoice{
		Customer:            customer,
		Date:                date,
		Year:                year,
		Number:              row[2],
		Amount:              parseAmount(row[3]),
		PaidAmount:          parseAmount(row[4]),
		PaymentDate:         parseDate(row[5]),
		PaymentMethod:       parseString(row[6]),
		ExpectedPaymentDate: parseDate(row[7]),
		Note:                parseString(row[8]),
	}, nil
}

func FindCustomerFromName(customers *[]data.Customer, name string) (data.Customer, error) {
	for _, customer := range *customers {
		if customer.Name == name {
			return customer, nil
		}
	}
	log.Infof("No customer with name '%s' found", name)
	return data.Customer{}, errors.New("No customer with name " + name + " found")
}

func parseDate(date string) *time.Time {
	if date == "" {
		return nil
	}
	parsedDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Info("Date not parsable:", date)
		return nil
	}
	return &parsedDate
}

func parseAmount(amount string) int {
	if amount == "" {
		return 0
	}
	amount = strings.Replace(amount, ",", ".", -1)
	parsedAmount, err := strconv.ParseFloat(strings.TrimSpace(amount), 64)
	if err != nil {
		log.Info("Amount not parsable:", amount)
		return 0
	}
	return int(parsedAmount * 100)
}

func parseString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
