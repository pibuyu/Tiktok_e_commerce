package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/checkout/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
