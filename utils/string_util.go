package utils

import "strings"

func ParseString(valueFrom string) *string {
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
