[toc]



### 使用channel等待任务结束

在前面的内容中很多地方使用到了:
```time.Sleep(time.Millisecond)```
如：
```go
func chanDemo(){
	var channels [10]chan int  //创建channel数组
	for i := 0; i < 10; i++{
		channels[i] = creatworker(i)
	}
	for i := 0; i < 10; i++{
		channels[i] <- 'a' + 1
	}
	time.Sleep(time.Millisecond)
}
```
```go
func bufferedChannel(){
	ch := make(chan int, 3)
	go worker(0, ch)
	for i := 0; i < 10; i++{
		ch <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}
```
```go
func channelClose(){
	ch := make(chan int)
	go worker(0, ch)
	ch <- 'a'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	close(ch)
	time.Sleep(time.Millisecond)
}
```
**在这些方法里面我们很容易知道他们运行所消耗的时间，但事实上，很多程序的时间是不能预估的，我们不能一直是用time包来对程序运行的时间进行预估，是不靠谱的，所以这里我们有了“使用channel等待任务结束”**。

先看这段代码：
```go
func chanDemo(){
   var channels [10]chan int //创建channel数组
   for i := 0; i < 10; i++{
      channels[i] = creatworker(i)
   }
   for i := 0; i < 10; i++{
      channels[i] <- 'a' + 1
  }
   time.Sleep(time.Millisecond)
}
```
我们需要使用channel并发的打印10个字母，为了让字母完整打印，我们对程序运行时间进行了预估，让程序运行1毫秒就结束；下面我们需要使用



