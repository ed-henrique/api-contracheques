package internal

import (
	"time"
)

type Employee struct {
	Id                         int
	Name                       string
	Surname                    string
	Document                   string
	Sector                     string
	GrossWage                  int
	AdmissionDate              time.Time
	HasHealthcare              bool
	HasDentalcare              bool
	HasTransportationAllowance bool
}

const (
	CONTRIBUTION_LEVEL_1 = 104500
	CONTRIBUTION_LEVEL_2 = 208960
	CONTRIBUTION_LEVEL_3 = 313440
	CONTRIBUTION_LEVEL_4 = 610106

	PERCENTAGE_1 = 75
	PERCENTAGE_2 = 9
	PERCENTAGE_3 = 12
	PERCENTAGE_4 = 14
)

func (e Employee) deductINSS() int {
	if e.GrossWage <= CONTRIBUTION_LEVEL_1 {
		return e.GrossWage * PERCENTAGE_1 / 1000
	}

	deductedValue := (CONTRIBUTION_LEVEL_1 * PERCENTAGE_1) / 1000 // 7837

	if e.GrossWage <= CONTRIBUTION_LEVEL_2 {
		return deductedValue + ((e.GrossWage - CONTRIBUTION_LEVEL_1) * PERCENTAGE_2) / 100
	}

	deductedValue += ((CONTRIBUTION_LEVEL_2 - CONTRIBUTION_LEVEL_1) * PERCENTAGE_2) / 100 // 9401

	if e.GrossWage <= CONTRIBUTION_LEVEL_3 {
		return deductedValue + ((e.GrossWage - CONTRIBUTION_LEVEL_2) * PERCENTAGE_3) / 100
	}

	deductedValue += ((CONTRIBUTION_LEVEL_3 - CONTRIBUTION_LEVEL_2) * PERCENTAGE_3) / 100 // 12537

	if e.GrossWage <= CONTRIBUTION_LEVEL_4 {
		return deductedValue + ((e.GrossWage - CONTRIBUTION_LEVEL_3) * PERCENTAGE_4) / 100
	}

	deductedValue += ((CONTRIBUTION_LEVEL_4 - CONTRIBUTION_LEVEL_3) * PERCENTAGE_4) / 100 // 41533

	return deductedValue
}
