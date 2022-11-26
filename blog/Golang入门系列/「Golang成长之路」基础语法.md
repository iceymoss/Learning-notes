[toc]

### 文章介绍

本系列文章将带你快速入门go这门编程语言，重点以实际应用为主，本文将介绍:变量、常量与枚举、数据类型、调节语句、循环语句、函数、指针、go的传值方式。



### 变量定义

Golang 的变量定义相比 c/c++，其最为特殊之处在于 c/c++ 是将变量类型放在变量的前面，而 Go 语言则是将变量类型放在变量的后面，如下：
这是 c/c++:

```cpp
#include <iostream>
using namespace std

int main(){
int a;
int b;
float c;
}
```

Go语言变量定义：

```go
package main
import "fmt"
//求和
func add(k int , y int ){
   var a int
   var b int 
   a = k + y
   b = k - y
   fmt.Println(a)
   fmt.Println(a)
```



### 常量与枚举

#### 常量 

说完变量 来说说常量，如下：

```go
package main
import ("fmt"
     "math"
     )
func consts(){
const (
   filename = "abc.txt"
       a, b = 3, 4
)
var c int 
c = int(math.Sqrt((3*3 + 4*4)))
fmt.Print(c)
fmt.Print("\n")
fmt.Print(filename,c)
}
func main(){
  consts()
}

//输出：
//5
//abc.txt  5
```



#### 枚举(iota）：

```go
package main
import "fmt"
func consts(){
const (
   cpp = 0
   java = 1
   python = 2
   golang = 3
)
fmt.Print(cpp, java, python, golang)
}
func main(){
    consts()
}

//输出：0 1 2 3
```

在 Golang 中可以使用 iota 来实现自增：

```go
package main
import "fmt"
//枚举自增
func enums(){
   const(
   cpp = iota
   java
   python 
   golang  
   )
fmt.Print(cpp, java, python, golang)
}
func main(){
     enums()
}

//输出：0 1 2 3 
```

但是iota的用法较灵活：没新的一行 iota累加1

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值， 会继续累加：iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1  后续会没有赋值，默认为100，直到出现iota为止
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}


//输出：0 1 3 ha ha 100 100 8
```



### 条件语句

#### if语句

```go
var x int = 59  //对于Go语言变量定义也可是：x := 59 可以不对类型进行划分
if x < 60 {  //x < 60 为true；x >= 60 为false
 fmt.Print("不及格")
}
```



#### if-else语句

```go
const filename = "abc1.txt"
pass, err  := ioutil.ReadFile(filename) //pass接受filename信息，err接受出错信息，如果读取正常则err == nil
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("%s\n",pass)
}
```



#### switch语句

在 Go 语言中每一个 case 下是不需要 break 的，它会自动的结束。

```go
func evel1(a,b int , op string) int {
   switch op{
    case "+":
       return a+b
    case "-":
       return a-b
    case "*":
       return a*b
    case "/":
       return a/b
    default:
       return 0
}
}
```



### 循环语句

在 Go 中循环语句最不同的就是没有 `while` 循环，对于 Go 语言，只是用 for 循环就够用了。

#### for 

```go
for{  //不使用如何条件，就是死循环
  fmt.Print("Golang")
}
```

```go
//求100的阶乘
func add2() int {
  var sum int = 0
  for i := 0; i <= 100; i++{
      sum = sum + i
  }
  return sum
}

//sum:5050
```



#### for -range 

```go
package main
import "fmt"
func sum(numbers ...int) int{
   s:= 0
   for i := range numbers{
       s += numbers[i]
   }
   return s
}
func main(){
     fmt.Print(sum(0, 1, 2, 3, 4, 5))
}
```



#### while的替代

前面说过 go 没有 `while`，其实只需要使用 `for` 循环就可以实现 `while`

```go
//没有起始条件和递增条件，就是while循环
for i < 10 {
  fmt.Print("Golang")
}
```



### 函数

在前面的内容中，其实就已经出现函数了，在 Go 中函数的返回值可以有多个，函数在定义是是将返回类型放在函数名的后面。定义函数是需要使用关键字 `func`。

**特点：go的函数可以有多个返回值**

1. 函数的返回类型放在函数名的后面

   ```go
   //求100的阶乘
   func add2() int {  //将int放在函数名之后
   var sum int = 0
   for i := 0; i <= 100; i++{
      sum = sum + i
   }
   return sum
   }
   ```

2. 多个返回值

   ```go
   //交换数值函数
   func sawp(a, b int ) (int,int ){  //返回两个int类型的值
      return b ,a
   }
   ```

   ```go
   //求商，求余
   func div(a, b int) (int, int){
      return a/b ,a%b
   }
   ```

   ```go
   package main
   import "fmt"
   func sum(numbers ...int) int{
      s:= 0
      for i := range numbers{
          s += numbers[i]
      }
      return s
   }
   func main(){
        fmt.Print(sum(0, 1, 2, 3, 4, 5))
   }
   
   //输出：15
   ```



### 指针

#### 什么是指针

在 Go 语言中指针并没有像 c/c++ 中那么麻烦

指针：一个指针变量指向了一个值的内存地址。（也就是我们声明了一个指针之后，可以像变量赋值一样，把一个值的内存地址放入到指针当中。）类似于变量和常量。
下面来看一个下例子：

```go
package main
import "fmt"
//go指针
func pointer() {
   var a int = 5   
   var pa* int  //声明一个指针
   pa = &a      //将指针指向a的地址
   *pa = 1      //更改指针指向地址的值
   fmt.Println(a)
   }
func main(){
     pointer()
  }

//输出：1
```



#### Go函数的传参方式：

##### 值传递：
使用普通变量作为函数参数的时候，在传递参数时只是对变量值得拷贝，就是将实参的值复制给变参，当函数对变参进行处理时，并不会影响原来实参的值。

```go
package main
import (
 "fmt"
)
func swap(a int, b int) {
   var temp int
   temp = a
   a = b
   b = temp
}
func main() {
 x := 5
 y := 10
 swap(x, y)
 fmt.Print(x, y)
}

//输出x,y值并没有变化
//5 10
```

传递给 swap 的是 x，y 的值得拷贝，函数对拷贝的值做了交换，但却没有改变 x,y 的值。



##### "指针传递"

函数的变量不仅可以使用普通变量，还可以使用指针变量，使用指针变量作为函数的参数时，在进行参数传递时将是一个地址拷贝，即将实参的内存地址复制给变参，这时对变参的修改也将会影响到实参的值。

```go
package main
import (
 "fmt"
)
func swap(a *int, b *int) {
   var temp int
   temp = *a
   *a = *b
   *b = temp
}
func main() {
     x := 5
     y := 10
     swap(&x, &y)
     fmt.Print(x, y)
}

//输出：10 5
```















