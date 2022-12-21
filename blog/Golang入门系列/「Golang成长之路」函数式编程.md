[toc]

### 介绍

#### 特点：

1. 在go语言中，函数是一等公民，可以作为参数，变量，返回值来使用
2. 高阶函数(在函数中定义函数，传入函数，返回函数)
3. 闭包

附加：正统式函数式编程
1.不可变性：不能有状态，只能有常量和函数
2.函数只有一个参数
对于正统式函数式编程(较为复杂，阅读性不高）我们不做太多介绍

好下面我们看一个例子：
```go
 func adder() func(v int) int{
   sum := 0
   return func(v int) int{
      sum += v
	  return sum
    }
 }
```
 这里我们写了一个求和的函数(这里将匿名函数func(v int)作为返回值，并将匿名函数func的定义直接写在函数adder的return后面)。

接着看下面：
我们在main函数中测试一下
```go
package main

import "fmt"

func adder() func(v int) int{
   sum := 0
   return func(v int) int{
      sum += v
      return sum
    }
 }

func main(){
	a := adder()
	fmt.Println(a(1))
	fmt.Println(a(2))
	fmt.Println(a(9))
	}
 //输出结果为：1、3、12

```
看到这里是不是有点模糊了，没事我来解释一遍：
首先 a := adder 是将adder()的返回值给到a
再仔细看adder的返回值是一个匿名函数，所以a就相当于函数
func(v int) int{}的实例，a(1) 、 a(2) 、a(3)自然就能理解了，就是对应v

是不是不明白输出结果，下面就来介绍一下闭包



#### 闭包：

定义：Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量，因此，简单的说：函数 + 引用环境 = 闭包
我的理解:就是指函数在第一次调用时的返回值，返回后，其返回值不会立即被删除(储存在某个地方)，当下一次调用该函数时会将上一次的返回值作为这次调用的参数；例如上面的求和函数：
```go
func adder() func(v int) int{
   sum := 0
   return func(v int) int{
      sum += v
      return sum
    }
 }
 
func main(){
 a := adder()
 fmt.Println(a(1))  //第一次调用时sum = 0 ， sum = 0 + 1，返回值为1
 fmt.Println(a(2)) //第二次调用时sum = 1 ，sum = 1 +2，返回值为3
 fmt.Println(a(9))  //第三次调用时sum = 9，sum = 3 + 9，返回值为12
 }
 
```
这是不是看懂了闭包
```go
func adder() func(v int) int{
   sum := 0
  return func(v int) int{
      sum += v
      return sum
   }

func main() {
   a := adder()
   for i := 0; i <= 5; i++{
      fmt.Println(a1(i))
   }
   //输出结果为：0、1、3、6、10、15
```



### 函数式编程的应用

#### 斐波那契数列

数列：0、1、1、2、3、5、8、13、21……
前两个数之和是第三个数，0+1 =1、1+1=2、1+2=3、2+3=5、3+5=8、5+8=12……
这里我们需要使用函数式编程的模式将其实现

```go
package main

func fibonacci() func() int{
  a :=0, b := 1
	return func() int{
	a , b = b, a+b
	return a
	}
}

func main(){
	f := fibonacci()
	
	fmt.Println(f())  //1
	fmt.Println(f())  //1
	fmt.Println(f())  //2
	fmt.Println(f())  //3
	fmt.Println(f())  //5
	fmt.Println(f())  //8
}
```
斐波那契数列生成器：
这里我们需要像读取文件一样，就需要实现read的接口
我们将打印函数拿过来：

```go
func printContentFile(reader io.Reader) {
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {    //
        println(scanner.Text())
    }
}
```
x想要直接调用printContentFile()自动生成斐波那契数，就必须实现read接口，定义一个type：函数。使其实现Reader接口。

```go
//函数实现接口
//定义一个函数，使用type修饰。可以实现接口，也就是说只要是被type修饰的东西都可以实现接口。

type IntGen func() int
//实现read接口
func (g IntGen) Read(p []byte) (n int, err error) {
    next := g()

    if next > 10000 { //这里因为不设置退出会一直打印下去，所以做了限制
        return 0, io.EOF
    }
    s := fmt.Sprintf("%d\n", next)
    return strings.NewReader(s).Read(p)
}
```
函数返回为fibonacci()
```go
/**
    函数生成斐波那契数列（每个数字都是由前两个数字相加得到）
 */
func fibonacci() IntGen {
    a, b := 0, 1
    return func() int {
        //在这里，生成关键
        // 1 1 2 3 5 8
        //   a b
        //     a b
        a, b = b, a+b
        return a
    }
}
```

最后调用：
```go
func main() {

    fun := fibonacci()

    printContentFile(fun)
}
```

打印结果：
```sh
1
1
2
3
5
8
13
21
34
55
89
144
233
377
610
987
1597
2584
4181
6765
```
从这个例子中可以发现，函数还可以实现接口，使用在go语言中函数是一等公民



#### 函数式编程实现中序遍历 + 多业务

```go
package tree1
import "fmt"
//结构体Node
type Node struct{
     Value int
     Left, Right *Node
}
//打印
func (node *Node)Println(){
     print(node.Value)
}
func (node *Node)Traverse(){    //中序遍历+打印
	 node.TraverseFun( func(*Node){
		node.print()
	 })
 	fmt.Println()
}
```
```go
func (node *Node)TraverseFun(f func(*node){   //中序遍历+……
     if node == nil{
       return
     }
node.Left.TraverseFun(f)
f(node)
node.Right.TraverseFun(f)
}
```

```go
func main() {
   //向树中添加数据
   root := tree1.Node{Value:0}
   root.Left = &tree1.Node{1, nil, nil}
   root.Right = &tree1.Node{2, nil,nil}
   root.Left.Left = &tree1.Node{3, nil,nil}
   root.Left.Right = &tree1.Node{4, nil, nil}
   root.Right.Left = &tree1.Node{5, nil,nil}
   root.Right.Left = &tree1.Node{6, nil, nil}
   //中序遍历
   root.Traverse()

   TraverseCount := 0 //打印数据个数
   root.TraverseFunc(func(node *tree1.Node) {
      TraverseCount++
       })
    fmt.Println(TraverseCount)
}
```
看到是没有问题的。且比之前的`Traverse`更加灵活了,在中序遍历是可以做很多东西了

