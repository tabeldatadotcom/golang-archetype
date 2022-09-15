package model

import "time"

type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Salary    int64
	CommissionPct float32
	HireDate  time.Time
}
