package product

import (
	"go-template-project/dto"
	"go-template-project/module/product/entity"
)

type Repository interface {
	ProductList(input dto.ProductsRequest) ([]*entity.MSProduct, int64, dto.ResponseError)
	GetIds(ids []uint64) ([]uint64, dto.ResponseError)
}

type UseCase interface {
	ProductList(input dto.ProductsRequest) ([]*entity.MSProduct, int64, dto.ResponseError)
}
