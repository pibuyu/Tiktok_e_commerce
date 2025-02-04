package main

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal"
	"log"
	"time"
)

func main() {
	dal.Init()

	//dal.DB.Create(&model.User{
	//	Email:    "123456789@qq.com",
	//	Password: "123456",
	//	Username: "zhangsan",
	//})

	if err := dal.Redis.Set(context.Background(), "name", "hhf", 30*time.Second).Err(); err != nil {
		log.Fatal("set key-value err")
	}
	result, err := dal.Redis.Get(context.Background(), "name").Result()
	if err != nil {
		log.Fatal("get key-value err")
	}
	log.Println("result:", result)
}
