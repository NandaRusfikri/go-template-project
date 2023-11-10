package usecase

import (
	"go-template-project/module/product"
	"go-template-project/module/product/entity"
	"go-template-project/schemas"
)

type ProductUseCase struct {
	product_repo product.ProductRepository
}

func InitProductUseCase(repo product.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		product_repo: repo,
	}
}

func (u *ProductUseCase) ProductList(input schemas.ProductsRequest) ([]*entity.MSProduct, int64, schemas.ResponseError) {

	res, count, err := u.product_repo.ProductList(input)

	return res, count, err
}
