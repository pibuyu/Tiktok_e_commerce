package category

import (
	"context"

	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/biz/service"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/biz/utils"
	category "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "category", resp)
}
