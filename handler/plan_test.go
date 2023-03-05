package handler

import (
	"doremi/models"
	"testing"
)

func TestUser_AddPlan(t *testing.T) {
	user := User{Id: 1}
	user.SubsDate = "01-03-2023"
	planFree := models.FREE
	planPersonal := models.PERSONAL
	categoryMusic := models.MUSIC
	categoryVideo := models.VIDEO

	err := user.AddPlan(categoryMusic, planFree)
	if err != nil {
		t.Errorf("Failed to add plan %v", err)
	}
	// Check that the plan was added
	if len(user.Plan) != 1 {
		t.Errorf("Expected user to have 1 plan, but got %d", len(user.Plan))
	}
	// test adding a plan with duplicate category
	err = user.AddPlan(categoryMusic, planPersonal)
	if err == nil {
		t.Error("Expected AddPlan to return an error for duplicate category, but it didn't")
	}
	// test adding a plan without date
	user.SubsDate = ""
	err = user.AddPlan(categoryVideo, planFree)
	if err == nil {
		t.Error("Expected AddPlan to return an error for invalid date")
	}
}
