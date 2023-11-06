package driver

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type UserDvier interface {
	GetUserById(userId string) (*GetUserByIdQuery, error)
}

type UserDvierImpl struct{}

type User struct {
	Id   string
	Name string
}

type GetUserByIdQuery struct {
	User User `graphql:"user(id: $id)"`
}

func (_ UserDvierImpl) GetUserById(userId string) (*GetUserByIdQuery, error) {
	client := graphql.NewClient("https://graphqlzero.almansi.me/api", nil)

	var query GetUserByIdQuery

	variables := map[string]interface{}{
		"id": graphql.ID(userId),
	}

	err := client.Query(context.Background(), &query, variables)

	if err != nil {
		return nil, err
	}

	return &query, nil
}
