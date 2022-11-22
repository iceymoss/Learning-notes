[toc]

### 文章介绍

本文我们将从零开始介绍Gin的安装，Gin的简单入门，基于Gin框架的登录/注册表单验证实例，Gin中间件的原理分析，Gin返回html，静态文件的挂载和Gin优雅的退出

### 什么是Gin？

***官方：Gin 是一个用 Go (Golang) 编写的 HTTP Web 框架。 它具有类似 Martini 的 API，但性能比 Martini 快 40 倍。如果你需要极好的性能，使用 Gin 吧。***

Gin 是 Go语言写的一个 web 框架，它具有运行速度快，分组的路由器，良好的崩溃捕获和错误处理，非常好的支持中间件和 json。总之在 Go语言开发领域是一款值得好好研究的 Web 框架，开源网址：https://github.com/gin-gonic/gin

### 安装

```
go get -u github.com/gin-gonic/gin
```

### 快速入门

实例1:

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//handle方法
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "ice_moss",
		"age":    18,
		"school": "家里蹲大学",
	})
}

func main() {
	//初始化一个gin的server对象
	//Default实例化对象具有日志和返回状态功能
	r := gin.Default()
  //注册路由，并编写处理方法
	r.GET("/ping", Pong)
	//监听端口：默认端口listen and serve on 0.0.0.0:8080
	r.Run(":8083")
}
```

接下来我们在浏览器中访问：http://localhost:8083/ping

可以访问到：

```json
"name":   "ice_moss",
"age":    18,
"school": "家里蹲大学",
```



实例二：gin的GET和POST方法

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"name": "ice_moss",
	})
}

func GinPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "您好",
	})
}
func main() {
	router := gin.Default()
	router.GET("/GinGet", GinGet)
	router.POST("/GinPost", GinPost)
	router.Run(":8083")
}
```

我们看到GinGet和GinPost这两个方法中的c.JSON()第二个参数不一样，原因：```gin.H```{}本质就是一个```map[string]interface{}```

```go
//H is a shortcut for map[string]interface{}
type H map[string]any
```

然后我们就可以访问：http://localhost:8083/GinGet

这里需要注意我们不能直接在浏览器中访问：http://localhost:8083/GinPost，因为他是POST方法

所以我们可以使用postman,来发送POST请求



### 路由分组

Gin为我们做了很好的路由分组，这样我们可以方便，对路由进行管理

实例三：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"矿泉水":  [5]string{"娃哈哈", "2元", "500"},
		"功能饮料": [3]string{"红牛", "6元", "200"},
	})
}

func Prouduct1(c *gin.Context) {
	req := c.Param("haha")

	c.JSON(http.StatusOK, gin.H{
		"矿泉水":   [5]string{"娃哈哈矿泉水", "2元", "500"},
		"token": req,
	})
}

func CreateProduct(c *gin.Context) {}

//路由分组
func main() {
	router := gin.Default()
  
  //未使用路由分组
	//获取商品列表
	//router.GET("/ProductList", ProductLists)
	//获取某一个具体商品信息
	//router.GET("/ProductList/1", Prouduct1)
	//添加商品
	//router.POST("ProductList/Add", CreateProduct)

	//路由分组
	ProductList := router.Group("/Produc")
	{
		ProductList.GET("/list", ProductLists)
		ProductList.GET("/1", Prouduct1)
		ProductList.POST("/Add", CreateProduct)
	}
	router.Run(":8083")
}
```



### URL值的提取

很多时候我们需要对URL中数据的提取，或者动态的URL，我们不可能将

实例四：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"矿泉水":  [5]string{"娃哈哈矿泉水", "2元", "500"},
		"功能饮料": [3]string{"红牛", "6元", "200"},
	})
}

func Prouduct1(c *gin.Context) {
	//获取url中的参数
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"矿泉水":    [5]string{"娃哈哈矿泉水", "2元", "500"},
		"id":     id,
		"action": action,
	})
}

func CreateProduct(c *gin.Context) {}

//url取值
func main() {
	router := gin.Default()
	//路由分组
	ProductList := router.Group("/Product")
	{
		ProductList.GET("", ProductLists)
		//使用"/:id"动态匹配参数
		ProductList.GET("/:id/:action", Prouduct1)
		ProductList.POST("", CreateProduct)
	}
	router.Run(":8083")
}
```

