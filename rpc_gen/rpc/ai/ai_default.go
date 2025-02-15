package ai

import (
	"context"
	ai "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func QueryOrder(ctx context.Context, req *ai.OrderQueryRequest, callOptions ...callopt.Option) (resp *ai.OrderQueryResponse, err error) {
	resp, err = defaultClient.QueryOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AutoOrder(ctx context.Context, req *ai.AutoOrderRequest, callOptions ...callopt.Option) (resp *ai.AutoOrderResponse, err error) {
	resp, err = defaultClient.AutoOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AutoOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
