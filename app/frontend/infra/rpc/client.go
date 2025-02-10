package rpc

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/conf"
	frontUtils "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/utils"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient() //不要忘记初始化各个client.忘记初始化checkoutClient导致结算完无法跳转到waiting页面
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

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)

	frontUtils.MustHandleError(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)

	frontUtils.MustHandleError(err)
}
