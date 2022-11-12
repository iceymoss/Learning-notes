[toc]

### 文章介绍

本文我们将以golang的角度来了解配置中心，nacos的安装，如何用nacos做配置，读取nacos中配置文件。

### 什么是配置中心

**用来统一管理项目中所有配置的系统**。 虽然听起来很简单，但也不要小瞧了这个模块。 如果一个中型互联网项

目，不采用配置中心的模式，一大堆的各类配置项，各种不定时的修改需求，一定会让开发同学非常头疼且管理

十分混乱。

配置中心的服务流程如下： 用户在配置中心更新配置信息。 服务A和服务B及时得到配置更新通知，从配置中心

获取配置。 总得来说，配置中心就是一种统一管理各种应用配置的基础服务组件。 在系统架构中，配置中心是

整个微服务基础架构体系中的一个组件，它的功能看上去并不起眼，无非就是配置的管理和存取，但它是整个微

服务架构中不可或缺的一环。

### nacos是什么

#### Nacos

`Nacos`是`Naming and Configuration Service`的缩写，从名字上能看出它重点关注的两个领域是`Naming`即注

册中心和`Configuration`配置中心。

业务上的配置，功能开关，服务治理上对弱依赖的降级，甚至数据库的密码等，都可能用到动态配置中心。 在没

有专门的配置中心组件时，我们使用硬编码、或配置文件、或数据库、缓存等方式来解决问题。 硬编码修改配置

时需要重新编译打包，配置文件需要重启应用，数据库受限于性能，缓存丧失了及时性。

#### Nacos的配置模型

namespace + group + dataId  唯一确定一个配置

- namespace：与client绑定，一个clinet对应到一个namespace，可用来隔离环境或区分租户
- group：分组，区分业务
- dataId：配置的id

来看一下是如何在实际场景使用的

例如：一个电商网站其中有这几个模块：用户模块、商品模块、订单模块、库存模块

这几个模块都需要进行配置且它们的配置不同，这是我们为每一个模块做一个```namespace```， 每一个模块都需要有

开发阶段的配置，和上线后的配置，使用我们使用dev,和pro两个分组来进行区分，对于dataId，不管是dev还是

pro都必须填写。



### Nacos的安装

这里我们直接使用docker进行安装

### Nacos的安装(docker)

```
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
```

访问：http://192.168.1.103:8848/nacos/index.html 用户名/密码：nacos/nacos

配置开机启动：

```
docker container update --restart=always xxx
```



### 用nacos做配置

nacos成功启动后访问：http://192.168.1.103:8848/nacos/index.html 

可以创建```namespace```, 我们新建一个用户模块```user``` , 创建成功后可以看到有对应的id， 例如：7ae18f62-e2b9-48bd-bff2-a49e7443f5bc

然后我们在user命名空间下新建一个配置文件，并填写对应的名称(dataId)和分组，这里我们新建一个josn的配置文件：

```json
{
    "name": "user-web",
    "host": "10.2.106.169",
    "port": 9091,
    "tags":["iceymoss", "goods", "golang", "web"],
    "user_srv":{
        "name": "user-srv",
        "host": "10.2.106.169",
        "port": 8081
    },
    "jwt":{
        "key": "dfijdfjidhfjijdfbdfdFwohPd6XmVCdnQi"
    },
    "sms":{
        "key": "mykey",
        "secret": "mysecret"
    },
    "params":{
        "sign_name": "生鲜小店",
        "code": "SMS_244610581"
    },
    "redis":{
        "host": "127.0.0.1",
        "port": 6379,
        "expir": 300
    },
    "verify":{
        "width": 5
    },
    "consul":{
        "host": "10.2.106.169",
        "port": 8500
    },
    "tracing":{
        "host": "127.0.0.1",
        "port": 6831,
        "name": "shopping"
    }
}
```

这样整个配置文件就配置完成了



### 读取nacos中配置文件

#### 拉取依赖

我们使用go来读取配置文件，使用需要拉取nacos的sdk：

```shell
go get github.com/nacos-group/nacos-sdk-go/clients
```

```shell
go get github.com/nacos-group/nacos-sdk-go/common/constant
```

```shell
go get github.com/nacos-group/nacos-sdk-go/vo
```



#### 读取配置

在读取配置之前我们先编写一个用来做配置映射的结构体

目录结构：

```shell
nacos_test
├── config
│   └── config.go
└── main.go

```

编写config时需要注意的是我们需要保持tag名和配置文件中的名称一致

```go
package config

//UserSerConfig 映射用户配置
type UserSerConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//JWTConfig 映射token配置
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

//AliSmsConfig 阿里秘钥
type AliSmsConfig struct {
	Apikey    string `mapstructure:"key" json:"key"`
	ApiSecret string `mapstructure:"secret" json:"secret"`
}

//ParamsConfig 短信模板配置
type ParamsConfig struct {
	SignName     string `mapstructure:"sign_name" json:"sign_name"`
	TemplateCode string `mapstructure:"code" json:"code"`
}

//RedisConfig redis数据库配置
type RedisConfig struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Expir int    `mapstructure:"expir" json:"expir"`
}

//Verifier 手机验证长度
type Verifier struct {
	Width int `mapstructure:"width" json:"width"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//ServerConfig  映射服务配置
type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	UserSerInfo UserSerConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSms      AliSmsConfig  `mapstructure:"sms" json:"sms"`
	Params      ParamsConfig  `mapstructure:"params" json:"params"`
	Redis       RedisConfig   `mapstructure:"redis" json:"redis"`
	Verify      Verifier      `mapstructure:"verify" json:"verify"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}
```



下面进行配置文件读取：

```go
package main

import (
	"StudyGin/nacos/config"
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	//服务端配置, nacos运行的socket
	sc := []constant.ServerConfig{
		{
			IpAddr: "10.2.81.102",
			Port:   8848,
		},
	}

	//客服端配置
	cc := constant.ClientConfig{
		NamespaceId:         "7ae18f62-e2b9-48bd-bff2-a49e7443f5bc", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel: "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	//获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}
	Config := &config.ServerConfig{}
  
  //将配置信息读取到config.ServerConfig{}对象中
	err = json.Unmarshal([]byte(content), &Config)
	if err != nil {
		panic(err)
	}
	fmt.Println(Config)

}
```

输出：

```
&{user-web 9091 {user-srv 10.2.106.169 8081} {dfijdfjidhfjijdfbdfdFwohPd6XmVCdnQi} {mykey mysecret} {生鲜小店 SMS_244610581} {127.0.0.1 6379 300} {5} {10.2.106.1600}}
```

当然配置中心和viper都提供实时监控配置

可以这样写：

```go
	//监听配置变化
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web",
		Group:  "DEV",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	time.Sleep(3000 * time.Second)
```





