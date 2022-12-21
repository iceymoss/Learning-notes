[toc]



### 并发任务的控制

在前面的内容中我们了解到了goroutine、channel和go常见的并发模式，这本文中我们来讲解一下Go语言并发任务的控制
>1. 非阻塞等待
>
>  非阻塞等待经常用来判断哪一个channel数据传输更快，我们使用select来判断channel是否为阻塞状态，如果阻塞则返回空串，和false，如果费阻塞则返回从channel里拿到的值和true：
```go
//非阻塞等待
func nonBlockWait(c chan string) (string, bool){
   select{
   case m := <- c:
      return m,true
 default:
      return "", false
    }
}
```
调用：
```go
func main() {
   m1 := msgGen("service1")
   m2 := msgGen("sercive2")
   for{
      fmt.Println(<-m1)
      if m, ok := nonBlockWait(m2); ok{
         fmt.Println(m)
      }else{
         fmt.Println("no mssage from sercive2")
      }
   }
}
```

>mssGen方法：
```go
func msgGen(name string) chan string {
   c := make(chan string)
   go func(){
      i := 0
  for {
         time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
         c <- fmt.Sprintf("service: %s, message ：%d", name, i)
         i++
      }
   }()
   return c
}
```
打印结果：
service: service1, message ：0
no mssage from sercive2
service: service1, message ：1
service: sercive2, message ：0
service: service1, message ：2
service: sercive2, message ：1
service: service1, message ：3
no mssage from sercive2
service: service1, message ：4
no mssage from sercive2
service: service1, message ：5
no mssage from sercive2

>2.超时等待
>其逻辑和非阻塞等待类似，我们需要使用time.after通道来控制时间：
```go
func timeoutWait(c chan string, tm time.Duration) (string, bool){
   select{
   case m := <-c:
      return m, true
 case <- time.After(tm):
      return "", false
  }
}
```
调用：
```go
func main() {
   m1 := msgGen("service1")
   m2 := msgGen("sercive2")
   for{
      fmt.Println(<-m1)
      if m, ok := timeoutWait(m2, time.Second); ok{
         fmt.Println(m)
      }else{
         fmt.Println("timeout")
      }
   }
}
```
mssGen方法不变
打印结果：
service: service1, message ：0
service: sercive2, message ：0
service: service1, message ：1
timeout
service: service1, message ：2
service: sercive2, message ：1
service: service1, message ：3
timeout
service: service1, message ：4
service: sercive2, message ：2

> 3. 任务中断/退出
> 当main结束时，会将所有的goroutine都杀掉，但是如果中途一下任务没有做完，我们希望外层可以通知我们，“说你要退出了”，那么我们如何实现呢？这里我们需要再次引入done用于，主goroutine给goroutine发信息
> 这是任务中断/退出的核心代码和逻辑：

```go
func msgGen(name string, done chan struct{}) chan string {  //done也可使用bool，struct{}内部没数据，比bool更节约空间
  c := make(chan string)
   go func(){
      i := 0
  for {
         select {
         case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):  //每隔time.Duration(rand.Intn(5000)) * time.Millisecond的时间就会向c中发数据
            c <- fmt.Sprintf("service: %s, message ：%d", name, i)
         case <-done:  //当main中向done中发数据，程序就会执行该case，最后退出
            fmt.Println("cleaing up")
            return
  }
         i++
      }

   }()
   return c
}
```

main：
```go
func main() {
   done := make(chan struct{})
   m1 := msgGen("service1", done)
   for i := 0; i < 5; i++{   //打印五遍就结束
      fmt.Println(<-m1)
      if m, ok := timeoutWait(m1, time.Second); ok{
         fmt.Println(m)
      }else{
         fmt.Println("timeout")
      }
   }
   //main通知其他goroutine，main将结束向done中发数据
   done <- struct{}{} //第一个{}是结构定义，第二个{}是初始化
  time.Sleep(time.Second)   //防止main瞬间退出
}
```
打印结果：
service: service1, message ：0
timeout
service: service1, message ：1
timeout
service: service1, message ：2
timeout
service: service1, message ：3
timeout
service: service1, message ：4
timeout
cleaing up

>4. 优雅退出
>在3.中讲到任务中断，在这个过程中，只是main(主goroutine)给其他goroutine发通知，说我要退了，没有真正的交互，这就显得不那么优雅，现在我们需要优雅的退出，如：男：我走了；女：好的，你走吧。这是不是很优雅，相互应答，好下面就来做这件事：
* 方法一：在“中断任务”的基础上增加一个```chan boo```类型的dones，用来回应，看代码：

```go
func msgGen(name string, done chan struct{}, dones chan bool) chan string {  //done也可使用bool，struct{}内部没数据，比bool更节约空间
  c := make(chan string)
   go func(){
      i := 0
  for {
         select {
         case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
            c <- fmt.Sprintf("service: %s, message ：%d", name, i)
         case <-done:
            fmt.Println("cleaing up")
            time.Sleep(2 * time.Second)
            fmt.Println("cleaing done")
            dones <- true
 return  }
         i++
      }

   }()
   return c
}
```
```go
func main() {
   done := make(chan struct{})
   dones := make(chan bool)
   m1 := msgGen("service1", done, dones)
   for i := 0; i < 5; i++{
      fmt.Println(<-m1)
      if m, ok := timeoutWait(m1, time.Second); ok{
         fmt.Println(m)
      }else{
         fmt.Println("timeout")
      }
   }
   done <- struct{}{} //第一个{}是结构定义，第二个{}是初始化
  <-dones
}
```
打印结果：
service: service1, message ：0
timeout
service: service1, message ：1
timeout
service: service1, message ：2
service: service1, message ：3
service: service1, message ：4
timeout
service: service1, message ：5
service: service1, message ：6
cleaing up
cleaing done

> * 方法二
> 不需要增加参数，只需要将```chan struct{}```类型的done在
> ```go func{……}()```中添加：```done <- struct{}{}```即可，看代码

```go
func msgGen(name string, done chan struct{}) chan string {  //done也可使用bool，struct{}内部没数据，比bool更节约空间
  c := make(chan string)
   go func(){
      i := 0
  for {
         select {
         case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
            c <- fmt.Sprintf("service: %s, message ：%d", name, i)
         case <-done:
            fmt.Println("cleaing up")
            time.Sleep(2 * time.Second)
            fmt.Println("cleaing done")
            done <- struct{}{}
 return  }
         i++
      }

   }()
   return c

}
```
```go
func main() {
   done := make(chan struct{})
   dones := make(chan bool)
   m1 := msgGen("service1", done, dones)
   for i := 0; i < 5; i++{
      fmt.Println(<-m1)
      if m, ok := timeoutWait(m1, time.Second); ok{
         fmt.Println(m)
      }else{
         fmt.Println("timeout")
      }
   }
   done <- struct{}{} //第一个{}是结构定义，第二个{}是初始化
  <-done
}
```
打印结果：
service: service1, message ：0
timeout
service: service1, message ：1
timeout
service: service1, message ：2
service: service1, message ：3
service: service1, message ：4
timeout
service: service1, message ：5
service: service1, message ：6
cleaing up
cleaing done

golang并发编程的内容我们就介绍到这里了，并发编程和和传统编程不同，并发编程的跳跃性较大，更难理解，大家加油！