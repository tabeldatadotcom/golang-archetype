package model

import "time"

type Employee struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Salary    int64     `json:"salary"`
	HireDate  time.Time `json:"hireDate"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
