package internal

import "time"

type Employee struct {
	Id                         int
	Name                       string
	Surname                    string
	Document                   string
	Sector                     string
	GrossWage                  float64
	AdmissionDate              time.Time
	HasHealthcare              bool
	HasDentalcare              bool
	HasTransportationAllowance bool
}
