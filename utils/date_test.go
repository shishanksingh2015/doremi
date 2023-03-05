package utils

import (
	"testing"
)

func TestGetDate(t *testing.T) {
	date := "05-02-2022"
	noOfMonth := 1
	expectedDate := "23-02-2022"
	result, err := GetDate(date, noOfMonth)
	if err != nil {
		t.Errorf("No error was expected %v", err)
	}
	if result != expectedDate {
		t.Errorf("Expecting date to be %v but received %v", expectedDate, result)
	}
	invalidDate := "01-41-2023"
	_, err = GetDate(invalidDate, noOfMonth)
	if err == nil {
		t.Errorf("There should be error for wrong date format")
	}
}
