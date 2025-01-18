package main

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/biz/service"
	payment "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)
	//todo:1.首先需要在本地生成预支付流水单
	return resp, err
}
