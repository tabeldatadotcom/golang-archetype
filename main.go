package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
	"log"
	"os"
	"tabeldatadotcom/archetype/backend-api/router"
)

func main() {
	// read configuration
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// initialize web application
	app := fiber.New()

	// enable logging file
	file, err := os.OpenFile("./logs/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		Output: file,
	}))

	apiV1 := app.Group("/api/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})

	apiV1.Get("/hello", router.HelloWorld)

	// employees mapping
	apiV1Employees := apiV1.Group("/employees")
	apiV1Employees.Get("/findBy/:id", router.FindEmployeeById)

	appPort := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + appPort))
}
