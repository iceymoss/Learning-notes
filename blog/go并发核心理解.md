[toc]

# 文章介绍

本文将以实战的模式来介绍go核心channel，首先我们将介绍go的协程，如何开启go协程，然后通过实例来介绍go并发核心goroutine和channel。

# go的协程

## 协程(Coroutine)

**百科：当有多个线程在操作时，如果系统只有一个 CPU, 则它根本不可能真正同时进行一个以上的线程，它只能把 CPU 运行时间划分成若干个时间段，再将时间 段分配给各个线程执行，在一个时间段的线程代码运行时，其它线程处于挂起状。. 这种方式我们称之为并发 **。

首先我们先了解什么是协程(Coroutine)，我们可以称为协作式线程，他的粒度比线程更轻，作用于用户态，操作系统内核态感受不到协程的存在的，协程可以帮助我们高并发的去完成一些指定任务，例如网络IO，文件IO等。

## go协程(goroutine)

其实协程的核心都一样，我们知道go原生的支持并发，而这一核心就是goroutine，在go中，main 和创建的 goroutine 是相互作用的，相互给予控制权，就像两个人，各做各的事，并且他们也相互通信，我们来看一下如何创建协程。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
  //使用go关键字创建协程
  //子协程1
	go func() {
		for {
			fmt.Println("我是goroutine1")
		}
	}()

  //子协程2
	go func() {
		for {
			fmt.Println("我是goroutine2")
		}
	}()

	fmt.Println("我是主goroutine")

	time.Sleep(4* time.Millisecond)  //go协程运行4毫秒
}
```

这样我们就启动两两个协程，他们不断的打印消息，直到4ms后主协程退出，将子协程kill掉

继续看实例：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func RwGlobalMap(){
	if value,exists := m.Load("name"); exists {
		fmt.Println("value:", value)
	}else{
		m.Store("name", "iceymoss")
	}
}
func main() {
	go RwGlobalMap()
	go RwGlobalMap()
	go RwGlobalMap()
	go RwGlobalMap()

	time.Sleep(time.Second)
}
```

我们可以想一想程序输出是什么？

其实输出不定：

```sh
➜  share_memonry go run main.go
value: iceymoss
➜  share_memonry go run main.go
➜  share_memonry go run main.go
value: iceymoss
value: iceymoss
➜  share_memonry go run main.go
➜  share_memonry go run main.go
value: iceymoss
value: iceymoss
value: iceymoss
value: iceymoss
```

# csp并发模型

CSP模型是上个世纪七十年代提出的，用于描述两个独立的并发实体通过共享的通讯 channel(管道)进行通信的并发模型。 CSP中channel是第一类对象，它不关注发送消息的实体，而关注与发送消息时使用的channel。

Golang 就是借用CSP模型的一些概念为之实现并发进行理论支持，其实从实际上出发，go语言并没有完全实现了CSP模型的所有理论，仅仅是借用了 process和channel这两个概念。process是在go语言上的表现就是 goroutine 是实际并发执行的实体，每个实体之间是通过channel通讯来实现数据共享。



# go并发核心

channel管道的意思，在go中指用来通讯的管道，用来传输数据，通过协程之间的通讯，下面来看看如何使用channel。

```go
var ch chan int  //声明一个整型的管道
```

也可以使用make函数进行内存分配和初始化

```go
ch := make(chan int)
```

下面使用channel来实现协程之间的通信，直接来看实例：

```go
package main

import (
	"fmt"
	"time"
)

//生产数据
func send(ch chan int){
	i := 0
	for i < 100 {
		ch <- i
		i++
	}
}

//接收并处理数据
func receive(ch chan int, rch chan string){
	for {
		num := <-ch
		fmt.Println("接收到值：", num)
		rch <- fmt.Sprintf("处理后的编号：%d", num+1)
	}
}

func main() {
	ch := make(chan int)
	rch := make(chan string)
	go receive(ch, rch)  //此时ch没有数据，子协程rec阻塞，直到其他协程向管道放入数据，阻塞解除
	go send(ch)         
	go func(chan string) {
		for {
			fmt.Println("收到结果：", <-rch)
		}
	}(rch)
  
	time.Sleep(time.Second)  //主协程等待1s
}
```

可以想想输出是什么？

