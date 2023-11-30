package rest

import (
	"fake_bff/driver"
	"fake_bff/gateway"
	"fake_bff/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetUser(c *fiber.Ctx) error {
	userId := utils.CopyString(c.Params("userId"))
	userGateway := gateway.UserGateway{
		Driver: driver.UserDvierImpl{},
	}

	user, err := usecase.GetUserById(userId, userGateway)

	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(user)
}
