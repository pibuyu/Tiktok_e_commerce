package service

import (
	"context"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
)

type UpdateOrderInfoService struct {
	ctx context.Context
} // NewUpdateOrderInfoService new UpdateOrderInfoService
func NewUpdateOrderInfoService(ctx context.Context) *UpdateOrderInfoService {
	return &UpdateOrderInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderInfoService) Run(req *order.UpdateOrderInfoReq) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
