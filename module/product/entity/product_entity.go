package entity

import (
	"go-template-project/constant"
)

type MSProduct struct {
	Id       uint64 `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Price    uint64 `gorm:"column:price" json:"price"`
	Quantity uint64 `gorm:"column:quantity" json:"quantity"`
}

func (entity *MSProduct) TableName() string {
	return constant.TABLE_MS_PRODUCT
}
