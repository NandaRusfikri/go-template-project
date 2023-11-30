package user

import (
	"go-template-project/dto"
	userEntity "go-template-project/module/user/entity"
)

type RepositoryInterface interface {
	GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	Insert(input userEntity.User) dto.ResponseError
	FindByEmail(email string) (*userEntity.User, dto.ResponseError)
	FindByUsername(username string) (*userEntity.User, dto.ResponseError)
	FindById(input uint64) (*userEntity.User, dto.ResponseError)
	Update(input userEntity.User) (*userEntity.User, dto.ResponseError)
	ChangePassword(userId uint64, newPassword string) dto.ResponseError
}

type UseCaseInterface interface {
	GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	Insert(input dto.UserInsert) dto.ResponseError
	Update(input dto.UserUpdate) (*userEntity.User, dto.ResponseError)
	ChangePassword(input dto.ChangePassword) dto.ResponseError
}
