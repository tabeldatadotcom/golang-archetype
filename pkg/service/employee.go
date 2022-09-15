package service

import (
	"tabeldatadotcom/archetype/backend-api/pkg/model"
)

type Service interface {
	InsertEmployee(employee *model.Employee) (*model.Employee, error)
	FindEmployeeById(employeeId string) (*[]model.Employee, error)
	UpdateEmployee(employee *model.Employee) (*model.Employee, error)
	DeleteEmployeeById(employeeId string) error
}

type service struct {
}
