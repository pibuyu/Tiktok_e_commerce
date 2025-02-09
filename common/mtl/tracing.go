package mtl

import "github.com/kitex-contrib/obs-opentelemetry/provider"

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		//关闭opentelemetry的指标功能
		provider.WithEnableMetrics(false),
	)
	return p
}
