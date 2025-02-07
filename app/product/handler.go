package main

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/service"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// CreateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp, err = service.NewCreateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *common.Empty, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// UpdateProductInfo implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProductInfo(ctx context.Context, req *product.UpdateProductInfoReq) (resp *common.Empty, err error) {
	resp, err = service.NewUpdateProductInfoService(ctx).Run(req)

	return resp, err
}
