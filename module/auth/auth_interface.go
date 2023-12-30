package auth

import (
	"go-template-project/dto"
	authEntity "go-template-project/module/auth/entity"
)

type RepositoryInterface interface {
	ForgotPassword(userId uint64, token string) (*authEntity.ForgotPassword, dto.ErrorResponse)
	ResetPassword(input dto.ResetPassword) dto.ErrorResponse
}

type UseCaseInterface interface {
	ForgotPassword(input dto.ForgotPassword) dto.ErrorResponse
	ResetPassword(input dto.ResetPassword) dto.ErrorResponse
	Login(input dto.LoginRequest) (*dto.LoginResponse, dto.ErrorResponse)
}
