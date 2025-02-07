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

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	req := &product.SearchProductsReq{
		Query: "vivo",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
