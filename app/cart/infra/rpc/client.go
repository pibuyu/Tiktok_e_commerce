package rpc

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/conf"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/utils"
	"github.com/Blue-Berrys/Tiktok_e_commerce/common/clientsuite"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	serviceName   string
	registryAddr  string
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	//utils.MustHandleError(err)
	//opts = append(opts, client.WithResolver(r))

	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)
}
