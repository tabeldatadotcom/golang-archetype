package repository

import (
	"github.com/jmoiron/sqlx"
	"tabeldatadotcom/archetype/backend-api/pkg/model"
	"time"
)

type EmployeesRepository interface {
	InsertEmployee(employee *model.Employee) (*model.Employee, error)
	FindEmployeeById(employeeId string) (*model.Employee, error)
	UpdateEmployee(employee *model.Employee) (*model.Employee, error)
	DeleteEmployeeById(employeeId string) error
}

func NewRepo(db *sqlx.DB) EmployeesRepository {
	return &repository{
		Connect: db,
	}
}

func (r repository) InsertEmployee(value *model.Employee) (*model.Employee, error) {
	const query = `INSERT INTO employees(first_name, last_name, salary, commission_pct, created_by)
values ($1, $2, $3, $4, $5)
returning id`
	tx, _ := r.Connect.Begin()
	preparedQuery, _ := tx.Prepare(query)
	defer preparedQuery.Close()
	returning := preparedQuery.QueryRow(value.FirstName, value.LastName, value.Salary, value.CommissionPct, time.Now())

	if err := returning.Scan(&value.ID); err != nil {
		return nil, err
	}
	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, errCommit
	}
	return value, nil
}

func (r repository) FindEmployeeById(id string) (*model.Employee, error) {
	const query = `select id,
       first_name,
       last_name,
       salary,
       commission_pct,
       hire_date
from employees
where id = $1`
	aRow := r.Connect.QueryRow(query, id)

	value := model.Employee{}
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
