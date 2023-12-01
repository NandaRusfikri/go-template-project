package user

import (
	"go-template-project/dto"
	userEntity "go-template-project/module/user/entity"
)

type RepositoryInterface interface {
	GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	Insert(input userEntity.Users) dto.ResponseError
	FindByEmail(email string) (*userEntity.Users, dto.ResponseError)
	FindByUsername(username string) (*userEntity.Users, dto.ResponseError)
	FindById(input uint64) (*userEntity.Users, dto.ResponseError)
	Update(input userEntity.Users) (*userEntity.Users, dto.ResponseError)
	ChangePassword(userId uint64, newPassword string) dto.ResponseError
}

type UseCaseInterface interface {
	GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	Insert(input dto.UserInsert) dto.ResponseError
	Update(input dto.UserUpdate) (*userEntity.Users, dto.ResponseError)
	ChangePassword(input dto.ChangePassword) dto.ResponseError
}
