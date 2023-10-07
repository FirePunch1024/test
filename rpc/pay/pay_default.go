package pay

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	pay "pay/kitex_gen/pay"
)

func BuyGoods(ctx context.Context, request *pay.BuyGoodsReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	resp, err = defaultClient.BuyGoods(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BuyGoods call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CallBack(ctx context.Context, request *pay.CallBackReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	resp, err = defaultClient.CallBack(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CallBack call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Notify(ctx context.Context, request *pay.NotifyReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	resp, err = defaultClient.Notify(ctx, request, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Notify call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
