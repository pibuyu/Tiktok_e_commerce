package service

import (
	"context"
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

	req := &ai.AutoOrderRequest{Message: "我现在需要购买vivo x200 手机和Notebook，请为我自动下单"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
