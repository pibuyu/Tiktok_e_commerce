package service

import (
	"context"
	"errors"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/model"
	user "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	//基础校验
	if req.Password == "" || req.Email == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is nil")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password does not match with confirmPassword")
	}
	//加密并存储
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("bcrypt password err:" + err.Error())
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(bcryptedPassword),
	}
	err = model.CreateUser(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{
		UserId: int32(newUser.ID),
	}, nil
}
