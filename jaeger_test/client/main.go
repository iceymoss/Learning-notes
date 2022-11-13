package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"StudyGin/jaeger_test/otgrpc"
	"StudyGin/jaeger_test/proto"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	timepb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	//日志追踪初始化配置
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: "mxshop",
	}

	//初始化跟踪连
	Tracer, Closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}

	//将tracer设置为全局
	opentracing.SetGlobalTracer(Tracer)
	defer Closer.Close()

	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		log.Panicln("连接失败", err)
	}
	defer clientConn.Close()
	client := proto.NewGreeterClient(clientConn)
	res, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name:   "kuangyang",
		Url:    "learinku.com",
		Gender: proto.Gender_MALE,
		M: map[string]string{
			"来自": "贵州",
			"现居": "无锡",
		},
		AddTime: timepb.New(time.Now()),
	})

	if err != nil {
		panic(err)
	}
	//pong, err := client.Ping(context.Background(), &proto.IsEmpty{})
	//if err != nil {
	//	panic(err)
	//}
	fmt.Printf("返回结果: %v", res)
}
