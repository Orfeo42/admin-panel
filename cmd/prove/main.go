package main

func main() {
	/*csvData := utils.ReadCsvFile("../preload/resources/invoices.csv")
	//result := []repositories.Invoice{}
	for number, row := range csvData {
		if number == 0 {
			continue
		}
		csvRowToInvoice(row)
	}*/
}

/*func csvRowToInvoice(row []string) repositories.Invoice {

	//Cliente,Data Fattura,Numero Fattura,Importo Fattura,Importo Pagato,Data Pagamenti,Modalità pagamento,Data Pagamento prevista,Note

	return repositories.Invoice{
		Customer: repositories.Customer{
			Name: row[0],
		},
		Date:                parseDate(row[1]),
		Number:              row[2],
		Amount:              parseAmount(row[3]),
		PaidAmount:          parseAmount(row[4]),
		PaymentDate:         parseDate(row[5]),
		PaymentMethod:       row[6],
		ExpectedPaymentDate: parseDate(row[7]),
		Note:                row[8],
	}
}

func parseDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	parsedDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Info("Data non parsata:", date)
		return time.Time{}
	}
	return parsedDate
}

func parseAmount(amount string) int {
	if amount == "" {
		return 0
	}
	amount = strings.Replace(amount, ",", ".", -1)
	parsedAmount, err := strconv.ParseFloat(strings.TrimSpace(amount), 64)
	if err != nil {
		log.Info("Importo non parsato:", amount)
		return 0
	}
	return int(parsedAmount * 100)
}
*/
