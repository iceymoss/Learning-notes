[toc]

### 概况

前面我们学习了api的开发及api的测试，但是也提到了api权限问题，那么这里将来介绍jwt授权和鉴权以及Md5盐值加密。

```
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





### JWT

#### 什么是jwt

jwt全称：json web token 

其实就是以json格式颁发的web服务使用的令牌，有了令牌就拥有了访问一些被保护资源的权利。



#### 分类

##### 对称加密

 指使用同一把钥匙开门，授权和鉴权都使用同一份秘钥进行验证

例如 ： 

>秘钥为：dhdifidfhiadfdhifhdf
>
>载入信息：
>
>```json
>{
>  "sub": "1234567890",
>  "name": "ice_moss",
>  "userId":"14" 
>}
>```
>
>加密方式、类型：
>
>```json
>{
>  "alg": "HS256",
>  "typ": "JWT"
>}
>
>```

>生成token:
>
>```
>eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImljZV9tb3NzIiwidXNlcklkIjoiMTQifQ.uPug9OgNZXZ0NutD26u81Q4HNQPIWfYauY6vvkZLiQM
>```

我们使用这段token就可以验证， 当前验证人是合法的。



##### 非对称加密

使用私钥加密，使用公钥进行验证

> 加密算法和类型：
>
> ```json
> {
>   "alg": "RS512",
>   "typ": "JWT"
> }
> ```
>
> 载入信息：
>
> ```json
> {
>   "sub": "1234567890",
>   "name": "ice_moss"
> }
> ```
>
> 

有两分秘钥：公钥和私钥

>公钥：
>
>```
>-----BEGIN PUBLIC KEY-----
>MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu1SU1LfVLPHCozMxH2Mo
>4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0/IzW7yWR7QkrmBL7jTKEn5u
>+qKhbwKfBstIs+bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyeh
>kd3qqGElvW/VDL5AaWTg0nLVkjRo9z+40RQzuVaE8AkAFmxZzow3x+VJYKdjykkJ
>0iT9wCS0DRTXu269V264Vf/3jvredZiKRkgwlL9xNAwxXFg0x/XFw005UWVRIkdg
>cKWTjpBP2dPwVZ4WWC+9aGVd+Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbc
>mwIDAQAB
>-----END PUBLIC KEY-----
>```
>
>
>
>私钥：
>
>```
>-----BEGIN PRIVATE KEY-----
>MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC7VJTUt9Us8cKj
>MzEfYyjiWA4R4/M2bS1GB4t7NXp98C3SC6dVMvDuictGeurT8jNbvJZHtCSuYEvu
>NMoSfm76oqFvAp8Gy0iz5sxjZmSnXyCdPEovGhLa0VzMaQ8s+CLOyS56YyCFGeJZ
>qgtzJ6GR3eqoYSW9b9UMvkBpZODSctWSNGj3P7jRFDO5VoTwCQAWbFnOjDfH5Ulg
>p2PKSQnSJP3AJLQNFNe7br1XbrhV//eO+t51mIpGSDCUv3E0DDFcWDTH9cXDTTlR
>ZVEiR2BwpZOOkE/Z0/BVnhZYL71oZV34bKfWjQIt6V/isSMahdsAASACp4ZTGtwi
>VuNd9tybAgMBAAECggEBAKTmjaS6tkK8BlPXClTQ2vpz/N6uxDeS35mXpqasqskV
>laAidgg/sWqpjXDbXr93otIMLlWsM+X0CqMDgSXKejLS2jx4GDjI1ZTXg++0AMJ8
>sJ74pWzVDOfmCEQ/7wXs3+cbnXhKriO8Z036q92Qc1+N87SI38nkGa0ABH9CN83H
>mQqt4fB7UdHzuIRe/me2PGhIq5ZBzj6h3BpoPGzEP+x3l9YmK8t/1cN0pqI+dQwY
>dgfGjackLu/2qH80MCF7IyQaseZUOJyKrCLtSD/Iixv/hzDEUPfOCjFDgTpzf3cw
>ta8+oE4wHCo1iI1/4TlPkwmXx4qSXtmw4aQPz7IDQvECgYEA8KNThCO2gsC2I9PQ
>DM/8Cw0O983WCDY+oi+7JPiNAJwv5DYBqEZB1QYdj06YD16XlC/HAZMsMku1na2T
>N0driwenQQWzoev3g2S7gRDoS/FCJSI3jJ+kjgtaA7Qmzlgk1TxODN+G1H91HW7t
>0l7VnL27IWyYo2qRRK3jzxqUiPUCgYEAx0oQs2reBQGMVZnApD1jeq7n4MvNLcPv
>t8b/eU9iUv6Y4Mj0Suo/AU8lYZXm8ubbqAlwz2VSVunD2tOplHyMUrtCtObAfVDU
>AhCndKaA9gApgfb3xw1IKbuQ1u4IF1FJl3VtumfQn//LiH1B3rXhcdyo3/vIttEk
>48RakUKClU8CgYEAzV7W3COOlDDcQd935DdtKBFRAPRPAlspQUnzMi5eSHMD/ISL
>DY5IiQHbIH83D4bvXq0X7qQoSBSNP7Dvv3HYuqMhf0DaegrlBuJllFVVq9qPVRnK
>xt1Il2HgxOBvbhOT+9in1BzA+YJ99UzC85O0Qz06A+CmtHEy4aZ2kj5hHjECgYEA
>mNS4+A8Fkss8Js1RieK2LniBxMgmYml3pfVLKGnzmng7H2+cwPLhPIzIuwytXywh
>2bzbsYEfYx3EoEVgMEpPhoarQnYPukrJO4gwE2o5Te6T5mJSZGlQJQj9q4ZB2Dfz
>et6INsK0oG8XVGXSpQvQh3RUYekCZQkBBFcpqWpbIEsCgYAnM3DQf3FJoSnXaMhr
>VBIovic5l0xFkEHskAjFTevO86Fsz1C2aSeRKSqGFoOQ0tmJzBEs1R6KqnHInicD
>TQrKhArgLXX4v3CddjfTRJkFWDbE/CkvKZNOrcf1nhaGCPspRJj2KUkj1Fhl9Cnc
>dn/RsYEONbwQSjIfMPkvxF+8HQ==
>-----END PRIVATE KEY-----
>```
>
>
>
>生成token：
>
>```
>eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImljZV9tb3NzIn0.K8ZCGpv2e7B6YUBtxZwLwh9YEQqHe6FxUHXhBn99EAmOoPsnnn7Jam5A34TGbWKbQHnrTvXkpy14WMIA9fFW3TQBbMPD4LSRMG_T2nyxRJspaWGWaVnYAqx16FCTcHr_UNjQ1AMc-GkN5InlkbRZiwck1GLgR6B2kSnHc6RqfGDRiM1Z5gG-quPmukzFvW9aGq0XoNpb73itI_6CH-OzaykwgNMWSj4vM-frvXWwKFD9ICYcXk2wZ1p3-3v59G75s9JqWV581pVNHXH7mA-x074sPar7oOg9HEUotGugc3LxUzEIqh4GkMS0f3ANaMtH2yy6m759P0E5p8_t3A6oGA
>```

使用私钥生成token，验证时使用公钥进行验证



我们的项目中使用非对称加密进行授权和鉴权。



### 项目集成jwt

#### 拉取依赖

```
go get github.com/dgrijalva/jwt-go
```

项目中我们分为两步：1. 授权； 2.鉴权

#### 授权

token的生成，我们在middlewear目录下新建一个jwt.go文件

```go
var (
	TokenExpired = errors.New("Token is expired")
)

