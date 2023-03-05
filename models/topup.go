package models

import (
	"doremi/messages"
	"errors"
	"time"
)

type TopUp struct {
	Type      string
	Amount    int
	Time      int
	CreatedAt time.Time
}
type Detail struct {
	Amount int
	Time   int
}

const (
	FOUR_DEVICE string = "FOUR_DEVICE"
	TEN_DEVICE  string = "TEN_DEVICE"
)

var topUpPlan = make(map[string]Detail)

func init() {
	topUpPlan[FOUR_DEVICE] = Detail{Amount: 50, Time: 1}
	topUpPlan[TEN_DEVICE] = Detail{Amount: 100, Time: 1}
}

func (tp *TopUp) GetTopUpDetails(topUpType string) (Detail, error) {
	if details, ok := topUpPlan[topUpType]; ok {
		return Detail{details.Amount, details.Time}, nil
	}
	return Detail{}, errors.New(messages.NO_TOP_UP_AVAILABLE)
}
