package service

import "tabeldatadotcom/archetype/backend-api/pkg/model"

type EmployeeService interface {
	SaveAnEmployee(employee *model.Employee) (*model.Employee, error)
	FindAnEmployeeById(employeeId string) (*model.Employee, error)
	UpdateEmployee(value *model.Employee) (*model.Employee, error)
	DeleteEmployeeById(value string) error
}
