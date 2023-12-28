package usecase

import (
	"go_trunk_based/dto"
	"go_trunk_based/module/product"
	"go_trunk_based/module/product/entity"
)

type ProductUseCase struct {
	productRepo product.RepositoryInterface
}

func InitProductUseCase(repo product.RepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		productRepo: repo,
	}
}

func (u *ProductUseCase) GetList(input dto.ProductsRequest) ([]*entity.Products, int64, dto.ResponseError) {

	res, count, err := u.productRepo.GetList(input)

	return res, count, err
}
