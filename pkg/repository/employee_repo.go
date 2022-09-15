package repository

import (
	"database/sql"
	"log"
	"tabeldatadotcom/archetype/backend-api/pkg/model"
	"time"
)

type EmployeesRepository interface {
	InsertEmployee(employee *model.Employee) (*model.Employee, error)
	FindEmployeeById(employeeId string) (*model.Employee, error)
	UpdateEmployee(employee *model.Employee) (*model.Employee, error)
	DeleteEmployeeById(employeeId string) error
}

func NewRepo(db *sql.DB) EmployeesRepository {
	return &repository{
		Database: db,
	}
}

func (r repository) InsertEmployee(value *model.Employee) (*model.Employee, error) {
	const query = "INSERT INTO employees(first_name, last_name, salary, commission_pct, created_by) values ($1, $2, $3, $3, $4) returning id"
	res, err := r.Database.Query(query, value.FirstName, value.LastName, value.Salary, value.CommissionPct, time.Now())
	if err != nil {
		return nil, err
	}
	defer res.Close()

	log.Println(res)
	if res.Next() {
		if err := res.Scan(&value.ID); err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (r repository) FindEmployeeById(id string) (*model.Employee, error) {
	const query = "select id,\n       first_name,\n       last_name,\n       salary,\n       commission_pct,\n       hire_date\nfrom employees\nwhere id = $1"
	aRow := r.Database.QueryRow(query, id)

	value  := model.Employee{}
	err := aRow.Scan(&value.ID, &value.FirstName, &value.LastName, &value.Salary, &value.CommissionPct, &value.HireDate)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (r repository) UpdateEmployee(employee *model.Employee) (*model.Employee, error) {
	return nil, nil
}

func (r repository) DeleteEmployeeById(employeeId string) error {
	return nil
}
