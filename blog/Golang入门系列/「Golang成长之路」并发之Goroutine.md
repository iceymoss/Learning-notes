[toc]

### 并发

#### goroutine

并发编程表现为程序有若干个自主的活动单元组成，在今天的互联网中，一个web服务器可能一次处理上千个请求，而平板电脑和手机在渲染用户界面的同时，后端还同步进行着计算和处理网络请求等。
在go语言中每一个并发执行的活动称之为goroutine，而在我们最常见的main函数中其实也是一个goroutine(主goroutine），在此之前，我们所见的介绍语法或者程序都是顺序执行的，但是在并发编程的领域里，如果从顺序编程获取的直觉可能让我们加倍迷茫。
在此我们需要理解什么是并发：
***百科：当有多个线程在操作时,如果系统只有一个CPU,则它根本不可能真正同时进行一个以上的线程，它只能把CPU运行时间划分成若干个时间段,再将时间 段分配给各个线程执行，在一个时间段的线程代码运行时，其它线程处于挂起状。.这种方式我们称之为并发(Concurrent)。***

![「Golang成长之路」并发之Goroutine](https://cdn.learnku.com/uploads/images/202108/23/69310/Hl0SHIe7Xc.png!large)
（这个是普通函数的运行模式，main将其控制权交给调用的函数，最后在返回给main）

我的理解：举个例子，我一边在上网课，一边在做笔记，我在做笔记的时候需要将网课暂停，等我的笔记做完后又继续看网课，这两件事就可以看作是一个并发的过程。




![「Golang成长之路」并发之Goroutine](https://cdn.learnku.com/uploads/images/202108/23/69310/0dymQkQASp.png!large)

（协程则是：main和创建的goroutine是相互作用的，相互给予控制权，就像两个人，各做各的事，并且他们也相互通信）

#### 进程

进程是操作系统调度和资源分配的基本单位。

#### 线程

是系统调用的最小单位，是进程里实践运行的实体，一个进程可以有多个线程。

#### 协程

跟轻量级的线程。



#### 线程和协程的区别

1. 由于协程的特性, 适合执行大量的**I/O 密集型任务**, 而线程在这方面弱于协程。

2. 协程涉及到函数的切换, 多线程涉及到线程的切换, 所以都有**执行上下文**, 但是协程不是被操作系统内核所管理, 而完全是由程序所控制（也就是在**用户态**执行）, 这样带来的好处就是性能得到了很大的提升, 不会像线程那样需要**在内核态进行上下文切换**来消耗资源，因此**协程的开销远远小于线程的开销**。

3. 同一时间, 在多核处理器的环境下, **多个线程是可以并行的**，但是**运行的协程的函数却只能有一个**，**其他的协程的函数都被suspend**, 即**协程是并发的**。

4. 由于协程在同一个线程中, 所以不需要用来守卫临界区段的同步性原语（primitive）比如互斥锁、信号量等，并且**不需要来自操作系统的支持。**

5. 在协程之间的切换不需要涉及任何系统调用或任何阻塞调用。

6. **通常的线程是抢先式(即由操作系统分配执行权)**, 而协程是**由程序分配执行权**。

   

### go关键字

当一个程序启动时，只有一个goroutine来调用```main```，称他为主goroutine，新的goroutine通过```go```关键字来进行创建，看下面例子：
```go
func main() {
   for i := 0; i < 3; i++ {
   //使用go关键字创建goroutine
   //匿名函数
   	go func() {
          for j := 3; j >= 0; j-- {
            fmt.Println("gorouting", j)
          }
       }()
    fmt.Println("gorouting main:", i)
	//控制程序运行时间
    time.Sleep(time.Microsecond)
   }
}
```
当然，也可以将匿名函数拿出来：
```go
package main

import (
   "fmt"
 "time"
)

func doWorker() {
     for j := 3; j >= 0; j-- {
        fmt.Println("gorouting", j)
     }
}

func main() {
   for i := 0; i < 3; i++ {
   //使用go关键字创建goroutine
       go doWroker()
    fmt.Println("gorouting mian:", i)
	//控制程序运行时间
    time.Sleep(time.Microsecond)
   }
}
```

先来看看打印结果：
```go
gorouting mian: 0
gorouting 3
gorouting 2
gorouting 1
gorouting 0
gorouting mian: 1
gorouting 3
gorouting 2
gorouting 1
gorouting 0
gorouting mian: 2
gorouting 3
gorouting 2
gorouting 1
gorouting 0
```
第二遍运行：
```go
gorouting mian: 0
gorouting mian: 1
gorouting 3
gorouting 2
gorouting 1
gorouting 0
gorouting 3
gorouting 2
gorouting 1
gorouting 0
gorouting mian: 2
gorouting 3
gorouting 2
gorouting 1
gorouting 0
```
***程序解释：***其实可以看出，每一遍的运行结果都是不一样的，当程序加入第一个for ， i=0时，运行到关键字```go```时，新的goroutine就创建了，但是程序不会立即执行新的goroutine，它会进行执行main中其余的代码，在我们的这个例子中，第一遍运行，时就是这个结果，第二遍运行时也是同样的。


```go
func main(){
	 var a[10] int
	 for i := 0; i < 10;i++{
	     //使用go关键字创建goroutine
         //匿名函数
      	go func(i int){
         	for{
          	  a[i]++
        	 }
     	 }(i)
 	  }
    time.Sleep(time.Microsecond)
    fmt.Println(a)
}
```
打印结果：
```go
[39904 0 0 9122 0 7123 0 0 0 0]
```
***程序解释：***在for循环中一共创建了10个goroutine，有的goroutine执行了，有的还处于等待状态，但是我们给主goroutine的时间不多，所以有一些还来不及执行就被主goroutine杀掉了，这样我们就能理解打印结果了。



### goroutine切换点(可能)

***goroutine可能切换:***

1. I/O、select
2. channel
3. 等待锁
4. 函数调用(有时)
5. runtime.Gosched