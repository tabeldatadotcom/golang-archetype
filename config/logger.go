package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"os"
)

func SetupLogger(app *fiber.App, file *os.File) {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ==> [${ip}:${port}] ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		Output: file,
		TimeFormat: "Jan-02-2006 15:04:05",
	}))
}
