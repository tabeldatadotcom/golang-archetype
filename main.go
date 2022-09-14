package main

import (
	"github.com/gofiber/fiber/v2"
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
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app_port := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + app_port))
}
