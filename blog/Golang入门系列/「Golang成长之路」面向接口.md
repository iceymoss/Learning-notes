[toc]

### 文章介绍

本文来聊聊go接口(interface) 的概念、如何定义接口、接口的使用场景、接口的组合和空接口(interface{})。



### 接口的概念

* 接口是一种抽象类型，是对其他类型行为的概括与抽象，从语法角度来看，接口是一组方法定义的集合。
* 很多面向对象的语言都有接口这个概念，但 Go 语言接口的独特之处在于它是隐式实现。
* 对于一个具体的类型，无须声明它实现了哪些接口，只要提供接口所必需的方法即可。
* 这种设计让编程人员无须改变已有类型的实现就可以为这些类型创建新的接口 —— 对于那些不能修改包的类型，这一点特别有用。



### 定义接口及实现

```go
type 接口名 interface {
   方法1
   方法2
   ……
}
```

这个就是接口的定义
接口是由使用者定义，由实现者实现。

来看一个小例子：

```go
type Tester interface {
  Eat()
  Sleep()
  Wake()
}
```

接口定义好了，下面来实现它：

```go
type Cat struct {
  name string 
  age int
}

func (c *Cat) Eat(){
  ……
}

func (c *Cat) Sleep(){
  ……
}

func (c *Cat) Wake(){
  ……
}
```

这样Cat就实现了Tester接口，go语言中的接口很神奇，不需要我们导入任何路径，go自动会隐式的将二者关联起来。



### 接口的使用场景(接口完成多态)

