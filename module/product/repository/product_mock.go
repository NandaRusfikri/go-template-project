package repository

import (
	"github.com/stretchr/testify/mock"
	"go-template-project/module/product/entity"
	"go-template-project/schemas"
)

type ItemRepositoryMock struct {
	Mock mock.Mock
}

func (r *ItemRepositoryMock) GetIds(ids []uint64) ([]uint64, schemas.ResponseError) {
	args := r.Mock.Called(ids)
	var item []uint64
	var databaseError schemas.ResponseError
	if customerArg := args.Get(0); customerArg != nil {
		item = customerArg.([]uint64)
	}
	if errorArg := args.Get(1); errorArg != nil {
		databaseError = errorArg.(schemas.ResponseError)
	}

	return item, databaseError
}

func (r *ItemRepositoryMock) ProductList(input schemas.ProductsRequest) ([]*entity.MSProduct, int64, schemas.ResponseError) {

	args := r.Mock.Called(input)

	customer := []*entity.MSProduct{}
	totalItem := int64(0)
	databaseError := schemas.ResponseError{}

	if len(args) > 0 {
		customer = args.Get(0).([]*entity.MSProduct)
	}
	if len(args) > 1 {
		totalItem = args.Get(1).(int64)
	}
	if len(args) > 2 {
		databaseError = args.Get(2).(schemas.ResponseError)
	}

	return customer, totalItem, databaseError
}
