package model

import (
	"context"
	"gorm.io/gorm"
)

type OrderItem struct {
	Base
	ProductId    uint32
	ProductName  string
	Picture      string
	OrderIdRefer string `gorm:"size:256;index"`
	Quantity     int32
	Cost         float32
}

func (OrderItem) TableName() string {
	return "order_item"
}

func GetOrderItem(ctx context.Context, db *gorm.DB, orderId string) (oi []OrderItem, err error) {
	err = db.WithContext(ctx).Model(&[]OrderItem{}).
		Where("order_id_refer = ?", orderId).Find(&oi).Error
	return
}