访问：http://localhost:8083/Product/01/product1

返回:

```json
{"action":"product1","id":"01","矿泉水":["娃哈哈矿泉水","2元","500","",""]}
```

当我们访问：http://localhost:8083/Product/100/product2000

返回：

```json
{"action":"product2000","id":"100","矿泉水":["娃哈哈矿泉水","2元","500","",""]}
```



### 构体体声明并做约束

实例五：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//结构体声明，并做一些约束
type Porsen struct {
	ID   int    `uri:"id" binding:"required"`    //uri指在client中的名字为id，binding:"required指必填
	Name string `uri:"name" binding:"required"`  //同理
}

//url参数获取
func main() {
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var porsen Porsen
		if err := c.ShouldBindUri(&porsen); err != nil {
			c.Status(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": porsen.Name,
			"id":   porsen.ID,
		})
	})
	router.Run(":8083")
}
```

当我们访问：http://localhost:8083/test100/2000

返回：

```go
{"id":2000,"name":"test100"}
```

但是我们这样访问：http://localhost:8083/100/test2000

返回：

```
找不到 localhost 的网页找不到与以下网址对应的网页：http://localhost:8083/100/test2000
HTTP ERROR 404
```

这和我们约束条件一致



### URL参数的提取

###### GET请求参数获取

URL参数的提取是GET方法常用的方法，URL中有需要的参数，例如我们访问百度图片：https://image.baidu.com/search/index?ct=201326592&tn=baiduimage&word=%E5%9B%BE%E7%89%87%E5%A3%81%E7%BA%B8&pn=&spn=&ie=utf-8&oe=utf-8&cl=2&lm=-1&fr=ala&se=&sme=&cs=&os=&objurl=&di=&gsm=1e&dyTabStr=MCwzLDYsMSw0LDIsNSw3LDgsOQ%3D%3D

以```?```后的都是参数参数间用```&```分隔：?ct=201326592&tn=baiduimage&word=%E5%9B%BE%E7%89%87%E5%A3%81%E7%BA%B8&pn=&spn=&ie=utf-8&oe=utf-8&cl=2&lm=-1&fr=ala&se=&sme=&cs=&os=&objurl=&di=&gsm=1e&dyTabStr=MCwzLDYsMSw0LDIsNSw3LDgsOQ%3D%3D

实例六：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcom(c *gin.Context) {
	//DefaultQuery根据字段名获取client请求的数据，client未提供数据则可以设置默认值
	first_name := c.DefaultQuery("first_name", "未知")
	last_mame := c.DefaultQuery("last_name", "未知")
	c.JSON(http.StatusOK, gin.H{
		"firstname": first_name,
		"lastname":  last_mame,
		"work":      [...]string{"公司：Tencent", "职位：Go开发工程师", "工资：20000"},
	})
}

//url参数获取
func main() {
	//实例化server对象
	router := gin.Default()
	router.GET("/welcom", Welcom)
	router.Run(":8083")
}
```

接着我们在浏览器中访问：http://localhost:8083/welcom?first_name=moss&last_name=ice

返回：这样我们的后台就拿到了client提供的参数，并做业务处理，然后返回client

```json
{"firstname":"moss","lastname":"ice","work":["公司：Tencent","职位：Go开发工程师","工资：20000"]}
```



###### POST表单提交及其数据获取

在很多时候我们都需要使用post方法来传输数据，列如，用户登录/注册都需要提交表单等

下面我们来看看简单的表单提交：

实例七：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Postform(c *gin.Context) {
	UserName := c.DefaultPostForm("username", "unkown")
	PassWord := c.DefaultPostForm("password", "unkown")
	if UserName == "ice_moss@163.com" && PassWord == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"name":        "ice_moss",
			"username":    UserName,
			"tokenstatus": "认证通过",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"tokenstatus": "认证未通过",
		})
	}
}

