[toc]

### Log

Go语言内置的log包实现了简单的日志服务 log包为我们封装了一系列日志相关方法。本文介绍了标准库log的基本使用。

### 使用Logger

log包定义了Logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。

下面我们直接来使用这些方法：

```go
package main

import "log"

func main() {
	log.Println("我是一条日志")
	log.Printf("我是%d条日志\n", 24)
	log.Print("我是日志")

	logger := log.Default()
	logger.Println("我是一条logger的日志")

	log.Fatal("这是一个故障")
	log.Panicln("异常")
}
```

默认会将日志信息打印到终端：

```
2023/03/06 21:48:21 我是一条日志
2023/03/06 21:48:21 我是24条日志
2023/03/06 21:48:21 我是日志
2023/03/06 21:48:21 我是一条logger的日志
2023/03/06 21:48:21 这是一个故障
```

logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误；Fatal系列方法填写日志后会调用os.Exit(1)程序退出；Panic系列函数会在写入日志信息后panic。



### logger的配置

在一些应用场景中，我们只需要看到日志内容，日期，时间等，我们需要看到更多的信息来对程序进行追踪，例如记录该日志的文件名和行号等。

```go
//log标准库中的Flags函数会返回标准logger的输出配置
func Flags() int
//SetFlags函数用来设置标准logger的输出配置
func SetFlags(flag int)
```

log标准库提供了如下的flag选项，它们是一系列定义好的常量

```go
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

#### 如何使用

我们在记录日志前，先配置日志

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(log.Llongfile, log.Lmicroseconds,log.Ldate )
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("我是一条很普通的日志。")
}
```

日志输出：

```
2023/03/06 22:01:33.028467 /Users/feng/go/src/gostudy/StandardPackage/learn_log/main.go:23: 这是一条很普通的日志。
8 4 1
```



### 配置日志前缀

```go
func Prefix() string
func SetPrefix(prefix string)
```

#### 如何使用

和flags类似，其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。

```go
package main

import "log"

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("【小黑子】")
	log.Println("这是一条很普通的日志。")
}
```

日志输出：

```
2023/03/06 22:08:43.403669 /Users/feng/go/src/gostudy/StandardPackage/learn_log/main.go:26: 这是一条很普通的日志。
【小黑子】2023/03/06 22:08:43.403794 /Users/feng/go/src/gostudy/StandardPackage/learn_log/main.go:28: 这是一条很普通的日志。
```

这样我们很容易的给日志前缀加上了对应内容，这样更方便我们对日志信息检索和处理。



### 日志输出的位置

前面我们使用的日志输出都是将日志向控制台输出，但实际项目或者工作中，我们需要对应用的日志进行记录，我们可以将对应日志写入对应文件中。

```go
func SetOutput(w io.Writer)
```



#### 如何使用

将日志写入文件：

```go
func LogToFile(){
  //打开文件给定相应权限
	file, err := os.OpenFile("./learn_log/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("打开文件失败", err)
	}

  //将日志写入文件
	log.SetOutput(file)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("【小黑子】")
	log.Println("这是一条普通日志")
}
```

在./learn_log/test.log文件中可以看到：

```
【小黑子】2023/03/06 22:20:27.332350 /Users/feng/go/src/gostudy/StandardPackage/learn_log/main.go:17: 这是一条普通日志
```

当然，如果我们从开始就想要将日志写入日志文件的话，我们可以在init方法中配置

```go
package main

import (
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("./learn_log/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("打开文件失败", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("【小黑子】")
}

func main() {
	log.Println("这是一条测试日志")
}

```

这样后面的日志都将写入文件中



### 创建logger

log包为我们提供了new()方法：

```go
func New(out io.Writer, prefix string, flag int) *Logger
```

入参分别是：写入日志位置，日志前缀内容，参数flag定义日志的属性（时间、文件等等）

实例代码：

```go
//写入标准输出
logger := log.New(os.Stdout, "【ikun之家】", log.Llongfile | log.Lmicroseconds | log.Ldate)
logger.Println("这是一条测试日志")
```

go的标准库log包为我们也就提供了这些日志方法，其实相关功能有限，如果遇到一些更复杂的场景，log包可能不能满足开发者的需求；当然我们也可以选择第三方日志库：如 zap。



