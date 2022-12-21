[toc]

### 并发编程模式

在上两篇文章中我们主要介绍了并发goroutine和channel，现在我们来介绍一下golang的并发模式，不像golang的设计模式，这里来介绍一下常用的并发模式：
#### 生成器

```go
package main
import (
	"fmt"
	"math/rand"
	"time"
)
//生成器msgGen
func msgGen() chan string {
	c := make(chan string)
	//启动并发，真正生成数据
	go func(){
		i := 0
		for {
		//生成时间在范围：0~2000毫秒
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("message ：%d", i)
			i++
		}
	}()
	return c
}
func main() {
	  m1 := msgGen()
  	for{
	 	fmt.Println(<-m1)
  	}
}
```
程序分析：msgGen()将c := make(chan string)返回给m1,在for中等待并发启动发送数据给m1,m1立即将数据送出并打印。

#### 服务/任务

看下面代码：

```go
func main() {
   m1 := msgGen()  //开启任务m1
   M2 := msgGen()  //开启任务m2
   for{
      fmt.Println(<-m1)
      fmt.Println(<-m2)
   }
}
```
在生成器的基础之上可以提供多个服务/任务，如上面代码中的m1,m2是使用同一个生成器的两个服务/任务，而m1和m2是两个独立的服务/任务，我们如果拿到m1j就可以和m1j交互，拿到m2就可以和m2进行交互。
```go
package main
import (
   "fmt"
 "math/rand" "time")
func msgGen(name string) chan string {
   c := make(chan string)
   go func(){
      i := 0
  for {
         time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
         c <- fmt.Sprintf("service: %s, message ：%d", name, i)
         i++
      }
   }()
   return c
}
func main() {
   m1 := msgGen("service1")
   M2 := msgGen("sercive2")
   for{
      fmt.Println(<-m1)
      fmt.Println(<-m2)
   }
}
```
打印结果：
service: service1, message ：0
service: sercive2, message ：0
service: service1, message ：1
service: sercive2, message ：1
service: service1, message ：2
service: sercive2, message ：2
service: service1, message ：3
service: sercive2, message ：3
service: service1, message ：4
service: sercive2, message ：4
service: service1, message ：5
service: sercive2, message ：5
service: service1, message ：6
service: sercive2, message ：6
……
……
……



#### 同时等待多任务：两种方法

从上面的打印结果可以看出两个任务是一起进行的，现在我们需要将两个结果交替打印：

##### 方法一：

>>将两个channel的数据放进一个节点中，然后在发送到第三个channel中
>>![「Golang成长之路」并发之并发模式篇](https://cdn.learnku.com/uploads/images/202110/06/69310/wfTjsL2ZnE.png!large)
>>下面来看是如何实现的：
```go
func fanIn(c1, c2 chan string) chan string{
   c := make(chan string)
   go func() {
      for{
         c <- <-c1
      }
   }()
      go func() {
         for{
            c <- <-c2
         }
      }()
   return c
}

```
完整代码：

```go
package main

import (
   "fmt"
 "math/rand" "time")

func msgGen(name string) chan string {
   c := make(chan string)
   go func(){
      i := 0
  for {
         time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
         c <- fmt.Sprintf("service: %s, message ：%d", name, i)
         i++
      }

   }()
   return c

}

func fanIn(c1, c2 chan string) chan string{
   c := make(chan string)
   go func() {
      for{
         c <- <-c1
      }
   }()
      go func() {
         for{
            c <- <-c2
         }
      }()
   return c

}
func main() {
   m1 := msgGen("service1")
   m2 := msgGen("sercive2")
   m := fanIn(m1, m2)
   for{
      fmt.Println(<-m)
   }

}
```



##### 方法二：
使用select对多个channel同时接收，此时只需要开一个goroutine即可,这里我们叫做fanInSelect

```go
func fanInSelect(c1, c2 chan string) chan string{
   c := make(chan string)
   go func() {
      for{
         select{
         case m := <- c1:
            c <- m
         case m := <- c2:
            c <- m
         }
      }
   }()
   return c
```
   ```go
   func main() {
   m1 := msgGen("service1")
   m2 := msgGen("sercive2")
   m := fanInSelect(m1, m2)
   for{
      fmt.Println(<-m)
   }
}
   ```
 >  打印结果：
 >  service: sercive2, message ：0
 >  service: service1, message ：0
 >  service: sercive2, message ：1
 >  service: service1, message ：1
 >  service: sercive2, message ：2
 >  service: sercive2, message ：3
 >  service: sercive2, message ：4
 >  service: service1, message ：2
 >  service: sercive2, message ：5
 >  service: sercive2, message ：6
 >  service: service1, message ：3
 >  ……
 >  ……



   ##### 方法一，方法二对比

   对比两种方法，方法一想要开两个goroutine(如果有多个参数就需要开很多goroutine），而方法二只需要开一个goroutine即可；当我们知道有具体的参数时(channel),使用方法二会更好，在不知道具体有多少个goroutine的情况下使用方法一更好。下面来看看方法一的优化在哪里：

```go

func fanIn(chs ...chan string) chan string{   //chs ...参数限制，可随意增减
	c := make(chan string)
	for _, ch := range chs{  //第一个for将每一个参数取出,每一个channel需要开一个goroutine
			go func() {
				for { //第二个for源源不断的将数据传出
					c <- <-ch
				}
			}()
		}
	return c
}
```
打印结果：
service: sercive2, message ：0
service: sercive2, message ：1
service: sercive2, message ：2
service: sercive2, message ：3
service: sercive2, message ：4
service: sercive2, message ：5
service: sercive2, message ：6
service: sercive2, message ：7
service: sercive2, message ：8
service: sercive2, message ：9

注意：打印结果全是service2
原因：我们每次从chs中取出一个channel给ch，运行到关键字"go"就return ch，接着进行将chs中的第二个channel的给ch，接着return，此时第一个ch已经被迭代成了第二ch了，所以当goroutine真正的运行时，传入c中的数据都来自于最新的ch。
解决方法:增加变量来存储chs中的数据，如 ```chcapy := ch```
```go
func fanIn(chs ...chan string) chan string{
   c := make(chan string)
   for _, ch := range chs{  //第一个for将每一个参数取出,每一个channel需要开一个goroutine
  chcapy := ch
         go func() {
            for { //第二个for源源不断的将数据传出
  c <- <- chcapy
            }
         }()
      }
   return c
}
```
或者将在函数式func()增加参数(这里需要了解go语言的传参，见文章[「Golang成长之路」基础语法](https://learnku.com/articles/57716)  如 ```go func(in chan string) {……}()```
```go
func fanIn(chs ...chan string) chan string{
   c := make(chan string)
   for _, ch := range chs{  //第一个for将每一个参数取出,每一个channel需要开一个goroutine
  go func(in chan string) {
            for { //第二个for源源不断的将数据传出
  c <- <- in
            }
         }(ch)
      }
   return c
}
```
再来看看调用：
可随意增加参数
>三个参数
```go
func main() {
   m1 := msgGen("service1")
   m2 := msgGen("sercive2")
   m3 := msgGen("service3")
   m := fanIn(m1, m2, m3)
   for{
      fmt.Println(<-m)
   }
}
```

>四个参数:
```go
func main() {
   m1 := msgGen("service1")
   m2 := msgGen("sercive2")
   m3 := msgGen("service3")
   m4 := msgGen("service4")
   m := fanIn(m1, m2, m3, m4)
   for{
      fmt.Println(<-m)
   }
}
```