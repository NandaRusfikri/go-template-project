package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
	"go-template-project/module/auth/repository"
	userrepo "go-template-project/module/user/repository"
	"go-template-project/pkg"
	"testing"
)

var authRepository = &repository.AuthRepositoryMock{
	Mock: mock.Mock{},
}
var userRepository = &userrepo.UserRepositoryMock{
	Mock: mock.Mock{},
}

var mockSMTP = new(pkg.SMTP)
var authUsecase = InitAuthUseCase(authRepository, userRepository, mockSMTP)

func TestResetPassword(t *testing.T) {

	testInput := dto.ResetPassword{
		Token:       "testToken",
		NewPassword: "newPassword",
		Email:       "test@example.com",
	}

	authRepository.Mock.On("ResetPassword", testInput).Return(dto.ErrorResponse{})

	actualError := authUsecase.ResetPassword(testInput)

	authRepository.Mock.AssertExpectations(t)
	assert.Equal(t, dto.ErrorResponse{}, actualError)

}
