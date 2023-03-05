package handler

import (
	"doremi/messages"
	"doremi/models"
	"errors"
	"fmt"
)

type PlanHandler interface {
	AddPlan(planCategory string, planName string) error
	CheckIfPlanExist(planCategory string) bool
}

func (u *User) AddPlan(planCategory string, planName string) error {
	if u.SubsDate == "" {
		fmt.Println(messages.ADD_SUBSCRIPTION_FAILED, messages.INVALID_DATE)
		return errors.New(messages.INVALID_DATE)
	}
	if u.Plan != nil && len(u.Plan) > 0 {
		exists := u.CheckIfPlanExist(planCategory)
		if exists {
			fmt.Println(messages.ADD_SUBSCRIPTION_FAILED, messages.DUPLICATE_CATEGORY)
			return errors.New(messages.ADD_SUBSCRIPTION_FAILED)
		}
	}
	plan := models.Plan{PlanCategory: planCategory, Type: planName}
	planDetail, err := plan.GetPlanDetails()
	if err != nil {
		return err
	}
	u.Plan = append(u.Plan, planDetail)
	return nil
}

func (u *User) CheckIfPlanExist(planCategory string) bool {
	for _, v := range u.Plan {
		if v.Category == planCategory {
			return true
		}
	}
	return false
}
