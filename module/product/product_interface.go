package product

import (
	"go_trunk_based/dto"
	"go_trunk_based/module/product/entity"
)

type RepositoryInterface interface {
	GetList(input dto.ProductsRequest) ([]*entity.Products, int64, dto.ResponseError)
	GetIds(ids []uint64) ([]uint64, dto.ResponseError)
}

type UseCaseInterface interface {
	GetList(input dto.ProductsRequest) ([]*entity.Products, int64, dto.ResponseError)
}
