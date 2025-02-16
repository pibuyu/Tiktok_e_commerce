package order

import (
	"context"
	"fmt"

	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/biz/service"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/biz/utils"
	common "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	fmt.Println("返回的数据为：\n", resp)

	c.HTML(consts.StatusOK, "order", utils.WarpResponse(ctx, c, resp))
}
