package employees

import (
	"github.com/gofiber/fiber/v2"
	dto "tabeldatadotcom/archetype/backend-api/api/model"
	"tabeldatadotcom/archetype/backend-api/pkg/model"
	"tabeldatadotcom/archetype/backend-api/pkg/repository"
)

func FindEmployeeById(repo repository.EmployeesRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		employee, err := repo.FindEmployeeById(id)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(employee)
	}
}

func CreateNewEmployee(repo repository.EmployeesRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		emp := new(dto.EmployeeDtoNew)
		if err := c.BodyParser(emp); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		value := model.Employee{
			FirstName:     emp.FirstName,
			LastName:      emp.LastName,
			Salary:        emp.Salary,
			CommissionPct: emp.CommissionPct,
			HireDate:      emp.HireDate,
		}
		employee, err := repo.InsertEmployee(&value)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(employee)
	}
}
