package rest

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Pong string `json:"pong"`
}

func Ping(c *fiber.Ctx) error {
	return c.JSON(Response{Pong: "ok"})
}
