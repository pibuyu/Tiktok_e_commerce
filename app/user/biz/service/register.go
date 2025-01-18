package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/consts"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/utils"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// init resp
	resp = &user.RegisterResp{}

	// validate params
	if !utils.VerifyEmailFormat(req.Email) {
		return nil, errors.New("invalid format of email or password")
	}

	// check if email is repeat
	userModel := &model.User{}
	if userModel.IsExistByField("email", req.Email) {
		resp.UserId = -1
		return resp, errors.New("existing email")
	}

	// register user info into database
	registerUserStruct := &model.User{
		Email:    req.Email,
		Password: utils.EncodePassword(req.Email, req.Password),
		Salt:     fmt.Sprintf("%s%s", consts.SECRET_SALT, req.Email),
	}
	if success := registerUserStruct.Create(); !success {
		resp.UserId = -1
		return resp, fmt.Errorf("failed to register due to : %v", err)
	}

	resp.UserId = int32(registerUserStruct.ID)
	return resp, nil
}
