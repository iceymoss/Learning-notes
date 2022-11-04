下面我们来学习一下基于MongoDB数据库的实战学习，我们这一章内容将来介绍微信小程序如何实现用户登录唯一id标识，简单的来说就是，用户登录小程序后，我们向数据库存入该用户的唯一id然后，当用户再一次登录时，我们可以直接从数据库中拿到该用户的唯一id标识，来证明该用户的身份，具体流程图如下：

![20210706153655258](/Users/feng/Desktop/20210706153655258.png)

下面我们需要来定义服务，这里我们使用GRPC来实现api，引领开发：

auth.proto:定义服务

```
syntax = "proto3";
package auth.v1;
option go_package="coolcar/auth/api/gen/v1;authpb";

message LoginRequest{
    string code = 1;  //向服务器发送code
}

message LoginResponse{   //服务器上传code到微信api返回服务器，服务器自定义token和token有效时间
    string accss_token = 1; 
    int32 expires_in = 2;
}
//接口
service AuthService{
    rpc Login(LoginRequest) returns (LoginResponse);
}
```

auth.yaml: 对外暴露接口

```
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: auth.v1.AuthService.Login
      post: /v1/auth/login
      body: "*"
```

接着我们使用命令：

```
protoc -I=. --go_out=plugins=grpc,paths=source_relative:gen/go auth.proto
```

然后GRPC就会给我生成客户端和服务端的api:

我们只需要来实现接口：

```go
// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}
```

这里我们就来实现 "客户端上传code" -> "返回自定义登录态"

实现接口：auth服务层，面向客户端

```go
package auth

import (
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"fmt"
	"time"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	TokenGenerate  TokenGenerate
	TokenExpire    time.Duration
	Mongo          *dao.Mongo
	Logger         *zap.Logger
}


//将客户端上传的code，和小程序ID和秘钥上传至微信api换取openID
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

//生成token接口
type TokenGenerate interface {
	GeneratorToken(accountID string, expire time.Duration) (string, error)
}

//后台事件处理方法
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code",zap.String("code", req.Code))
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "获取不到openID")
	}
	//将openID存入数据库，返回对应_id
	accountID, err := s.Mongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("不能解析到accountID", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	//使用accountID生成token
	token, err := s.TokenGenerate.GeneratorToken(accountID.String(), s.TokenExpire)
	if err != nil {
		s.Logger.Error("不能生成token")
		return nil, status.Errorf(codes.Internal, "")
	}

	fmt.Printf("openID: %v", openID)
	return &authpb.LoginResponse{
		AccssToken: token,
		ExpiresIn:  int32(s.TokenExpire.Seconds()),
	}, nil
}

```

在auth中声明了两个接口：

```go
//将客户端上传的code，和小程序ID和秘钥上传至微信api换取openID
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

//生成token接口
type TokenGenerate interface {
	GeneratorToken(accountID string, expire time.Duration) (string, error)
}
```

相应功能都有注释

1. ```type OpenIDResolver interface{}```

```go
package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v2"
)

//Service is a wechat auth service
type Service struct {
	AppId     string
	Appsecret string
}

//将客户端上传的code，和小程序ID和秘钥上传至微信api换取openID
func (s *Service) Resolve(code string) (string, error) {
	resp, err := weapp.Login(s.AppId, s.Appsecret, code)
	if err != nil {
		return "", fmt.Errorf("weapp login: %v", err)
	}
	if err = resp.GetResponseError(); err != nil {
		return "", fmt.Errorf("weapp response error: %v", err)
	}
	return resp.OpenID, nil
}
```





2. ```type TokenGenerate interface{}```

```go
package token

import (
	"crypto/rsa"
	"time"
	"github.com/dgrijalva/jwt-go"
)

//生成一个jwt的token
type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	newFunc    func() time.Time
}

//构造函数
func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		newFunc:    time.Now,
		privateKey: privateKey,
	}
}

//Generator(accountID string, expire time.Duration)(string, error)
func (t *JWTTokenGen) GeneratorToken(accountID string, expire time.Duration) (string, error) {
	newSec := t.newFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.issuer,
		IssuedAt:  newSec,
		ExpiresAt: newSec + int64(expire.Seconds()),
		Subject:   accountID,
	})

	return tkn.SignedString(t.privateKey)
}
```

在auth.go中：

```go
//将openID存入数据库，返回对应_id
	accountID, err := s.Mongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("不能解析到accountID", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
```

这里我们需要将微信api返回的openID存入数据库中，然后返回该文档的_id

```go
package dao

import (
	"context"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
  openidfield = "open_id"
  IDFieldName = "_id"
  
)


//定义一个 Mongo 类型
type Mongo struct {
	col *mongo.Collection
}

//初始化数据库， 类似构造函数
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

//将openID存入数据库，返回对应_id给用户
func (m *Mongo) ResolveAccountID(c context.Context, openID string) (id.AccountID, error) {
	//先生成一个primitive.ObjectID类型作为文档索引
	insertedID := mgo.NewObjID()
	//然后再去查找openID，如果查到原来的openID,没有则插入我们固定的insertedID,然后将对应_id返回出来
	res := m.col.FindOneAndUpdate(c, bson.M{
		openidfield: openID,
	}, SetInsert(bson.M{
		mgo.IDFieldName: insertedID,
		openidfield:     openID,
	}), options.FindOneAndUpdate().SetUpsert(true).
		SetReturnDocument(options.After))
	//检测是否返回成功
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}
	var row mgo.IDField
	//解码
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot Decode result: %v", err)
	}
	return objid.ToAccountID(row.ID), nil
}


func SetInsert(V interface{}) bson.M {
	return bson.M{
		"$setOnInsert": V,
	}
}
```

