package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/model"
	product "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)

	//遍历列表
	for _, v1 := range c {
		//遍历列表的商品
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
		}
	}
	return resp, nil
}
