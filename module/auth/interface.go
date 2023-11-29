package auth

import (
	"go-template-project/dto"
	auth_entity "go-template-project/module/auth/entity"
)

type Repository interface {
	ForgotPassword(user_id uint64, token string) (*auth_entity.ForgotPassword, dto.ResponseError)
	ResetPassword(input dto.ResetPassword) dto.ResponseError
}

type UseCase interface {
	ForgotPassword(input dto.ForgotPassword) dto.ResponseError
	ResetPassword(input dto.ResetPassword) dto.ResponseError
	Login(input dto.LoginRequest) (*dto.LoginResponse, dto.ResponseError)
}
