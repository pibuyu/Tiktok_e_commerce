package rpc

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/conf"
	frontUtils "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/utils"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)

	frontUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)

	frontUtils.MustHandleError(err)
}
