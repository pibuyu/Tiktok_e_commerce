package service

import (
	"context"
	"errors"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/model"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserInfoService struct {
	ctx context.Context
} // NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *common.Empty, err error) {
	//1.基础校验
	if req.Email == "" || req.OldPassword == "" || req.NewPassword == "" {
		return nil, errors.New("email or password is nil")
	}
	//2.比对输入的旧密码与原始密码
	u, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.OldPassword))
	if err != nil {
		return nil, errors.New("incorrect old password")
	}

	//3.更新密码
	newHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	err = model.UpdatePassword(mysql.DB, req.Email, string(newHashedPassword))
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}
