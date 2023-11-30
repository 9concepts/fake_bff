package gateway_test

import (
	"fake_bff/domain"
	"fake_bff/driver"
	"fake_bff/gateway"
	"fake_bff/mock/mock_driver"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserById(t *testing.T) {
	t.Run("should return user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserDriver := mock_driver.NewMockUserDvier(ctrl)
		mockUserDriver.EXPECT().GetUserById("1").Return(&driver.GetUserByIdQuery{
			User: driver.User{
				Id:   "1",
				Name: "Jane Doe",
			},
		}, nil).Times(1)

		expected := &domain.User{Id: "1", Name: "Jane Doe"}

		actual, err := gateway.UserGateway{
			Driver: mockUserDriver,
		}.GetUserById("1")

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
