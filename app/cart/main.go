package main

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/biz/dal"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/infra/rpc"
	"github.com/Blue-Berrys/Tiktok_e_commerce/common/mtl"
	"github.com/Blue-Berrys/Tiktok_e_commerce/common/serversuite"
	"github.com/joho/godotenv"
	"net"
	"time"

	"github.com/Blue-Berrys/Tiktok_e_commerce/app/cart/conf"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	//init metrics.注意需要放在init dal和rpc之前,后者可能依赖前者
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)

	//init tracing
	p := mtl.InitTracing(ServiceName)
	defer func() { //退出前上传剩余链路数据
		_ = p.Shutdown(context.Background())
	}()

	//init database
	_ = godotenv.Load()
	dal.Init()

	//init rpc client
	rpc.InitClient()

	opts := kitexInit()
	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}

	//server info,整合 service registry
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
