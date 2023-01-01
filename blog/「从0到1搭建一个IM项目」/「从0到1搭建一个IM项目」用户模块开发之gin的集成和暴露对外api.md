[toc]



### 概况

在上一篇中完成了用户数据库表设计及dao层的开发，完成了底层基础功能；在本篇中将Gin框架集成到项目中， 以及对外实现api的开发， 目前项目目录结构：

```sh
HiChat   
    ├── common    //放置公共文件
    │  
    ├── config    //做配置文件
    │  
    ├── dao//数据库crud
    │     |——user.go
    |
    ├── global    //放置各种连接池，配置等
    │   		|——global.go
    |
    ├── initialize  //项目初始化文件
    │  			|——db.go
    |				|——logger.go
    |
    ├── middlewear  //放置web中间件
    │ 
    ├── models      //数据库表设计
    │   		|——user_basic.go
    |
    ├── router   		//路由
    │   
    ├── service     //对外api
    │   
    ├── test        //测试文件
    │  
    ├── main.go     //项目入口
    ├── go.mod			//项目依赖管理
    ├── go.sum			//项目依赖管理
```



### 集成Gin框架

#### 安装Gin

```go
go get -u github.com/gin-gonic/gin
```



#### 快速入门

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

接下来我们在浏览器中访问：[localhost:8083/ping](http://localhost:8083/ping)

可以访问到：

```json
"name":   "ice_moss",
"age":    18,
"school": "家里蹲大学",
```

您可能需要学习更多的gin知识：[GoWeb框架Gin学习总结](https://learnku.com/articles/69259)



#### gin的集成

在router目录下新建文件router.go

```go
package router

import (
	"HiChat/middlewear"
	"HiChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//初始化路由
	router := gin.Default()

	//v1版本
	v1 := router.Group("v1")

	//用户模块，后续有个用户的api就放置其中
	user := v1.Group("user")
	{
		user.GET("/list", service.List)
	}
	
	return router
}
```



然后在main.go中调用该方法， 并启动gin服务

```go
package main

import (
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

	router := router.Router()
	router.Run(":8000")
}
```

这样gin成功集成进项目中了。



### 用户模块api的开发

现在开始正式编写完整的api了，首先我们在service目录下新建user.go文件

现在需要思考元一下项目中需要哪些用户模块的api

1. 用户列表
2. 密码登录
3. 注册用户
4. 更新用户信息
5. 账号注销



要下面我们来一一实现



#### 用户列表

这是一个get方法， 可以提供管理员，目前还没有进行管理员认证鉴权，所以此时所有用户都可以调用这个api

```go
func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1, //0 表示成功， -1 表示失败
			"message": "获取用户列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, list)
}
```

这里将用户所有的信息都无脑返回了，其实是不符合我们接口设计的，有兴趣的小伙伴可以将不该返回的数据屏蔽。



#### 密码登录

这里就应该是一个post方法了，用户提交用户名和密码，然后需要对密码和用户名进行核验。

```go
func LoginByNameAndPassWord(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	data, err := dao.FindUserByName(name)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1, //0 表示成功， -1 表示失败
			"message": "登录失败",
		})
		return
	}

	if data.Name == "" {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名不存在",
		})
		return
	}

  //由于数据库密码保存是使用md5密文的， 所以验证密码时，是将密码再次加密，然后进行对比，后期会讲解md:common.CheckPassWord
	ok := common.CheckPassWord(password, data.Salt, data.PassWord)
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}

	Rsp, err := dao.FindUserByNameAndPwd(name, data.PassWord)
	if err != nil {
		zap.S().Info("登录失败", err)
	}

  //这里使用jwt做权限认证，后面将会介绍
	token, err := middlewear.GenerateToken(Rsp.ID, "yk")
	if err != nil {
		zap.S().Info("生成token失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"tokens":  token,
		"userId":  Rsp.ID,
	})
}
```



#### 注册用户

使用post方法， 用户输入注册用户名、两次输入密码

```go
func NewUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Request.FormValue("name")
	password := ctx.Request.FormValue("password")
	repassword := ctx.Request.FormValue("Identity")

	if user.Name == "" || password == "" || repassword == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "用户名或密码不能为空！",
			"data":    user,
		})
		return
	}

	//查询用户是否存在
	_, err := dao.FindUser(user.Name)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户已注册",
			"data":    user,
		})
		return
	}

	if password != repassword {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "两次密码不一致！",
			"data":    user,
		})
		return
	}

	//生成盐值
	salt := fmt.Sprintf("%d", rand.Int31())
  
  //加密密码
	user.PassWord = common.SaltPassWord(password, salt)
	user.Salt = salt
	t := time.Now()
	user.LoginTime = &t
	user.LoginOutTime = &t
	user.HeartBeatTime = &t
	dao.CreateUser(user)
	ctx.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "新增用户成功！",
		"data":    user,
	})
}
```





#### 更新用户信息

可以使用post方法或者put方法，这里我们依然使用post方法吧

```go
func UpdataUser(ctx *gin.Context) {
	user := models.UserBasic{}

	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	if err != nil {
		zap.S().Info("类型转换失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "注销账号失败",
		})
		return
	}
	user.ID = uint(id)
	Name := ctx.Request.FormValue("name")
	PassWord := ctx.Request.FormValue("password")
	Email := ctx.Request.FormValue("email")
	Phone := ctx.Request.FormValue("phone")
	avatar := ctx.Request.FormValue("icon")
	gender := ctx.Request.FormValue("gender")
	if Name != "" {
		user.Name = Name
	}
	if PassWord != "" {
		salt := fmt.Sprintf("%d", rand.Int31())
		user.Salt = salt
		user.PassWord = common.SaltPassWord(PassWord, salt)
	}
	if Email != "" {
		user.Email = Email
	}
	if Phone != "" {
		user.Phone = Phone
	}
	if avatar != "" {
		user.Avatar = avatar
	}
	if gender != "" {
		user.Gender = gender
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		zap.S().Info("参数不匹配", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "参数不匹配",
		})
		return
	}

	Rsp, err := dao.UpdateUser(user)
	if err != nil {
		zap.S().Info("更新用户失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "修改信息失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "修改成功",
		"data":    Rsp.Name,
	})
}
```



#### 账号注销

```go
func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	if err != nil {
		zap.S().Info("类型转换失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "注销账号失败",
		})
		return
	}

	user.ID = uint(id)
	err = dao.DeleteUser(user)
	if err != nil {
		zap.S().Info("注销用户失败", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "注销账号失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "注销账号成功",
	})
}
```





### 配置路由

```go
package router

import (
	"HiChat/middlewear"
	"HiChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//初始化路由
	router := gin.Default()

	//v1版本
	v1 := router.Group("v1")

	//用户模块，后续有个用户的api就放置其中
	user := v1.Group("user")
	{
		user.GET("/list", service.List)
		user.POST("/login_pw", service.LoginByNameAndPassWord)
		user.POST("/new", service.NewUser)
		user.DELETE("/delete", service.DeleteUser)
		user.POST("/updata", service.UpdataUser)
	}
	return router
}
```

这样整个api就完成了：



### 测试api

#### 启动服务

```
go run main.go
```

可以看到：

```
[GIN-debug] GET    /v1/user/list             --> HiChat/service.List (4 handlers)
[GIN-debug] POST   /v1/user/login_pw         --> HiChat/service.LoginByNameAndPassWord (3 handlers)
[GIN-debug] POST   /v1/user/new              --> HiChat/service.NewUser (4 handlers)
[GIN-debug] DELETE /v1/user/delete           --> HiChat/service.DeleteUser (4 handlers)
[GIN-debug] POST   /v1/user/updata           --> HiChat/service.UpdataUser (4 handlers)
```



#### 测试

##### /v1/user/list

这里需要您之前写入一些数据，用slq或者到层的CreatUser()方法

当我们访问：http://127.0.0.1:8000/v1/user/list

```json
[ 
  {
        "ID": 8,
        "CreatedAt": "2022-12-22T19:17:16.365+08:00",
        "UpdatedAt": "2022-12-22T19:17:16.365+08:00",
        "DeletedAt": null,
        "Name": "ice_moss",
        "PassWord": "d41d8cd98f00b204e9800998ecf8427e$1298498081",
        "Avatar": "",
        "Gender": "male",
        "Phone": "",
        "Email": "",
        "Identity": "",
        "ClientIp": "",
        "ClientPort": "",
        "Salt": "1298498081",
        "LoginTime": "2022-12-22T19:17:16.363+08:00",
        "HeartBeatTime": "2022-12-22T19:17:16.363+08:00",
        "LoginOutTime": "2022-12-22T19:17:16.363+08:00",
        "IsLoginOut": false,
        "DeviceInfo": ""
    },
    {
        "ID": 9,
        "CreatedAt": "2022-12-22T19:30:34.893+08:00",
        "UpdatedAt": "2022-12-22T19:30:34.893+08:00",
        "DeletedAt": null,
        "Name": "ice_moss1",
        "PassWord": "d41d8cd98f00b204e9800998ecf8427e$1298498081",
        "Avatar": "",
        "Gender": "male",
        "Phone": "",
        "Email": "",
        "Identity": "",
        "ClientIp": "",
        "ClientPort": "",
        "Salt": "1298498081",
        "LoginTime": "2022-12-22T19:30:34.892+08:00",
        "HeartBeatTime": "2022-12-22T19:30:34.892+08:00",
        "LoginOutTime": "2022-12-22T19:30:34.892+08:00",
        "IsLoginOut": false,
        "DeviceInfo": ""
    },
    {
        "ID": 10,
        "CreatedAt": "2022-12-22T19:37:19.508+08:00",
        "UpdatedAt": "2022-12-24T16:38:56.717+08:00",
        "DeletedAt": null,
        "Name": "ice_moss2",
        "PassWord": "0192023a7bbd73250516f069df18b500$1298498081",
        "Avatar": "https://mxshopfiles.oss-cn-shanghai.aliyuncs.com/work/103800kbdgbv2zdv1vnnrd.jpeg",
        "Gender": "male",
        "Phone": "",
        "Email": "",
        "Identity": "9fce97499eea554562d27d086da558e3",
        "ClientIp": "",
        "ClientPort": "",
        "Salt": "1298498081",
        "LoginTime": "2022-12-22T19:37:19.507+08:00",
        "HeartBeatTime": "2022-12-22T19:37:19.507+08:00",
        "LoginOutTime": "2022-12-22T19:37:19.507+08:00",
        "IsLoginOut": false,
        "DeviceInfo": ""
    },
    {
        "ID": 11,
        "CreatedAt": "2022-12-24T16:51:53.418+08:00",
        "UpdatedAt": "2022-12-24T18:26:06.611+08:00",
        "DeletedAt": null,
        "Name": "ice_moss4",
        "PassWord": "0192023a7bbd73250516f069df18b500$1298498081",
        "Avatar": "",
        "Gender": "male",
        "Phone": "",
        "Email": "",
        "Identity": "5993bcfc7b16b8a84e22aefc6b42a528",
        "ClientIp": "",
        "ClientPort": "",
        "Salt": "1298498081",
        "LoginTime": "2022-12-24T16:51:53.417+08:00",
        "HeartBeatTime": "2022-12-24T16:51:53.417+08:00",
        "LoginOutTime": "2022-12-24T16:51:53.417+08:00",
        "IsLoginOut": false,
        "DeviceInfo": ""
    }
 ]
```



##### /v1/user/login_pw

post方法我们不能直接使用浏览器，这里推荐使用postman

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-28%20%E4%B8%8B%E5%8D%885.55.10.png)

其他api的测试方法都一样， 这里就不重复了。



### 总结

到这里，我们用户模块的功能api就开发完成了，其实也很简单，就是获取到请求参数，然后进行判断，最后调用dao层的方法，但是依然不完整，例如我们的获取用户列表，更新，删除等功能，我们直接调用api就可以完成，那岂不是每一个人都能操作我们是数据库了，所以接下来，需要做的就是限制用户请求，只能需要修改他当前用户下的数据，称为鉴权，在测试登录用户时，返回了一个字段tokens,我们便是是token来实现，鉴权的，当然还需要介绍一下密码的加密功能。

