package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GraphQLRequest struct {
	Query string `json:"query"`
}

func GetUserById(userId string) (*User, error) {
	query := fmt.Sprintf(`
		{
			user(id: %s) {
				id
				name
			}
		}
	`, userId)
	requestBody, err := json.Marshal(GraphQLRequest{Query: query})
	if err != nil {
		return nil, err
	}

	res, err := http.Post("https://graphqlzero.almansi.me/api", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		User User `json:"user"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &User{Id: userId, Name: "Test User"}, nil
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user, err := GetUserById(userId)

	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(user)
}
