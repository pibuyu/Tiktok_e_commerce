package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/infra/rpc"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//var resp = make(map[string]any)
	//items := []map[string]any{
	//	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
	//	{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.jpeg"},
	//	{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.jpeg"},
	//	{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/notebook.jpeg"},
	//	{"Name": "T-shirt-5", "Price": 140, "Picture": "/static/image/t-shirt-1.jpeg"},
	//	{"Name": "T-shirt-6", "Price": 150, "Picture": "/static/image/t-shirt.jpeg"},
	//}
	//resp["title"] = "Hot Sales"
	//resp["items"] = items

	//查询数据库中的结果在首页进行展示
	homeProducts, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{"title": "Hot Sale", "items": homeProducts.Products}, nil
}
