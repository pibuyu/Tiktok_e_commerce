package service

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
	"gorm.io/gorm"
)

type AutoOrderService struct {
	ctx context.Context
} // NewAutoOrderService new AutoOrderService
func NewAutoOrderService(ctx context.Context) *AutoOrderService {
	return &AutoOrderService{ctx: ctx}
}

// Run create note info
func (s *AutoOrderService) Run(req *ai.AutoOrderRequest) (resp *ai.AutoOrderResponse, err error) {
	// Finish your business logic.
	//	text := `-- 插入订单中Notebook的产品记录
	//INSERT INTO order_item (created_at, updated_at, product_id, order_id_refer, quantity, cost)
	//VALUES (CURRENT_TIMESTAMP(3), CURRENT_TIMESTAMP(3),
	//        (SELECT id FROM product WHERE name = 'vivo x200'),
	//        '1111', 2,
	//        (SELECT price FROM product WHERE name = 'vivo x200') * 2)`
	//
	//	text = strings.ReplaceAll(text, "[]", "`")

	user, _ := model.GetById(req.UserId)
	sqlStrings, state := util.GenAutoOrderSQL(req.Message, user)

	if state == 1 {
		return &ai.AutoOrderResponse{
			Data: " 请提供收货地址，包括street_address、city、state、country、zip_code等信息，目前缺少收货地址相关信息，无法完成自动下单",
		}, nil
	}

	if err = executeMultipleSQL(sqlStrings); err != nil || state == 0 {
		resp = &ai.AutoOrderResponse{
			Data: "自动下单失败，请稍后重试",
		}
		return resp, err
	}

	resp = &ai.AutoOrderResponse{
		Data: "自动下单成功，可以点击下面链接查看所有订单：\n http://localhost:8080/order",
	}

	return resp, nil

}

// 在事务中运行所有sql语句
func executeMultipleSQL(sqlStatments []string) error {
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		for _, sql := range sqlStatments {
			if result := tx.Exec(sql); result.Error != nil {
				return result.Error // 返回错误会出发回滚
			}
		}
		return nil
	})
}
