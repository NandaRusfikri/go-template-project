package user

import (
	"go-template-project/dto"
	userEntity "go-template-project/module/user/entity"
)

type RepositoryInterface interface {
	GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ErrorResponse)
	Insert(input userEntity.Users) dto.ErrorResponse
	FindByEmail(email string) (*userEntity.Users, dto.ErrorResponse)
	FindByUsername(username string) (*userEntity.Users, dto.ErrorResponse)
	FindById(input uint64) (*userEntity.Users, dto.ErrorResponse)
	Update(input userEntity.Users) (*userEntity.Users, dto.ErrorResponse)
	ChangePassword(userId uint64, newPassword string) dto.ErrorResponse
}

type UseCaseInterface interface {
	GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ErrorResponse)
	Insert(input dto.UserInsert) dto.ErrorResponse
	Update(input dto.UserUpdate) (*userEntity.Users, dto.ErrorResponse)
	ChangePassword(input dto.ChangePassword) dto.ErrorResponse
}
