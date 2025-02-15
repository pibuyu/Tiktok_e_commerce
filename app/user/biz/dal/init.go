package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
