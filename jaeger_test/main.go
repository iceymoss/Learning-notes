package main

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	//初始化配置
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

	//初始化跟踪链
	Tracer, Closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}

	//将跟踪连设置为项目全局变量
	opentracing.SetGlobalTracer(Tracer)

	//span1 := opentracing.StartSpan("funcA")
	//time.Sleep(time.Second * 1)
	//span1.Finish()

	defer Closer.Close()

	span1 := Tracer.StartSpan("funcA") //开始追踪
	time.Sleep(time.Second * 1)
	span1.Finish() //完成追踪

	span2 := Tracer.StartSpan("funcC") //开始追踪
	time.Sleep(time.Second * 2)
	span2.Finish() //完成追踪
}