//url参数获取
func main() {
	//实例化server对象
	router := gin.Default()
	router.POST("/Postform", Postform)
	router.Run(":8083")
}
```

由于是post请求我们使用postman提交表单：http://localhost:8083/Postform

![queue1.png](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/vp6Ppr23Lv.png%21large.png)



后台输出：

```
username ice_moss password 18dfdf
[GIN] 2022/06/23 - 19:24:24 | 500 |     108.029µs |             ::1 | POST     "/Postform"
```





###### GET和POST混合使用

直接实例，实例八：

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "未知的")
	name := c.DefaultPostForm("name", "未知的")
	password := c.PostForm("password")
	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"page":     page,
		"name":     name,
		"password": password,
	})
}

//url参数获取
func main() {
	//实例化server对象
	router := gin.Default()

	//GET和POST混合使用
	router.POST("/Post", GetPost)
	router.Run(":8083")
}
```

因为是post方法使用，访问：http://localhost:8083/Post?id=1&page=2

![post.png](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/aXQ0r7FJEA.png%21large.png)



返回：

```json
{
    "id": "1",
    "name": "ice_moss@163.com",
    "page": "2",
    "password": "123456"
}
```



### 数据格式JSON和ProtoBuf

我们知道前后端数据交互大多数都是以json的格式，Go也同样满足，我们知道GRPC的数据交互是以ProtoBuf格式的

下面我们来看看Go是如何处理json的，如何处理ProtoBuf的

实例九

```go
ackage main

import (
	"net/http"

	"StudyGin/HolleGin/ch07/proto"

	"github.com/gin-gonic/gin"
)

func moreJSON(c *gin.Context) {
	var msg struct {
		Nmae    string `json:"UserName"`
		Message string
		Number  int
	}
	msg.Nmae = "ice_moss"
	msg.Message = "This is a test of JSOM"
	msg.Number = 101

	c.JSON(http.StatusOK, msg)
}

//使用ProtoBuf
func returnProto(c *gin.Context) {
	course := []string{"python", "golang", "java", "c++"}
	user := &proto.Teacher{
		Name:   "ice_moss",
		Course: course,
	}
	//返回protobuf
	c.ProtoBuf(http.StatusOK, user)
}

//使用结构体和JSON对结构体字段进行标签，使用protobuf返回值
func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)
	router.Run(":8083")
}
```

访问：http://localhost:8083/moreJSON

返回：

```go
{"UserName":"ice_moss","Message":"This is a test of JSOM","Number":101}
```

访问：http://localhost:8083//someProtoBuf

返回：然后会将someProtoBuf返回的数据下载，当然我们可以使用GRPC中的方法将数据接收解析出来



![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/eoOzPib2wb.png%21large.png)



### Gin解析特殊字符

我们很多时候需要处理特殊的字符，比如：JSON会将特殊的HTML字符替换为对应的unicode字符，比如<替换为\u003c，如果想原样输出html，则使用PureJSON

实例十：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//通常情况下，JSON会将特殊的HTML字符替换为对应的unicode字符，比如<替换为\u003c，如果想原样输出html，则使用PureJSON
func main() {
	router := gin.Default()
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>您好，世界!</b>",
		})
	})

	router.GET("/pureJSON", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<div><b>您好，世界!</b></div>",
		})
	})
	router.Run(":8083")
}
```

访问：http://localhost:8083/json

返回：

```
{"html":"\u003cb\u003e您好，世界!\u003c/b\u003e"}
```



访问：http://localhost:8083/pureJSON

返回：

```
{"html":"<div><b>您好，世界!</b></div>"}
```



### Gin翻译器的实现

在这段代码中，我们是将注册代码实现翻译功能

实例十一：

```go
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"github.com/gin-gonic/gin"
)

// 定义一个全局翻译器T
var trans ut.Translator

//Login登录业务，字段添加tag约束条件
type Login struct {
	User     string `json:"user" binding:"required"`     //必填
	Password string `json:"password" binding:"required"` //必填
}

//SignUp注册业务，字段添加tag约束条件
type SignUp struct {
	Age        int    `json:"age" binding:"gte=18"`                            //gte大于等于
	Name       string `json:"name" binding:"required"`                         //必填
	Email      string `json:"email" binding:"required,email"`                  //必填邮件
	Password   string `json:"password" binding:"required"`                     //必填
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //RePassword和Password值一致
}

