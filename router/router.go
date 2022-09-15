package router

import (
	"github.com/gofiber/fiber/v2"
	"tabeldatadotcom/archetype/backend-api/handler"
	"tabeldatadotcom/archetype/backend-api/handler/employees"
)

func SetupRouters(app *fiber.App) {
	apiV1 := app.Group("/api/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})

	// example mapping
	apiV1.Get("/hello", handler.HelloWorld)

	// employees mapping
	apiV1Employees := apiV1.Group("/employees")
	apiV1Employees.Get("/findBy/:id", employees.FindEmployeeById)
}
