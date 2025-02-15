package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/infra/rpc"
	rpcProduct "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	category "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	listResp, err := rpc.ProductClient.ListProducts(h.Context, &rpcProduct.ListProductsReq{
		CategoryName: req.Category,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": listResp.Products,
	}, nil
}
