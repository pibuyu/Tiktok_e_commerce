// Code generated by Kitex v0.9.1. DO NOT EDIT.
package paymentservice

import (
	payment "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler payment.PaymentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler payment.PaymentService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
