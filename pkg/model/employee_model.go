package model

import "time"

type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Salary    float64
	CommissionPct float32
	HireDate  time.Time
}
