package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMarkOrderPaid_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMarkOrderPaidService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	req := &order.MarkOrderPaidReq{
		UserId:  1,
		OrderId: "f95bff50-889f-49ff-861a-151fd68436a4",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
