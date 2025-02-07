package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	//根据req.categories(类别名称)找到商品所属的所有类别
	var categories []model.Category
	for _, categoryName := range req.Categories {
		//根据每个类别的name查到对应的category
		category, err := categoryQuery.GetCategoryById(categoryName)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *category)
	}
	//构造商品
	productItem := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories,
	}
	err = productQuery.CreateProduct(productItem)

	if err != nil {
		return nil, err
	}
	return &product.CreateProductResp{Id: 1}, nil
}
