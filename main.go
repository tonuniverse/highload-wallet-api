package main

import (
	"highload-wallet-api/src/api"
	"highload-wallet-api/src/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Configure()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${path} ${method} ${status}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
	}))

	app.Post("/transfer", api.Transfer)

	app.Listen(config.Cfg.Server.Host + ":" + config.Cfg.Server.Port)
}
