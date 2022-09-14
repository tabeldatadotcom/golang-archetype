package router

import (
	"github.com/gofiber/fiber/v2"
	"tabeldatadotcom/archetype/backend-api/model"
	"time"
)

func FindEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	dimasMaryanto := model.Employee{
		ID:        id,
		FirstName: "Dimas",
		LastName:  "Maryanto",
		Salary:    10000,
		HireDate:  time.Date(2015, time.August, 25, 0, 0, 0, 0, time.Local),
	}
	return c.Status(fiber.StatusOK).JSON(dimasMaryanto)
}
