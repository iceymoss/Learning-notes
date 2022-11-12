### 文章介绍

本文我们将来介绍一下如何从第三方包viper来获取配置文件的配置信息，以及开发模式下的配置和上线产品如何使用进行viper做配置映射

### viper的安装

```shell
go get github.com/spf13/viper
```



### 快速开始

实例目录结构：

```shell
ch1
├── config.yaml    //配置文件
└── main.go  			 //viper读取配置
```

config.yaml:

```yaml
name: 'user-web'
mysql:
  host: '127.0.0.1'
  port: 3306
```

main.go:

```go
package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type mysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

//ServerConfig 将配置文件内容映射为struct
type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	SqlConfig  mysqlConfig `mapstructure:"mysql"`
}

func main() {
	//new对象
	v := viper.New()

	//读取配置文件
	v.SetConfigFile("viper_test/ch1/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	ServerConfig := ServerConfig{}
	v.Unmarshal(&ServerConfig)
	fmt.Println(ServerConfig)
}
```

在上述代码中我们定义了两个结构体：

```go
type mysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

//ServerConfig 将配置文件内容映射为struct
type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	SqlConfig  mysqlConfig `mapstructure:"mysql"`
}
```

这两个结构体用来做配置映射，这里使用```mapstructure```为配置做tag

完整代码：

```go
package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type mysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

//ServerConfig 将配置文件内容映射为struct
type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	SqlConfig  mysqlConfig `mapstructure:"mysql"`
}

func main() {
	//new一个viper对象
	v := viper.New()

	//获取配置文件
	v.SetConfigFile("viper_test/ch1/config.yaml")
  
  //将配置读入viper对象中
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	ServerConfig := ServerConfig{}
  
  //将读取到的配置信息映射至ServerConfig中
	v.Unmarshal(&ServerConfig)
	fmt.Println(ServerConfig)
}
```

输出：

```shell
{user-web {127.0.0.1 3306}}
```



### 实时监控配置

viper为我们提供了实时监控配置的功能

```go
v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置变化：", e)
		_ = v.ReadInConfig()
		v.Unmarshal(&ServerConfig)
		fmt.Println(ServerConfig)
	})
```

事例：

```go
package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type mysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

//ServerConfig 将配置文件内容映射为struct
type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	SqlConfig  mysqlConfig `mapstructure:"mysql"`
}

func main() {
	//new对象
	v := viper.New()

	//读取配置文件
	v.SetConfigFile("viper_test/ch1/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	ServerConfig := ServerConfig{}
	v.Unmarshal(&ServerConfig)
	fmt.Println(ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置变化：", e)
		_ = v.ReadInConfig()
		v.Unmarshal(&ServerConfig)
		fmt.Println(ServerConfig)
	})
  
  //延迟程序退出时间
	time.Sleep(300 * time.Second)
}
```

运行程序后我们对port进行更改，会有如下输出：

```sh
{user-web {127.0.0.1 3306}}
配置变化： WRITE         "viper_test/ch1/config.yaml"
{user-web {127.0.0.1 3303}}
配置变化： WRITE         "viper_test/ch1/config.yaml"
{user-web {127.0.0.1 3306}}
```



### 线上线下环境隔离

目录结构：

```go
ch02
├── config-debug.yaml   //开发环境
├── config-por.yaml			//线上环境
└── main.go
```

config-debug.yaml：

```yaml
name: 'user-web-debug'
mysql:
  host: '127.0.0.1'
  port: 3306
```

config-por.yaml:

```yaml
name: 'user-web'
mysql:
  host: '195.1.43.12'
  port: 2000
```

根据条件读写线上\线下配置并做监控

```go
package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//将配置文件嵌套映射
type MysqlServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

//将配置文件内容映射为struct
type ServerConfig struct {
	ServerName string            `mapstructure:"name"`
	Mysql      MysqlServerConfig `mapstructure:"mysql"`
}

//通过环境变量，将线上线下环境隔离
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
  
  //默认为线上配置
	configFileName := fmt.Sprintf("viper_test/ch02/%s-por.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("viper_test/ch02/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	ServerConfig := ServerConfig{}
	if err := v.Unmarshal(&ServerConfig); err != nil {
		panic(err)
	}

	fmt.Println(ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置变化：", e)
		_ = v.ReadInConfig()
		v.Unmarshal(&ServerConfig)
		fmt.Println(ServerConfig)
	})
	time.Sleep(300 * time.Second)
}
```

