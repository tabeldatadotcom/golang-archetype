package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"tabeldatadotcom/archetype/backend-api/config"
	"tabeldatadotcom/archetype/backend-api/router"
)

func main() {
	app := fiber.New()

	// read configuration
	config.SetupEnvironment()

	// using cors
	app.Use(cors.New())

	// enable logging file
	file, err := os.OpenFile("./logs/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// setting logger management
	config.SetupLogger(app, file)

	// setting routing url
	router.SetupRouters(app)

	appPort := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + appPort))
}
