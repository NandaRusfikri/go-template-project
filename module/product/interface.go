package product

import (
	"go-template-project/module/product/entity"
	"go-template-project/schemas"
)

type ProductRepository interface {
	ProductList(input schemas.ProductsRequest) ([]*entity.MSProduct, int64, schemas.ResponseError)
	GetIds(ids []uint64) ([]uint64, schemas.ResponseError)
}

type ItemUseCase interface {
	ProductList(input schemas.ProductsRequest) ([]*entity.MSProduct, int64, schemas.ResponseError)
}
