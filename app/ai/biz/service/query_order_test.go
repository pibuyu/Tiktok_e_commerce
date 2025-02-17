package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util/chat_ai"
	ai "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestQueryOrder_Run(t *testing.T) {

	ctx := context.Background()
	s := NewQueryOrderService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	// init req and assert value
	_ = godotenv.Load()
	dal.Init()
	chat_ai.Init()

	req := &ai.OrderQueryRequest{
		UserId:  1,
		Message: "查询 2024 年 1 月 1 日之后创建的订单的订单 ID 和对应的产品名称，以及价格",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	fmt.Println("data:\n", resp.Data)

	// todo: edit your unit test

}
