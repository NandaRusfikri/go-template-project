package usecase

import (
	"fmt"
	"go-template-project/module/auth"
	user_repo "go-template-project/module/user/repository"
	"go-template-project/pkg"
	"go-template-project/schemas"
	"go-template-project/util"
)

type AuthUseCase struct {
	auth_repo auth.AuthRepository
	user_repo user_repo.UserRepository
	SMTP      *pkg.SMTP
}

func InitAuthUseCase(auth_repo auth.AuthRepository, smtp *pkg.SMTP) *AuthUseCase {
	return &AuthUseCase{auth_repo: auth_repo, SMTP: smtp}
}

func (u *AuthUseCase) Login(input schemas.LoginRequest) (*schemas.LoginResponse, schemas.ResponseError) {

	res, err := u.auth_repo.Login(input)

	return res, err
}

func (u *AuthUseCase) RequestForgotPassword(input schemas.ForgotPassword) schemas.ResponseError {

	CheckEmail, _ := u.user_repo.CheckEmail(input.Email)
	if CheckEmail == nil {
		return schemas.ResponseError{Error: fmt.Errorf("email not found"), Code: 400}
	}
	res, err := u.auth_repo.RequestForgotPassword(CheckEmail.ID, util.RandomInt(6))
	if err != (schemas.ResponseError{}) {
		return err
	}

	message := fmt.Sprintf("Forgot Password Link \n http://disewa.id/nama_web/reset-password/%v/%v", input.Email, res.Token)
	SendEmail := u.SMTP.SendEmail([]string{input.Email}, []string{}, []string{}, "Forgot Password", "text/plain", message, []string{})
	if SendEmail != nil {
		return schemas.ResponseError{
			Error: fmt.Errorf("failed send email"),
			Code:  500,
		}
	}
	return schemas.ResponseError{}
}

func (u *AuthUseCase) ResetPassword(input schemas.ResetPassword) schemas.ResponseError {

	err := u.auth_repo.ResetPassword(input)
	return err
}
