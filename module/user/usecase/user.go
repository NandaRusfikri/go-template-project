package usecase

import (
	"errors"
	"fmt"
	"go-template-project/dto"
	"go-template-project/module/user"
	userEntity "go-template-project/module/user/entity"
	"go-template-project/pkg"
)

type UserUseCase struct {
	userRepo user.UserRepository
	SMTP     *pkg.SMTP
}

func InitUserUseCase(repository user.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repository}
}

func (u *UserUseCase) UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError) {

	res, count, err := u.userRepo.UserList(input)
	return res, count, err
}

func (u *UserUseCase) UserInsert(input dto.UserInsert) dto.ResponseError {

	CheckEmail, _ := u.userRepo.CheckEmail(input.Email)
	if CheckEmail != nil {
		return dto.ResponseError{Error: fmt.Errorf("email already exist"), Code: 400}
	}

	EntityUser := userEntity.EntityUser{
		Name: input.Name,
		//Username: input.Username,
		Email:    input.Email,
		Password: pkg.HashPassword(input.Password),
	}

	err := u.userRepo.UserInsert(EntityUser)

	return err

}

func (u *UserUseCase) UserUpdate(input dto.UserUpdate) (*userEntity.EntityUser, dto.ResponseError) {

	entity := userEntity.EntityUser{
		ID:    input.Id,
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	if input.IsActive != nil {
		entity.IsActive = input.IsActive
	}

	res, err := u.userRepo.UserUpdate(entity)
	return res, err
}

func (u *UserUseCase) ChangePassword(input dto.ChangePassword) dto.ResponseError {

	DataUser, err := u.userRepo.GetById(input.UserId)
	if err.Error != nil {
		return dto.ResponseError{Code: 404, Error: errors.New("user not found")}
	}

	CheckPassword := pkg.ComparePassword(DataUser.Password, input.OldPassword)
	if CheckPassword != nil {
		return dto.ResponseError{Error: CheckPassword, Code: 401}
	}

	err = u.userRepo.ChangePassword(input.UserId, pkg.HashPassword(input.NewPassword))
	return err
}
