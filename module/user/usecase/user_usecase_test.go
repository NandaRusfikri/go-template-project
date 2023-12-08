package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-template-project/dto"
	userrepo "go-template-project/module/user/repository"
	"testing"
)

var userRepository = &userrepo.UserRepositoryMock{Mock: mock.Mock{}}

var userUsecase = InitUserUseCase(userRepository)

type expectedGetList struct {
	res   []*dto.UsersResponse
	count int64
	error dto.ResponseError
}

func TestGetList(t *testing.T) {

	tests := []struct {
		name     string
		request  dto.UsersRequest
		expected expectedGetList
	}{
		{
			name:    "Get two data",
			request: dto.UsersRequest{Limit: 1, Page: 1},
			expected: expectedGetList{
				res: []*dto.UsersResponse{
					{Id: 1, Name: "nanda", Email: "nanda@gmail.com", IsActive: true},
					{Id: 2, Name: "rusfikri", Email: "rusfikri@gmail.com", IsActive: true},
				},
				count: 2,
				error: dto.ResponseError{},
			},
		},
		{
			name:    "No Data",
			request: dto.UsersRequest{Limit: 10, Page: 1},
			expected: expectedGetList{
				res:   []*dto.UsersResponse{},
				count: 0,
				error: dto.ResponseError{},
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			userRepository.Mock.On("GetList", test.request).Return(test.expected.res, test.expected.count, test.expected.error)

			res, count, err := userUsecase.GetList(test.request)

			assert.Equal(t, test.expected.count, count)
			assert.Equal(t, len(test.expected.res), len(res))
			assert.Equal(t, test.expected.error, err)
		})

	}

}
