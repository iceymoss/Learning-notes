[toc]

### 一、变量定义
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

Go 语言变量定义：

```go
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



###  二、Golang 的数据类型

#### 数据类型

* (u) int 型

```(u)int,(u)int8, (u)int16, (u)int32, (u)int64```

* float 型

  float32, float64 //不能只是用float,在go语言中，不存在单独的float
  string 型
  在 go 中 string 和其他语言一样

```go
package main
import "fmt"

func main(){
     var a string = "abcdefg"
     fmt.Print(a)
}
```

输出如下：

```go
abcdefg
```



* bool 型
  对于 bool 来说，go 中没有什么不同的，和其他编程语言一样，只有 true 和 false

* rune
  rune 在其他编程语言中叫 char—— 字符型；其长度为 32 位。

* complex
  分为两类：

  complex64, complex128

* byte

  

#### 类型转换
类型转换就是将一种类型装换为另一种类型，如：



```go
package main

import ("fmt"
        "math"
       )
func exchange(g1 int ,g2 int ){
   var x  int
   x = int(math.Sqrt(float64(g1 * g1 + g2 * g2))) //这里开根号一定是浮点数 需要将float64(g1 * g1 + g2 * g2)转换为float类型，又因为x在定义是为 int 类型，所以又需要将最后转换为int类型
   fmt.Print(x)
}

//最后输出：5
```





### 三、常量与枚举

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
// 5
// abc.txt  5
```


#### 枚举
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

//输出：
// 0 1 2 3
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
输出：
0 1 2 3
```

#### 条件语句
* if 语句

  ```go
  var x int = 59 //对于Go语言变量定义也可是：x := 59 可以不对类型进行划分
  if x < 60 {  //x < 60 为true；x >= 60 为false
   fmt.Print("不及格")
  }
  
  ```

* if - else 语句

```go
const filename = "abc1.txt"
pass, err  := ioutil.ReadFile(filename) //pass接受filename信息，err接受出错信息，如果读取正常则err == nil
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("%s\n",pass)
}
```

* switch 语句
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



### 三、 循环语句

在 Go 中循环语句最不同的就是没有 while 循环，对于 Go 语言，只是用 for 循环就够用了。

#### for 死循环

```go
for{  //不使用如何条件，就是死循环
  fmt.Print("Golang")
} 
```

while 循环的取代
前面说过 go 没有 while，其实只需要使用 for 循环就可以实现 while

```go
//没有起始条件和递增条件，就是while循环
for i < 10 {
  fmt.Print("Golang")
}
for 循环
//求100的阶乘
func add2() int {
  var sum int = 0
  for i := 0; i <= 100; i++{
      sum = sum + i
  }
  return sum
}
sum = 5050
```







### 四、函数

在前面的内容中，其实就已经出现函数了，在 Go 中函数的返回值可以有多个，函数在定义是是将返回类型放在函数名的后面。定义函数是需要使用关键字 func。

#### 函数的返回类型放在函数名的后面

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

#### Go 函数可以有多个返回值

```go
//交换数值函数
func sawp(a, b int ) (int,int ){  //返回两个int类型的值
   return b ,a
}
```

又如：

```go
//求商，求余
func div(a, b int) (int, int){
   return a/b ,a%b
}
```

函数的另一种求和

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

```
输出：
15
```









### 五、指针
在 Go 语言中指针并没有像 c/c++ 中那么麻烦
#### 指针：一个指针变量指向了一个值的内存地址（也就是我们声明了一个指针之后，可以像变量赋值一样，把一个值的内存地址放入到指针当中）类似于变量和常量

```go
package main

import "fmt"

func main() {
    // a,b 是一个值
    a := 5
    b := 6

    fmt.Println("a的值：", a)

    // 指针变量 c 存储的是变量 a 的内存地址
    c := &a
    fmt.Println("a的内存地址：", c)

    // 指针变量不允许直接赋值，需要使用 * 获取引用
    //c = 4

    // 将指针变量 c 指向的内存里面的值设置为4
    *c = 4
    fmt.Println("a的值：", a)

    // 指针变量 c 现在存储的是变量 b 的内存地址
    c = &b
    fmt.Println("b的内存地址：", c)

    // 将指针变量 c 指向的内存里面的值设置为8
    *c = 8
    fmt.Println("a的值：", a)
    fmt.Println("b的值：", b)

    // 把指针变量 c 赋予 c1, c1 是一个引用变量，存的只是指针地址，他们两个现在是独立的了
    c1 := c
    fmt.Println("c的内存地址：", c)
    fmt.Println("c1的内存地址：", c1)

    // 将指针变量 c 指向的内存里面的值设置为9
    *c = 9
    fmt.Println("c指向的内存地址的值", *c)
    fmt.Println("c1指向的内存地址的值", *c1)

    // 指针变量 c 现在存储的是变量 a 的内存地址，但 c1 还是不变
    c = &a
    fmt.Println("c的内存地址：", c)
    fmt.Println("c1的内存地址：", c1)
}

```

打印出：

```go
a的值： 5
a的内存地址： 0xc000016070
a的值： 4
b的内存地址： 0xc000016078
a的值： 4
b的值： 8
c的内存地址： 0xc000016078
c1的内存地址： 0xc000016078
c指向的内存地址的值 9
c1指向的内存地址的值 9
c的内存地址： 0xc000016070
c1的内存地址： 0xc000016078

```

那么 a，b 是一个值变量，而 c 是指针变量，c1 是引用变量。

如果 & 加在变量 a 前：c := &a，表示取变量 a 的内存地址，c 指向了 a，它是一个指针变量。

当获取或设置指针指向的内存的值时，在指针变量前面加 *，然后赋值，如：*c = 4，指针指向的变量 a 将会变化。

如果将指针变量赋予另外一个变量：c1 := c，那另外一个变量 c1 可以叫做引用变量，它存的值也是内存地址，内存地址指向的也是变量 a，这时候，引用变量只是指针变量的拷贝，两个变量是互相独立的。

值变量可以称为值类型，引用变量和指针变量都可以叫做引用类型。

如何声明一个引用类型的变量（也就是指针变量）呢？

我们可以在数据类型前面加一个 * 来表示：

    var d *int
Copy to clipboardErrorCopied
我们以后只会以值类型，和引用类型来区分变量。



下面来看一个下例子：

```go
package main
import "fmt"
//go指针
func pointer() {
   var a int = 5
   var pa* int
   pa = &a
   *pa = 1
   fmt.Println(a)
   }
func main(){
     pointer()
  }

//输出：
//1
```



#### Golang 函数的传参方式

* 值传递
  使用普通变量作为函数参数的时候，在传递参数时只是对变量值得拷贝，即将实参的值复制给变参，当函数对变参进行处理时，并不会影响原来实参的值。

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

//输出结果：
//5 10
```

传递给 swap 的是 x，y 的值得拷贝，函数对拷贝的值做了交换，但却没有改变 x,y 的值。



* 指针的值传递

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

//输出结果：
//10 5
```



