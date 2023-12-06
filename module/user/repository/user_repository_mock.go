package repository

import (
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
	"go-template-project/module/user/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) GetList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError) {
	args := r.Mock.Called(input)

	if args.Get(0) != nil {
		return args.Get(0).([]*dto.UsersResponse), args.Get(1).(int64), args.Get(2).(dto.ResponseError)
	}

	return []*dto.UsersResponse{}, 0, dto.ResponseError{}
}

func (r *UserRepositoryMock) Insert(input entity.Users) dto.ResponseError {
	args := r.Mock.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(dto.ResponseError)
	}
	return dto.ResponseError{}
}
func (r *UserRepositoryMock) FindByEmail(email string) (*entity.Users, dto.ResponseError) {
	args := r.Mock.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ResponseError)
	}
	return &entity.Users{}, dto.ResponseError{}
}
func (r *UserRepositoryMock) FindByUsername(username string) (*entity.Users, dto.ResponseError) {
	args := r.Mock.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ResponseError)
	}
	return &entity.Users{}, dto.ResponseError{}
}
func (r *UserRepositoryMock) FindById(input uint64) (*entity.Users, dto.ResponseError) {
	args := r.Mock.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ResponseError)
	}

	return &entity.Users{}, dto.ResponseError{}
}
func (r *UserRepositoryMock) Update(input entity.Users) (*entity.Users, dto.ResponseError) {
	args := r.Mock.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Users), args.Get(1).(dto.ResponseError)
	}
	return &entity.Users{}, dto.ResponseError{}
}
func (r *UserRepositoryMock) ChangePassword(userId uint64, newPassword string) dto.ResponseError {
	args := r.Mock.Called(userId, newPassword)
	if args.Get(0) != nil {
		return args.Get(0).(dto.ResponseError)
	}
	return dto.ResponseError{}
}
