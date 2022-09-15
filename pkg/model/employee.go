package model

import "time"

type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Salary    int64
	HireDate  time.Time
}
