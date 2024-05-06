package validation

import (
	"errors"
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
	invoiceList, err := ValidateInvoiceCsv()
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

func ValidateInvoiceCsv() (*[]data.Invoice, error) {

	var invList []data.Invoice

	log.Info("Starting validating invoice 2022")
	invList2022, err := validateInvoiceSingleCsv("resources/invoices - 2022.csv")
	if err != nil {
		log.Info("Error in adding 2022 invoices")
		return nil, err
	}
	invList = append(invList, invList2022...)

	log.Info("Starting validating invoice 2023")
	invList2023, err := validateInvoiceSingleCsv("resources/invoices - 2023.csv")
	if err != nil {
		log.Infof("Error in adding 2023 invoices")
		return nil, err
	}
	invList = append(invList, invList2023...)

	log.Info("Starting validating invoice 2024")
	invList2024, err := validateInvoiceSingleCsv("resources/invoices - 2024.csv")
	if err != nil {
		log.Infof("Error in adding 2024 invoices")
		return nil, err
	}
	invList = append(invList, invList2024...)

	return &invList, nil
}

func validateInvoiceSingleCsv(csvPath string) ([]data.Invoice, error) {

	var invList []data.Invoice

	csvData, err := utils.ReadCsvFile(csvPath)
	if err != nil {
		return invList, err
	}
	for number, row := range csvData {
		if number == 0 {
			continue
		}
		if row[2] == "" {
			continue
		}
		invoice, err := csvRowToInvoice(row)
		if err != nil {
			log.Errorf("Error in converting csv to invoice in row %d: %+v", number, err)
		}
		invList = append(invList, invoice)
	}

	return invList, nil
}

func csvRowToInvoice(row []string) (data.Invoice, error) {
	date := utils.ParseDate(row[1])
	year := date.Year()
	return data.Invoice{
		Customer:            data.Customer{Name: row[0]},
		Date:                date,
		Year:                year,
		Number:              row[2],
		Amount:              utils.ParseAmount(row[3]),
		PaidAmount:          utils.ParseAmount(row[4]),
		PaymentDate:         utils.ParseDate(row[5]),
		PaymentMethod:       utils.ParseString(row[6]),
		ExpectedPaymentDate: utils.ParseDate(row[7]),
		Note:                utils.ParseString(row[8]),
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
