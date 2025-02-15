package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/model"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/common"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *common.Empty, err error) {
	if req.UserId == 0 {
		return
	}
	err = model.DeleteUserById(mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	return &common.Empty{}, nil
}
