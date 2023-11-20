package user

import (
	"go-template-project/dto"
	user_entity "go-template-project/module/user/entity"
)

type UserRepository interface {
	UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	CreateUser(input user_entity.EntityUser) dto.ResponseError
	CheckEmail(email string) (*user_entity.EntityUser, dto.ResponseError)
	CheckUsername(username string) (*user_entity.EntityUser, dto.ResponseError)
	GetById(input uint64) (*user_entity.EntityUser, dto.ResponseError)
	UpdateUser(input user_entity.EntityUser) (*user_entity.EntityUser, dto.ResponseError)
	ChangePassword(user_id uint64, new_password string) dto.ResponseError
}

type UserUseCase interface {
	ListUserService(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	CreateNewUserService(input dto.UserInsert) dto.ResponseError
	UpdateUserService(input dto.UserUpdate) (*user_entity.EntityUser, dto.ResponseError)
	ChangePasswordService(input dto.ChangePassword) dto.ResponseError
}
