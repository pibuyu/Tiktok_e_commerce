package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/biz/dal"
	payment "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	orderid, _ := uuid.NewRandom()
	req := &payment.ChargeReq{
		Amount:  10.00,
		OrderId: orderid.String(),
		UserId:  1,
		CreditCard: &payment.CreditCardInfo{
			//CreditCardNumber及CreditCardCvv需要符合Luhn算法,需要注意:过期时间需要在当前时间之后
			CreditCardNumber:          "4111111111111111",
			CreditCardExpirationYear:  2030,
			CreditCardExpirationMonth: 02,
			CreditCardCvv:             123,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