#### 使用channel等待任务结束
仍然打印字母(打印20个）部分内容见代码注释
使用'chan bool'的通道来共享通讯来使用内存，告诉main任务结束

```go
package main

import (
   "fmt"
)

//定义一个结构体
//包含一个 'chan int ' 的in和 'chan bool'的done
type Worker struct{
   in chan int
   done chan bool  //对done的接和收作为结束的信号
}

func DoWork( id int, c chan int, done chan bool){
   for {
      n := <-c //接受channel的内容
      fmt.Printf("worker %d received %c\n", id, n)
      done <- true  //done发数据true
  }
}

func createWorker(id int) Worker {
   //建立Worker的一个对象w
   w := Worker{make(chan int), make(chan bool)}

   go DoWork(id, w.in, w.done) //此处启动goroutine，即并发
   return w
}

func ChanDemo() {

   var workers [10]Worker  //创建10个抽象类型Worker

  for i := 0; i < 10; i++{ 
      workers[i] = createWorker(i) //创建10个Worker的对象，并返回给workers[i]
   }

   for i := 0; i < 10; i++{  //可使用range
      workers[i].in <- 'a' + i   //workers[i].in接受数据
   }

   for _, worker := range workers {
      <-worker.done   //将done的数据发给mian，告知main该任务结束
   }

   for i := 0; i < 10; i++{ //可使用range
      workers[i].in <- 'A'+ i  //workers[i].in接受数据
   }

   for _, worker := range workers{
      <- worker.done  //将done的数据发给mian，告知main该任务结束
   }
}

func main(){
   ChanDemo()
}
```
打印结果：
worker 0 received a
worker 5 received f
worker 1 received b
worker 6 received g
worker 4 received e
worker 9 received j
worker 8 received i
worker 2 received c
worker 7 received h
worker 3 received d
worker 6 received G
worker 2 received C
worker 3 received D
worker 7 received H
worker 1 received B
worker 4 received E
worker 5 received F
worker 0 received A
worker 9 received J
worker 8 received I

从打印结果看出：是按顺序打印的(先小写后大写)
这里还有一种方法：
```go
func ChanDemo() {

   var workers [10]Worker

  for i := 0; i < 10; i++{
      workers[i] = createWorker(i)
   }
   for i := 0; i < 10; i++{
      workers[i].in <- 'a' + i
   }
   for i := 0; i < 10; i++{
      workers[i].in <- 'A'+ i
   }
   //将两个<- worker.done 放在一起
   for _, worker := range workers{
      <- worker.done   
	  <- worker.done
   }
}
```
但是需要注意的是：
我们一共创建了10个goroutine，在第二个for中就已经向所有channel中发送数据，接着第三个for，又向channel中发送数据，这样会死锁，因为第一次channel中的数据没有人来接，然后又向channel发数据。
解决方法：在```DoWork```函数增加并发,让``` done <- true```处于并发执行状态，可随时向main发数据。
```go
func DoWork( id int, c chan int, done chan bool){
   for {
      n := <-c //接受channel的内容
      fmt.Printf("worker %d received %c\n", id, n)
      go func() {
	    done <- true
	 }()
  }
}
```



#### 使用系统提供的 'WaitGroup'等待任务结束
WaitGroup提供了：Add()、Wait()、Done()方法

```go
package main

import (
   "fmt"
 "sync")

type Worker struct{
   in chan int
   wg *sync.WaitGroup  //引用需要指针
}

func DoWork( id int, c chan int, wg *sync.WaitGroup){
   for {
      n := <-c //接受channel的内容
      fmt.Printf("worker %d received %c\n", id, n)
       wg.Done()  //接和收结束信息
   }
}

func createWorker(id int, wg *sync.WaitGroup) Worker {
   //建立Worker的一个对象w
   w := Worker{make(chan int), wg,}

   go DoWork(id, w.in, wg)
   return w
}

func ChanDemo() {
   var wg sync.WaitGroup

   var workers [10]Worker
   for i := 0; i < 10; i++{
      workers[i] = createWorker(i, &wg)  //指针
   }
   wg.Add(20)  //20个任务
   for i := 0; i < 10; i++{
      workers[i].in <- 'a' + i
   }
   for i := 0; i < 10; i++{
      workers[i].in <- 'A'+ i
   }
   wg.Wait()  //任务结束
}

func main(){
   ChanDemo()
}
```
打印结果：
worker 0 received a
worker 4 received e
worker 6 received g
worker 2 received c
worker 9 received j
worker 7 received h
worker 0 received A
worker 8 received i
worker 5 received f
worker 3 received d
worker 1 received b
worker 1 received B
worker 9 received J
worker 2 received C
worker 3 received D
worker 4 received E
worker 7 received H
worker 8 received I
worker 6 received G
worker 5 received F



### select

之前的内容中，我们使用channel都是一个一个的收数据，如果我们需要把多个channel同时收，该怎么办？
答案是：Go语言引入了***select***语句
> 下面来具体介绍一下***select***:
> select的逻辑和switch的逻辑类似，他们都有多个case分支和default，但select是针对channel的，其逻辑是：在多个含有case分支的select里面，当某时刻相应的channel满足发发出数据，让外面接收，就能满足对应case，接下来就会执行该case对应的语句块,如果多个case同时都满足条件，则会随机选择其中一个case，如果所有case都不满足则会执行default
> 例如：
```go
var activeWorker chan<- int
n := 0
select {
  //c1, c2 为chan int类型
  case n = <-c1:
     fmt.Printf("this is c1:%d\n", n)
  case n = <-c2:
     fmt.Printf("this is c2:%d\n", n)
  case activeWorker <- n:
     hasValue = false
  default:
     fmt.Println("not find channel")
    return
  }
```
在执行select时，程序会将所有的case分析一遍，先来看第一个case，如果此时c1发出数据，则第一个case可被执行，再看第二个case，如果此时c2发出数据，则第二个case可被执行，再看第三个case，如果此时n有值就会将其值发给activeWorker,最后来看default，当上面所有的case都不满足时，就会执行default的语句块。

下面来看一个完整的select的应用：
```go
package main

import (
   "fmt"
 "math/rand" "time")

//控制时间，向channel里面发送消息
func generator() chan int{
   out := make(chan int)
   go func() {
      i := 0
     for{
         //控制发送数据时间间隔
         time.Sleep(time.Duration( rand.Intn(1500) ) * time.Microsecond)
         out <- i
         i++
      }
   }()
   return out
}

//channel接受和打印信息
func DoWork( id int, c chan int){
   for {
    n := <- c  //接受channel的内容
    time.Sleep(time.Second)  //控制打印时间间隔
    fmt.Printf("worker %d received %d\n", id, n)
   }
}

//建立channel，启动并发
func createWorker(id int) chan<- int{
   w := make(chan int)
   go DoWork(id, w)
   return w
}

func main() {
   var c1, c2 = generator(), generator()
   var worker = createWorker(0)
   n := 0
   var Values []int //动态缓存数据
   //tm为程序总时间  tm := time.After(1 * time.Second)
   tick := time.Tick(time.Second)

   for{
      var activeWorker chan<- int
      var activeValue int
      if len(Values) > 0 {
          activeWorker = worker
          activeValue = Values[0]
      }
      select {
      case n = <-c1:
         Values = append(Values, n)
      case n = <-c2:
         Values = append(Values, n)
      case activeWorker <- activeValue:
         Values = Values[1:]
      case <- time.After(2000 * time.Microsecond):  //每两次发送数据时间差超过800毫秒执行一次
         fmt.Println("timeout")
      case <- tick:   //使用tick反映系统状态
         fmt.Println("queue len is:",len(Values))
      case <- tm:  //使用tm控制总时间
         fmt.Println("bey")
         return
     }
   }
}
```
打印结果：
worker 0 received 132
worker 0 received 133
worker 0 received 134
worker 0 received 136
worker 0 received 135
worker 0 received 136
worker 0 received 137
worker 0 received 138
worker 0 received 137
worker 0 received 138
worker 0 received 139
timeout
worker 0 received 140
worker 0 received 139
timeout
worker 0 received 140
worker 0 received 141
worker 0 received 142
worker 0 received 143
worker 0 received 141
timeout
worker 0 received 142
worker 0 received 144
worker 0 received 145
worker 0 received 143
worker 0 received 144
worker 0 received 145
worker 0 received 146
worker 0 received 146
worker 0 received 147
worker 0 received 147
worker 0 received 148
worker 0 received 148
worker 0 received 149
worker 0 received 149
worker 0 received 150
worker 0 received 151
timeout
worker 0 received 152
worker 0 received 150
worker 0 received 153
worker 0 received 151
worker 0 received 152
worker 0 received 153
worker 0 received 154
worker 0 received 154
timeout
worker 0 received 155
worker 0 received 156
worker 0 received 155
worker 0 received 157
worker 0 received 156
worker 0 received 158
worker 0 received 157
timeout
worker 0 received 158
worker 0 received 159
worker 0 received 159
worker 0 received 160
worker 0 received 161
worker 0 received 162
worker 0 received 160
timeout
worker 0 received 161
worker 0 received 163
worker 0 received 162
timeout
worker 0 received 164
worker 0 received 163
worker 0 received 164
worker 0 received 165
worker 0 received 166
worker 0 received 167
worker 0 received 165
worker 0 received 168
timeout
worker 0 received 166
worker 0 received 169
worker 0 received 167
worker 0 received 170
timeout
worker 0 received 171
worker 0 received 168
timeout
worker 0 received 172
worker 0 received 169
worker 0 received 170
worker 0 received 171
worker 0 received 173
timeout
worker 0 received 174
worker 0 received 172
worker 0 received 175
……
……
……
worker 0 received 2352
worker 0 received 2386
timeout
worker 0 received 2387
worker 0 received 2353
timeout
worker 0 received 2388
worker 0 received 2389
worker 0 received 2354
worker 0 received 2390
worker 0 received 2355
bey
这个程序充分体现了select的实际应用



### 在这里总结了几个常见的问题

#### 提问

```go
func ChanDemo() {

var workers [10]Worker

  for i := 0; i < 10; i++{
workers[i] = createWorker(i)
}

for i := 0; i < 10; i++{
workers[i].in <- 'a' + i
}
```
在第一个for中，第一步workers[0] = createWorker(0)

然后就进入这里
```go
func createWorker(id int) Worker {
//建立Worker的一个对象w
  w := Worker{make(chan int),
     make(chan bool),
     }

go DoWork(id, w.in, w.done)
return w
}
```
在这个函数中我们开了一个goroutine，同时我们会将w返回给workers[0]，然后就进入：
```go
for i := 0; i < 10; i++{
workers[i] = createWorker(i)
}
```
的第二次，第三次循环……

直到循环结束。

但是这里就有问题了，在这个途中我们一共开了10 goroutine，但是这10 goroutine都处于等待状态(因为我们还没有给channel任何内容，从我们的输出结果可以看出）

1.  那么这里的10个goroutine是处于等待状态是不是因为，我们channel没有接受到任何信息，所以就会造成goroutine的等待？
2.  还有这里：

```go
 func DoWork( id int, c chan int, done chan bool){
    for {
         n := <-c //接受channel的内容
         fmt.Printf("id: %v, chan:%c\n", id, n)
         done <- true
      }
    }
```
这个死循环，为什么在函数调用后只循环了一遍？ 当然这里我知道他是其中一个goroutine

3\. 当然还有一个问题，就是我们在前两个问题在基础上，调用函数DoWork()时，也会对应的将true发送给与之对应的workers[i].done中，然后：
```go
for i := 0; i < 10; i++{
workers[i].in <- 'a' + i
}

for _, worker := range workers {
<- worker.done
}
```
在这里的第二个for中，这里<- worker.done全为true，我们是不是从这里就可以了解到前面的10个goroutine结束了？

4\. 也正是因为这样我们才不需要time包，来预计程序的运行时间了？

#### 回答

1.  是的，它们此时都在等待，等别人从in中发送任务数据。

2.  这是个死循环，一般我们goroutine中常会这么写，只要有任务就做。视频里实际上大写字母，小写字母，一共执行两遍。执行多少遍取决于外界，这里是main函数，到底发送了多少任务给我这个worker[i]。

3.  这里的true方向同学搞错了，是worker通知main函数，说我做完了。<- worker.done这里是main函数接收worker.done的数据，如果收到，就说明这个worker的事情做完了。

4.  是的，理想情况下应该不需要time包来预计运行时间。预计的时间会不靠谱。