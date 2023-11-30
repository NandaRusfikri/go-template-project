package usecase

import (
	"go-template-project/dto"
	"go-template-project/module/product"
	"go-template-project/module/product/entity"
)

type ProductUseCase struct {
	productRepo product.RepositoryInterface
}

func InitProductUseCase(repo product.RepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		productRepo: repo,
	}
}

func (u *ProductUseCase) GetList(input dto.ProductsRequest) ([]*entity.MSProduct, int64, dto.ResponseError) {

	res, count, err := u.productRepo.GetList(input)

	return res, count, err
}
