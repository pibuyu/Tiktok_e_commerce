package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal"
	product "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteProductService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	//待删除的商品id应该先确认存在于product表中
	req := &product.DeleteProductReq{
		Id: 12,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	
}
