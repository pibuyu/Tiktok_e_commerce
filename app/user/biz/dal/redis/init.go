package redis

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/conf"
	"github.com/redis/go-redis/v9"
	"github.com/sethvargo/go-retry"
	"log"
	"time"
)

var (
	RedisClient *redis.Client
)

//func Init() {
//	RedisClient = redis.NewClient(&redis.Options{
//		Addr:     conf.GetConf().Redis.Address,
//		Username: conf.GetConf().Redis.Username,
//		Password: conf.GetConf().Redis.Password,
//		DB:       conf.GetConf().Redis.DB,
//	})
//	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
//		panic(err)
//	}
//}

func ReturnsInstance() *redis.Client {
	var err error
	b := retry.NewFibonacci(10 * time.Second)
	ctx := context.Background()
	if err := retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:         conf.GetConf().Redis.Address,
			Password:     conf.GetConf().Redis.Password, // no password set
			DB:           0,                             // use default DB
			PoolSize:     20,                            // max active conn num
			MinIdleConns: 5,
		})
		_, err = RedisClient.Ping(context.Background()).Result()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Cannot connect to Redis after 5 retries,ERR INFO : %v \n", err)
	}
	return RedisClient
}
