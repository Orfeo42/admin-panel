package utils

import (
	"strconv"
	"strings"
	"time"
)

func FormatAmount(amount int) string {
	result := float64(amount) / 100
	return strconv.FormatFloat(result, 'f', 2, 64)
}

func StringToUint(valueFrom string) *uint {
	if valueFrom == "" {
		return nil
	}
	value, err := strconv.ParseUint(valueFrom, 10, 32)
	if err != nil {
		return nil
	}
	result := uint(value)
	return &result
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

const formDateFormat = "2006-01-02"

func FormatTimeToForm(valueFrom *time.Time) string {
	if valueFrom == nil {
		return ""
	}
	return valueFrom.Format(formDateFormat)
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

func FormatUintToFormString(valueFrom *uint) string {
	if valueFrom == nil {
		return ""
	}
	u64 := uint64(*valueFrom)
	return strconv.FormatUint(u64, 10)
}
