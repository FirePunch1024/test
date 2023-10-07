package pay

import (
	"context"
	pay "pay/kitex_gen/pay"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"pay/kitex_gen/pay/payservice"
)

type RPCClient interface {
	KitexClient() payservice.Client
	Service() string
	BuyGoods(ctx context.Context, request *pay.BuyGoodsReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error)

	CallBack(ctx context.Context, request *pay.CallBackReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error)

	Notify(ctx context.Context, request *pay.NotifyReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := payservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient payservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() payservice.Client {
	return c.kitexClient
}

func (c *clientImpl) BuyGoods(ctx context.Context, request *pay.BuyGoodsReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	return c.kitexClient.BuyGoods(ctx, request, callOptions...)
}

func (c *clientImpl) CallBack(ctx context.Context, request *pay.CallBackReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	return c.kitexClient.CallBack(ctx, request, callOptions...)
}

func (c *clientImpl) Notify(ctx context.Context, request *pay.NotifyReq, callOptions ...callopt.Option) (resp *pay.NilResponse, err error) {
	return c.kitexClient.Notify(ctx, request, callOptions...)
}
