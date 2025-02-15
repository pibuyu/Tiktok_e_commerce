package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/hertz_gen/frontend/product"
	rpcProduct "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	getProductResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcProduct.GetProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	//这里的item字段名不要写错了.否则前端无法正确渲染商品详情页面数据
	return utils.H{
		"item": getProductResp.Product,
	}, nil
}
