package usecase_test

import (
	"fake_bff/domain"
	"fake_bff/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserByIdSuccessfully(t *testing.T) {
	expected := domain.User{Id: "1", Name: "Jane Doe"}
	mockUserGateway := new(MockUserPort)
	mockUserGateway.On("GetUserById", "1").Return(&expected, nil)

	actual, err := usecase.GetUserById("1", mockUserGateway)

	assert.NoError(t, err)
	assert.Equal(t, &expected, actual)
	mockUserGateway.AssertExpectations(t)
}

type MockUserPort struct {
	mock.Mock
}

func (m *MockUserPort) GetUserById(userId string) (*domain.User, error) {
	args := m.Called(userId)
	return args.Get(0).(*domain.User), args.Error(1)
}
