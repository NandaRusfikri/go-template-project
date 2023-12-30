package repository

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"go-template-project/dto"
	authEntity "go-template-project/module/auth/entity"
	userEntity "go-template-project/module/user/entity"
	"go-template-project/pkg"
	"gorm.io/gorm"

	"time"
)

type AuthRepository struct {
	db *gorm.DB
}

func InitAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) ForgotPassword(userId uint64, token string) (*authEntity.ForgotPassword, dto.ErrorResponse) {

	var entity authEntity.ForgotPassword
	entity.DeletedAt = &gorm.DeletedAt{Valid: true, Time: time.Now()}
	Delete := r.db.Where("user_id = ?", userId).Updates(&entity)
	if Delete.Error != nil {
		log.Errorln("❌ Error when delete to database ==> ", Delete.Error.Error())
		return nil, dto.ErrorResponse{Error: Delete.Error, Code: 500}
	}

	Create := authEntity.ForgotPassword{
		UserId: userId,
		Token:  token,
	}

	Update := r.db.Create(&Create)
	if Update.Error != nil {
		log.Errorln("❌ Error when update to database ==> ", Update.Error.Error())
		return nil, dto.ErrorResponse{Error: Update.Error, Code: 500}
	}

	return &Create, dto.ErrorResponse{}
}

func (r *AuthRepository) ResetPassword(input dto.ResetPassword) dto.ErrorResponse {

	var entity authEntity.ForgotPassword
	Find := r.db.Where("token = ?", input.Token).First(&entity)
	if Find.Error != nil {
		if errors.Is(Find.Error, gorm.ErrRecordNotFound) {
			log.Errorln("❌ token not found ==> ", Find.Error.Error())
			return dto.ErrorResponse{Error: errors.New("token reset password not found"), Code: 404}
		}
		log.Errorln("❌ Error when query to database ==> ", Find.Error.Error())
		return dto.ErrorResponse{Error: Find.Error, Code: 500}
	}

	var user userEntity.Users
	FindUser := r.db.Where("id = ?", entity.UserId).Where("email = ?", input.Email).First(&user)
	if FindUser.Error != nil {
		if errors.Is(FindUser.Error, gorm.ErrRecordNotFound) {
			log.Errorln("❌ Record not found ==> ", FindUser.Error.Error())
			return dto.ErrorResponse{Error: FindUser.Error, Code: 401}
		}
		log.Errorln("❌ Error when query to database ==> ", FindUser.Error.Error())
		return dto.ErrorResponse{Error: FindUser.Error, Code: 500}
	}

	Password := pkg.HashPassword(input.NewPassword)
	user.Password = Password

	tx := r.db.Begin()
	Update := tx.Updates(&user)
	if Update.Error != nil {
		tx.Rollback()
		log.Errorln("❌ Error when update to database ==> ", Update.Error.Error())
		return dto.ErrorResponse{Error: Update.Error, Code: 500}
	}
	entity.DeletedAt = &gorm.DeletedAt{Time: time.Now(), Valid: true}
	DeleteReset := tx.Updates(&entity)
	if DeleteReset.Error != nil {
		tx.Rollback()
		log.Errorln("❌ Error when delete to database ==> ", DeleteReset.Error.Error())
		return dto.ErrorResponse{Error: DeleteReset.Error, Code: 500}
	}

	tx.Commit()

	return dto.ErrorResponse{}
}
