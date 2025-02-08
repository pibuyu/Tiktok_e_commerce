package service

import (
	"context"
	"errors"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/mysql"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// 基础校验
	if req.UserId == 0 || req.OrderId == "" {
		err = errors.New("user_id or order_id is required")
		return
	}

	//更新订单状态
	_, err = model.GetOrder(s.ctx, mysql.DB, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.GetOrder.err:%v", err)
		return nil, err
	}
	err = model.UpdateOrderState(s.ctx, mysql.DB, req.UserId, req.OrderId, model.OrderStatePaid)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	resp = &order.MarkOrderPaidResp{}
	return
}
