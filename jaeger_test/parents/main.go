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
			LocalAgentHostPort: "10.2.118.0:6831",
		},
		ServiceName: "test",
	}

	//初始化跟踪连
	Tracer, Closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	defer Closer.Close()

	Parents := Tracer.StartSpan("main")
	span1 := Tracer.StartSpan("funcD", opentracing.ChildOf(Parents.Context()))
	time.Sleep(time.Second * 1)
	span1.Finish()

	span2 := Tracer.StartSpan("funcE", opentracing.ChildOf(Parents.Context()))
	time.Sleep(time.Second * 2)
	span2.Finish()
}
