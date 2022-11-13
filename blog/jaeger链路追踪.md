[toc]

### 文章介绍

这里将来介绍什么是链路追踪，为什么需要链路追踪， 用jaeger做服务链路追踪

### 什么是链路追踪

#### 链路 

这里的链路指的是客户端向服务发起一个请求，该请求所经过的路线，也可以说是该请求经过的流量

例如: 客户端发起一个下订单的请求其流量过程：

```shell
request—>service—>order-web—>order_srv—>mysql—>order_srv—>order-web—>service—>response
```

这就一个请求的完整链路



#### 链路追踪

指我们通过一些手段将链路进行监控， 对于系统调试和维护链路追踪是非常重要的，尤其微服务中，我们知道各

个微服务部署在不同的服务器上，并且每一个微服务可能是不同的人开发的，如果我们不做链路追踪，微服务之

间相互调用，假如有的微服务出问题了，整个系统都会受影响，那么我们怎么知道是哪一个微服务出的问题，找

谁维护等一系列问题。



### Jaeger的安装(docker)

```shell
docker run \
  --rm \
  --name jaeger \
  -p6831:6831/udp \
  -p16686:16686 \
  jaegertracing/all-in-one:latest
```

启动成功后可以访问：localhost:16686

架构：

https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/xianliu/7.png

### Jaeger的快速使用

使用jaeger就行链路追踪

```go
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
	
	span1 := Tracer.StartSpan("funcA")			//开始追踪
	time.Sleep(time.Second * 1)
	span1.Finish()			//完成追踪

	span2 := Tracer.StartSpan("funcC")         //开始追踪
	time.Sleep(time.Second * 2)
	span2.Finish()			//完成追踪
}
```

我们在ui界面可以看到

两个Traces即：funcC和funcA,里面有运行的时间等信息

<img src="/Users/feng/Desktop/截屏2022-11-13 下午8.36.40.png" alt="截屏2022-11-13 下午8.36.40" style="zoom:50%;" />



funcC:

<img src="/Users/feng/Desktop/截屏2022-11-13 下午8.36.53.png" alt="截屏2022-11-13 下午8.36.53" style="zoom:50%;" />

层级关系：

```go
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

```

可以看到

main: 

​		funcD:---------1s---------

​											 funcE ----------------------------2s---------------------

![截屏2022-11-13 下午8.55.53](/Users/feng/Desktop/截屏2022-11-13 下午8.55.53.png)



### 将jaeger集成到grpc服务中

我们需要向获取Jaeger-grpc的代码：https://github.com/iceymoss/Learning-notes 的jaeger_test中

这里我们对Grpc不做介绍了，如果您不了解Grpc您需要先了解[Grpc详解](https://learnku.com/articles/68089)

Server:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"StudyGin/jaeger_test/proto"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type HelloSerivce struct{}

func (h *HelloSerivce) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	//接收context中的内容
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		fmt.Println("get metadata err", ok)
	}
	for key, value := range md {
		fmt.Printf("%v: %v\n", key, value)
	}

	return &proto.HelloReply{
		Id:      "123456789",
		Request: req,
	}, nil
}
func (h *HelloSerivce) Ping(context.Context, *proto.IsEmpty) (*proto.Pong, error) {
	return nil, nil
}

func main() {
	//实例化server
	lit, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicln("监听失败", err)
	}
	//注册处理逻辑
	//NewServer创建一个未注册服务且尚未开始接受请求的 gRPC 服务器。
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &HelloSerivce{})
	log.Println(s.Serve(lit))
}

```

Client:

```go
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
			"来自": "无锡",
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

```

我们可以到jaeger的ui界面查看本次调用的链路数据

