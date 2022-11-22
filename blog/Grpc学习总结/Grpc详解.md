[toc]



### 文章介绍

在阅读本文之前建议您先阅读[RPC核心概念理解](https://learnku.com/articles/67348)和[GRPC之protobuf理解](https://learnku.com/articles/68062)，在本文中我们不再过多介绍GRPC的基础部分，主要介绍：grpc的metadata机制、GRPC拦截器、通过拦截器和metadata实现GRPC的auth认证、grpc状态码、grpc的超时机制和protoc文件生成的go源码解析



### GRPC的metadata机制

* ###### metadata的建立

  1. 使用metadata.New()

     这里需要注意：返回的metadata.New()返回的是类型：```type MD map[string][]string```，而我们的metadata.New()方法是入参是：```map[string]string```,下面我们来看看metadata.New()方法的实现:

     ```go
     type MD map[string][]string
     
     func New(m map[string]string) MD {
     	md := MD{}
     	for k, val := range m {
     		key := strings.ToLower(k)
     		md[key] = append(md[key], val)
     	}
     	return md
     }
     ```

     示例：

     ```go
     md := metadata.New(map[string]string{
     		"name":     "ice_moss",
     		"password": "ice_12345",
     	})
     ```



   2. 使用metadata.Pairs()

      同样我们可以来查看一下源码：

      ```go
      type MD map[string][]string
      
      func Pairs(kv ...string) MD {
      	if len(kv)%2 == 1 {
      		panic(fmt.Sprintf("metadata: Pairs got the odd number of input pairs for metadata: %d", len(kv)))
      	}
      	md := MD{}
      	for i := 0; i < len(kv); i += 2 {
      		key := strings.ToLower(kv[i])
      		md[key] = append(md[key], kv[i+1])
      	}
      	return md
      }
      ```

      仔细看metadata.Pairs()的源码可以看出我们的入参个数必须是偶数，并且每两两为键值对key-value

      示例：

      ```go
      md := metadata.Pairs(
      	"name", "ice_moss",
      	"password", "ice_moss12345"
      )
      ```

      即:

      ```json
      name: ice_moss
      password: ice_moss12345
      ```





* ###### metadata数据的发送和接收

  在此之前我们先来定义一个proto文件，来满足我们下面的需要的metadata数据使用场景

  Proto:

  ```protobuf
  syntax = "proto3";
  
  option go_package="/.;proto";
  //引入protobuf的内置类型
  import "google/protobuf/timestamp.proto";
  
  //定义接口
  service Greeter {
      rpc SayHello (HelloRequest) returns (HelloReply);
  }
  
  //枚举类型
  enum Gender{
      MALE = 0;
      FE_MALE = 1;
  }
  
  message HelloRequest {
      string name = 1;
      string url = 2;
      Gender gender = 3;
      map<string, string> m = 4;  //proto map类型
      google.protobuf.Timestamp addTime = 5;  //protobuf的内置类型
  }
  
  message HelloReply {
      string id = 1;
      HelloRequest request = 2;
  }
  ```

  以上代码有不理解的可以阅读：[GRPC之protobuf理解](https://learnku.com/articles/68062)

  这里我们来模拟客服端client和服务端server进行交互，来进行metadata数据的发送和接收,这里我们需要知道的是metadata是通过go语言中的context的上下文来传输的

  ###### server:

  ```go
  package main
  
  import (
  	"context"
  	"log"
  	"net"
  	"rpcstudy/metadata_test/proto"
  
  	"google.golang.org/grpc"
  )
  
  type HelloSerivce struct{}
  
  //SayHello实现对用户数据和id绑定
  func (h *HelloSerivce) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    //通过metadata.FromIncomingContext()方法获取context的上下文
    md, ok := metadata.FromIncomingContext(c)
  	if !ok {
  		log.Fatal("FromIncomingContext error")
  	}
  	for key, value := range md {
  		fmt.Printf("key:%v, value:%v\n", key, value)
  	}
  	return &proto.HelloReply{
  		Id:      "123456789",
  		Request: req,
  	}, nil
  }
  
  func main() {
  	//实例化server
  	lit, err := net.Listen("tcp", ":9090")
  	if err != nil {
  		log.Panicln("监听失败", err)
  	}
  
    //创建grpc服务器
  	s := grpc.NewServer()  
    //注册处理逻辑
  	proto.RegisterGreeterServer(s, &HelloSerivce{})
  
  	//开启服务，这里理解为将连接连入grpc服务器中，就可以完成通讯
  	log.Println(s.Serve(lit))
  }
  ```

  写完server并运行server,等待请求进入

  ###### client:

  ```go
  package main
  
  import (
  	"context"
  	"fmt"
  	"log"
  	"rpcstudy/metadata_test/proto"
  	"time"
  
  	"google.golang.org/grpc/metadata"
  
  	"google.golang.org/grpc"
  	timepb "google.golang.org/protobuf/types/known/timestamppb"
  )
  
  func main() {
    //使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
  	clientConn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
  	if err != nil {
  		log.Panicln("连接失败", err)
  	}
  	defer clientConn.Close()
  
  	md := metadata.New(map[string]string{
  		"name":     "ice_moss",
  		"password": "ice_12345",
  	})
  
  	//发送metadata
  	//NewOutgoingContext创建一个新的上下文，并附加了md放入context中，并返回context
  	//即：新建一个有metada的context
  	ctx := metadata.NewOutgoingContext(context.Background(), md)
  
  	//连接grpc服务器
  	client := proto.NewGreeterClient(clientConn)
  	//单向传送至server，将带有md的上下文ctx入参
  	res, err := client.SayHello(ctx, &proto.HelloRequest{
  		Name:   "kuangyang",
  		Url:    "https://learnku.com/blog/yangkuang",
  		Gender: proto.Gender_MALE,
  		M: map[string]string{
  			"来自": "北京",
  			"现居": "上海",
  		},
  		AddTime: timepb.New(time.Now()),
  	})
  	if err != nil {
  		panic(err)
  	}
  	fmt.Printf("返回结果: %v", res)
  
  }
  ```

  运行client,最后client返回：

  ```shell
  返回结果: id:"123456789" request:{name:"kuangyang" url:"https://learnku.com/blog/yangkuang" m:{key:"来自" value:"北京"} m:{key:"现居" value:"上海"} addTime:{seconds:165s:929482000}}
  ```

  server输出：

  ```shell
  key:user-agent, value:[grpc-go/1.46.2]
  key:name, value:[ice_moss]
  key:password, value:[ice_12345]
  key::authority, value:[localhost:9090]
  key:content-type, value:[application/grpc]
  ```

  输出的内容中我们就可以看到客服端通过metadata请求发过来的数据，当然里面还有一些其他的数据，其实就是我们的handler，和http的handler类似

  

  >#### 总结一下：
  >
  >>创建metadata：
  >
  >>```go
  >>metadata.New()
  >>metadata.Pairs()
  >>```
  >
  >>将metadata数据放入context中：
  >
  >>```go
  >>//NewOutgoingContext(ctx, md)
  >>```
  >
  >>//函数声明
  >>func NewOutgoingContext(ctx context.Context, md MD) context.Context {
  >>	return context.WithValue(ctx, mdOutgoingKey{}, rawMD{md: md})
  >>}
  >>
  >>```
  >>
  >>```
  >
  >>将metadata数据从context中拿出:
  >
  >>```go
  >>metadata.FromIncomingContext(c)
  >>```
  >
  >>//函数声明
  >>func FromIncomingContext(ctx context.Context) (MD, bool) {
  >>	md, ok := ctx.Value(mdIncomingKey{}).(MD)
  >>	if !ok {
  >>		return nil, false
  >>	}
  >>	out := MD{}
  >>	for k, v := range md {
  >>		// We need to manually convert all keys to lower case, because MD is a
  >>		// map, and there's no guarantee that the MD attached to the context is
  >>		// created using our helper functions.
  >>		key := strings.ToLower(k)
  >>		s := make([]string, len(v))
  >>		copy(s, v)
  >>		out[key] = s
  >>	}
  >>	return out, true
  >>}
  >>
  >>```
  >>
  >>```



### GRPC的拦截器

我们在开发的过程中验证用户身份、验证token、验证有效时间等，我们又不想将我们在验证机制加入到我们的业务逻辑的代码中，并且我们的业务逻辑不可能每一个业务都进行验证，这时就需要拦截器来解决这类问题。而grpc为我们提供了拦截器的实现，可以方便开发者更好的开发，在执行业务逻辑之前，服务端会先执行我们实现的拦截器，对客服端数据的提取和验证。

在metadata机制中，我们获取metadata数据是通过我们的业务代码SayHell()方法来实现的，这显然是影响到了我们的业务逻辑，业务代码中我们只想做业务的开发，并且不可能每一个业务都写入验证机制，下面我们直接来实现一个简单的拦截器：

我们依然使用metadata机制中的proto代码，下面我们来写server和client的代码

###### server:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"rpcstudy/grpc_interpretor/proto"

	"google.golang.org/grpc"
)

type HelloSerivce struct{}

func (h *HelloSerivce) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	//此时我们不需要在 SayHello中写拦截器逻辑了
  //接收client发送的handle消息
	//md, ok := metadata.FromIncomingContext(c)
	//if !ok {
	//	fmt.Println("get metadata err")
	//}
	//
	//for key, value := range md {
	//	fmt.Printf("%v: %v\n", key, value)
	//}
  
	return &proto.HelloReply{
		Id:      "123456789",
		Request: req,
	}, nil
}

