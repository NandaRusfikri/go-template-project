package user

import (
	user_entity "go-template-project/module/user/entity"
	"go-template-project/schemas"
)

type UserRepository interface {
	UserList(input schemas.UsersRequest) ([]*schemas.UsersResponse, int64, schemas.ResponseError)
	CreateUser(input user_entity.EntityUser) schemas.ResponseError
	CheckEmail(email string) (*user_entity.EntityUser, schemas.ResponseError)
	CheckUsername(username string) (*user_entity.EntityUser, schemas.ResponseError)
	GetById(input uint64) (*user_entity.EntityUser, schemas.ResponseError)
	UpdateUser(input user_entity.EntityUser) (*user_entity.EntityUser, schemas.ResponseError)
	ChangePassword(user_id uint64, new_password string) schemas.ResponseError
}

type UserUseCase interface {
	ListUserService(input schemas.UsersRequest) ([]*schemas.UsersResponse, int64, schemas.ResponseError)
	CreateNewUserService(input schemas.UserInsert) schemas.ResponseError
	UpdateUserService(input schemas.UserUpdate) (*user_entity.EntityUser, schemas.ResponseError)
	ChangePasswordService(input schemas.ChangePassword) schemas.ResponseError
}
