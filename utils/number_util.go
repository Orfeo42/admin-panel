package utils

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

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

func StringToInt(valueFrom string) *int {
	if valueFrom == "" {
		return nil
	}
	value, err := strconv.Atoi(strings.TrimSpace(valueFrom))
	if err != nil {
		return nil
	}
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
