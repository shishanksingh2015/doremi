package utils

import (
	"time"
)

func GetDate(dateStr string, months int) (string, error) {
	reminderTenDaysBefore := -10
	layout := "02-01-2006"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return "", err
	}
	newDate := date.AddDate(0, months, 0)
	newDate = newDate.AddDate(0, 0, reminderTenDaysBefore)
	return newDate.Format(layout), nil
}
