package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
	"log"
	"os"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app_port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + app_port))
}
