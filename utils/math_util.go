package utils

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
)

func formatIntegerWithCustomSeparator(integerPart, separator string) string {
	parts := []rune(integerPart)
	result := make([]rune, 0, len(parts)+(len(parts)-1)/3)
	for i, part := range parts {
		if i > 0 && (len(parts)-i)%3 == 0 {
			result = append(result, []rune(separator)...)
		}
		result = append(result, part)
	}

	// Ritorna la stringa formattata
	return string(result)
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

func IntToString(number int) string {
	result := float64(number) / 100
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

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
