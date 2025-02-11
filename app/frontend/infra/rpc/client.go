package rpc

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/conf"
	frontUtils "github.com/Blue-Berrys/Tiktok_e_commerce/app/frontend/utils"
	"github.com/Blue-Berrys/Tiktok_e_commerce/common/clientsuite"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	err            error

	ServiceName  = frontUtils.ServiceName
	MetricsPort  = conf.GetConf().Hertz.MetricsPort
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient() //不要忘记初始化各个client.忘记初始化checkoutClient导致结算完无法跳转到waiting页面
		initOrderClient()
	})
}

func initUserClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontUtils.MustHandleError(err)
	//
	//opts = append(opts, client.WithResolver(r))
	//
	//UserClient, err = userservice.NewClient("user", opts...)

	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontUtils.MustHandleError(err)
}

func initProductClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontUtils.MustHandleError(err)
	//
	//opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontUtils.MustHandleError(err)
}

func initCartClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontUtils.MustHandleError(err)
	//
	//opts = append(opts, client.WithResolver(r))

	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontUtils.MustHandleError(err)
}

func initCheckoutClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontUtils.MustHandleError(err)
	//
	//opts = append(opts, client.WithResolver(r))
	//
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontUtils.MustHandleError(err)
}

func initOrderClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	//frontUtils.MustHandleError(err)
	//opts = append(opts, client.WithResolver(r))

	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontUtils.MustHandleError(err)
}
