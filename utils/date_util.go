package utils

import (
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

func StringToTimePtr(valueFrom string) *time.Time {
	if valueFrom == "" {
		return nil
	}

	value, err := time.Parse(formDateFormat, valueFrom)
	if err != nil {
		return nil
	}

	return &value
}