这样第一条来回的线路我们就走通了



然后我们来实现第二条路线：

这里我们需要实现一个拦截器，来获取从客户端发的数据

```go
package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHandle = "authorization"
	BearerProfix        = "Bearer "
)

// Intercetor创建一个auth的拦截器
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("不能打开文件: %v", err)
	}

	rea, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("不能读取到文件: %v", f)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(rea)
	if err != nil {
		return nil, fmt.Errorf("不能解析key: %v", err)
	}

	i := &interceptor{
		Verifier: &token.JWTTokenVerifier{
			PublicKey: key,
		},
	}
	fmt.Printf("Intercetor结束:\n")

	return i.HandleReq, nil
}

//声明接口
type toekenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	Verifier toekenVerifier
}

//ctx请求，req请求内容， info帮助文档，handle接下来要做的
func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("HandleReq结束:\n")
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token已经过期: %v", err)
	}
	aid, err := i.Verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token已经过期: %v", err)
	}
	//把accountID放入ctx中

	fmt.Printf("HandleReq结束:\n")

	return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
}

//解析数据
func tokenFromContext(c context.Context) (string, error) {
	//使用 metadata.FromIncomingContext 方法进行读取,创建写入ctx的数据类似m
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "解析数据失败")
	}
	tkn := ""
	//将token分离出来
	for _, v := range m[authorizationHandle] {
		if strings.HasPrefix(v, BearerProfix) {
			tkn = v[len(BearerProfix):]
		}
	}
	if tkn == "" {
		return "", status.Errorf(codes.Unauthenticated, "tkn仍为空串")
	}
	fmt.Printf("tokenFromContex结束:\n")
	return tkn, nil
}

type accountKeyID struct{}

//ContextWithAccountID将数据放入context中
func ContextWithAccountID(c context.Context, aid id.AccountID) context.Context {
	fmt.Printf("ContextWithAccountID结束:\n")
	return context.WithValue(c, accountKeyID{}, aid)
}

//AccountIDWithContext将context中的数据aid拿出
func AccountIDFromContext(c context.Context) (id.AccountID, error) {
	v := c.Value(accountKeyID{})
	aid, ok := v.(id.AccountID)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "不能获取aid")
	}
	fmt.Printf("AccountIDFromContext结束:\n")
	return aid, nil
}

```

同样我们声明了接口：

```go
type toekenVerifier interface {
	Verify(token string) (string, error)
}
```

接口的实现:

```go
package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

//解析token，返回accountID
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("不能解析token: %v", err)
	}

	if !t.Valid {
		return "", fmt.Errorf("无效的token")
	}

	cli, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claims 不是一个standardClims: %v", ok)
	}
	//验证Claims,里面的所有的字段，例如： "exp"
	if err := cli.Valid(); err != nil {
		return "", fmt.Errorf("无效的Cliams: err")
	}
	return cli.Subject, nil
}

```

这样我们在后面的服务中，就可以直接拿出我们的accountID了，下面我们来启动服务：

main.go:

```go
package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
）
  
func main(){
    //使用zap包，打印日志
    logger, err := zap.NewDevelopment()
    if err != nil {
      log.Fatalf("cannot creat logger: %v", err)
    }
    //tcp,监听8081端口
    list, err := net.Listen("tcp", ":8081")
    if err != nil {
      logger.Fatal("连接失败：%v", zap.Error(err))
    }

    //连接Mongo数据库
    c := context.Background()
    mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar"))
    if err != nil {
      logger.Fatal("不能连接Mongo数据库: %v", zap.Error(err))
    }
    //打开读取文件
    pkfile, err := os.Open("../auth/private.key")
    if err != nil {
      logger.Fatal("打开文件失败: %v", zap.Error(err))
    }
    pkByte, err := ioutil.ReadAll(pkfile)
    if err != nil {
      logger.Fatal("读取失败: %v", zap.Error(err))
    }
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkByte)
    if err != nil {
      logger.Fatal("解析失败: %v", zap.Error(err))
    }
    s := grpc.NewServer()                              //创建s
    authpb.RegisterAuthServiceServer(s, &auth.Service{ //注册服务
      OpenIDResolver: &wechat.Service{
        AppId:     "***********",   //微信小程序id
        Appsecret: "***********",   //微信小程序秘钥
      },
      Mongo:         dao.NewMongo((mongoClient.Database("coolcar"))),
      Logger:        logger,
      TokenExpire:   2 * time.Hour,
      TokenGenerate: token.NewJWTTokenGen("coolcar/auth", privateKey),
    })
    err = s.Serve(list) //开启服务
    logger.Fatal("connot sever", zap.Error(err))
}
```



