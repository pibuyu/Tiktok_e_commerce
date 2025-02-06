package model

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).
		Where("name like ? or description like ?", "%"+name+"%", "%"+name+"%").
		Preload("Products").
		Find(&categories).Error
	return
}

// GetCategoryById 根据分类id获取分类
func (c CategoryQuery) GetCategoryById(name string) (result *Category, err error) {
	err = c.db.WithContext(c.ctx).Debug().Model(&Category{}).Where("name = ?", name).Find(&result).Error
	return
}
func (c CategoryQuery) CreateCategory(name string) (err error) {
	cate := &Category{
		Name:        name,
		Description: name,
	}
	err = c.db.WithContext(c.ctx).Debug().Model(&Category{}).Create(&cate).Error
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}
