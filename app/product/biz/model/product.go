package model

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category;constraint:OnDelete:CASCADE"`
}

func (Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// GetById 根据商品id获取商品信息
func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).
		First(&product, productId).Error
	return
}

// SearchProducts 根据关键词查询商品
func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}

// CreateProduct 创建商品
func (p ProductQuery) CreateProduct(pro *Product) (err error) {
	err = p.db.WithContext(p.ctx).Model(Product{}).Create(&pro).Error
	return
}

// DeleteProduct 根据商品id删除商品
func (p ProductQuery) DeleteProduct(id int) (err error) {
	var product Product
	//方案1：级联删除。但是没有正确地删除product_category的关联项，先仅删除product表项
	//err = p.db.Debug().WithContext(p.ctx).Select(clause.Associations).Delete(&product, id).Error

	//方案2：根据product_id找到product_category表中对应的关联项，然后先后删除该表以及product表项
	err = p.db.WithContext(p.ctx).Delete(&product, id).Error
	return
}

// UpdateProductInfo 修改商品信息
func (p ProductQuery) UpdateProductInfo(pro *Product) (err error) {
	//使用Session确保级联更新product_category
	err = p.db.WithContext(p.ctx).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&pro).Error
	return
}

// NewProductQuery 构造NewProductQuery
func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}
