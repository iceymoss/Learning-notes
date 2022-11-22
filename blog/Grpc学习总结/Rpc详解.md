[toc]



### 文章介绍

本文主要介绍 RPC 是什么， 为什么要使用 RPC，使用 RPC 需要解决问题及 RPC 使用实例

### RPC 是什么
#### RPC

Rpc（Remote Procedure Call）远程过程调用协议，一种通过网络从远程计算机上请求服务，而不需要了解底层

网络技术的协议。RPC 它假定某些协议的存在，例如 TPC/UDP 等，为通信程序之间携带信息数据。在 OSI 网络

七层模型中，RPC 跨越了传输层和应用层，RPC 使得开发，包括网络分布式多程序在内的应用程序更加容易。

简单一点说：就是向远程服务器发送请求，做业务处理或任务计算等 (即程序)，就是想要把调用远程服务器中方

法的过程和像调用本地方法一样简单。

#### 为什么要使用 RPC

当我们无法在一个进程内，甚至通过本第地调用的方式满足我们的需求时，比如我们的通讯系统，甚至不同的组

织间的通讯，由于计算能力需要横向扩展，需要在多台机器组成的集群上部署应用，这才能是我们进行远程的通

讯；又或者说电商系统，我们不可能靠着本地调用就能满足我们的需求，例如提交订单的处理方法，库存处理方

法等在本地调用，在肯定是不行的，这是用户发起相应的请求由我们远程服务器去做业务处理和计算的。所以 

RPC 就显得如此重要了。



#### 使用 RPC 需要解决问题
这里我们来看，我们想象一下当我们在远程调用 (也就是使用 RPC) 的过程中，我们需要执行的函数或者方法在远

程机器上，例如我们要调用远程计算机的 add 方法，下面就会有这几个问题：

1. **Call ID 映射**。 我们要怎么告诉远程计算机我们需要调用的是 add 方法，而不是 reduce 方法，mult 方法，

   divi 方法等，在本地调用中，函数体是直接通过函数指针来指定的，当我们调用本地 add 方法时，编译器会

   自动给我们调用到它相应的函数指针，而在远程调用中，使用指针明显是不行的，因为两个进程的地址不

   同。所以，在 RPC 中，所有函数都必须有一个唯一的 ID，这个 ID 在所有进程中都是唯一确定的，客户端在

   调用时都必须附加上这个 ID，然后我们还需要在客户端和服务端分别维护一个 **{函数 <—–> Call ID }** 的对应

   表，两者的表不一定需要完全相同，但相同的函数对应的 Call ID 必须相同，当客户端需要进行远程调用时，

   它就查一下这个表，找出相应的 Call ID，然后把它传给服务端，服务端也通过查表，来确定客户端需要调用

   的函数，然后执行相应函数的代码。

   

2. **序列化及反序列化**。客户端怎么把参数值传给远程的函数呢？在本地调用中，我们只需要把参数压到栈里，

   然后让函数自己去栈里读就行。但是在远程过程调用时，客户端跟服务端是不同的进程，不能通过内存来传

   递参数。甚至有时候客户端和服务端使用的都不是同一种语言（比如服务端用 C++，客户端用 Java 或者 

   Go）。这时候就需要客户端把参数先转成一个字节流，传给服务端后，再把字节流转成自己能读取的格式。

   这个过程叫序列化和反序列化。同理，从服务端返回的值也需要序列化反序列化的过程。

   整个流程如下：

   ![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/5glZiYH2P6.png)

   

   

3. **网络传输**。远程调用往往用在网络上，客户端和服务端是通过网络连接的。所有的数据都需要通过网络传

   输，因此就需要有一个网络传输层。网络传输层需要把 Call ID 和序列化后的参数字节流传给服务端，然后再

   把序列化后的调用结果传回客户端。只要能完成这两者的，都可以作为传输层使用。因此，它所使用的协议

   其实是不限的，能完成传输就行。尽管大部分 RPC 框架都使用 TCP 协议，但其实 UDP 也可以，而 gRPC 干

   脆就用了 HTTP2。Java 的 Netty 也属于这层的东西。

   



