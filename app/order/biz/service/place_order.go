package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/mysql"
	order "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp = &order.PlaceOrderResp{}
	// 基础校验
	if len(req.OrderItems) == 0 {
		err = kerrors.NewGRPCBizStatusError(500001, "items is empty")
		return
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		//1.创建order
		orderId, _ := uuid.NewRandom()

		o := &model.Order{
			OrderId:      orderId.String(),
			OrderState:   model.OrderStatePlaced,
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		//填充收货地址信息
		if req.Address != nil {
			addr := req.Address
			o.Consignee.StreetAddress = addr.StreetAddress
			o.Consignee.City = addr.City
			o.Consignee.State = addr.State
			o.Consignee.Country = addr.Country
			o.Consignee.ZipCode = addr.ZipCode
		}
		if err = tx.Debug().Create(o).Error; err != nil {
			klog.Errorf("插入订单err:%v", err)
			return err
		}

		//2.创建order与item的映射
		var items []*model.OrderItem
		for _, v := range req.OrderItems {
			items = append(items, &model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
		}
		if err := tx.Debug().Create(&items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		//事务结束
		return nil
	})

	return
}