// 指定加密密钥
var jwtSecret = []byte("ice_moss")

//Claims 是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

//GenerateToken 根据用户的用户名和密码产生token
func GenerateToken(userId uint, iss string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(48 * 30 * time.Hour)

	claims := Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: iss,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
```



#### 鉴权

下面就是，用户在http请求中将token带上，然后我们就需要进行验证。

```go
var (
	TokenExpired = errors.New("Token is expired")
)

// 指定加密密钥
var jwtSecret = []byte("ice_moss")

//Claims 是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}


func JWY() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		user := c.Query("userId")
		userId, err := strconv.Atoi(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "您userId不合法",
			})
			c.Abort()
			return
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "请登录",
			})
			c.Abort()
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "token失效",
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				err = TokenExpired
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "授权已过期",
				})
				c.Abort()
				return
			}

			if claims.UserID != uint(userId) {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "您的登录不合法",
				})
				c.Abort()
				return
			}

			fmt.Println("token认证成功")
			c.Next()
		}
	}
}


//ParseToken 根据传入的token值获取到Claims对象信息（进而获取其中的用户id）
func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

```



#### jwt.go完整代码

```go
package middlewear

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	TokenExpired = errors.New("Token is expired")
)

// 指定加密密钥
var jwtSecret = []byte("ice_moss")

