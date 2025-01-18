package service

import (
	"context"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
	"log"
	"os"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx)

	//chdir to adjust config file path
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("chdir err : %v", err)
	}

	// init req and assert value
	req := &user.LoginReq{
		Email:    "123456789@qq.com",
		Password: "123456abcde",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
