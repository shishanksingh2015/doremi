package models

import (
	"time"
)

type Plan struct {
	PlanCategory string
	Type         string
	CreateAt     time.Time
}
type PlanDetail struct {
	Name     string
	Category string
	Amount   int
	Time     int
}

var planDetails = make(map[string]map[string]PlanDetail)

const (
	FREE     string = "FREE"
	PERSONAL string = "PERSONAL"
	PREMIUM  string = "PREMIUM"
	VIDEO    string = "VIDEO"
	MUSIC    string = "MUSIC"
	PODCAST  string = "PODCAST"
)

func init() {
	planDetails[FREE] = make(map[string]PlanDetail)
	planDetails[PERSONAL] = make(map[string]PlanDetail)
	planDetails[PREMIUM] = make(map[string]PlanDetail)

	planDetails[FREE][MUSIC] = PlanDetail{Amount: 0, Time: 1}
	planDetails[FREE][VIDEO] = PlanDetail{Amount: 0, Time: 1}
	planDetails[FREE][PODCAST] = PlanDetail{Amount: 0, Time: 1}

	planDetails[PERSONAL][MUSIC] = PlanDetail{Amount: 100, Time: 1}
	planDetails[PERSONAL][VIDEO] = PlanDetail{Amount: 200, Time: 1}
	planDetails[PERSONAL][PODCAST] = PlanDetail{Amount: 100, Time: 1}

	planDetails[PREMIUM][MUSIC] = PlanDetail{Amount: 250, Time: 3}
	planDetails[PREMIUM][VIDEO] = PlanDetail{Amount: 500, Time: 3}
	planDetails[PREMIUM][PODCAST] = PlanDetail{Amount: 300, Time: 3}
}

func (p *Plan) GetPlanDetails() (PlanDetail, error) {
	category := p.PlanCategory
	planName := p.Type
	if details, ok := planDetails[planName][category]; ok {
		return PlanDetail{planName, category, details.Amount, details.Time}, nil
	}
	return PlanDetail{}, nil
}
