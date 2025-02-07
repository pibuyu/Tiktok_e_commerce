package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/product/biz/dal"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	product "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product"
)

func TestUpdateProductInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateProductInfoService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	_ = godotenv.Load()
	dal.Init()

	req := &product.UpdateProductInfoReq{
		Id:          13,
		Name:        "vivo x200 手机",
		Description: "vivo X200 Pro 12GB+256GB 宝石蓝 国家补贴 蔡司2亿APO超级长焦 蓝晶×天玑9400 6000mAh蓝海电池手机",
		Categories:  []string{"vivo phone"},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
