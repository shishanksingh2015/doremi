package handler

import (
	"doremi/models"
	"testing"
)

func TestUser_AddTopUp(t *testing.T) {
	user := User{Id: 1}
	user.SubsDate = "01-03-2023"
	planFree := models.FREE
	categoryMusic := models.MUSIC
	err := user.AddPlan(categoryMusic, planFree)
	if err != nil {
		t.Errorf("Failed to add plan %v", err)
	}
	topUpFour := models.FOUR_DEVICE
	topUpTen := models.TEN_DEVICE

	// add topup without date
	user.SubsDate = ""
	err = user.AddTopUp(topUpTen, "2")
	if err == nil {
		t.Errorf("It should give us error")
	}

	user.SubsDate = "01-03-2023"
	err = user.AddTopUp(topUpFour, "2")
	if err != nil {
		t.Errorf("Failed to add topUp %v", err)
	}

	// add another topup
	err = user.AddTopUp(topUpTen, "2")
	if err == nil {
		t.Errorf("It should give us error")
	}

}
