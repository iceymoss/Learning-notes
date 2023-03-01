[toc]

# 文章介绍

学习每一门编程语言都应该将其常用标准库掌握，这样能使我们更好的使用编程语言，同时也避免了重复造轮子，学习标准库是很有必要的，fmt包实现了类似C语言printf和scanf的格式化I/O。这里我们主要分为向外输出内容和获取输入内容两大部分。



# 向外输出

标准库fmt提供了以下几种输出相关函数。

## Print

向外输出，Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。

```go
func Print(a ...interface{}) (n int, err error)
```

```go
func Printf(format string, a ...interface{}) (n int, err error)
```

```go
func Println(a ...interface{}) (n int, err error)
```

下面来看实例：

```go
fmt.Print("吃饭了嘛？")
name := "iceymoss"
fmt.Printf("%s吃饭了吗\n", name)
fmt.Println("吃饭了嘛？")
```

输出：

```
吃饭了嘛？iceymoss吃饭了吗
吃饭了嘛？
```



## Fprint

Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。这里需要注意：只要满足io.Writer接口的类型都支持写入

方法定义：

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

接下来看实例：

```go
//os.Stdout：标准输出  os.Stdin：标准输入
fmt.Fprint(os.Stdout, "向标准输出中写入内容")
fileOpj, err := os.OpenFile("learn_fmt/test.txt", os.O_WRONLY, 0466)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
}
name := "iceymoss"
fmt.Fprintf(fileOpj, "姓名：%s", name)
```

控制台输出：

```
向标准输出中写入内容
```

然后对应的txt文件：

```tex
姓名：iceymoss
```

实例2：

```go
func RredAndWrite(){
  //读文件，读取为字符串
	str, err := os.ReadFile("learn_fmt/test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(str))
  
  //创建文件
	file, err := os.Create("text2.txt")
	if err != nil {
		panic(err)
	}

  //将读取到的字符串写入创建的文件中
	fmt.Fprintf(file, string(str))
}
```



## Sprint

Sprint系列函数会把传入的数据生成并返回一个字符串。

方法定义：

```go
func Sprintln(a ...interface{}) string

func Sprint(a ...interface{}) string

func Sprintf(format string, a ...interface{}) string
```

来看实例：

```go
gender := fmt.Sprintln("男")
name := fmt.Sprint("iceymoss")
ip := fmt.Sprintf("%s:%d", "127.0.0.1", 8080)
fmt.Print(gender, name, ip)
```

输出：

```
男
iceymoss127.0.0.1:8080
```



## Errorf

Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

方法定义：

```go
func Errorf(format string, a ...interface{}) error
```

实例：

```go
err := fmt.Errorf("这是一个%s", "错误")
	if err != nil {
	fmt.Println("错误：", err)
  return
}
```

输出：

```
错误： 这是一个错误
```



## 格式化占位符

### 通用占位符

| 占位符 | 说明                               |
| ------ | ---------------------------------- |
| %v     | 值的默认格式表示                   |
| %+v    | 类似%v，但输出结构体时会添加字段名 |
| %#v    | 值的Go语法表示                     |
| %T     | 打印值的类型                       |
| %%     | 百分号                             |

```go
fmt.Printf("%v\n", 100)
fmt.Printf("%v\n", false)
o := struct{ name string }{"iceymoss"}
fmt.Printf("%v\n", o)
fmt.Printf("%#v\n", o)
fmt.Printf("%T\n", o)
fmt.Printf("100%%\n")
```

输出：

```
100
false
{iceymoss}
struct { name string }{name:"iceymoss"}
struct { name string }
100%
```



### 布尔型

| 占位符 | 说明        |
| ------ | ----------- |
| %t     | true或false |



### 整型

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %b     | 表示为二进制                                                 |
| %c     | 该值对应的unicode码值                                        |
| %d     | 表示为十进制                                                 |
| %o     | 表示为八进制                                                 |
| %x     | 表示为十六进制，使用a-f                                      |
| %X     | 表示为十六进制，使用A-F                                      |
| %U     | 表示为Unicode格式：U+1234，等价于”U+%04X”                    |
| %q     | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |

实例代码：

