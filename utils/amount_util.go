package utils

import (
	"errors"
	"strconv"
	"strings"
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
		return 0
	}
	return int(parsedAmount * 100)
}

func formatFloat(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

func FormatAmount[T int | int64](number T) string {
	numberAsString := float64(number) / 100
	return formatFloat(numberAsString)
}

func addThousandsSeparator(numberAsString string) string {

	parts := strings.Split(numberAsString, decimalSeparator)
	integerPart := parts[0]

	formattedIntegerPart := formatIntWithSeparator(integerPart, thousandsSeparator)

	if len(parts) > 1 {
		return formattedIntegerPart + decimalSeparator + parts[1]
	}

	return formattedIntegerPart
}

func AmountToString(number float64) string {
	numberAsString := formatFloat(number)
	return addThousandsSeparator(numberAsString)
}

func AmountIntegerToString[T int | int64](number T) string {
	numberAsString := FormatAmount(number)
	return addThousandsSeparator(numberAsString)
}

func FormatIntToAmount[T int | int64](valueFrom *T) string {
	if valueFrom == nil {
		return ""
	}
	return FormatAmount(*valueFrom)
}

func formatIntWithSeparator(integerPart, separator string) string {
	parts := []rune(integerPart)
	result := make([]rune, 0, len(parts)+(len(parts)-1)/3)
	for i, part := range parts {
		if i > 0 && (len(parts)-i)%3 == 0 {
			result = append(result, []rune(separator)...)
		}
		result = append(result, part)
	}
	return string(result)
}

func StringToAmount(valueFrom string) (int, error) {
	v, err := StringToAmountPtr(valueFrom)
	if err != nil {
		return 0, err
	}
	return *v, nil
}

func StringToAmountValidEmpty(valueFrom string) (int, error) {
	valueFrom = strings.TrimSpace(valueFrom)
	if valueFrom == "" {
		return 0, nil
	}
	v, err := StringToAmountPtr(valueFrom)
	if err != nil {
		return 0, err
	}
	return *v, nil
}

func StringToAmountPtr(valueFrom string) (*int, error) {
	valueFrom = strings.TrimSpace(valueFrom)
	if valueFrom == "" {
		return nil, errors.New("empty can't be converted to int")
	}
	valueFrom = strings.Replace(valueFrom, ",", ".", -1)
	valueFloat, err := strconv.ParseFloat(valueFrom, 64)
	if err != nil {
		return nil, err
	}
	valueFloat = valueFloat * 100

	value := int(valueFloat)
	return &value, nil
}

func StringToAmountPtrNoErr(valueFrom string) *int {
	amount, err := StringToAmountPtr(valueFrom)
	if err != nil {
		return nil
	}
	return amount
}
