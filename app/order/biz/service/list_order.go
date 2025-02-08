package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"strconv"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Cost: oi.Cost,
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
			})
		}
		zipCode, _ := strconv.ParseInt(v.Consignee.ZipCode, 10, 32)
		o := &order.Order{
			OrderItems:   nil,
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       int32(zipCode),
			},
			CreatedAt: int32(v.CreatedAt.Unix()),
		}
		orders = append(orders, o)
	}

	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