func main() {
	//grpc在调用业务逻辑之前会先执行inerpertor,进行数据的拦截
	//实例化server
	lit, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panicln("监听失败", err)
	}

	//实现拦截器的处理逻辑
	interpretor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    
    ma, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Fatal("FromIncomingContext error")
		}
		fmt.Println("metadata数据: \n", ma)
    
		fmt.Println("拦截到一条新消息:\n", req)

		//我们想继续做时间统计，不让直接返回
		res, err := handler(ctx, req)
		fmt.Println("该请求已经完成")
		return res, err
	}

	//grpc拦截器的使用，返回一个ServerOption
	sopt := grpc.UnaryInterceptor(interpretor)

	//注册处理逻辑
	//创建grpc服务, NewServer的参数为ServerOption类型的可选参数
	//即创建服务器时使用可放入拦截器
	s := grpc.NewServer(sopt)
	proto.RegisterGreeterServer(s, &HelloSerivce{})

	//开启服务
	log.Println(s.Serve(lit))

}
```

运行server，等待请求进入

###### Client：

```go
package main

import (
	"context"
	"fmt"
	"log"
	"rpcstudy/grpc_interpretor/proto"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	timepb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	//client拦截器,client的拦截器和server拦截器逻辑一致，代码有点出入，具体不过多介绍，具体看这部分注释：
	//interpretor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	md := metadata.New(map[string]string{
	//		"appid":  "ice_moss",
	//		"appkey": "ice_12345",
	//	})
	//	//NewOutgoingContext创建一个新的上下文context，并附加了md放入context中，并返回这个context
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//
	//	//调用者invoker
	//	err := invoker(ctx, method, req, reply, cc)
	//	fmt.Printf("消耗时间：%s\n", time.Since(start))
	//	return err
	//}
	//
	//opt := grpc.WithUnaryInterceptor(interpretor)

	clientConn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	defer clientConn.Close()
  
	//连接grpc服务器
	client := proto.NewGreeterClient(clientConn)

	md := metadata.New(map[string]string{
		"appid":  "ice_moss",
		"appkey": "ice_12345",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.SayHello(ctx, &proto.HelloRequest{
		Name:   "kuangyang",
		Url:    "https://learnku.com/blog/yangkuang",
		Gender: proto.Gender_MALE,
		M: map[string]string{
			"来自": "北京",
			"现居": "上海",
		},
		AddTime: timepb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("返回结果: %v", res)

}
```

这里我们只介绍server拦截器，client返回结果：

```shell
返回结果: id:"123456789" request:{name:"kuangyang" url:"https://learnku.com/blog/yangkuang" m:{key:"来自" value:"北京"} m:{key:"现居" value:"上海"} addTime:{seconds:165s:686279000}}
```

server输出:

```shell
metadata数据: 
 map[:authority:[localhost:9090] appid:[ice_moss] appkey:[ice_12345] content-type:[application/grpc] user-agent:[grpc-go/1.46.2]]
拦截到一条新消息:
 name:"kuangyang"  url:"https://learnku.com/blog/yangkuang"  m:{key:"来自"  value:"北京"}  m:{key:"现居"  value:"上海"}  addTime:{seconds:1652953608  nanos:686279000}
该请求已经完成

```

仔细看我们的server不仅拿到了metadata中的数据，也拿到了client上传的信息数据,接着我们就可以将这些数据进行有效验证了





### 拦截器和metadata实现GRPC的auth认证

学习了上面的内容，下面我们来实现对用户的身份和密码的验证这里我们叫auth认证，同样使用上面内容中的proto文件，直接上代码：

###### server:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"rpcstudy/grpc_interpretor/proto"

	"google.golang.org/grpc"
)

//实现接口的实现者
type HelloSerivce struct{}

//对用户数据和id进行绑定
func (h *HelloSerivce) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Id:      "123456789",
		Request: req,
	}, nil
}

func main() {
	//grpc在调用业务逻辑之前会先执行inerpertor,进行数据的拦截

	//实例化server
	lit, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Panicln("监听失败", err)
	}

	//拦截器处理逻辑
	interpretor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		//接收client发送的handle消息
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("get metadata err")
			return nil, status.Error(codes.Unauthenticated, "没有token认证")
		}

		var appid string
		var appkey string
		
    //将appid的value取出
		if value, ok := md["appid"]; ok {
			appid = value[0]
		}

     //将appkey的value取出
		if value, ok := md["appkey"]; ok {
			appkey = value[0]
		}

		if appid != "ice_moss" || appkey != "ice_12345" {
			fmt.Println("用户token未认证")
			return nil, status.Error(codes.Unauthenticated, "没有token认证")
		}

		fmt.Println("拦截到一条新消息:\n", req)

		//我们想继续做时间统计，不让直接返回
		res, err := handler(ctx, req)
		fmt.Println("该请求已经完成, 认证成功")
		return res, err
	}

	//grpc拦截器的使用，返回一个ServerOption
	sopt := grpc.UnaryInterceptor(interpretor)

	//注册处理逻辑
	//创建grpc服务, NewServer的参数为ServerOption类型的可选参数
	//即创建服务器时将拦截器放入服务器中
	s := grpc.NewServer(sopt)
	proto.RegisterGreeterServer(s, &HelloSerivce{})

	//开启服务
	log.Println(s.Serve(lit))

}
```

运行server, 等待请求进入



###### client:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"rpcstudy/grpc_interpretor/proto"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	timepb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	clientConn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	defer clientConn.Close()

	//连接grpc服务器
	client := proto.NewGreeterClient(clientConn)

	md := metadata.New(map[string]string{
		"appid":  "ice_moss",
		"appkey": "ice_12345",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.SayHello(ctx, &proto.HelloRequest{
		Name:   "kuangyang",
		Url:    "https://learnku.com/blog/yangkuang",
		Gender: proto.Gender_MALE,
		M: map[string]string{
			"来自": "北京",
			"现居": "上海",
		},
		AddTime: timepb.New(time.Now()),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("返回结果: %v", res)

}
```

client返回:

```shell
返回结果: id:"123456789" request:{name:"kuangyang" url:"https://learnku.com/blog/yangkuang" m:{key:"来自" value:"北京"} m:{key:"现居" value:"上海"} addTime:{seconds:165nanos:883406000}}
```

server输出:

```shell
拦截到一条新消息:
 name:"kuangyang"  url:"https://learnku.com/blog/yangkuang"  m:{key:"来自"  value:"北京"}  m:{key:"现居"  value:"上海"}  addTime:{seconds:1652954947  nanos:883406000}
该请求已经完成, 认证成功
```



接下来我们将metadata中数据修改一下：

```go
	md := metadata.New(map[string]string{
		"appid":  "ice_moss",
		"appkey": "ice",
	})
```

client返回:

```shell
panic: rpc error: code = Unauthenticated desc = 没有token认证
```

Server输出:

```shell
用户token未认证
```

这就是一个简单的auth认证过程，当然这里也欢迎阅读我的微信小程序用户登录token认证:   [微信小程序登录服务后端实战](https://learnku.com/articles/67053)





### grpc状态码

其实在上一部分内容中我们就已经接触到了grpc状态码

```go
md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("get metadata err")
			return nil, status.Error(codes.Unauthenticated, "没有token认证")  
		}
```

```status.Error(codes.Unauthenticated, "没有token认证") ```这既是一个状态码

又或者是:

```go
if req.Name == "" {
		return nil, status.Errorf(codes.NotFound, "未找到name：%s", req)
	}
```

不做过多介绍，关于[更多状态码](https://github.com/grpc/grpc/blob/master/doc/statuscodes.md)



### grpc的超时机制

我们在使用grpc调用过程中，如果某一个请求没跟上，或者出现一些情况，导致我们的请求一直没有返回，这肯定是不行的，例如：我们请求A服务，然后A会去调用其他服务，如：A -> B -> C -> D，但是这个过程D一直没有返回，这肯定会导致我们的C回不来，就处于一直等待的状态了，这不仅占用资源，这也使得调用者体验不好，所以需要使用grpc的超时机制来防止此类问题的发送，而grpc为我们提供了这样方法，我们只需要只需要在客服端使用：```func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)```

我们以auth认证代码为例，只需要在客服端调用```func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)```

client:

```go
	//创建一个有时间限定context的上下文
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	//使用metadata机制和拦截器进行用户身份验证
	md := metadata.New(map[string]string{
		"appid":  "ice_moss",
		"appkey": "12345",
	})
	//将md附加到context上下文中
	ctx = metadata.NewOutgoingContext(ctx, md)

```

在server中我们我们加入:

```go
func (h *HolleService) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	//验证超时机制
	time.Sleep(5 * time.Second)
	……
	……
	……
}
```

我们给这个请求2秒的时间，但是我们执行 SayHello服务需要执行5秒时间，显然这个请求已经超时了，加入超时机制后，他不会一直等待，2秒后他就返回：

```shell
2022/05/19 18:38:41 cannot get id for messagerpc error: code = DeadlineExceeded desc = context deadline exceeded
```



### protoc文件生成的go源码解析

我们以metadata机制内容中的proto文件为例:

```protobuf
syntax = "proto3";

option go_package="/.;proto";
//引入protobuf的内置类型
import "google/protobuf/timestamp.proto";

//定义接口
service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply);
}

//枚举类型
enum Gender{
    MALE = 0;
    FE_MALE = 1;
}

message HelloRequest {
    string name = 1;
    string url = 2;
    Gender gender = 3;
    map<string, string> m = 4;  //proto map类型
    google.protobuf.Timestamp addTime = 5;  //protobuf的内置类型
}

message HelloReply {
    string id = 1;
    HelloRequest request = 2;
}
```



执行```protoc -I . holle.proto --go_out=plugins=grpc:. ```后生成：hello.pb.go文件



```go
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: md.proto

package proto

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//枚举类型
type Gender int32

const (
	Gender_MALE    Gender = 0
	Gender_FE_MALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FE_MALE",
	}
	Gender_value = map[string]int32{
		"MALE":    0,
		"FE_MALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_md_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_md_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_md_proto_rawDescGZIP(), []int{0}
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url     string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Gender  Gender               `protobuf:"varint,3,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	M       map[string]string    `protobuf:"bytes,4,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //proto map类型
	AddTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=addTime,proto3" json:"addTime,omitempty"`                                                                             //protobuf的内置类型
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_md_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_md_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_md_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HelloRequest) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

func (x *HelloRequest) GetM() map[string]string {
	if x != nil {
		return x.M
	}
	return nil
}

func (x *HelloRequest) GetAddTime() *timestamp.Timestamp {
	if x != nil {
		return x.AddTime
	}
	return nil
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Request *HelloRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_md_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_md_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_md_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HelloReply) GetRequest() *HelloRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_md_proto protoreflect.FileDescriptor

var file_md_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x01, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x01, 0x6d, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x34, 0x0a,
	0x06, 0x4d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x1f, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x46, 0x45, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x32, 0x31, 0x0a, 0x07, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0a,
	0x5a, 0x08, 0x2f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_md_proto_rawDescOnce sync.Once
	file_md_proto_rawDescData = file_md_proto_rawDesc
)

