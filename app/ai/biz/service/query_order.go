package service

import (
	"context"
	ai "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
)

type QueryOrderService struct {
	ctx context.Context
} // NewQueryOrderService new QueryOrderService
func NewQueryOrderService(ctx context.Context) *QueryOrderService {
	return &QueryOrderService{ctx: ctx}
}

// Run create note info
func (s *QueryOrderService) Run(req *ai.OrderQueryRequest) (resp *ai.OrderQueryResponse, err error) {
	// Finish your business logic.

	return
}