//RemoveTopStruct去除以"."及其左部分内容
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, value := range fields {
		res[field[strings.Index(field, ".")+1:]] = value
	}
	return res
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func main() {
	res := map[string]string{
		"ice_moss.habbit": "打球",
		"ice_moss.from":   "贵州 中国",
	}
	fmt.Println(RemoveTopStruct(res))

	//初始化翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器失败", err)
		return
	}

	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			fmt.Println(err.Error())
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errs.Translate(trans),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "验证通过",
		})
	})

	router.POST("/signupJSON", func(c *gin.Context) {
		var signup SignUp
		if err := c.ShouldBind(&signup); err != nil {
			fmt.Println(err.Error())
			//获取validator.ValidationErrors类型的error
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			//validator.ValidationErrors类型错误则进行翻译
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": RemoveTopStruct(errs.Translate(trans)),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	router.Run(":8083")
}
```

当我们访问：http://localhost:8083/signupJSON

如果参数不满足 tag 中的条件，则会返回如下结果：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/QMwRltZo11.png%21large.png)



当我们输入满足 tag 中的条件，就成功返回了：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/d6GM1eyNwm.png%21large.png)







### Gin中间件原理及自定义中间件

在此之前我们先来看一下Gin实例化server,我们在之前是使用```router := gin.Default()```，但其实我们是可以直接使用```router := gin.New()```, 那么在之前是实例中我们为什么不使用```gin.New()```呢？

别急，我们先来看看```gin.Default()```的源码：

```go
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```

我们可以看到Default()其实是对gin.New()做了一层封装，并且做了其他事情，***这里的其他事情就有“中间件”***

即：```engine.Use(Logger(), Recovery())```他调用两两个中间件（ Logger()用来输出日志信息，Recovery 中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌，中间件将会写入500）



###### Gin中间件原理

我们来看看：engine.Use()

```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}
```

入参是```HandlerFunc```类型，那么我们接着往下看```HandlerFunc```

```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)
```

其实```HandlerFunc ```是```func(*Context)```类型

到这里中间件我们就可以自定义了

###### 自定义中间件

我们定义一个监控服务运行时间，运行状态的中间件：

实例十二：

```go
//自定义中间件，这里我们以函数调用的形式，对中间件进一步封装
func MyTimeLogger() gin.HandlerFunc {
	return func(c *gin.Context) {  		//真正的中间件类型
		t := time.Now()
		c.Set("msg", "This is a test of middleware")
		//它执行调用处理程序内链中的待处理处理程序
		//让原本执行的逻辑继续执行
		c.Next()

		end := time.Since(t)
		fmt.Printf("耗时：%D\n", end.Seconds())
		status := c.Writer.Status()
		fmt.Println("状态监控:", status)
	}
}
```

我们在main函数中：

```go
func main() {
	router := gin.Default()
	router.Use(MyTimeLogger())   //这里使用函数调用
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Pong",
		})
	})
	router.Run(":8083")
}
```

访问：http://localhost:8083/ping

返回  ：

```json
{"msg":"Pong"}
```



###### 中间件实际应用

基于中间件模拟登录：

实例十三：

```go
//自定义中间件
func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
    //从请求头中获取数据
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
		}
		fmt.Println(token)
		if token != "ice_moss" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "认证未通过",
			})
      
      //return在这里不会有被执行
			c.Abort()   //这里先不用理解，后面会讲解，这里先理解为return
		}
    //继续往下执行该执行的逻辑
		c.Next()
	}
}
```

将中间件加入gin中：

```go
func main() {
	router := gin.Default()
	router.Use(TokenRequired())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Pong",
		})
	})
	router.Run(":8083")
}
```



我们在postman中进行请求：http://localhost:8083/ping

将Headers，增加字段：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/vX75tL6Gwe.png%21large.png)



正确参数返回：

```json
{
    "msg": "Pong"
}
```



不正确参数返回：

```json
{
    "msg": "认证未通过"
}
```



不正确参数返回：

```json
{
    "msg": "认证未通过"
}
```

这里我们需要对中间件原理进一步剖析：
现在我们将围绕两个方法来解释
`c.Abort() 和 c.Next()`

###### c.Abort()

在**实例十三**中我们看到：

```go
func TokenRequired() gin.HandlerFunc {
   return func(c *gin.Context) {
      var token string
 for k, v := range c.Request.Header {
         if k == "X-Token" {
            token = v[0]
         } else {
            fmt.Println(k, v)
         }
      }
      fmt.Println(token)
      if token != "ice_moss" {
         c.JSON(http.StatusUnauthorized, gin.H{
            "msg": "认证未通过",
  })
         //return 不会被执行，需要使用c.Abort()来结束当前
         c.Abort()
      }
      //继续执行该执行的逻辑
      c.Next()
   }
}
```



1. 为什么 return 不能直接返回，而是使用 c.Abort ()

**原因**：当我们启动服务后，Gin 会有一个类似于任务队列将所有配置的中间件和在注册处理方法压入队列中：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/loO5d6WBkD.png%21large.png)



在处理业务代码之前，会将所有注册路由中的中间件以队列的执行方式执行，比如上面我们：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/WLWtzFqCOC.png%21large.png)





当我们在实例十三中执行 return 他只是将当前函数返回，但是后面的方法仍然是按逻辑执行的，很显然这不是我们想要的结果，不满足验证条件的情况，应该将对此时的 client 终止服务，如果要终止服务就应该将图中的箭头跳过所有方法：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/aSmEdHBReP.png%21large.png)

这样整个服务才是真正的终止，下面再来看看 `Abort()`:

```go
func (c *Context) Abort() {
   c.index = abortIndex
}
```

当代码执行到 Abort () 时，index 被赋值为 abortIndex，`abortIndex` 是什么？

```go
const abortIndex int8 = math.MaxInt8 >> 1
```

可以看到，最后 index 指向任务末端，这就是 ```const abortIndex int8 = math.MaxInt8 >> 1``` 作用的效果

**c.Next()**
理解了 ```Abort()```,  ```Next()```自然就好理解了，我们来看看 ```Next()``` 定义

```go
func (c *Context) Next() {
   c.index++
   for c.index < int8(len(c.handlers)) {
      c.handlers[c.index](c)
      c.index++
   }
}
```

执行过程：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/EYVerBi4j0.png%21large.png)



![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/Ct3AXNzBVg.png%21large.png)



![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/QwAbcXGxTx.png%21large.png)



### Gin 返回 html 模板

我们使用 html 模板，将后端获取到的数据，直接填充至 html 中
我们先来编写一个 html (实例为 tmpl, 无影响)

```html
<!DOCTYPE html>
<html lang="en">
<head>
 <meta charset="UTF-8">
 <title>{{ .title }}</title>
