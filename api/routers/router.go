package routers

import (
	"github.com/gofiber/fiber/v2"
	"tabeldatadotcom/archetype/backend-api/api/handlers"
	"tabeldatadotcom/archetype/backend-api/api/handlers/employees"
)

func SetupRouters(app *fiber.App) {
	apiV1 := app.Group("/api/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})

	// example mapping
	apiV1.Get("/hello", handlers.HelloWorld)

	// employees mapping
	apiV1Employees := apiV1.Group("/employees")
	apiV1Employees.Get("/findBy/:id", employees.FindEmployeeById())
}
