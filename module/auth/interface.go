package auth

import (
	"go-template-project/dto"
	auth_entity "go-template-project/module/auth/entity"
)

type AuthRepository interface {
	RequestForgotPassword(user_id uint64, token string) (*auth_entity.EntityForgotPassword, dto.ResponseError)
	ResetPassword(input dto.ResetPassword) dto.ResponseError
}

type AuthUseCase interface {
	RequestForgotPassword(input dto.ForgotPassword) dto.ResponseError
	ResetPassword(input dto.ResetPassword) dto.ResponseError
	Login(input dto.LoginRequest) (*dto.LoginResponse, dto.ResponseError)
}
