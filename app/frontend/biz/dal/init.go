package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_conmmerce/app/frontend/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_conmmerce/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
