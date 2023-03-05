package handler

import (
	"geektrust/models"
	"log"
	"testing"
)

func TestUser_StartSubscription(t *testing.T) {
	user := User{}
	// test with a valid date
	validDate := "01-03-2023"
	invalidDate := "13-19-2022"
	invalidDateFormat := "01/12/2022"
	emptyDate := ""
	// valid date
	err := user.StartSubscription(validDate)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if user.SubsDate != validDate {
		t.Errorf("expected subscription date %s,got %s", validDate, user.SubsDate)
	}

	//Invalid date
	err = user.StartSubscription(invalidDate)
	if err == nil {
		t.Errorf("It should give error for parsing error")
	}
	log.Printf("Error is %v", err.Error())
	//InvalidDateFormat
	err = user.StartSubscription(invalidDateFormat)
	if err == nil {
		t.Errorf("It should give error for invalid DateFormat")
	}
	log.Printf("Error is %v", err.Error())

	err = user.StartSubscription(emptyDate)
	if err == nil {
		t.Errorf("It should give error for empty date")
	}
	log.Printf("Error is %v", err.Error())
}

func TestUser_GetRenewalDate(t *testing.T) {
	user := User{Id: 1}
	user.SubsDate = "01-03-2023"
	user.Plan = []models.PlanDetail{{
		Name:     models.PERSONAL,
		Category: models.MUSIC,
		Amount:   100,
		Time:     1,
	}}
	// get details without topup
	err := user.GetRenewalDetails()
	if err != nil {
		t.Errorf("Failed to get renewal Details %v", err)
	}
	// get details with topup
	user.TopUp = &models.TopUp{
		Type:   models.FOUR_DEVICE,
		Amount: 50,
		Time:   1,
	}
	err = user.GetRenewalDetails()
	if err != nil {
		t.Errorf("Failed to get renewal Details %v", err)
	}
	// get detail with no plans added
	user.Plan = nil
	err = user.GetRenewalDetails()
	if err == nil {
		t.Errorf("It should give us error")
	}
}
