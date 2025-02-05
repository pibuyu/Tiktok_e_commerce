package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
