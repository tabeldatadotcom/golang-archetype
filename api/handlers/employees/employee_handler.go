package employees

import (
	"github.com/gofiber/fiber/v2"
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
