package rpc

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/checkout/conf"
	"github.com/Blue-Berrys/Tiktok_e_commerce/common/clientsuite"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	err           error
	serviceName   string
	registryAddr  string
	CommonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		CommonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})
		initCartClient()
		initProductClient()
		initPaymentClient()
	})
}

// todo:下面的几个initClient应该需要修改为clientSuite方式，待定
func initCartClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	//if err != nil {
	//	panic(err)
	//}
	//opts = append(opts, client.WithResolver(r))

	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	CartClient, err = cartservice.NewClient("cart", CommonSuite)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	//if err != nil {
	//	panic(err)
	//}
	//
	//opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	ProductClient, err = productcatalogservice.NewClient("product", CommonSuite)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	//var opts []client.Option
	//r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	//if err != nil {
	//	panic(err)
	//}
	//
	//opts = append(opts, client.WithResolver(r))
	//opts = append(opts,
	//	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	//	client.WithTransportProtocol(transport.GRPC),
	//	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	//)
	PaymentClient, err = paymentservice.NewClient("payment", CommonSuite)
	if err != nil {
		panic(err)
	}
}
