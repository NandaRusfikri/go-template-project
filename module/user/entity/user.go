package entity

import (
	"go-template-project/constant"
	"time"
)

type User struct {
	ID         uint64     `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt  time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `gorm:"updated_at" json:"-"`
	DeletedAt  *time.Time `gorm:"deleted_at" json:"-"`
	Name       string     `gorm:"name" json:"name"`
	Email      string     `gorm:"email" json:"email"`
	Phone      string     `gorm:"phone" json:"phone,omitempty"`
	Password   string     `gorm:"password" json:"-"`
	IsActive   *bool      `gorm:"is_active;default:true" json:"is_active,omitempty"`
	AvatarPath string     `gorm:"avatar_path" json:"avatar_path,omitempty"`
}

func (entity *User) TableName() string {
	return constant.TABLE_MS_USER
}
