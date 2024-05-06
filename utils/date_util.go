package utils

import (
	"github.com/labstack/gommon/log"
	"time"
)

func ParseDate(date string) *time.Time {
	if date == "" {
		return nil
	}
	parsedDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Info("Date not parsable:", date)
		return nil
	}
	return &parsedDate
}
