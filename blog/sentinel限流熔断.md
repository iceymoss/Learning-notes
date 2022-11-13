[toc]



### 文章介绍

本文将介绍什么是限流和熔断， alibaba开源的Sentinel、安装、实战



### 什么是限流\熔断\降级

**限流**：在我们的后天系统中，如果那一天突然进入大量流量，我们服务原本最高只能处理同时2k的请求，突然一

下就来来了5k的请求，这对服务器的压力是不是很要命，这很可能直接导致服务器宕机，崩溃，导致原本2K的处

理量都不能处理了，这时候我们需要限流，限流的作用就是保持访问量到达服务器最高的情况下，对多余的请求

不做处理，相比之下，比服务器直接挂掉是好很多的。例如在双十一的时候，我们要下单就会看到类似"请求繁

忙，请稍后重试！"。

**熔断**:相信大家对断路器并不陌生，它就相当于一个开关，打开后可以阻止流量通过。比如保险丝，当电流过大

时，就会熔断，从而避免元器件损坏。

服务熔断是指调用方访问服务时通过断路器做代理进行访问，断路器会持续观察服务返回的成功、失败的状态，

当失败超过设置的阈值时断路器打开，请求就不能真正地访问到服务了。

**使用场景**

- 服务故障或者升级时，让客户端快速失败
- 失败处理逻辑容易定义
- 响应耗时较长，客户端设置的`read timeout`会比较长，防止客户端大量重试请求导致的连接、线程资源不能释放

**降级 **:服务降级是从整个系统的负荷情况出发和考虑的，对某些负荷会比较高的情况，为了预防某些功能（业务

场景）出现负荷过载或者响应慢的情况，在其内部暂时舍弃对一些非核心的接口和数据的请求，而直接返回一个

提前准备好的fallback（退路）错误处理信息。这样，虽然提供的是一个有损的服务，但却保证了整个系统的稳

定性和可用性。

### 什么是Sentinel

Sentinel是阿里开源的项目，提供了流量控制、熔断降级、系统负载保护等多个维度来保障服务之间的稳定性。
官网：https://github.com/alibaba/Sentinel/wiki

2012年，Sentinel诞生于阿里巴巴，其主要目标是流量控制。2013-2017年，Sentinel迅速发展，并成为阿里巴巴所有微服务的基本组成部分。 它已在6000多个应用程序中使用，涵盖了几乎所有核心电子商务场景。2018年，Sentinel演变为一个开源项目。2020年，Sentinel Golang发布。

**特点** ：

**丰富的应用场景** ：Sentinel 承接了阿里巴巴近 10 年的双十一大促流量的核心场景，例如秒杀（即
突发流量控制在系统容量可以承受的范围）、消息削峰填谷、集群流量控制、实时熔断下游不可用应用等。
**完备的实时监控** ：Sentinel 同时提供实时的监控功能。您可以在控制台中看到接入应用的单台机
器秒级数据，甚至 500 台以下规模的集群的汇总运行情况。

**生态广广泛 ** 



### Sentinel 的历史

