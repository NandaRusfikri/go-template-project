package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
	"go-template-project/module/auth/repository"
	"testing"
)

var authRepository = &repository.AuthRepositoryMock{
	Mock: mock.Mock{},
}
var authUseCase = authUseCase{
	authRepo: authRepository,
}

func TestResetPassword(t *testing.T) {
	authRepository.Mock.On("ResetPassword", dto.ResetPassword{}).Return(nil)
	errResp := authUseCase.ResetPassword(dto.ResetPassword{Email: "nanda@gmail.com", NewPassword: "admin"})
	assert.Equal(t, dto.ResponseError{}, errResp)
}
