package model

import "time"

type EmployeeDtoNew struct {
	FirstName     string    `json:"firstName" `
	LastName      string    `json:"lastName"`
	Salary        float64     `json:"salary"`
	HireDate      time.Time `json:"hireDate"`
	CommissionPct float32   `json:"commissionPct"`
}

type EmployeeDtoUpdate struct {
	ID        string    `json:"id"`
	FirstName     string    `json:"firstName" `
	LastName      string    `json:"lastName"`
	Salary        float64     `json:"salary"`
	HireDate      time.Time `json:"hireDate"`
	CommissionPct float32   `json:"commissionPct"`
}
