package main

import (
	"context"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal"
	"github.com/Blue-Berrys/Tiktok_e_commerce/common/mtl"
	"github.com/joho/godotenv"
	consul "github.com/kitex-contrib/registry-consul"
	"net"
	"time"

	_ "github.com/Blue-Berrys/Tiktok_e_commerce/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/user/conf"
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
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
	opts := kitexInit()

	//init metrics.注意需要放在init dal和rpc之前,后者可能依赖前者
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)

	//init tracing
	p := mtl.InitTracing(ServiceName)
	defer func() { //退出前上传剩余链路数据
		_ = p.Shutdown(context.Background())
	}()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	//init mysql & redis
	//初始化总是失败记得检查下是否读取了环境变量!!!
	_ = godotenv.Load()
	dal.Init()

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
	opts = append(opts, server.WithServiceAddr(addr))

	//registry info
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
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
