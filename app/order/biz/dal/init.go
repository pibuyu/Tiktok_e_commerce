package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
	model.Init(mysql.DB)
}
