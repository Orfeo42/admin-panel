package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

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

func FloatToString(amount float64) string {
	return strconv.FormatFloat(amount, 'f', -1, 64)
}

func FormatAmount(amount int) string {
	result := float64(amount) / 100
	return strconv.FormatFloat(result, 'f', 2, 64)
}

func StringToUint(valueFrom string) (*uint, error) {
	if valueFrom == "" {
		return nil, errors.New("empty can't be converted to uint")
	}
	value, err := strconv.ParseUint(valueFrom, 10, 32)
	if err != nil {
		return nil, err
	}
	result := uint(value)
	return &result, nil
}

func StringToTime(valueFrom string) *time.Time {
	if valueFrom == "" {
		return nil
	}

	value, err := time.Parse("2006-01-02", valueFrom)

	if err != nil {
		return nil
	}

	return &value
}

func StringToInt(valueFrom string) *int {
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

func StringToString(valueFrom string) *string {
	if strings.TrimSpace(valueFrom) == "" {
		return nil
	}
	return &valueFrom
}

func FormatStringToForm(valueFrom *string) string {
	if valueFrom == nil {
		return ""
	}
	return strings.TrimSpace(*valueFrom)
}

func FormatIntToFormNumber(valueFrom *int) string {
	if valueFrom == nil {
		return ""
	}
	return FormatAmount(*valueFrom)
}

func IntToString(amount int) string {
	result := float64(amount) / 100
	return strconv.FormatFloat(result, 'f', 2, 64)
}

func FormatUintToFormString(valueFrom *uint) string {
	if valueFrom == nil {
		return ""
	}
	return UintToString(*valueFrom)
}

func UintToString(valueFrom uint) string {
	u64 := uint64(valueFrom)
	return strconv.FormatUint(u64, 10)
}
