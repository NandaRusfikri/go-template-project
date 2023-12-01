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
	userRepo user.RepositoryInterface
	SMTP     *pkg.SMTP
}

func InitUserUseCase(repository user.RepositoryInterface) *UserUseCase {
	return &UserUseCase{userRepo: repository}
}

func (u *UserUseCase) GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError) {
	return u.userRepo.GetList(input)
}

func (u *UserUseCase) Insert(input dto.UserInsert) dto.ResponseError {

	GetByEmail, _ := u.userRepo.FindByEmail(input.Email)
	if GetByEmail != nil {
		return dto.ResponseError{Error: fmt.Errorf("email already exist"), Code: 400}
	}

	EntityUser := userEntity.Users{
		Name:     input.Name,
		Email:    input.Email,
		Password: pkg.HashPassword(input.Password),
	}

	err := u.userRepo.Insert(EntityUser)

	return err

}

func (u *UserUseCase) Update(input dto.UserUpdate) (*userEntity.Users, dto.ResponseError) {

	entity := userEntity.Users{
		ID:    input.Id,
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	if input.IsActive != nil {
		entity.IsActive = input.IsActive
	}

	res, err := u.userRepo.Update(entity)
	return res, err
}

func (u *UserUseCase) ChangePassword(input dto.ChangePassword) dto.ResponseError {

	DataUser, err := u.userRepo.FindById(input.UserId)
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