解决了上面三个问题，我们就能实现 RPC 了，现在我们再看看，客户端和服务端在 RPC 中的工作是什么：

**客户端 (client):**

```go
1. 将这个调用映射为Call ID。这里假设用最简单的字符串当Call ID的方法，例如这里可以：http://127.0.0.1:8080/add?a=1&b=1，即直接使用add作为path，又或者http://127.0.0.1:8080/?method=add&a=1&b=1等

2. 将Call ID，a和b序列化。可以直接将它们的值以二进制形式打包

3. 把2中得到的数据包发送给ServerAddr，这需要使用网络传输层

4. 等待服务器返回结果

4. 如果服务器调用成功，那么就将结果反序列化，并赋给total
```

**服务端 (service)：**

```go
1. 在本地维护一个Call ID到函数指针的映射call_id_map，可以用dict完成

2. 等待请求，包括多线程的并发处理能力

3. 得到一个请求后，将其数据包反序列化，得到Call ID

4. 通过在call_id_map中查找，得到相应的函数指针

5. 将a和rb反序列化后，在本地调用add函数，得到结果

6. 将结果序列化后通过网络返回给Client

```



**注意：**

* Call ID 可以是字符串，也可以是整数 ID，映射其实是一个哈希表
* 序列化与反序列化可以自己写，当然也可以使用 Protobuf 或者 FlatBuffers 之类的。
* 网络传输库可以自己实现 socket，也可以使用 asio，ZeroMQ，Netty 之类。



### RPC 实例

我们来简单模拟一下 RPC 的调用过程，实例介绍：客户端需要调用远程计算机上的计算方法，这里以加法为例 (当然可以非常复杂的大型计算)

客户端的调用时无需关心远程服务端的 add 方法是怎么实现的，我们只需要在客户端对远程调用的 add 方法的过程进行封装：

**service:**

由于是简单的例子，就不做完整的错误处理了

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

func main() {
    //http://127:0:0:1:8000/add?a=1&b=2
    //Call ID 使用request.URL.PATH
    http.HandleFunc("/add",
        func(writer http.ResponseWriter, request *http.Request) {
            //解析参数
            err := request.ParseForm()
            if err != nil {
                panic("解码失败")
            }
            fmt.Println("path:", request.URL.Path)
            //取出参数,做类型转换
            a, err := strconv.Atoi(request.Form["a"][0])
            if err != nil {
                panic("转换失败")
            }
            b, err := strconv.Atoi(request.Form["b"][0])
            if err != nil {
                panic(err)
            }
            //返回的数据格式：json{"data":3}
            //使用json编码，即序列化
            writer.Header().Set("Content-Type", "application/json")
            //序列化
            jData, err := json.Marshal(map[string]int{
                "data": a + b,
            })
            if err != nil {
                panic(err)
            }
            _, err = writer.Write(jData)
            if err != nil {
                panic("写入失败")
            }
        })

    //监听端口
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic("监听失败")
    }
}
```

启动服务端



**client**:

这里同样不做完整的错误处理了
这里我们使用第三方包：`github.com/kirinlabs/HttpRequest`
使用以下命令获取:

```
go get github.com/kirinlabs/HttpRequest
```

当然你也可以使用其他相关的方法来连接我们的 service

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/kirinlabs/HttpRequest"
)

//解析结构
type ResponseData struct {
    data int `json:"data"`
}

//对add进行封装
func add(a, b int) int {
    //生成一个实例
    req := HttpRequest.NewRequest()
    res, err := req.Get(fmt.Sprintf("http://127.0.0.1:8080/add?a=%d&b=%d", a, b))
    if err != nil {
        panic("连接失败")
    }
    body, err := res.Body()
    if err != nil {
        panic(err)
    }

    var resData ResponseData
    err = json.Unmarshal(body, &resData)
    if err != nil {
        panic("解码失败")
    }
    return resData.data
}

func main() {
  fmt.Println(add(1, 4))
}
```

打印结果：

```json
{"data":5}
```



当然我们也可以直接在浏览器里访问：

http://127.0.0.1:8080/add?a=1&b=4

返回：

```json
{"data":5}
```

