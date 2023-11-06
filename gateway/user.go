package gateway

import (
	"fake_bff/domain"
	"fake_bff/driver"
)

type UserGateway struct {
	Driver driver.UserDvier
}

func (g UserGateway) GetUserById(userId string) (*domain.User, error) {
	user, err := g.Driver.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return &domain.User{
		Id:   user.User.Id,
		Name: user.User.Name,
	}, nil
}
