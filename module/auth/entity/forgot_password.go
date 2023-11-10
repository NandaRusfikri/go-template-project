package entity

import (
	"go-template-project/constant"
	entity_user "go-template-project/module/user/entity"
	"gorm.io/gorm"
	"time"
)

type EntityForgotPassword struct {
	ID        uint64                  `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time               `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt *time.Time              `gorm:"updated_at" json:"-"`
	DeletedAt *gorm.DeletedAt         `gorm:"deleted_at" json:"-"`
	UserId    uint64                  `gorm:"column:user_id" json:"user_id"`
	User      *entity_user.EntityUser `gorm:"foreignKey:user_id" json:"user"`
	Token     string                  `gorm:"column:token" json:"token"`
}

func (entity *EntityForgotPassword) TableName() string {
	return constant.TABLE_TR_FORGOT_PASSWORD
}

func (entity *EntityForgotPassword) BeforeUpdate(db *gorm.DB) error {
	time := time.Now().Local()
	entity.UpdatedAt = &time
	return nil
}
