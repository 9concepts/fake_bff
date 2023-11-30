package port

import "fake_bff/domain"

type UserPort interface {
	GetUserById(userId string) (*domain.User, error)
}
