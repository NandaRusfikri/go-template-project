package repository

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
	authentity "go-template-project/module/auth/entity"
)

type AuthRepositoryMock struct {
	Mock mock.Mock
}

func (r *AuthRepositoryMock) ResetPassword(input dto.ResetPassword) dto.ErrorResponse {
	args := r.Mock.Called(input)
	if args.Get(0) == nil {
		return args.Get(0).(dto.ErrorResponse)
	} else {
		return dto.ErrorResponse{}
	}
}

func (r *AuthRepositoryMock) ForgotPassword(userId uint64, token string) (*authentity.ForgotPassword, dto.ErrorResponse) {
	args := r.Mock.Called(userId, token)
	if args.Get(0) != nil {
		return args.Get(0).(*authentity.ForgotPassword), args.Get(1).(dto.ErrorResponse)
	}
	return nil, dto.ErrorResponse{Code: 500, Error: errors.New("unexpected error")}
}
