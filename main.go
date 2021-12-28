package main

import (
	"highload-wallet-api/src/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Configure()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Listen(config.Cfg.Server.Host + ":" + config.Cfg.Server.Port)
}
