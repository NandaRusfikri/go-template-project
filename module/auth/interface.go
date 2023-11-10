package auth

import (
	auth_entity "go-template-project/module/auth/entity"
	"go-template-project/schemas"
)

type AuthRepository interface {
	RequestForgotPassword(user_id uint64, token string) (*auth_entity.EntityForgotPassword, schemas.ResponseError)
	ResetPassword(input schemas.ResetPassword) schemas.ResponseError
}

type AuthUseCase interface {
	RequestForgotPassword(input schemas.ForgotPassword) schemas.ResponseError
	ResetPassword(input schemas.ResetPassword) schemas.ResponseError
	Login(input schemas.LoginRequest) (*schemas.LoginResponse, schemas.ResponseError)
}
