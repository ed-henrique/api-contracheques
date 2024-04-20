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
	INSS_CONTRIBUTION_LEVEL_1 = 104500
	INSS_CONTRIBUTION_LEVEL_2 = 208960
	INSS_CONTRIBUTION_LEVEL_3 = 313440
	INSS_CONTRIBUTION_LEVEL_4 = 610106

	INSS_PERCENTAGE_1 = 75
	INSS_PERCENTAGE_2 = 9
	INSS_PERCENTAGE_3 = 12
	INSS_PERCENTAGE_4 = 14

	IRPF_LEVEL_1 = 190398
	IRPF_LEVEL_2 = 282665
	IRPF_LEVEL_3 = 375105
	IRPF_LEVEL_4 = 466468

	IRPF_DEDUCTION_1 = 14280
	IRPF_DEDUCTION_2 = 35480
	IRPF_DEDUCTION_3 = 63613
	IRPF_DEDUCTION_4 = 86936

	IRPF_PERCENTAGE_1 = 75
	IRPF_PERCENTAGE_2 = 15
	IRPF_PERCENTAGE_3 = 225
	IRPF_PERCENTAGE_4 = 275

	HEALTHCARE_DEDUCTION = 1000
)

func (e Employee) deductINSS() int {
	if e.GrossWage <= INSS_CONTRIBUTION_LEVEL_1 {
		return e.GrossWage * INSS_PERCENTAGE_1 / 1000
	}

	deductedValue := (INSS_CONTRIBUTION_LEVEL_1 * INSS_PERCENTAGE_1) / 1000 // 7837

	if e.GrossWage <= INSS_CONTRIBUTION_LEVEL_2 {
		return deductedValue + ((e.GrossWage - INSS_CONTRIBUTION_LEVEL_1) * INSS_PERCENTAGE_2) / 100
	}

	deductedValue += ((INSS_CONTRIBUTION_LEVEL_2 - INSS_CONTRIBUTION_LEVEL_1) * INSS_PERCENTAGE_2) / 100 // 9401

	if e.GrossWage <= INSS_CONTRIBUTION_LEVEL_3 {
		return deductedValue + ((e.GrossWage - INSS_CONTRIBUTION_LEVEL_2) * INSS_PERCENTAGE_3) / 100
	}

	deductedValue += ((INSS_CONTRIBUTION_LEVEL_3 - INSS_CONTRIBUTION_LEVEL_2) * INSS_PERCENTAGE_3) / 100 // 12537

	if e.GrossWage <= INSS_CONTRIBUTION_LEVEL_4 {
		return deductedValue + ((e.GrossWage - INSS_CONTRIBUTION_LEVEL_3) * INSS_PERCENTAGE_4) / 100
	}

	deductedValue += ((INSS_CONTRIBUTION_LEVEL_4 - INSS_CONTRIBUTION_LEVEL_3) * INSS_PERCENTAGE_4) / 100 // 41533

	return deductedValue
}

func (e Employee) deductIRPF() int {
	switch {
	case e.GrossWage > IRPF_LEVEL_4:
		return (e.GrossWage * IRPF_PERCENTAGE_4) / 1000 - IRPF_DEDUCTION_4
	case e.GrossWage > IRPF_LEVEL_3:
		return (e.GrossWage * IRPF_PERCENTAGE_3) / 1000 - IRPF_DEDUCTION_3
	case e.GrossWage > IRPF_LEVEL_2:
		return (e.GrossWage * IRPF_PERCENTAGE_2) / 100 - IRPF_DEDUCTION_2
	case e.GrossWage > IRPF_LEVEL_1:
		return (e.GrossWage * IRPF_PERCENTAGE_1) / 1000 - IRPF_DEDUCTION_1
	default:
		return 0
	}
}

func (e Employee) deductHealthcare() int {
	return HEALTHCARE_DEDUCTION
}
