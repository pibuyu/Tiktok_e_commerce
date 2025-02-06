package model

import (
	"context"
	"gorm.io/gorm"
)

type ProductCategory struct {
	CategoryId int `json:"category_id"`
	ProductId  int `json:"product_id"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}

type ProductCategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductCategoryQuery(ctx context.Context, db *gorm.DB) *ProductCategoryQuery {
	return &ProductCategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

func (pcq ProductCategoryQuery) GetByProductId(productId int) (result *ProductCategory, err error) {
	err = pcq.db.WithContext(pcq.ctx).Model(&ProductCategory{}).Where("product_id = ?", productId).Find(result).Error
	return
}

func (pcq ProductCategoryQuery) DeleteByProductId(productId int) (err error) {
	err = pcq.db.WithContext(pcq.ctx).Model(&ProductCategory{}).Where("product_id = ?", productId).Delete(&ProductCategory{}).Error
	return
}
