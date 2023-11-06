package usecase_test

import (
	"fake_bff/domain"
	"fake_bff/mock/mock_port"
	"fake_bff/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserById(t *testing.T) {
	t.Run("should return user", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockUserPort := mock_port.NewMockUserPort(ctrl)
		mockUserPort.EXPECT().GetUserById("1").Return(&domain.User{Id: "1", Name: "Jane Doe"}, nil).Times(1)

		expected := domain.User{Id: "1", Name: "Jane Doe"}

		actual, err := usecase.GetUserById("1", mockUserPort)

		assert.NoError(t, err)
		assert.Equal(t, &expected, actual)
	})
}

func NewMockUserPort(ctrl *gomock.Controller) {
	panic("unimplemented")
}
