package usecase

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-template-project/constant"
	"go-template-project/dto"
	"go-template-project/module/auth"
	"go-template-project/module/user"
	"go-template-project/pkg"
	"go-template-project/util"
	"os"
	"time"
)

type AuthUseCase struct {
	auth_repo auth.AuthRepository
	user_repo user.UserRepository
	SMTP      *pkg.SMTP
}

func InitAuthUseCase(auth_repo auth.AuthRepository, user_repo user.UserRepository, smtp *pkg.SMTP) *AuthUseCase {
	return &AuthUseCase{auth_repo: auth_repo, user_repo: user_repo, SMTP: smtp}
}

func (u *AuthUseCase) Login(input dto.LoginRequest) (*dto.LoginResponse, dto.ResponseError) {

	user, err := u.user_repo.CheckEmail(input.Email)

	CheckPassword := pkg.ComparePassword(user.Password, input.Password)
	if CheckPassword != nil {
		return nil, dto.ResponseError{Error: errors.New("password invalid"), Code: 401}
	}

	if user.IsActive != nil && !*user.IsActive {
		return nil, dto.ResponseError{Error: fmt.Errorf("user not active"), Code: 401}
	}

	expiredAt := time.Now().Add(time.Hour * time.Duration(constant.DURATION_TOKEN))
	claims := dto.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, ErrorJWT := pkg.Sign(claims, jwtSecretKey)
	if ErrorJWT != nil {
		return nil, dto.ResponseError{Code: 500}
	}

	res := dto.LoginResponse{
		Id:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AvatarPath:  user.AvatarPath,
		AccessToken: token,
		ExpiredAt:   expiredAt,
	}

	return &res, err
}

func (u *AuthUseCase) ForgotPassword(input dto.ForgotPassword) dto.ResponseError {

	CheckEmail, _ := u.user_repo.CheckEmail(input.Email)
	if CheckEmail == nil {
		return dto.ResponseError{Error: fmt.Errorf("email not found"), Code: 400}
	}
	res, err := u.auth_repo.ForgotPassword(CheckEmail.ID, util.RandomInt(6))
	if err != (dto.ResponseError{}) {
		return err
	}

	message := fmt.Sprintf("Forgot Password Link \n http://disewa.id/nama_web/reset-password/%v/%v", input.Email, res.Token)

	SendEmail := u.SMTP.SendEmail(dto.SendEmail{
		To:         []string{input.Email},
		Cc:         []string{},
		Bcc:        []string{},
		Subject:    "Forgot Password",
		BodyType:   "text/plain",
		Body:       message,
		Attachment: []string{},
	})
	if SendEmail != nil {
		return dto.ResponseError{
			Error: fmt.Errorf("failed send email"),
			Code:  500,
		}
	}
	return dto.ResponseError{}
}

func (u *AuthUseCase) ResetPassword(input dto.ResetPassword) dto.ResponseError {

	err := u.auth_repo.ResetPassword(input)
	return err
}
