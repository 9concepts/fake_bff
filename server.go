package main

import (
	"fake_bff/rest"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/systems/ping", rest.Ping)

	app.Listen(":3000")
}