```go
接收到值： 0
接收到值： 1
收到结果： 处理后的编号：1
收到结果： 处理后的编号：2
接收到值： 2
收到结果： 处理后的编号：3
接收到值： 3
接收到值： 4
收到结果： 处理后的编号：4
收到结果： 处理后的编号：5
接收到值： 5
收到结果： 处理后的编号：6
接收到值： 6
接收到值： 7
收到结果： 处理后的编号：7
收到结果： 处理后的编号：8
接收到值： 8
收到结果： 处理后的编号：9
接收到值： 9
收到结果： 处理后的编号：10
```



接着使用channel实现文件读写协程的交互，看实例：将三个文件中的内容并发的读取，并发写入一个文件中

```go
age main

import (
	"bufio"
	"io"
	"log"
	"os"
)

var (
	content = make(chan string, 1000) //用于传输内容
	readOk = make(chan struct{}, 3)   //用于通知消费者
	writeOk = make(chan struct{})     //用户通知主协程
)

//ReadToSend 生产者
func ReadToSend(filename string){
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("打开文件失败：", err)
	}
	defer file.Close()

	//读文件
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == nil {
			content <-line
		}else {
			if err == io.EOF {
				if len(line) != 0 {
					content <- line +"\n"
				}
				break
			} else {
				log.Fatal("其他错误：", err)
			}
		}
	}
	//当生产者协程结束后需要通知消费者协程
	//当readTag管道为空时，则关闭content，从而通知消费者协程停止
	<-readOk
	if len(readOk) == 0 {
		close(content)
	}
}

func WriteFile(filename string){
	//打开文件
	fout, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 666)
	if err != nil {
		log.Fatal("打开文件失败：", err)
	}
	defer fout.Close()

	//写文件
	w := bufio.NewWriter(fout)

	for line := range content{
		w.WriteString(line)
	}
	w.Flush()

	//通知主协程write结束
	writeOk <-struct{}{}
}

func main() {
	for i := 0; i < 3; i++ {
		readOk <-struct{}{}
	}

	go ReadToSend("/data/data1.txt")
	go ReadToSend("/data/data2.txt")
	go ReadToSend("/data/data3.txt")

	go WriteFile("/data/data5.txt")

	//当执行writeOk没有数据时，会将主协程阻塞，直到writeOk有数据过来，然后程序运行结束
	<-writeOk
}
```

这个程序显示更加合理，我们没有像之前那样使用time.sleep方法来预估程序执行时间了(预估往往是不靠谱的)，当读协程完成后，就通知写协程：文件读完了并关闭管道；当写协程被收到读协程的结束消息后，要退出循环，然后通知主协程已经写完了，当主协程收到读完消息后，就退出主协程，程序结束。

接着再看实例：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{}, 1)  //初始化一个管道
	ch <-struct{}{}     //向管道添加数据

	go func(){
		//子协程1 5秒管道数据拿走，解除主协程的阻塞
		time.Sleep(5*time.Second)
		<- ch
		fmt.Println("goroutine1 over")
	}()

	ch <- struct{}{}  //管道容量满了，此时主协程将在此行阻塞，等待子协程1将数据拿走，然后放入数据
	fmt.Println("main goroutine run")

	go func() {
		ch <- struct{}{} //此时子协程2，将阻塞在此行
		fmt.Println("goroutine2 over")
	}()

	//3秒后主协程退出，子协程2一直阻塞，直到主协程退出
	time.Sleep(3*time.Second)
	fmt.Println("main exit")
}
```

思考一下程序输出什么：

```
goroutine1 over
main goroutine run
main exit
```

接着看实例：

```go
package main

import "fmt"

func traveseChannal(){
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)   //如果不关闭ch,在主协程中ch数据将空读，造成主协程永远阻塞，即死锁
	}()
  
  //遍历ch并拿走ch中的数据
	for item := range ch {
		fmt.Println(item)
	}
	fmt.Println("bye")
}

func main() {
	traveseChannal()
}
```

再思考一下程序输出什么?

```
1
2
3
bye
```

当程序执行到```for item := range ch ```主协程会读阻塞，然后等待子协程向管道中放入数据

# 总结

* channel满了，就会阻塞写，channel空了就会阻塞读。
* 阻塞之后会交出CPU，去执行其他协程，希望其他协程能够帮助自己解除阻塞状态。
* 如果阻塞发送在main协程里面，并且没有其他子协程可以执行，那就可以确定，“希望永远等不来”，然后main协程就会自己把自己杀掉，然后报：fatal error: deadlock， 即：死锁。
* 如果阻塞发生在子协程里面，就不会发生死锁，因为至少main协程是一个值得等待的“希望”，会一直等(阻塞)下去。

