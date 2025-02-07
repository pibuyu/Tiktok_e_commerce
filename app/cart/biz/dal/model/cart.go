package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null"`
	Qty       uint32 `gorm:"type:int(11);not null"`
}

func (Cart) TableName() string {
	return "cart"
}

// AddItem 向购物车添加商品
func AddItem(ctx context.Context, db *gorm.DB, c *Cart) error {
	var row Cart
	//先查询是否存在
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//如果存在，直接更新该商品的数量
	if row.ID > 0 {
		return db.WithContext(ctx).Model(&Cart{}).
			Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty + ?", c.Qty)).
			Error
	}
	//不存在，则向购物车添加商品
	return db.WithContext(ctx).Create(c).Error
}

// EmptyCart 清空购物车
func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	return db.WithContext(ctx).Debug().Where("user_id = ?", userId).Delete(&Cart{}).Error
}

// GetCartBtUserId 获取用户的购物车信息
func GetCartBtUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	if userId == 0 {
		return nil, errors.New("userId is required")
	}
	err := db.WithContext(ctx).Debug().Model(&Cart{}).Where("user_id = ?", userId).
		Find(&rows).Error
	return rows, err
}
