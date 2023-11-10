package repository

import (
	"errors"
	"fmt"
	"go-template-project/constant"
	"go-template-project/module/product"
	"go-template-project/module/product/entity"
	"go-template-project/schemas"
	"go-template-project/util"
	"gorm.io/gorm"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func InitProductRepository(dbCon *gorm.DB) product.ProductRepository {
	return &ProductRepository{
		db: dbCon,
	}
}

func (r *ProductRepository) GetIds(ids []uint64) ([]uint64, schemas.ResponseError) {

	var ArrayId []uint64
	Check := r.db.Model(&entity.MSProduct{}).Select("id").Where("id IN ?", ids).Find(&ArrayId)

	if Check.Error != nil {
		if errors.Is(Check.Error, gorm.ErrRecordNotFound) {
			return nil, schemas.ResponseError{Error: Check.Error, Code: 404}
		}
		return nil, schemas.ResponseError{Error: Check.Error, Code: 500}
	}

	return ArrayId, schemas.ResponseError{}

}

func (r *ProductRepository) ProductList(input schemas.ProductsRequest) ([]*entity.MSProduct, int64, schemas.ResponseError) {

	var ListItem []*entity.MSProduct

	db := r.db.Model(&entity.MSProduct{})

	if input.SearchText != "" {
		search := "%" + strings.ToLower(input.SearchText) + "%"
		db = db.Where(fmt.Sprintf("LOWER(%v.name) LIKE ? ", constant.TABLE_MS_PRODUCT), search)
	}

	var count int64

	dbCount := db.Model(entity.MSProduct{}).Count(&count)
	if dbCount.Error != nil && !errors.Is(dbCount.Error, gorm.ErrRecordNotFound) {
		return nil, 0, schemas.ResponseError{Error: dbCount.Error, Code: 500}
	}

	if count < 1 {
		return nil, 0, schemas.ResponseError{}
	}

	orderByQuery := ""
	if input.OrderField != "" {
		orderColumn, orderType := util.SplitOrderQuery(input.OrderField)
		switch orderColumn {
		case "id":
			orderByQuery += fmt.Sprintf(" %v.id %v", constant.TABLE_MS_PRODUCT, orderType)
		case "name":
			orderByQuery += fmt.Sprintf(" %v.name %v", constant.TABLE_MS_PRODUCT, orderType)
		default:
			orderByQuery += fmt.Sprintf(" %v.id DESC", constant.TABLE_MS_PRODUCT)
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
			return nil, 0, schemas.ResponseError{Error: Find.Error, Code: 401}
		}
		return nil, 0, schemas.ResponseError{Error: Find.Error, Code: 500}
	}

	return ListItem, count, schemas.ResponseError{}

}

func (r *ProductRepository) ProductInsert(input entity.MSProduct) schemas.ResponseError {

	Create := r.db.Create(&input)
	if Create.Error != nil {
		return schemas.ResponseError{Error: Create.Error, Code: 500}
	}
	if Create.RowsAffected < 1 {
		return schemas.ResponseError{Error: fmt.Errorf("Error Insert"), Code: 500}
	}

	return schemas.ResponseError{}

}