```go
n := 65
fmt.Printf("%b\n", n)
fmt.Printf("%c\n", n)
fmt.Printf("%d\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n)
fmt.Printf("%X\n", n)
```

输出：

```
1000001
A
65
101
41
41
```



### 浮点数与复数

| 占位符 | 说明                                                   |
| ------ | ------------------------------------------------------ |
| %b     | 无小数部分、二进制指数的科学计数法，如-123456p-78      |
| %e     | 科学计数法，如-1234.456e+78                            |
| %E     | 科学计数法，如-1234.456E+78                            |
| %f     | 有小数部分但无指数部分，如123.456                      |
| %F     | 等价于%f                                               |
| %g     | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） |
| %G     | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出） |

实例代码：

```go
f := 12.34
fmt.Printf("%b\n", f)
fmt.Printf("%e\n", f)
fmt.Printf("%E\n", f)
fmt.Printf("%f\n", f)
fmt.Printf("%g\n", f)
fmt.Printf("%G\n", f)
```

输出：

```
6946802425218990p-49
1.234000e+01
1.234000E+01
12.340000
12.34
12.34
```



### 字符串和[]byte

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %s     | 直接输出字符串或者[]byte                                     |
| %q     | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x     | 每个字节用两字符十六进制数表示（使用a-f                      |
| %X     | 每个字节用两字符十六进制数表示（使用A-F）                    |

实例代码：

```go
s := "淦饭"
fmt.Printf("%s\n", s)
fmt.Printf("%q\n", s)
fmt.Printf("%x\n", s)
fmt.Printf("%X\n", s)
```

输出：

```
淦饭
"淦饭"
e6b7a6e9a5ad
E6B7A6E9A5AD
```



### 指针

| 占位符 | 说明                           |
| ------ | ------------------------------ |
| %p     | 表示为十六进制，并加上前导的0x |

实例代码：

```go
a := 18
fmt.Printf("%p\n", &a)
fmt.Printf("%#p\n", &a)
```

输出：

```
0xc00001e108
c00001e108
```



## 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下

| 占位符 | 说明               |
| ------ | ------------------ |
| %f     | 默认宽度，默认精度 |
| %9f    | 宽度9，默认精度    |
| %.2f   | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
| %9.f   | 宽度9，精度0       |

示例代码如下：

```go
n := 88.88
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
```

输出结果如下：

```
    88.880000
    88.880000
    88.88
        88.88
           89
```





# 基本输入

## Scan

Go语言fmt包下有Scan、Scanf、Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。下面来看看如何使用

方法定义：

```go
func Scan(a ...interface{}) (n int, err error)
```

该方法默认如何的参数以空格为界，回车键结束输入。

实例代码：

```go
var(
		name string
		age int
		gender string
	)
fmt.Scan(&name, &age, &gender)
fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, gender)
```

输出：

```go
iceymoss 18 男 
扫描结果 name:iceymoss age:18 married:%!t(string=男) 
```



实例：

```go
var(
		name string
		age int
		gender string
)
fmt.Scanf("1:%s, 2:%d, 3:%s", &name, &age, &gender)
fmt.Printf("%s, %d, %s",name, age, gender)
```



实例：

```go
//fmt.Scanln遇到回车就结束扫描了，这个比较常用。
var(
		name string
		age int
		gender string
)
fmt.Scanln( &name, &age, &gender)
fmt.Printf("%s, %d, %s",name, age, gender)
```



实例：

```go
func ScanTest(){
	fmt.Println("请输入您的年龄:")
	var age int
	fmt.Scan(&age)
	for age < 18 {
		fmt.Println("您的年龄不合适进入网吧")
		fmt.Println("有请下一位")
		fmt.Println("请输入您的年龄:")
		fmt.Scan(&age)

	}
	fmt.Printf("欢迎光临！")
}
```



## 其他

有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。

实例代码：

```go
func bufioDemo() {
	//os.Stdin 标准输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)  //返回字符串 s 的一部分，删除了所有前导和尾随空格
	fmt.Printf("%#v\n", text)
}
```

输出：

```
请输入内容：我喜欢唱跳rap篮球
我喜欢唱跳rap篮球
```



# 参考

[go常用标准库](https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/fmt.html)

