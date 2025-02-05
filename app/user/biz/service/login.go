package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/model"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	//output to console
	klog.Infof("LoginReq:%+v", req)
	resp = &user.LoginResp{}

	//validate username
	userModel := &model.User{}
	if exist := userModel.IsExistByField("email", req.Email); !exist {
		resp.UserId = -1
		return resp, fmt.Errorf("invalid email")
	}

	//check password
	if correctPwd := userModel.ValidatePassword(req.Email, req.Password); !correctPwd {
		resp.UserId = -1
		return resp, fmt.Errorf("invalid password")
	}

	resp.UserId = int32(userModel.ID)
	return resp, nil
}
