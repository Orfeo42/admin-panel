package utils

import (
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"
)

const decimalSeparator = "."
const thousandsSeparator = "˙"

func ParseAmount(amount string) int {
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

func FormatAmount(number int) string {
	numberAsString := float64(number) / 100
	return formatFloat(numberAsString)
}

func AmountToString(number float64) string {
	numberAsString := formatFloat(number)
	return stringAmountToString(numberAsString)
}

func AmountIntegerToString(number int) string {
	numberAsString := FormatAmount(number)
	return stringAmountToString(numberAsString)
}

func stringAmountToString(numberAsString string) string {

	parts := strings.Split(numberAsString, decimalSeparator)
	integerPart := parts[0]

	formattedIntegerPart := formatIntegerWithCustomSeparator(integerPart, thousandsSeparator)

	if len(parts) > 1 {
		return formattedIntegerPart + decimalSeparator + parts[1]
	}

	return formattedIntegerPart
}

func formatFloat(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}
