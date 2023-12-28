package repository

import (
	"errors"
	"fmt"
	"go_trunk_based/constant"
	"go_trunk_based/dto"
	"go_trunk_based/module/product"
	"go_trunk_based/module/product/entity"
	"go_trunk_based/util"
	"gorm.io/gorm"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func InitProductRepository(dbCon *gorm.DB) product.RepositoryInterface {
	return &ProductRepository{
		db: dbCon,
	}
}

func (r *ProductRepository) GetIds(ids []uint64) ([]uint64, dto.ResponseError) {

	var ArrayId []uint64
	Check := r.db.Model(&entity.Products{}).Select("id").Where("id IN ?", ids).Find(&ArrayId)

	if Check.Error != nil {
		if errors.Is(Check.Error, gorm.ErrRecordNotFound) {
			return nil, dto.ResponseError{Error: Check.Error, Code: 404}
		}
		return nil, dto.ResponseError{Error: Check.Error, Code: 500}
	}

	return ArrayId, dto.ResponseError{}

}

func (r *ProductRepository) GetList(input dto.ProductsRequest) ([]*entity.Products, int64, dto.ResponseError) {

	var ListItem []*entity.Products

	db := r.db.Model(&entity.Products{})

	if input.SearchText != "" {
		search := "%" + strings.ToLower(input.SearchText) + "%"
		db = db.Where(fmt.Sprintf("LOWER(%v.name) LIKE ? ", constant.TABLE_PRODUCTS), search)
	}

	var count int64

	dbCount := db.Model(entity.Products{}).Count(&count)
	if dbCount.Error != nil && !errors.Is(dbCount.Error, gorm.ErrRecordNotFound) {
		return nil, 0, dto.ResponseError{Error: dbCount.Error, Code: 500}
	}

	if count < 1 {
		return nil, 0, dto.ResponseError{}
	}

	orderByQuery := ""
	if input.OrderField != "" {
		orderColumn, orderType := util.SplitOrderQuery(input.OrderField)
		switch orderColumn {
		case "id":
			orderByQuery += fmt.Sprintf(" %v.id %v", constant.TABLE_PRODUCTS, orderType)
		case "name":
			orderByQuery += fmt.Sprintf(" %v.name %v", constant.TABLE_PRODUCTS, orderType)
		default:
			orderByQuery += fmt.Sprintf(" %v.id DESC", constant.TABLE_PRODUCTS)
		}
	} else {
		orderByQuery += "id DESC"
	}
	db = db.Order(orderByQuery)

	if input.Limit > 0 && input.Page > 0 {
		offset := input.Page*input.Limit - input.Limit
		db = db.Limit(input.Limit).Offset(offset)
	}

	Find := db.Find(&ListItem)
	if Find.Error != nil {
		if errors.Is(Find.Error, gorm.ErrRecordNotFound) {
			return nil, 0, dto.ResponseError{Error: Find.Error, Code: 401}
		}
		return nil, 0, dto.ResponseError{Error: Find.Error, Code: 500}
	}

	return ListItem, count, dto.ResponseError{}

}

func (r *ProductRepository) ProductInsert(input entity.Products) dto.ResponseError {

	Create := r.db.Create(&input)
	if Create.Error != nil {
		return dto.ResponseError{Error: Create.Error, Code: 500}
	}
	if Create.RowsAffected < 1 {
		return dto.ResponseError{Error: fmt.Errorf("error Insert"), Code: 500}
	}

	return dto.ResponseError{}

}
