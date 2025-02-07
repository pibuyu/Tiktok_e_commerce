package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/infra/rpc"
	cart "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	//init productClient.单元测试时rpc.ProductClient还未被初始化，因此需要判断
	if rpc.ProductClient == nil {
		rpc.InitClient()
	}
	//使用productClient的GetProduct方法先获取商品信息
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})

	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(40004, "product not found")
	}

	//添加商品到购物车
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
