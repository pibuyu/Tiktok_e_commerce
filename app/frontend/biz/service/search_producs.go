package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/infra/rpc"
	rpcProduct "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProducsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProducsService(Context context.Context, RequestContext *app.RequestContext) *SearchProducsService {
	return &SearchProducsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProducsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	searchResp, err := rpc.ProductClient.SearchProducts(h.Context, &rpcProduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"q":     req.Q,
		"items": searchResp.Results,
	}, nil
}
