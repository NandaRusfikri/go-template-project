package user

import (
	"go-template-project/dto"
	user_entity "go-template-project/module/user/entity"
)

type Repository interface {
	UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	UserInsert(input user_entity.User) dto.ResponseError
	CheckEmail(email string) (*user_entity.User, dto.ResponseError)
	CheckUsername(username string) (*user_entity.User, dto.ResponseError)
	GetById(input uint64) (*user_entity.User, dto.ResponseError)
	UserUpdate(input user_entity.User) (*user_entity.User, dto.ResponseError)
	ChangePassword(user_id uint64, new_password string) dto.ResponseError
}

type UseCase interface {
	UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError)
	UserInsert(input dto.UserInsert) dto.ResponseError
	UserUpdate(input dto.UserUpdate) (*user_entity.User, dto.ResponseError)
	ChangePassword(input dto.ChangePassword) dto.ResponseError
}