- 2012 年，Sentinel 诞生，主要功能为入口流量控制。
- 2013-2017 年，Sentinel 在阿里巴巴集团内部迅速发展，成为基础技术模块，覆盖了所有的核心场景。Sentinel 也因此积累了大量的流量归整场景以及生产实践。
- 2018 年，Sentinel 开源，并持续演进。
- 2019 年，Sentinel 朝着多语言扩展的方向不断探索，推出 [C++ 原生版本](https://github.com/alibaba/sentinel-cpp)，同时针对 Service Mesh 场景也推出了 [Envoy 集群流量控制支持](https://github.com/alibaba/Sentinel/tree/master/sentinel-cluster/sentinel-cluster-server-envoy-rls)，以解决 Service Mesh 架构下多语言限流的问题。
- 2020 年，推出 [Sentinel Go 版本](https://github.com/alibaba/sentinel-golang)，继续朝着云原生方向演进。
- 2021 年，Sentinel 正在朝着 2.0 云原生高可用决策中心组件进行演进；同时推出了 [Sentinel Rust 原生版本](https://github.com/sentinel-group/sentinel-rust)。同时我们也在 Rust 社区进行了 Envoy WASM extension 及 eBPF extension 等场景探索。
- 2022 年，Sentinel 品牌升级为流量治理，领域涵盖流量路由/调度、流量染色、流控降级、过载保护/实例摘除等；同时社区将流量治理相关标准抽出到 [OpenSergo 标准](https://opensergo.io/)中，Sentinel 作为流量治理标准实现。



### Sentinel-go的安装

[Sentinel-go开源地址](https://github.com/alibaba/sentinel-golang)

[官网文档](https://sentinelguard.io/zh-cn/docs/golang/quick-start.html)

安装：```go get github.com/alibaba/sentinel-golang/api```





### Go限流实战

#### qps限流

```go
package main

import (
	"fmt"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func main() {
	//基于sentinel的qps限流
	//必须初始化
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}

	//配置限流规则：1秒内通过10次
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some_test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //超过直接拒绝
			Threshold:              10,          //请求次数
			StatIntervalInMs:       1000,        //允许时间内
		},
	})

	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}

	for i := 0; i < 12; i++ {
		e, b := sentinel.Entry("some_test", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			fmt.Println("限流了")
		} else {
			fmt.Println("检查通过")
			e.Exit()
		}
	}
}
```

打印结果：

```go
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
限流了
限流了
```

#### Thrnotting

```go
package main

import (
	"fmt"
	"log"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func main() {
	//基于sentinel的qps限流
	//必须初始化
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}

	//配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some_test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Throttling, //匀速通过
			Threshold:              10,              //请求次数
			StatIntervalInMs:       1000,            //允许时间内
		},
	})

	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}

	for i := 0; i < 12; i++ {
		e, b := sentinel.Entry("some_test", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			fmt.Println("限流了")
		} else {
			fmt.Println("检查通过")
			e.Exit()
		}
		time.Sleep(time.Millisecond * 100)
	}
}
```

```
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
检查通过
```



#### Warrm_up

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func main() {
	//先初始化sentinel
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("初始化sentinel 异常: %v", err)
	}

	var globalTotal int
	var passTotal int
	var blockTotal int
	ch := make(chan struct{})

	//配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			TokenCalculateStrategy: flow.WarmUp, //冷启动策略
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              1000,
			WarmUpPeriodSec:        30,
		},
	})

	if err != nil {
		log.Fatalf("加载规则失败: %v", err)
	}

	//我会在每一秒统计一次，这一秒只能 你通过了多少，总共有多少， block了多少, 每一秒会产生很多的block
	for i := 0; i < 100; i++ {
		go func() {
			for {
				globalTotal++
				e, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
				if b != nil {
					//fmt.Println("限流了")
					blockTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					passTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					e.Exit()
				}
			}
		}()
	}

	go func() {
		var oldTotal int //过去1s总共有多少个
		var oldPass int  //过去1s总共pass多少个
		var oldBlock int //过去1s总共block多少个
		for {
			oneSecondTotal := globalTotal - oldTotal
			oldTotal = globalTotal

			oneSecondPass := passTotal - oldPass
			oldPass = passTotal

			oneSecondBlock := blockTotal - oldBlock
			oldBlock = blockTotal

			time.Sleep(time.Second)
			fmt.Printf("total:%d, pass:%d, block:%d\n", oneSecondTotal, oneSecondPass, oneSecondBlock)
		}
	}()

	<-ch
}
```

打印结果：逐渐到达1k,在1k位置上下波动

```sh
total:11, pass:9, block:0
total:21966, pass:488, block:21420
total:21793, pass:339, block:21414
total:21699, pass:390, block:21255
total:21104, pass:393, block:20654
total:21363, pass:453, block:20831
total:21619, pass:491, block:21052
total:21986, pass:533, block:21415
total:21789, pass:594, block:21123
total:21561, pass:685, block:20820
total:21663, pass:873, block:20717
total:20904, pass:988, block:19831
total:21500, pass:996, block:20423
total:21769, pass:1014, block:20682
total:20893, pass:1019, block:19837
total:21561, pass:973, block:20524
total:21601, pass:1014, block:20517
total:21475, pass:993, block:20420
total:21457, pass:983, block:20418
total:21397, pass:1024, block:20320
total:21690, pass:996, block:20641
total:21526, pass:991, block:20457
total:21779, pass:1036, block:20677
```



### Go熔断实战

这里我们介绍一个错误数量的，[查看详细熔断机制](https://sentinelguard.io/zh-cn/docs/golang/circuit-breaking.html)



#### error_count

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/alibaba/sentinel-golang/util"
)

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %d, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func main() {
	//基于连接数的降级模式
	total := 0
	totalPass := 0
	totalBlock := 0
	totalErr := 0
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=10s, recoveryTimeout=3s, maxErrorCount=50
		{
			Resource:         "abc",
			Strategy:         circuitbreaker.ErrorCount,
			RetryTimeoutMs:   3000, //3s只有尝试回复
			MinRequestAmount: 10,   //静默数
			StatIntervalMs:   5000,
			Threshold:        50,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("[CircuitBreaker ErrorCount] Sentinel Go circuit breaking demo is running. You may see the pass/block metric in the metric log.")
	go func() {
		for {
			total++
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g1 blocked
				totalBlock++
				fmt.Println("协程熔断了")
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				totalPass++
				if rand.Uint64()%20 > 9 {
					totalErr++
					// Record current invocation as error.
					sentinel.TraceError(e, errors.New("biz error"))
				}
				// g1 passed
				time.Sleep(time.Duration(rand.Uint64()%20+10) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	go func() {
		for {
			total++
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g2 blocked
				totalBlock++
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				// g2 passed
				totalPass++
				time.Sleep(time.Duration(rand.Uint64()%80) * time.Millisecond)
				e.Exit()
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(totalErr)
		}
	}()
	<-ch
}

```































