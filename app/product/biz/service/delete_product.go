package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/model"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	product "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *common.Empty, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	productCategoryQuery := model.NewProductCategoryQuery(s.ctx, mysql.DB)
	//1.先删除关联表product_category
	err = productCategoryQuery.DeleteByProductId(int(req.Id))
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "delete product_category failed:"+err.Error())
	}
	//2.然后删除product表项
	err = productQuery.DeleteProduct(int(req.Id))
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "删除商品failed:"+err.Error())
	}
	return &common.Empty{}, nil
}
