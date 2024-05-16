package utils

import (
	"testing"
)

type stringToIntTestCase struct {
	input string
	want  int
}

type float64ToStringTestCase struct {
	input float64
	want  string
}

type intToStringTestCase struct {
	input int
	want  string
}

type stringToStringTestCase struct {
	input string
	want  string
}
type twoStringToStringTestCase struct {
	input1 string
	input2 string
	want   string
}

func TestParseAmount(t *testing.T) {
	testCases := []stringToIntTestCase{
		{input: "10", want: 1000},
		{input: "10,00", want: 1000},
		{input: "12", want: 1200},
		{input: "12,00", want: 1200},
		{input: "12,99", want: 1299},
		{input: "10002,3594", want: 1000235},
		{input: "1359782", want: 135978200},
		{input: "987456321", want: 98745632100},
		{input: "pippo", want: 0},
		{input: "", want: 0},
	}

	for _, item := range testCases {
		got := ParseAmount(item.input)
		if got != item.want {
			Fail(t, "input %s, got %d, wanted %d", item.input, got, item.want)
		} else {
			Pass("input %s, got %d, wanted %d", item.input, got, item.want)
		}
	}
}

func TestFormatFloat(t *testing.T) {

	testCases := []float64ToStringTestCase{
		{input: 1000, want: "1000.00"},
		{input: 10, want: "10.00"},
		{input: 0, want: "0.00"},
		{input: 0.5974, want: "0.60"},
	}

	for _, item := range testCases {
		got := formatFloat(item.input)
		if got != item.want {
			Fail(t, "input '%f', got '%s', wanted '%s'", item.input, got, item.want)
		} else {
			Pass("input '%f', got '%s', wanted '%s'", item.input, got, item.want)
		}
	}
}

func TestFormatAmount(t *testing.T) {

	testCases := []intToStringTestCase{
		{input: 1000, want: "10.00"},
		{input: 10, want: "0.10"},
		{input: 0, want: "0.00"},
		{input: 100, want: "1.00"},
		{input: 10000, want: "100.00"},
		{input: 29875, want: "298.75"},
		{input: 99999999, want: "999999.99"},
	}

	for _, item := range testCases {
		got := FormatAmount(item.input)
		if got != item.want {
			Fail(t, "input '%d', got '%s', wanted '%s'", item.input, got, item.want)
		} else {
			Pass("input '%d', got '%s', wanted '%s'", item.input, got, item.want)
		}
	}
}

func TestFormatIntWithSeparator(t *testing.T) {

	testCases := []twoStringToStringTestCase{
		{input1: "1000", input2: ".", want: "1.000"},
		{input1: "10", input2: "˙", want: "10"},
		{input1: "10000", input2: "...", want: "10...000"},
		{input1: "100", input2: "˙", want: "100"},
		{input1: "10000", input2: "˙", want: "10˙000"},
		{input1: "29875", input2: "˙", want: "29˙875"},
	}

	for _, item := range testCases {
		got := formatIntWithSeparator(item.input1, item.input2)
		if got != item.want {
			Fail(t, "input1 '%s', input2 '%s', got '%s', wanted '%s'", item.input1, item.input2, got, item.want)
		} else {
			Pass("input1 '%s', input2 '%s', got '%s', wanted '%s'", item.input1, item.input2, got, item.want)
		}
	}
}

func TestStringAmountToString(t *testing.T) {

	testCases := []stringToStringTestCase{
		{input: "1000", want: "1˙000"},
		{input: "10", want: "10"},
		{input: "0", want: "0"},
		{input: "100", want: "100"},
		{input: "10000", want: "10˙000"},
		{input: "29875", want: "29˙875"},
		{input: "99999999", want: "99˙999˙999"},
	}

	for _, item := range testCases {
		got := stringAmountToString(item.input)
		if got != item.want {
			Fail(t, "input '%s', got '%s', wanted '%s'", item.input, got, item.want)
		} else {
			Pass("input '%s', got '%s', wanted '%s'", item.input, got, item.want)
		}
	}
}

/*
func TestAmountToString(number float64) string {
	numberAsString := formatFloat(number)
	return stringAmountToString(numberAsString)
}

func TestAmountIntegerToString(number int) string {
	numberAsString := FormatAmount(number)
	return stringAmountToString(numberAsString)
}

func TestFormatIntToAmount(valueFrom *int) string {
	if valueFrom == nil {
		return ""
	}
	return FormatAmount(*valueFrom)
}

func TestStringToAmount(valueFrom string) *int {
	if valueFrom == "" {
		return nil
	}
	valueFloat, err := strconv.ParseFloat(strings.TrimSpace(valueFrom), 64)
	if err != nil {
		return nil
	}
	valueFloat = valueFloat * 100

	value := int(valueFloat)
	return &value
}
*/
