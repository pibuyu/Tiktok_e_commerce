package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
