package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/checkout/biz/dal"
	checkout "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckoutService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	cardInfo := &payment.CreditCardInfo{
		CreditCardNumber:          "4111111111111111",
		CreditCardExpirationYear:  2030,
		CreditCardExpirationMonth: 02,
		CreditCardCvv:             123,
	}
	deliveryAddress := &checkout.Address{
		StreetAddress: "外环东路52号",
		City:          "广州市",
		State:         "广东省",
		Country:       "中国",
		ZipCode:       "510006",
	}
	req := &checkout.CheckoutReq{
		UserId:     1,
		Firstname:  "feng",
		Lastname:   "hu",
		Email:      "3531095171@qq.com",
		Address:    deliveryAddress,
		CreditCard: cardInfo,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
