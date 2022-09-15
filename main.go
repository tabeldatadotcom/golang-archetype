package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"log"
	"os"
	"tabeldatadotcom/archetype/backend-api/api/routers"
	"tabeldatadotcom/archetype/backend-api/config"
	"tabeldatadotcom/archetype/backend-api/pkg/delivery"
)

func main() {
	// read configuration
	config.SetupEnvironment()

	// connect with database
	if err := delivery.SetupDatabase(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// using cors
	app.Use(cors.New())

	// enable logging file
	logLocation := viper.GetString("LOG_FILE_LOCATION")
	log.Printf("writing logfile in: %s", logLocation)
	file, err := os.OpenFile(logLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// setting logger management
	config.SetupLogger(app, file)

	// setting routing url
	routers.SetupRouters(app)

	appPort := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + appPort))
}
