package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal"
	cart "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	req := &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 1,
			Quantity:  1,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