//Claims 是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

func JWY() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		user := c.Query("userId")
		userId, err := strconv.Atoi(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "您userId不合法",
			})
			c.Abort()
			return
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "请登录",
			})
			c.Abort()
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "token失效",
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				err = TokenExpired
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "授权已过期",
				})
				c.Abort()
				return
			}

			if claims.UserID != uint(userId) {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "您的登录不合法",
				})
				c.Abort()
				return
			}

			fmt.Println("token认证成功")
			c.Next()
		}
	}
}

//GenerateToken 根据用户的用户名和密码产生token
func GenerateToken(userId uint, iss string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(48 * 30 * time.Hour)

	claims := Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: iss,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseToken 根据传入的token值获取到Claims对象信息（进而获取其中的用户id）
func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
```





### jwt加入gin中间件

之前我们看到在登录api代码中

```go
token, err := middlewear.GenerateToken(Rsp.ID, "yk")
	if err != nil {
		zap.S().Info("生成token失败", err)
		return
	}
```



现在我们要如何进行鉴权呢?

答案是：在路由中

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
		user.GET("/list",middlewear.JWY(), service.List)
		user.POST("/login_pw",middlewear.JWY(), service.LoginByNameAndPassWord)
		user.POST("/new",middlewear.JWY(), service.NewUser)
		user.DELETE("/delete",middlewear.JWY(), service.DeleteUser)
		user.POST("/updata",middlewear.JWY(), service.UpdataUser)
	}
	return router
}
```

这样我们现在访问执行api,需要在body中加入token了。当再次请求用户列表时没有加token：

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-28%20%E4%B8%8B%E5%8D%887.19.10.png" style="zoom:50%;" />

添加token后：

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-28%20%E4%B8%8B%E5%8D%887.21.40.png" style="zoom:50%;" />



现在整个用户模块api就完整了。



### MD5密码加密

之前的文章也看到了，在存储密码时没有使用明文存储，而是存储使用md加密后的密文，下面来看看这个加密过程。

#### 什么是MD5

参考文章：[MD5加密](https://learnku.com/articles/69126)



#### 项目中MD5的使用

将加密文件放置common目录下，命名为md5.go

```go
package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

//Md5encoder 加密后返回小写值
func Md5encoder(code string) string {
	m := md5.New()
	io.WriteString(m, code)
	return hex.EncodeToString(m.Sum(nil))
}

//Md5StrToUpper 加密后返回大写
func Md5StrToUpper(code string) string {
	return strings.ToUpper(Md5encoder(code))
}

//SaltPassWord 密码加盐
func SaltPassWord(pw string, salt string) string {
	saltPW := fmt.Sprintf("%s$%s", Md5encoder(pw), salt)
	return saltPW
}

//CheckPassWord 核验密码
func CheckPassWord(rpw, salt, pw string) bool {
	return pw == SaltPassWord(rpw, salt)
}
```



### 总结

到目前为止，完成了用户模块的基本开发，相关功能已经进行了测试，整套完整的用户服务api开发完成， 当然还有很多不做， 希望小伙伴们可以进行优化，这个模块涉及的知识点还有挺多的，如有不足，错误等欢迎评论区指正；接下来的聊天模块的任务就比较重了，也是我们HiChat项目的核心，大家一起加油！

