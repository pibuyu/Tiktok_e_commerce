package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	addr := &order.Address{
		StreetAddress: "外环东路52号",
		City:          "广州市",
		State:         "广东省",
		Country:       "中国",
		ZipCode:       "510006",
	}
	var itemList []*order.OrderItem
	item := &order.OrderItem{
		Item: &cart.CartItem{
			ProductId: 1,
			Quantity:  1,
		},
		Cost: 0,
	}
	itemList = append(itemList, item)
	req := &order.PlaceOrderReq{
		UserId: 1,
		//UserCurrency: "100",
		Address: addr,
		Email:   "3531095171@qq.com",
		//OrderItem:    itemList,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