</head>
<body>
<h1>{{ .menu }}</h1>
</body>
</html>
```

其中数据以 `{{ .title }}` 从 web 层填充进来
我们需要注意目录结构，程序的执行入口 `main` 需要和模板 `templates` 放置同一目录下，这样保证 `main` 能读取文件 html

```
ch11
├── main.go
└── templates
    └── index.tmpl
```

实例十四：
main:

```go
package main

import (
   "fmt"
 "net/http" "os" "path/filepath"
 "github.com/gin-gonic/gin")

func main() {
   router := gin.Default()
   //读取文件
   router.LoadHTMLFiles("templates/index.tmpl")
   router.GET("/index", func(c *gin.Context) {
       //写入数据，  key必须要tmpl一致
      c.HTML(http.StatusOK, "index.tmpl", gin.H{
         "title": "购物网",  
         "menu":  "菜单栏",
  })
   })
   router.Run(":8085")
}
```

当我们在浏览器中访问：[localhost:8085/index](http://localhost:8085/index)
获取到:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/05xgpzzzzd.png%21large.png)



当然 `router.LoadHTMLFiles()` 方法可以加载多个 html 文件

实例十五：

```go
package main

import (
   "fmt"
 "net/http" "os" "path/filepath"
 "github.com/gin-gonic/gin")

