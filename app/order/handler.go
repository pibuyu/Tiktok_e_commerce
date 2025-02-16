package main

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/service"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// UpdateOrderInfo implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderInfo(ctx context.Context, req *order.UpdateOrderInfoReq) (resp *common.Empty, err error) {
	resp, err = service.NewUpdateOrderInfoService(ctx).Run(req)

	return resp, err
}
