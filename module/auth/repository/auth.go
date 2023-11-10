package repository

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	entity_auth "go-template-project/module/auth/entity"
	entity_user "go-template-project/module/user/entity"
	"go-template-project/pkg"
	"go-template-project/schemas"
	"gorm.io/gorm"
	"os"
	"time"
)

type AuthRepository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Login(input schemas.LoginRequest) (*schemas.LoginResponse, schemas.ResponseError) {

	var user entity_user.EntityUser

	db := r.db.Where("email = ?", input.Email)

	err := db.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorln("❌ Email not found ==> ", err.Error())
			return nil, schemas.ResponseError{Error: err, Code: 401}
		}
		log.Errorln("❌ Error when query to database ==> ", err.Error())
		return nil, schemas.ResponseError{Error: err, Code: 500}
	}

	CheckPassword := pkg.ComparePassword(user.Password, input.Password)
	if CheckPassword != nil {
		return nil, schemas.ResponseError{Error: errors.New("password invalid"), Code: 401}
	}

	if user.IsActive != nil && !*user.IsActive {
		return nil, schemas.ResponseError{Error: fmt.Errorf("user not active"), Code: 401}
	}

	expiredAt := time.Now().Add(time.Hour * time.Duration(1440))
	claims := schemas.Claims{
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
		return nil, schemas.ResponseError{Code: 500}
	}

	resp := schemas.LoginResponse{
		Id:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AvatarPath:  user.AvatarPath,
		AccessToken: token,
		ExpiredAt:   expiredAt,
	}

	return &resp, schemas.ResponseError{}
}

func (r *AuthRepository) RequestForgotPassword(user_id uint64, token string) (*entity_auth.EntityForgotPassword, schemas.ResponseError) {

	var entity entity_auth.EntityForgotPassword
	entity.DeletedAt = &gorm.DeletedAt{Valid: true, Time: time.Now()}
	Delete := r.db.Where("user_id = ?", user_id).Updates(&entity)
	if Delete.Error != nil {
		log.Errorln("❌ Error when delete to database ==> ", Delete.Error.Error())
		return nil, schemas.ResponseError{Error: Delete.Error, Code: 500}
	}

	Create := entity_auth.EntityForgotPassword{
		UserId: user_id,
		Token:  token,
	}

	Update := r.db.Create(&Create)
	if Update.Error != nil {
		log.Errorln("❌ Error when update to database ==> ", Update.Error.Error())
		return nil, schemas.ResponseError{Error: Update.Error, Code: 500}
	}

	return &Create, schemas.ResponseError{}
}

func (r *AuthRepository) ResetPassword(input schemas.ResetPassword) schemas.ResponseError {

	var entity entity_auth.EntityForgotPassword
	Find := r.db.Where("token = ?", input.Token).First(&entity)
	if Find.Error != nil {
		if errors.Is(Find.Error, gorm.ErrRecordNotFound) {
			log.Errorln("❌ token not found ==> ", Find.Error.Error())
			return schemas.ResponseError{Error: errors.New("token reset passoword not found"), Code: 404}
		}
		log.Errorln("❌ Error when query to database ==> ", Find.Error.Error())
		return schemas.ResponseError{Error: Find.Error, Code: 500}
	}

	var user entity_user.EntityUser
	FindUser := r.db.Where("id = ?", entity.UserId).Where("email = ?", input.Email).First(&user)
	if FindUser.Error != nil {
		if errors.Is(FindUser.Error, gorm.ErrRecordNotFound) {
			log.Errorln("❌ Record not found ==> ", FindUser.Error.Error())
			return schemas.ResponseError{Error: FindUser.Error, Code: 401}
		}
		log.Errorln("❌ Error when query to database ==> ", FindUser.Error.Error())
		return schemas.ResponseError{Error: FindUser.Error, Code: 500}
	}

	Password := pkg.HashPassword(input.NewPassword)
	user.Password = Password

	tx := r.db.Begin()
	Update := tx.Updates(&user)
	if Update.Error != nil {
		tx.Rollback()
		log.Errorln("❌ Error when update to database ==> ", Update.Error.Error())
		return schemas.ResponseError{Error: Update.Error, Code: 500}
	}
	entity.DeletedAt = &gorm.DeletedAt{Time: time.Now(), Valid: true}
	DeleteReset := tx.Updates(&entity)
	if DeleteReset.Error != nil {
		tx.Rollback()
		log.Errorln("❌ Error when delete to database ==> ", DeleteReset.Error.Error())
		return schemas.ResponseError{Error: DeleteReset.Error, Code: 500}
	}

	tx.Commit()

	return schemas.ResponseError{}
}
