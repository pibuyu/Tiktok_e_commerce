package service

import (
	"context"
	payment "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		Amount:  10.00,
		OrderId: "1111",
		UserId:  1,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1111",
			CreditCardExpirationYear:  2025,
			CreditCardExpirationMonth: 11,
			CreditCardCvv:             1212,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
