package utils

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func StringToUintPtr(valueFrom string) (*uint, error) {
	value, err := StringToUint(valueFrom)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func StringToUint(valueFrom string) (uint, error) {
	if valueFrom == "" {
		return 0, errors.New("empty can't be converted to uint")
	}
	value, err := strconv.ParseUint(valueFrom, 10, 32)
	if err != nil {
		return 0, err
	}
	result := uint(value)
	return result, nil
}

func StringToInt(valueFrom string) (*int, error) {
	if valueFrom == "" {
		return nil, errors.New("empty can't be converted to int")
	}
	value, err := strconv.Atoi(strings.TrimSpace(valueFrom))
	if err != nil {
		return nil, err
	}
	return &value, nil
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