假如现在有一个任务：将 [www.imooc.com](https://www.imooc.com/) 的首页下载下来 (假如是一个很大的工程），我们需要先进行测试，然后再产品上线。

先定义一个接口：

```go
type Retriever interface{
     Get(cur string)string
}
```

下面来看看接口是如何实现的：
我们在另一个目录中实现 Get () 方法 (接口的实现)

1. 实现测试部分：

   ```go
   package mock
   
   //定义一个结构体
   type Retrievers struct{
      Context string
   }
   
   //实现接口
   func (r Retrievers)Get(cur string)string{
      return r.Context
   }
   ```

2. 产品上线部分：
   然后我们再去目录下对 Get () 方法的实现

   ```go
   package real
   
   import (
      "net/http"
    "net/http/httputil" "time")
   
   //定义一个结构体
   type Retrievers struct{
      UserAgent string
     TimeOut time.Duration
   }
   
   //Get接口的实现
   func (r Retrievers)Get(cur string)string{
      re, err :=http.Get(cur)
      if err != nil{
         panic(err)
      }
   
      result, err :=httputil.DumpResponse(re, true)
      if err != nil{
         panic(err)
      }
      return string(result)
   }
   ```

   开发流程：

   ```go
   package main
   
   import (
      "fmt"
     
    "interfacetest/retriever/mock" 
    "interfacetest/retriever/real"
   )
   
   //接口由使用者定义,接口的实现其实就是对象函数(方法)的实现
   //golang中duck type
   type Retriever interface {
      Get(cur string) string
   }
   
   func download(r Retriever) string{
      //下载https://www.imooc.com网页
      return r.Get("https://www.imooc.com")
   }
   
   //主函数
   func main() {
     //测试部分
     //mock.Retrievers{}表示来自于我们实现的mock包
      var q = mock.Retrievers{"my name"}
      fmt.Println(downloda(q))   //此时输出的是my name
   
     //开发部分
     //当测试部分通过过后就可以产品上线了
     //real.Retrievers{}表示来自于我们实现的real包
     UserAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
      var r Retriever= real.Retrievers{UserAgent, TimeOut: 3}
      fmt.Println(download(r))  //这时就可以获取https://www.imooc.com的首页内容
   }
   ```

上述实例就是描述了go接口的实际场景运用，并且这个过程就实现了多态，实例中为两种状态，即：测试态，上线态。

更详细的接口实战应用可阅读：[微信小程序登录服务后端实战](https://learnku.com/articles/67053)

### 接口断言

接口断言的类型有三种方法：

#### fmt.Printf("%T, %v\n", x, x)

例如：

```go
func main() {
var q = mock.Retrievers{"my name"}

var r Retriever= real.Retrievers{UserAgent: "hhh", TimeOut: 3}
fmt.Printf("%T, %v\n", r, r)
fmt.Printf("%T, %v\n", q, q)
```

输出：

```
real.Retrievers, {hhh 3ns}
mock.Retrievers, {my name}
```

由此可以判断出，两种类型分别属于什么类型



#### switch

```go
func inspect(r Retriever){
  fmt.Printf("%T, %v\n", r, r)
  switch v := r.(type){
    case mock.Retrievers:
       fmt.Println("mock.Retrievers:",v.Context)
    case real.Retrievers:
       fmt.Println("real.Retrievers:", v.UserAgent)
    case *mock.Retrievers:
       fmt.Println("*mock.Retrievers", v.Context)
   }
}
```



#### assertion(T.(type))
例如：

```go
//使用assertion断言类型
if note , ok := r.(real.Retrievers);ok {
	fmt.Println(note.UserAgent)
} else{
	fmt.Println("this not real Retrieers")
}
```

接口也可以是指针类型

```go
func (r *Retrievers)Get(cur string)string{
	return r.Context
}
```

```go
func main() {
//&取地址
var q = &mock.Retrievers{"my name"}
```



### 接口组合

在 Go 语言中，可以在接口 A 中组合其它的一个或多个接口（如接口 B、C），这种方式等价于在接口 A 中添加接口 B、C 中声明的方法。

例：

```go
//接口中可以组合其它接口，这种方式等效于在接口中添加其它接口的方法  
//读接口
type Reader interface {  
    read()  
}  

//写接口
type Writer interface {  
    write()  
}  

//定义上述两个接口的实现类  
type MyReadWrite struct{}  

//read接口的实现
func (mrw *MyReadWrite) read() {  
    fmt.Println("MyReadWrite...read")  
}  

//write接口的实现
func (mrw *MyReadWrite) write() {  
    fmt.Println("MyReadWrite...write")  
}  

//定义一个接口，组合了上述两个接口  
type ReadWriter interface {  
    Reader  
    Writer  
}  

//上述接口等价于：  
type ReadWriterV2 interface {  
    read()  
    write()  
}  

//ReadWriter和ReadWriterV2两个接口是等效的，因此可以相互赋值  
func interfaceTest() {  
    mrw := &MyReadWrite{}  
    //mrw对象实现了read()方法和write()方法，因此可以赋值给ReadWriter和ReadWriterV2  
    var rw1 ReadWriter = mrw  
    rw1.read()  
    rw1.write()  

    fmt.Println("------")  
    var rw2 ReadWriterV2 = mrw  
    rw2.read()  
    rw2.write()  

    //同时，ReadWriter和ReadWriterV2两个接口对象可以相互赋值  
    rw1 = rw2  
    rw2 = rw1  
}
```



### 空接口(interface{})

* 可能出现在函数声明的参数列表中，表示可以传入任意类型参数
* 空接口可以被赋予任意类型的值
* 实际的例子就是fmt.Println函数，通过使用空接口和类型查询，可以实现接受任意对象实例并进行处理

实例：

```go
OrderList := make([]interface{}, 0)   //make关于可以存放任何类型的切片
	for _, Item := range Rsp.Data {
		ItemMap := map[string]interface{}{}    //使用interface{}存放多任何类型的值
		ItemMap["id"] = Item.Id
		ItemMap["name"] = Item.Name
		ItemMap["total"] = Item.Total
		ItemMap["userId"] = Item.UserId
		ItemMap["status"] = Item.Status
		ItemMap["orderSn"] = Item.OrderSn
		ItemMap["address"] = Item.Address
		ItemMap["mobile"] = Item.Mobile
		ItemMap["post"] = Item.Post
		ItemMap["pay_type"] = Item.PayType
		ItemMap["add_time"] = Item.AddTime
		OrderList = append(OrderList, Item)
	}

	ReMap := gin.H{
		"total": Rsp.Total,
		"data":  OrderList,
	}
```





