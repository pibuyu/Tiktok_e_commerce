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

func TestCreateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateProductService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	req := &product.CreateProductReq{
		Name:        "vivo手机",
		Description: "vivo X200 Pro 12GB+256GB 宝石蓝 国家补贴 蔡司2亿APO超级长焦 蓝晶×天玑9400 6000mAh蓝海电池手机",
		Picture:     "https://img14.360buyimg.com/n1/s546x546_jfs/t1/256968/38/17274/105817/67a33a49F3c311c7b/ea839d1023887a7b.jpg",
		Price:       5149.00,
		Categories:  []string{"phone"},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
