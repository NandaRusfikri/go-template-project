package repository

import (
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) UserList(input dto.UsersRequest) ([]*dto.UsersResponse, int64, dto.ResponseError) {
	args := r.Mock.Called(input)

	if args.Get(0) != nil {
		return args.Get(0).([]*dto.UsersResponse), args.Get(1).(int64), args.Get(2).(dto.ResponseError)
	}

	return []*dto.UsersResponse{}, 0, dto.ResponseError{}
}
