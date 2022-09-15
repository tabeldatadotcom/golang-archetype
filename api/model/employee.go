package model

import "time"

type EmployeeDtoNew struct {
	FirstName string    `json:"firstName" `
	LastName  string    `json:"lastName"`
	Salary    int64     `json:"salary"`
	HireDate  time.Time `json:"hireDate"`
}

type EmployeeDtoUpdate struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Salary    int64     `json:"salary"`
	HireDate  time.Time `json:"hireDate"`
}
