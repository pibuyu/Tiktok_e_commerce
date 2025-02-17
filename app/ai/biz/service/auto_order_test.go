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

func TestAutoOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAutoOrderService(ctx)
	// init req and assert value

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	// init req and assert value
	_ = godotenv.Load()
	dal.Init()
	chat_ai.Init()

	question := "我现在需要购买vivo x200 手机和Notebook，请为我自动下单"
	req := &ai.AutoOrderRequest{
		UserId:  1,
		Message: question,
	}
	resp, _ := s.Run(req)

	fmt.Println("问题：", question)
	fmt.Println("回复：", resp.Data)
}

func TestAutoOrder_Run2(t *testing.T) {
	ctx := context.Background()
	s := NewAutoOrderService(ctx)
	// init req and assert value

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	// init req and assert value
	_ = godotenv.Load()
	dal.Init()
	chat_ai.Init()

	question := "我现在需要购买vivo x200 手机和Notebook，请为我自动下单, 邮政编码为'200000'，在中国上海市上海和平路123号"
	req := &ai.AutoOrderRequest{
		UserId:  1,
		Message: question,
	}
	resp, _ := s.Run(req)

	fmt.Println("问题：", question)
	fmt.Println("回复：", resp.Data)
}
