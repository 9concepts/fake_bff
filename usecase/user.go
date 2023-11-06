package usecase

import (
	"fake_bff/domain"
	"fake_bff/port"
)

func GetUserById(userId string, userGateway port.UserPort) (*domain.User, error) {
	return userGateway.GetUserById(userId)
}
