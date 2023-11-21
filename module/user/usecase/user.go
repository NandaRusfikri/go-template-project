package usecase

import (
	"errors"
	"fmt"
	"go-template-project/dto"
	"go-template-project/module/user"
	user_entity "go-template-project/module/user/entity"
	"go-template-project/pkg"
)

type UserUseCase struct {
	user_repo user.UserRepository
	SMTP      *pkg.SMTP
}

func InitUserUseCase(repository user.UserRepository) *UserUseCase {
	return &UserUseCase{user_repo: repository}
}

func (u *UserUseCase) UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError) {

	res, count, err := u.user_repo.UserList(input)
	return res, count, err
}

func (u *UserUseCase) UserInsert(input dto.UserInsert) dto.ResponseError {

	CheckEmail, _ := u.user_repo.CheckEmail(input.Email)
	if CheckEmail != nil {
		return dto.ResponseError{Error: fmt.Errorf("email already exist"), Code: 400}
	}

	EntityUser := user_entity.EntityUser{
		Name: input.Name,
		//Username: input.Username,
		Email:    input.Email,
		Password: pkg.HashPassword(input.Password),

		//Status: ""
	}

	err := u.user_repo.UserInsert(EntityUser)

	return err

}

func (u *UserUseCase) UserUpdate(input dto.UserUpdate) (*user_entity.EntityUser, dto.ResponseError) {

	entity := user_entity.EntityUser{
		ID:    input.Id,
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	if input.IsActive != nil {
		entity.IsActive = input.IsActive
	}

	res, err := u.user_repo.UserUpdate(entity)
	return res, err
}

func (u *UserUseCase) ChangePassword(input dto.ChangePassword) dto.ResponseError {

	DataUser, err := u.user_repo.GetById(input.UserId)
	if err.Error != nil {
		return dto.ResponseError{Code: 404, Error: errors.New("user not found")}
	}

	CheckPassword := pkg.ComparePassword(DataUser.Password, input.OldPassword)
	if CheckPassword != nil {
		return dto.ResponseError{Error: CheckPassword, Code: 401}
	}

	err = u.user_repo.ChangePassword(input.UserId, pkg.HashPassword(input.NewPassword))
	return err
}
