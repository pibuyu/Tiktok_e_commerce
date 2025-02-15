package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
)

type AutoOrderService struct {
	ctx context.Context
} // NewAutoOrderService new AutoOrderService
func NewAutoOrderService(ctx context.Context) *AutoOrderService {
	return &AutoOrderService{ctx: ctx}
}

// Run create note info
func (s *AutoOrderService) Run(req *ai.AutoOrderRequest) (resp *ai.AutoOrderResponse, err error) {
	// Finish your business logic.

	return
}
