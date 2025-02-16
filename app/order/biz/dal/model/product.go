package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
}

func (Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// ProductMutation 读写分离，给ProductMutation传递写操作
type ProductMutation struct {
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
	err = p.db.WithContext(p.ctx).Model(&Product{}).
		Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
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

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cachedClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  cachedClient,
		prefix:       "shop",
	}
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	klog.Infof("redisKey:%s", cachedKey)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)

	//捕获可能发生的任意错误
	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return err
		}
		//获取结果的byte
		cachedResultByte, err := cachedResult.Bytes()
		if err != nil {
			return err
		}
		//对byte结果反序列化
		err = json.Unmarshal(cachedResultByte, &product)
		if err != nil {
			return err
		}
		return nil
	}()

	//如果redis查询过程中出错，则使用MySQL查询
	if err != nil {
		product, err = c.productQuery.GetById(productId)
		//MySQL查询也错误,返回错误
		if err != nil {
			return Product{}, err
		}

		//序列化并放入redis
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	}
	return
}

func (c CachedProductQuery) SearchProducts(q string) (products []*Product, err error) {
	return c.productQuery.SearchProducts(q)
}
