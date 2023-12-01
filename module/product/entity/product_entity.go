package entity

import (
	"go-template-project/constant"
)

type Products struct {
	Id       uint64 `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Price    uint64 `gorm:"column:price" json:"price"`
	Quantity uint64 `gorm:"column:quantity" json:"quantity"`
}

func (entity *Products) TableName() string {
	return constant.TABLE_PRODUCTS
}
