package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
)

func TestUpdateUserInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateUserInfoService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	//init database
	_ = godotenv.Load()
	dal.Init()

	req := &user.UpdateUserInfoReq{
		Email:       "zhang@qq.com",
		OldPassword: "12345678",
		NewPassword: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