func file_md_proto_rawDescGZIP() []byte {
	file_md_proto_rawDescOnce.Do(func() {
		file_md_proto_rawDescData = protoimpl.X.CompressGZIP(file_md_proto_rawDescData)
	})
	return file_md_proto_rawDescData
}

var file_md_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_md_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_md_proto_goTypes = []interface{}{
	(Gender)(0),                 // 0: Gender
	(*HelloRequest)(nil),        // 1: HelloRequest
	(*HelloReply)(nil),          // 2: HelloReply
	nil,                         // 3: HelloRequest.MEntry
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_md_proto_depIdxs = []int32{
	0, // 0: HelloRequest.gender:type_name -> Gender
	3, // 1: HelloRequest.m:type_name -> HelloRequest.MEntry
	4, // 2: HelloRequest.addTime:type_name -> google.protobuf.Timestamp
	1, // 3: HelloReply.request:type_name -> HelloRequest
	1, // 4: Greeter.SayHello:input_type -> HelloRequest
	2, // 5: Greeter.SayHello:output_type -> HelloReply
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_md_proto_init() }
func file_md_proto_init() {
	if File_md_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_md_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_md_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_md_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_md_proto_goTypes,
		DependencyIndexes: file_md_proto_depIdxs,
		EnumInfos:         file_md_proto_enumTypes,
		MessageInfos:      file_md_proto_msgTypes,
	}.Build()
	File_md_proto = out.File
	file_md_proto_rawDesc = nil
	file_md_proto_goTypes = nil
	file_md_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "md.proto",
}
```

***message --> : 生成了对应的结构体***

```go
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url     string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Gender  Gender               `protobuf:"varint,3,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	M       map[string]string    `protobuf:"bytes,4,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //proto map类型
	AddTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=addTime,proto3" json:"addTime,omitempty"`                                                                             //protobuf的内置类型
}
```



我们可以看到这部分代码: 这部分为客服端调用

```go
//接口
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

