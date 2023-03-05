package handler

import (
	"doremi/messages"
	"doremi/models"
	"doremi/utils"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Id       int
	Plan     []models.PlanDetail
	TopUp    *models.TopUp
	SubsDate string
}

type SubscriptionHandler interface {
	StartSubscription(date string) error
	GetRenewalDetails() error
	GetRenewalAmount() error
}

func (u *User) StartSubscription(date string) error {
	layout := "02-01-2006"
	_, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(messages.INVALID_DATE)
		return err
	}
	u.SubsDate = date
	return nil
}

func (u *User) GetRenewalDetails() error {
	if u.SubsDate != "" && u.Plan != nil {
		for _, value := range u.Plan {
			date, err := utils.GetDate(u.SubsDate, value.Time)
			if err != nil {
				return err
			}
			fmt.Println(messages.RENEWAL_REMINDER, value.Category+" "+date)
		}
		return nil
	}
	fmt.Println(messages.SUBSCRIPTIONS_NOT_FOUND)
	return errors.New(messages.SUBSCRIPTIONS_NOT_FOUND)
}

func (u *User) GetRenewalAmount() error {
	totalRenewableAmount := 0
	if u.Plan != nil {
		for _, val := range u.Plan {
			totalRenewableAmount += val.Amount
		}
		if u.TopUp != nil {
			totalRenewableAmount += u.TopUp.Amount
		}
		fmt.Println(messages.RENEWAL_AMOUNT, totalRenewableAmount)
		return nil
	}
	return errors.New(messages.NOT_AVAILABLE)
}
