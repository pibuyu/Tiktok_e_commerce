package service

import (
	"context"
	"errors"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal/mysql"
	cart "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gorm.io/gorm"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	list, err := model.GetCartBtUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}

	var items []*cart.CartItem
	if len(list) > 0 {
		for _, productItem := range list {
			items = append(items, &cart.CartItem{
				ProductId: productItem.ProductId,
				Quantity:  int32(productItem.Qty),
			})
		}
	}

	return &cart.GetCartResp{Cart: &cart.Cart{
		UserId: req.UserId,
		Items:  items,
	}}, nil
}
