package usecase

import (
	"go-template-project/dto"
	"go-template-project/module/product"
	"go-template-project/module/product/entity"
)

type ProductUseCase struct {
	productRepo product.ProductRepository
}

func InitProductUseCase(repo product.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: repo,
	}
}

func (u *ProductUseCase) ProductList(input dto.ProductsRequest) ([]*entity.MSProduct, int64, dto.ResponseError) {

	res, count, err := u.productRepo.ProductList(input)

	return res, count, err
}
