[toc]



### 概况

到目前为止我们已经对I项目的基本业务开发完毕了，接下来我们需要将项目里涉及到的所有参数进行配置化，编辑配置文件，对于go开发做参数配置使用较高的就是viper，这里也将使用viper进行配置。

到目前为止项目的结构：

```
HiChat
|
├── common
│   ├── md5.go
│   ├── resp.go
│  
├── config
│   
├── dao
│   ├── community.go
│   ├── relation.go
│   └── user.go
|
├── global
│   └── global.go
|
├── initialize
│   ├── db.go
│   └── logger.go
├── middlewear
│   └── jwt.go
├── models
│   ├── community.go
│   ├── group_info.go
│   ├── message.go
│   ├── relation.go
│   └── userBasic.go
├── router
│   └── router.go
├── service
│   ├── attach_upload.go
│   ├── index.go
│   ├── relation.go
│   └── user.go
├── sql
├── test
|   ├── main.go
├── go.mod
├── go.sum
├── main.go
```



### 编写配置文件

这里配置文件yaml文件为例，在项目目录下新建文件config-debug.yaml，需要将项目中所有的参数进行编写

```yaml
port: '8000'
mysql:
  host: '127.0.0.1'
  port: '3306'
  name: 'hi_chat'
  user: 'root'
  password: 'Qq/201djfidf'
redis:
  host: '127.0.0.1'
  port: '6379'

```



### 编写配置结构体

在config目录下新建文件config.go，编写的结构体结构需要和yaml文件对应。

```mapstructure```是用来读取yaml文件字段名tag

```go
package config

//MysqlConfig mysql信息配置
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"Name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

//对应yaml文件结构
type ServiceConfig struct {
	Port    int         `mapstructure:"port" json:"port"`
	DB      MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisDB RedisConfig `mapstructure:"redis" json:"redis"`
}
```



### 添加全局变量

在global目录下的global.go文件中添加字段```ServiceConfig```

```go
package global

import (
	"HiChat/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
	RedisDB       *redis.Client
)
```



### 读取配置文件

在initialize目录下新建config.go文件，使用viper将配置文件读取到全局变量ServiceConfig中。

```go
package initialize

import (
	"HiChat/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	//实例化对象
	v := viper.New()

	configFile := "../HiChat/config-debug.yaml"

	//读取配置文件
	v.SetConfigFile(configFile)

	//读入文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//将数据放入global.ServerConfig 这个对象如何在其他文件中使用--全局变量
	if err := v.Unmarshal(&global.ServiceConfig); err != nil {
		panic(err)
	}

	zap.S().Info("配置信息", global.ServiceConfig)
}
```

最后在main.go中进行初始化：

```go
package main

import (
	"fmt"

	"HiChat/global"
	"HiChat/initialize"
	"HiChat/router"
)


func main() {
	//初始化日志
	initialize.InitLogger()
	//初始化配置
	initialize.InitConfig()
	//初始化数据库
	initialize.InitDB()
	initialize.InitRedis()

	router := router.Router()

	router.Run(fmt.Sprintf(":%d", global.ServiceConfig.Port))
}
```



### 总结

这样整个项目的参数就配置好了，整个IM项目的基础功能就完成了，整个项目被分成了：用户，关系，文件，消息等模块；这个虽然很小也很简单，但是对于缺少项目经验同学来说，可以快速了解实际项目的设计到落地，从中学到项目如何拆分，项目结构如何设计，如何设计api等；也是将学习到的心得分享出来，希望对大家也有帮助，同时有很多不足，欢迎大家指正，建议。