package repository

import (
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
	"go-template-project/module/user/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ErrorResponse) {
	args := r.Mock.Called(input)

	if args.Get(0) != nil {
		return args.Get(0).([]*dto.UsersResponse), args.Get(1).(int64), args.Get(2).(dto.ErrorResponse)
	}

	return []*dto.UsersResponse{}, 0, dto.ErrorResponse{}
}

func (r *UserRepositoryMock) Insert(input entity.Users) dto.ErrorResponse {
	args := r.Mock.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(dto.ErrorResponse)
	}
	return dto.ErrorResponse{}
}
func (r *UserRepositoryMock) FindByEmail(email string) (*entity.Users, dto.ErrorResponse) {
	args := r.Mock.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ErrorResponse)
	}
	return &entity.Users{}, dto.ErrorResponse{}
}
func (r *UserRepositoryMock) FindByUsername(username string) (*entity.Users, dto.ErrorResponse) {
	args := r.Mock.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ErrorResponse)
	}
	return &entity.Users{}, dto.ErrorResponse{}
}
func (r *UserRepositoryMock) FindById(input uint64) (*entity.Users, dto.ErrorResponse) {
	args := r.Mock.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ErrorResponse)
	}

	return &entity.Users{}, dto.ErrorResponse{}
}
func (r *UserRepositoryMock) Update(input entity.Users) (*entity.Users, dto.ErrorResponse) {
	args := r.Mock.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ErrorResponse)
	}
	return &entity.Users{}, dto.ErrorResponse{}
}
func (r *UserRepositoryMock) ChangePassword(userId uint64, newPassword string) dto.ErrorResponse {
	args := r.Mock.Called(userId, newPassword)
	if args.Get(0) != nil {
		return args.Get(0).(dto.ErrorResponse)
	}
	return dto.ErrorResponse{}
}
