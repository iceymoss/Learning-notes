[toc]

### flag

Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。

#### os.Args

如果你只是简单的想要获取命令行参数，可以像下面的代码示例一样使用os.Args来获取命令行参数。

```go
import (
	"fmt"
	"os"
)
//ArgsDemo 如果你只是简单的想要获取命令行参数，可以像下面的代码示例一样使用os.Args来获取命令行参数。
//os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。
func ArgsDemo(){
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

func main() {
	ArgsDemo()
}
```

然后进行编译：```go build main.go```后执行输出：

```
$ ./main 您好 吃饭了嘛？
args[0]=./main
args[1]=您好
args[2]=吃饭了嘛？
```



### 定义命令行flag参数

有两种定义flag参数方法

####  flag.Type()

方法定义：

```go
func String(name string, value string, usage string) *string
```

Type(flag名", "默认值", "帮助信息")，例如我们下面以，姓名，年龄，性别为例子：

```go
name := flag.String("name", "iceymoss", "姓名")
age := flag.Int("age", 18, "年龄")
gender := flag.Bool("gender", false, "性别")
```

这里需要注意的是，方法的返回值是指针，即对应类型的指针。



####  flag.TypeVar()

基本格式如下： flag.TypeVar(Type指针, flag名, 默认值, 帮助信息) 例如我们要定义姓名、年龄、性别三个命令行参数，我们可以按如下方式定义：

```go
var name string
var age int
var gender bool
flag.StringVar(&name, "name","iceymoss", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&gender,"gender", false, "性别")
```



#### flag.Parse()

通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。

支持的命令行参数格式有以下几种：

- -flag xxx （使用空格，一个-符号）
- --flag xxx （使用空格，两个-符号）
- -flag=xxx （使用等号，一个-符号）
- --flag=xxx （使用等号，两个-符号）

其中，布尔类型的参数必须使用等号的方式指定。

Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。



#### flag.Args()

返回命令行参数后的其他参数，以[]string类型



#### flag.NArg() 

flag.NArg() //返回命令行参数后的其他参数个数



#### flag.NFlag()

flag.NFlag() //返回使用的命令行参数个数



### 实例

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var age int
	var gender bool
	flag.StringVar(&name, "name","iceymoss", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&gender,"gender", false, "性别")

	//解析命令行参数
	flag.Parse()
	fmt.Printf("name:%s, age:%d, gender: %t", name, age, gender)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
```

或者

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String( "name","iceymoss", "姓名")
	age := flag.Int("age", 18, "年龄")
	gender := flag.Bool("gender", false, "性别")

	//解析命令行参数
	flag.Parse()
	fmt.Printf("name:%s, age:%d, gender: %t\n", *name, *age, *gender)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
```



### 如何使用

使用命令：```go build main.go```编译，然后使用命令：

```
$ ./main -help
Usage of ./main:
  -age int
        年龄 (default 18)
  -gender
        性别
  -name string
        姓名 (default "iceymoss")

```

使用命令行flag参数

```
$ ./main --name 土拨鼠 -age 22 -gender=true 
name:土拨鼠, age:22, gender: true
[]
0
3
```

使用非flag参数：

```
$ ./main a b c d
name:iceymoss, age:18, gender: false
[a b c d]
4
0
```