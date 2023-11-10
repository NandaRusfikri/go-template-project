package usecase

import (
	"errors"
	"fmt"
	"go-template-project/module/user"
	user_entity "go-template-project/module/user/entity"
	"go-template-project/pkg"
	"go-template-project/schemas"
)

type UserUseCase struct {
	user_repo user.UserRepository
	SMTP      *pkg.SMTP
}

func InitUserUseCase(repository user.UserRepository) *UserUseCase {
	return &UserUseCase{user_repo: repository}
}

func (u *UserUseCase) ListUserService(input schemas.UsersRequest) ([]*schemas.UsersResponse, int64, schemas.ResponseError) {

	res, count, err := u.user_repo.UserList(input)
	return res, count, err
}

func (u *UserUseCase) CreateNewUserService(input schemas.UserInsert) schemas.ResponseError {

	CheckEmail, _ := u.user_repo.CheckEmail(input.Email)
	if CheckEmail != nil {
		return schemas.ResponseError{Error: fmt.Errorf("email already exist"), Code: 400}
	}

	EntityUser := user_entity.EntityUser{
		Name: input.Name,
		//Username: input.Username,
		Email:    input.Email,
		Password: pkg.HashPassword(input.Password),

		//Status: ""
	}

	err := u.user_repo.CreateUser(EntityUser)

	return err

}

func (u *UserUseCase) UpdateUserService(input schemas.UserUpdate) (*user_entity.EntityUser, schemas.ResponseError) {

	entity := user_entity.EntityUser{
		ID:    input.Id,
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	if input.IsActive != nil {
		entity.IsActive = input.IsActive
	}

	res, err := u.user_repo.UpdateUser(entity)
	return res, err
}

func (u *UserUseCase) ChangePasswordService(input schemas.ChangePassword) schemas.ResponseError {

	DataUser, err := u.user_repo.GetById(input.UserId)
	if err.Error != nil {
		return schemas.ResponseError{Code: 404, Error: errors.New("user not found")}
	}

	CheckPassword := pkg.ComparePassword(DataUser.Password, input.OldPassword)
	if CheckPassword != nil {
		return schemas.ResponseError{Error: CheckPassword, Code: 401}
	}

	err = u.user_repo.ChangePassword(input.UserId, pkg.HashPassword(input.NewPassword))
	return err
}
