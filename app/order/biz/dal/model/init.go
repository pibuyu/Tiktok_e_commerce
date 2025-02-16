package model

import (
	"context"
	"gorm.io/gorm"
)

var (
	ProductMapper *ProductQuery
)

func Init(db *gorm.DB) {
	ctx := context.Background()
	ProductMapper = NewProductQuery(ctx, db)
}