func main() {
   router := gin.Default()

   //读取模板文件，按指定个读取
  router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")
   router.GET("/index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.tmpl", gin.H{
         "title": "shop",
         "menu":  "菜单栏",
  })
   })

   router.GET("goods", func(c *gin.Context) {
      c.HTML(http.StatusOK, "goods.html", gin.H{
         "title": "goods",
         "goods": [4]string{"矿泉水", "面包", "薯片", "冰淇淋"},
  })
   })
   router.Run(":8085")
}
```

这样就可以访问：localhost:8085/goods
或者：localhost:8085/index
返回结果：略

当然如果 html 文件很多，Gin 还提供了 ```
func (engine *Engine) LoadHTMLGlob(pattern string) {……}
我们只需要这样调用：

````go
//将"templates文件夹下所有文件加载
router.LoadHTMLGlob("templates/*")
````

对应二级目录，我们又是如何处理的呢？

```go
//加载templates目录下的目录中的所有文件
router.LoadHTMLGlob("templates/**/*")
```

实例十六：

```go
package main

import (
   "fmt"
 "net/http" "os" "path/filepath"
 "github.com/gin-gonic/gin")

func main() {
   router := gin.Default()
 router.LoadHTMLGlob("templates/**/*")
   router.GET("user/list", func(c *gin.Context) {
      c.HTML(http.StatusOK, "list.html", gin.H{
         "title": "shop",
         "list":  "用户列表",
  })
   })

   router.GET("goods/list", func(c *gin.Context) {
      c.HTML(http.StatusOK, "list.html", gin.H{
        "title": "shop",
        "list":  "商品列表",
  })
   })
   router.Run(":8085")
}
```

这样我们访问：[localhost:8085/goods/list](http://localhost:8085/goods/list) 或者 [http://localhost:8085/user/list](http://localhost:8085/user/list)都能访问到



### Gin静态文件的挂载

在 web 开发中经常需要将 js 文件和 css 文件，进行挂载，来满足需求
目录结构：

```
ch11
├── main.go
├── static
│   └── style.css
└── templates
    └── user
        └── list.html
```

Html文件：

```html
<!DOCTYPE html>
<html lang="en">
<head>
 <meta charset="UTF-8">
 <title>{{ .title }}</title>
 <link rel="stylesheet" href="/static/style.css">
</head>
<body>
<h1>{{ .list }}</h1>
</body>
</html>
```

css 文件:

```css
*{
    background-color: aquamarine;
}
```

静态文件挂载方法：

```go
router.Static("/static", "./static")
```


该方法会去在 html 文件中 <link> 标签中找到以 static 开头的链接，然后去找在当前 main 所在的目录下找到以第二个参数./static 名称的目录下找到静态文件，然后挂载

实例十七：

```go
package main

import (
    "fmt"
 	  "net/http" "os" "path/filepath"
    "github.com/gin-gonic/gin")

func main() {
   router := gin.Default()
   //挂载静态文件
   router.Static("/static", "./static")
   router.LoadHTMLGlob("templates/**/*")
   router.GET("user/list", func(c *gin.Context) {
      c.HTML(http.StatusOK, "list.html", gin.H{
         "title": "shop",
  "list":  "用户列表",
  })
   })

   router.Run(":8085")
}
```

然后访问：[localhost:8085/user/list](http://localhost:8085/user/list)
可以看到：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gin%E6%A1%86%E6%9E%B6/TsPHtlG2AF.png%21large.png)





### Gin优雅退出

在业务中，我们很多时候涉及到服务的退出，如：各种订单处理中，用户突然退出，支付费用时，程序突然退出，这里我们是需要是我们的服务合理的退出，进而不造成业务上的矛盾
实例十八：

```go
package main

import (
   "fmt"
   "net/http" "os" "os/signal" "syscall"
   "github.com/gin-gonic/gin")

func main() {
   router := gin.Default()
   router.GET("ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
         "msg": "ping",
  })
   })

   go func() {
      router.Run(":8085")
   }()

   qiut := make(chan os.Signal)
   //接收control+c
   //当接收到退出指令时，我们向chan收数据
  signal.Notify(qiut, syscall.SIGINT, syscall.SIGTERM)
   <-qiut

 //服务退出前做处理
   fmt.Println("服务退出中")
   fmt.Println("服务已退出")
}
```

在 terminal 中运行：go run main.go
服务启动后在 terminal 中退出 (control+c) 就可以看到:

```go
[GIN-debug] Listening and serving HTTP on :8085
服务退出中
服务已退出
```















