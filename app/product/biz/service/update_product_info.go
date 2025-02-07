package service

import (
	"context"
	"errors"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/model"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	product "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type UpdateProductInfoService struct {
	ctx context.Context
} // NewUpdateProductInfoService new UpdateProductInfoService
func NewUpdateProductInfoService(ctx context.Context) *UpdateProductInfoService {
	return &UpdateProductInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductInfoService) Run(req *product.UpdateProductInfoReq) (resp *common.Empty, err error) {
	//根据req构造出商品
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "product id is required")
	}
	//尝试查询到具体的商品
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	item, err := productQuery.GetById(int(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not exists")
		}
		return nil, kerrors.NewGRPCBizStatusError(2004002, "query product failed:"+err.Error())
	}
	//然后根据req中的新属性更新item，并重新插入到product表
	if len(req.Name) > 0 {
		item.Name = req.Name
	}
	if len(req.Description) > 0 {
		item.Description = req.Description
	}
	if len(req.Picture) > 0 {
		item.Picture = req.Picture
	}
	if req.Price != 0 {
		item.Price = req.Price
	}
	if len(req.Categories) > 0 {
		//更新product所属类别
		//1.先删除旧的product_category关联项
		productCategoryQuery := model.NewProductCategoryQuery(s.ctx, mysql.DB)
		err = productCategoryQuery.DeleteByProductId(item.ID)
		if err != nil {
			return nil, err
		}
		//2.插入新的product_category关联项
		var newCategories []model.Category
		categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
		for _, categoryName := range req.Categories {
			//根据每个类别的name查到对应的category
			category, err := categoryQuery.GetCategoryById(categoryName)
			if err != nil {
				return nil, err
			}
			if category.ID == 0 {
				//新建并再次查询
				klog.Info("category is nil,create new category")
				err = categoryQuery.CreateCategory(categoryName)
				if err != nil {
					return nil, err
				}
				category, err = categoryQuery.GetCategoryById(categoryName)
				if err != nil {
					return nil, err
				}
			}
			newCategories = append(newCategories, *category)
		}
		item.Categories = newCategories
	}
	//回写
	err = productQuery.UpdateProductInfo(&item)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "update product info failed:"+err.Error())
	}
	return &common.Empty{}, nil
}
