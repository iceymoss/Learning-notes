### 文章介绍

本文我们将介绍什么是consul，为什么需要consul， consul的安装， 服务注册， 健康检查， 服务发现

### 什么是consul

官方介绍：

Consul is a service networking solution to automate network configurations, discover services, and enable secure connectivity across any cloud or runtime

翻译：Consul是一种服务网络解决方案，可自动化网络配置、发现服务并支持跨任何云或运行时的安全连接。

这里列举一个consul使用场景：微服务中，会将整个服务拆分为多个微服务，当多个微服务之间进行调用时，就需要进行配置，这个过程更好对解决方案是使用consul对各个微服务做服务发现，例如A服务，B服务，C服务，D服务，E服务在服务启动时，我们将其注册到consul中，当ABCDE之间进行相互调用时，各个微服务就可以从consul中做服务发现，拿到需要调用的服务的信息，然后再去调用服务就可以了



###  consul的安装

这里我们使用docker安装

```
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp  consul consul agent  -dev -client=0.0.0.0
```

开机启动consul

```
docker container update --restart=always 容器名字
```

浏览器访问 127.0.0.1:8500



###   服务注册

这里我们直接使用go做服务注册

这里需要拉取```github.com/hashicorp/consul/api```

```shell
go get github.com/hashicorp/consul/api
```

注册http服务：

```go
package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

//Register 注册服务至注册中心
func Register(address string, port int, name string, tags []string, id string) error {
	//DefaultConfig 返回客户端的默认配置
	cfg := api.DefaultConfig()
  
  //安装consul的ip:port
	cfg.Address = "10.2.69.164:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://10.2.69.164:5001/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func main(){
  tags := []string{"ice_moss", "inventory", "server", "golang"}
  Register("10.2.69.164", 5001, "inventory", tags, "inventory-srv")
}
```

仔细看的话就会发现问题，我们整个过程并没有启动：```"http://10.2.69.164:5001/health```

如果我们注册的是grpc服务的话就需要这样配置：

```go
//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.Host, *Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}
```

我们到consul中虽然可以看到```inventory-srv```的信息，但是这个服务是挂掉的

所以我们还需要做健康检查

### 健康检查

这里我们之间修改上述代码：

```go
package main

import (
	"fmt"
	"net"

  "github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

//Register 注册服务至注册中心
func Register(address string, port int, name string, tags []string, id string) error {
	//DefaultConfig 返回客户端的默认配置
	cfg := api.DefaultConfig()
  
  //安装consul的ip:port
	cfg.Address = "10.2.69.164:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://10.2.69.164:5001/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func main(){
  tags := []string{"ice_moss", "inventory", "server", "golang"}
  Register("10.2.69.164", 5001, "inventory", tags, "inventory-srv")

	router := gin.Default()
	router.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "test",
		})
	})
	err := router.Run(":5001")
	if err != nil {
		panic(err)
	}
```

这样我们就可以成功注册服务到consul中了，这样consul会在http://10.2.69.164:5001/health一直轮询检查我们的服务



### 服务发现

当我们的注册服务时可以对consul做服务发现，看看这个有哪些服务注册到consul中，下面直接看代码：

```go
package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

//Register 注册服务至注册中心
func Register(address string, port int, name string, tags []string, id string) error {
	//DefaultConfig 返回客户端的默认配置
	cfg := api.DefaultConfig()
	cfg.Address = "10.2.69.164:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://10.2.69.164:5001/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

//AllServices 发现所有服务
func AllServices() {
	//配置
	cfg := api.DefaultConfig()
	cfg.Address = "10.2.69.164:8500"

	//将配置写入对象中
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, v := range data {
		fmt.Println("key:", key, "address:", v.Address, "port:", v.Port)
	}
}

func main() {
  //开启协程
	go AllServices()

	tags := []string{"ice_moss", "inventory", "server", "golang"}
	Register("10.2.69.164", 5001, "inventory", tags, "inventory-srv")

	Register("10.2.69.164", 5001, "user-web", []string{"user-web", "ice_moss"}, "user-web")

	router := gin.Default()
	router.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "test",
		})
	})
	err := router.Run(":5001")
	if err != nil {
		panic(err)
	}
}
```

其中输出：

```go
key: inventory-srv address: 10.2.69.164 port: 5001
key: user-web address: 10.2.69.164 port: 5001
```



当我们需要发现指定服务：

```go
//FilterSerivice 服务发现
func FilterSerivice() {
	cfg := api.DefaultConfig()
	cfg.Address = "10.2.69.164:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

  //服务name
	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}
```











