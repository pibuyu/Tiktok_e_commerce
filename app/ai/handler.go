package main

import (
	"context"
	ai "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// QueryOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QueryOrder(ctx context.Context, req *ai.OrderQueryRequest) (resp *ai.OrderQueryResponse, err error) {
	resp, err = service.NewQueryOrderService(ctx).Run(req)

	return resp, err
}

// AutoOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) AutoOrder(ctx context.Context, req *ai.AutoOrderRequest) (resp *ai.AutoOrderResponse, err error) {
	resp, err = service.NewAutoOrderService(ctx).Run(req)

	return resp, err
}
