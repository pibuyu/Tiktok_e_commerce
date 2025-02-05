package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
