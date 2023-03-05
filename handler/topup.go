package handler

import (
	"doremi/messages"
	"doremi/models"
	"errors"
	"fmt"
	"strconv"
)

type TopUpHandler interface {
	AddTopUp(topUpName string, noOfMonths string) error
}

func (u *User) AddTopUp(topUpName string, noOfMonths string) error {
	if u.SubsDate == "" {
		fmt.Println(messages.ADD_TOPUP_FAILED, messages.INVALID_DATE)
		return errors.New(messages.ADD_TOPUP_FAILED)
	}
	if u.TopUp == nil {
		u.TopUp = &models.TopUp{}
		topUpDetails, err := u.TopUp.GetTopUpDetails(topUpName)
		if err != nil {
			return err
		}
		month, _ := strconv.Atoi(noOfMonths)
		u.TopUp.Amount = month * topUpDetails.Amount
		u.TopUp.Type = topUpName
		return nil
	}
	fmt.Println(messages.ADD_TOPUP_FAILED, messages.DUPLICATE_TOPUP)
	return errors.New(messages.ADD_TOPUP_FAILED)
}
