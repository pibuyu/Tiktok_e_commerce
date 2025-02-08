package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := "DYMall:DYMall@tcp(8.138.149.242:3306)/DYMall?charset=utf8mb4&parseTime=True&loc=Local" //先给写死
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}
