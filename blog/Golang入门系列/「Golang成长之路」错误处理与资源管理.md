### defer调用

**介绍：**
确保在函数调用结束时发生
defer就是在程序或函数调用结束后，然后执行defer后面的语句
例如：

```go
package main

import "fmt"

func tryDefer(){

	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	//打印：1、2、3

	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//打印：3、2、1
	
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
	//打印：3、2、1
	
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
	//打印：3、2、1
	//panic: error occurred
}
func mian(){
	tryDefer()
}
```
1. 在这个例子第二个打印中，前面加上```defer```关键字后，其后面的 fmt.Println(1)和 fmt.Println(2)都会该程序所有语句执行结束后执行，而defer里面本身是一个栈(先进后出)，所有会先打印2，然后打印1。
2. 在这个例子的第三个打印中，我们加了关键字```return```其输出结果还是1、2、3，这里是因为return后程序就会退出，当程序退出前defer仍然会执行
3. 在这个例子的第四个打印中，我们使用到函数```panic()```程序会报错退出，但是在结束之前仍然会执行```defer```

还有一个例子：
```go
func writeFile(filename string){
	file, err := os.Create(filename)  //os.Create()建立文件
	if err != nil{  //err不为空，则文件建立失败，panic输出err信息
		panic(err)
	}
	
	defer File.close() //最后将文件关闭
	
	writer := bufio.NerWriter(file)
	defer writer.Flush()
	
	f := fibonacci()
	for i := 0; i < 20; i++{
		fmt.Fprintln(writer, f())
	}
}

func fibonacci() func() int {
	a := 0, b := 1
	return func() int{
		a, b = b, a+b
		return a
	}
}

func mian(){
	writerFile("fib.txt")
}
```



### 错误处理

前面我们提到了```panic()```函数其实它也是用来做错误处理的，但是，使用```panic```程序会挂掉，这是我们不希望看到的。
```go
file, err := os.Open("abc.txt")
if err != nil{
	panic(err)
}
```
像上面代码使用panic后程序就会把err中的内容输出，并且挂掉

```go
file, err := os.Open("abc.txt")
if err != nil{
	pathError, ok := err.(*os.PathError)
	if !ok{
		fmt.Println(pathError)
	}else {
		fmt.Println("unknown error ", err)
	}
}
```

使用这种处理方式更好



### panic&recover

#### panic

![「Golang成长之路」错误处理与资源管理篇](https://cdn.learnku.com/uploads/images/202108/14/69310/WDmKTuijKO.png!large)

#### recover

![「Golang成长之路」错误处理与资源管理篇](https://cdn.learnku.com/uploads/images/202108/14/69310/8575VU2ZrG.png!large)
```go
package main
import "fmt"

func tryRecover(){
	defer func(){
	r := recover()
	err, ok := r.(error)
	if ok{
		fmt.Println("Error occurred:", err)
	}else{
		panic("I don`t kown what to : %v", r)
	}
  }()

panic("123")

}

func mian(){
	tryRecover()

}

```



#### 统一处理

服务器统一出错：
web.go:
```go
package main

import (
   "learngo/errhandling/filelistingserver/handle"
 "log" 
  "net/http"
  "os"
)

//定义一个函数类型
type Apphandle func(write http.ResponseWriter,
  request *http.Request) error

//函数式编程,传入一个函数Apphandle类型的函数类型，返回一个匿名函数func(http.ResponseWriter,*http.Request)
func errwarappler(handler Apphandle) func(http.ResponseWriter, *http.Request) {
   return func(writer http.ResponseWriter, request *http.Request) {
      //匿名函数func(http.ResponseWriter,*http.Request)内部构造：

  defer func(){
         r := recover()
         if r != nil{
            log.Printf("panic:%v\n", r)
           //写入返回err信息
            http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
  }

      }()

      err := handler(writer, request)
      if err != nil{
         log.Printf("Error handle request:%s", err.Error())

          if userErr, ok := err.(userError); ok{
            http.Error(writer, userErr.Massage(), http.StatusBadRequest)
            return
  }

         code := http.StatusOK
  switch{
         case os.IsNotExist(err):  //情况一：文件不存在
  code = http.StatusNotFound
  case os.IsPermission(err): //情况二：没有权限
  code = http.StatusForbidden
  default:
            code = http.StatusInternalServerError

  }
         http.Error(writer, http.StatusText(code), code)
      }

   }
}

type userError interface{
   error
  Massage() string
}

func main() {
   //传入部分
  http.HandleFunc("/",errwarappler(handle.Handler))

   err := http.ListenAndServe(":8888" , nil)
   if err != nil{
      panic(err)
   }
}
```
handle.go:
```go
package handle

import (
   "fmt"
 "io/ioutil"
  "net/http"
  "os"
  "strings"
)

const perfix = "/list/"

type userError string

func (e userError)Error() string{
   return e.Massage()
}

func (e userError)Massage() string {
   return string(e)
}

func Handler(write http.ResponseWriter, Request *http.Request) error{

   if strings.Index(Request.URL.Path, perfix) != 0 {
      return userError(fmt.Sprintf("path %s must start" +
         "whit %s",Request.URL.Path, perfix))
   }

   path := Request.URL.Path[len(perfix):]
   file , err := os.Open(path)
   if err != nil{
      //将错误信息返回出去
  return err
   }
   defer file.Close()

   all, err := ioutil.ReadAll(file)
   if err != nil{
      return err
   }

   write.Write(all)
   return nil
}
```
**注意：error和panic**
在意料之外的情况：使用panic
在意料之内的情况：使用error