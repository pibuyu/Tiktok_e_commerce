package mysql

import (
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/biz/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	klog.Infof("mysql dsn:%s", dsn)
	klog.Infof("MYSQL_USER: %s, MYSQL_PASSWORD: %s, MYSQL_HOST: %s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	_ = DB.AutoMigrate(&model.PaymentLog{})
	if err != nil {
		panic(err)
	}
}
