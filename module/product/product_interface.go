package product

import (
	"go-template-project/dto"
	"go-template-project/module/product/entity"
)

type RepositoryInterface interface {
	GetList(input dto.ProductsRequest) ([]*entity.MSProduct, int64, dto.ResponseError)
	GetIds(ids []uint64) ([]uint64, dto.ResponseError)
}

type UseCaseInterface interface {
	GetList(input dto.ProductsRequest) ([]*entity.MSProduct, int64, dto.ResponseError)
}
