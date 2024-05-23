package utils

import (
	"errors"
	"time"
)

const (
	formDateFormat = "2006-01-02"
	dateFormat     = "02-01-2006"
)

func FormatTimePtrToForm(valueFrom *time.Time) string {
	if valueFrom == nil {
		return ""
	}
	return FormatTimeToForm(*valueFrom)
}

func FormatTimeToForm(valueFrom time.Time) string {
	return valueFrom.Format(formDateFormat)
}

func FormatTimePtrToTable(valueFrom *time.Time) string {
	if valueFrom == nil {
		return ""
	}
	return valueFrom.Format(dateFormat)
}

func FormatTimeToTable(valueFrom time.Time) string {
	return valueFrom.Format(dateFormat)
}

func StringToTimePtr(valueFrom string) (*time.Time, error) {
	if valueFrom == "" {
		return nil, errors.New("empty can't be converted to time")
	}

	value, err := time.Parse(formDateFormat, valueFrom)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
func StringToTimePtrNoErr(valueFrom string) *time.Time {
	date, err := StringToTimePtr(valueFrom)
	if err != nil {
		return nil
	}
	return date
}

/*func StringToUintPtr(valueFrom string) (*uint, error) {
	if valueFrom == "" {
		return nil, errors.New("empty can't be converted to uint")
	}
	value, err := strconv.ParseUint(valueFrom, 10, 32)
	if err != nil {
		return nil, err
	}
	result := uint(value)
	return &result, nil
}*/
