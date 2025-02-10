package model

import (
	"context"
	"gorm.io/gorm"
)

// Consignee 收货人信息
type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	//todo:需要统一idl中的ZipCode与数据库及代码中的ZipCode类型
	ZipCode string
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   OrderState
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var results []*Order
	err := db.WithContext(ctx).Debug().
		Model(&Order{}).
		Where("user_id = ?", userId).Find(&results).Error
	return results, err
}

func GetOrder(ctx context.Context, db *gorm.DB, userId uint32, orderId string) (o Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).
		Where("user_id = ? and order_id = ?", userId, orderId).Find(&o).Error
	return
}

func UpdateOrderState(ctx context.Context, db *gorm.DB, userId uint32, orderId string, newState OrderState) error {
	return db.WithContext(ctx).Model(&Order{}).
		Where("user_id = ? and order_id = ?", userId, orderId).
		UpdateColumn("order_state", newState).Error
}
