package mysql

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/conf"
	"github.com/sethvargo/go-retry"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

//func Init() {
//	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
//		&gorm.Config{
//			PrepareStmt:            true,
//			SkipDefaultTransaction: true,
//		},
//	)
//	if err != nil {
//		panic(err)
//	}
//}

func ReturnsInstance() *gorm.DB {
	//sql log config
	myLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent, //仅仅在控制台输出指定Debug的语句
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)
	b := retry.NewFibonacci(10 * time.Second)
	ctx := context.Background()
	if err := retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		dsn := conf.GetConf().MySQL.DSN
		//dsn := "DYMall:DYMall@tcp(8.138.149.242:3306)/DYMall?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: myLogger,
		})
		if err != nil {
			return err
		}
		if DB.Error != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Cannot connect to MySQL after 5 retries,ERR INFO : %v \n", err)
	}
	return DB
}
