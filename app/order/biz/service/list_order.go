package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
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
		v.OrderItems, _ = model.GetOrderItem(s.ctx, mysql.DB, v.OrderId)

		fmt.Println(v.OrderItems)
		for _, oi := range v.OrderItems {
			product, _ := model.ProductMapper.GetById(int(oi.ProductId))
			items = append(items, &order.OrderItem{
				Cost: oi.Cost,
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				ProductName: product.Name,
				Qty:         oi.Quantity,
				Picture:     product.Picture,
			})
		}
		o := &order.Order{
			OrderItems: items,
			OrderId:    v.OrderId,
			UserId:     v.UserId,
			Email:      v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
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
