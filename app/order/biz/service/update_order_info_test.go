package service

import (
	"context"
	"testing"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
)

func TestUpdateOrderInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOrderInfoService(ctx)
	// init req and assert value

	req := &order.UpdateOrderInfoReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