//接口的实现者
type greeterClient struct {
	cc grpc.ClientConnInterface
}

//构造一个greeterClient对象
func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

//实现接口的方法
func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Greeter/SayHello", in, out, opts...)  
	if err != nil {
		return nil, err
	}
	return out, nil
}

```



我们在client中要调用grpc我们定义的业务代码时调用了：

```go
func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}
```

入参是: ***这个接口中的两个方法主要就是为我们做序列化和反序列化操作***：

```go
type ClientConnInterface interface {
	// Invoke performs a unary RPC and returns after the response is received
	// into reply.
	Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...CallOption) error
	// NewStream begins a streaming RPC.
	NewStream(ctx context.Context, desc *StreamDesc, method string, opts ...CallOption) (ClientStream, error)
}
```

然后他返回了一个```greeterClient{cc}```的对象

***在代码中可以看到```err := c.cc.Invoke(ctx, "/Greeter/SayHello", in, out, opts...)d ```其中："/Greeter/SayHello"就是SayHello方法的映射ID***



接下来看这部分代码:

```go
// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}
```

里面有一个接口：

```go
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}
```

这个接口就需要我们来实现，他就是我们用来做业务逻辑处理的，我们在server中是这样实现它的:

```go
//接口的实现者
type HelloSerivce struct{}

//实现接口中的方法
func (h *HelloSerivce) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Id:      "123456789",
		Request: req,
	}, nil
}
```

然后再来看这个方法：

```go
func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}
```

我们在server中是这样调用的：

```go
//创建一个grpc服务器
s := grpc.NewServer()
proto.RegisterGreeterServer(s, &HelloSerivce{})

```

我们看这个```RegisterGreeterServer(s *grpc.Server, srv GreeterServer)```方法入参是```s *grpc.Server, srv ```	和```GreeterServer```类型，而```GreeterServer```是一个接口：

```go
type GreeterServer interface {
    SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}
```

为什么我们可以在调用的时候直接将```&HelloSerivce{}```传入？ 答案是：```HelloSerivce```结构体实现了```GreeterServer```接口，这个```&HelloSerivce{}```对象可以直接传入。 如果你对接口的知识不理解也可以阅读[「Golang成长之路」面向接口](https://learnku.com/articles/58910)

其实我们在hello.pb.go的源码中就需要理解这么多就行了。

***这里我们的GRPC的内容就介绍结束，感谢阅读！***

