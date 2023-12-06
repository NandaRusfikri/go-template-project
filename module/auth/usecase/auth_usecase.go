package usecase

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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
	authRepo auth.RepositoryInterface
	userRepo user.RepositoryInterface
	SMTP     *pkg.SMTP
}

func InitAuthUseCase(authRepo auth.RepositoryInterface, userRepo user.RepositoryInterface, smtp *pkg.SMTP) *AuthUseCase {
	return &AuthUseCase{authRepo: authRepo, userRepo: userRepo, SMTP: smtp}
}

func (u *AuthUseCase) ResetPassword(input dto.ResetPassword) dto.ResponseError {

	err := u.authRepo.ResetPassword(input)
	return err
}

func (u *AuthUseCase) Login(input dto.LoginRequest) (*dto.LoginResponse, dto.ResponseError) {

	entityUser, err := u.userRepo.FindByEmail(input.Email)

	checkPassword := pkg.ComparePassword(entityUser.Password, input.Password)
	if checkPassword != nil {
		return nil, dto.ResponseError{Error: errors.New("password invalid"), Code: 401}
	}

	if entityUser.IsActive != nil && !*entityUser.IsActive {
		return nil, dto.ResponseError{Error: fmt.Errorf("entityUser not active"), Code: 401}
	}

	expiredAt := time.Now().Add(time.Hour * time.Duration(constant.DURATION_TOKEN))
	claims := dto.Claims{
		Id:    entityUser.ID,
		Name:  entityUser.Name,
		Email: entityUser.Email,
	}
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiredAt),
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, errJWT := pkg.Sign(claims, jwtSecretKey)
	if errJWT != nil {
		return nil, dto.ResponseError{Code: 500}
	}

	res := dto.LoginResponse{
		Id:          entityUser.ID,
		Name:        entityUser.Name,
		Email:       entityUser.Email,
		AvatarPath:  entityUser.AvatarPath,
		AccessToken: token,
		ExpiredAt:   expiredAt,
	}

	return &res, err
}

func (u *AuthUseCase) ForgotPassword(input dto.ForgotPassword) dto.ResponseError {

	checkEmail, _ := u.userRepo.FindByEmail(input.Email)
	if checkEmail == nil {
		return dto.ResponseError{Error: fmt.Errorf("email not found"), Code: 400}
	}
	res, err := u.authRepo.ForgotPassword(checkEmail.ID, util.RandomInt(6))
	if err != (dto.ResponseError{}) {
		return err
	}

	message := fmt.Sprintf("Forgot Password Link \n https://disewa.id/nama_web/reset-password/%v/%v", input.Email, res.Token)

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
