/*
highload-wallet-api – API wrapper over high-load TON wallet smart contract

Copyright (C) 2021 Alexander Gapak

This file is part of highload-wallet-api.

highload-wallet-api is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

highload-wallet-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with highload-wallet-api.  If not, see <https://www.gnu.org/licenses/>.
*/

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
