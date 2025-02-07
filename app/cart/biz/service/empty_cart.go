package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal/mysql"
	cart "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	//需要注意：cart里用了gorm.Model，自带deleted_at字段，因此是逻辑删除
	err = model.EmptyCart(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}
