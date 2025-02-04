package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/mysql"
	redisClient "github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Init() {
	Redis = redisClient.ReturnsInstance()
	DB = mysql.ReturnsInstance()
}

var (
	DB    *gorm.DB
	Redis *redis.Client
)
