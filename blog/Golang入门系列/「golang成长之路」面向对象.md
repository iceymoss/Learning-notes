[toc]



### 文章介绍

本文将以golang的”面向对象“为核心展开介绍结构体、面向对象思想、继承、封装、多态(基于接口)和接收者问题。

### 概述

go的面向对象和其它的面向语言并不太相同。go是没有类的概念的，只需要使用```struct```就可以完成类的工作，在go中数据是数据，方法是方法，二者比较独立。其次go对继承的实现更像是“组合”。go的接口是鸭子类型：只要你像个鸭子，那么我就认为你是个鸭子，而不需要显式的实现接口。
面向对象的三个特征：继承、封装、多态 go都可以实现。（多态需要基于接口实现)。

### 结构体

在上面说过：“Go 语言中的 struct 和 其他编程语言中的 class 具有同等地位”
对结构体的定义使用如下关键字：

```go
type 标识符 struct{
       field1 type
       field2 type
 }
```

例如:

```go
// 构建一个二叉树的结构体
type Node struct{
     Value int
     Left, right *Node
}
```

那么怎么来创建 Node 呢？即实例化：

```go
package main

import "fmt"
// 构建一个二叉树的结构体
type Node struct{
     Value int
     Left, Right *Node
}

func mian(){
    var root Node //创建一个根节点
    //root ：= Node{} 
    root = Node{Value: 3}
    //root.left = &Node{}  没有给值时，则默认值为0
    //left和right在结构体中都是指针需要加&
    root.Left = &Node{0, nil, nil}
    root.Right = &Node{5, nil, nil}
    root.Left.Right = &Node{2, nil, nil}
    root.Rigth.Left = new(Node) //先new再赋值
    root.Right.Left = &Node{0, nil, nil}

}
```



### Go的面向对象思想

#### Go 面向对象的特性

* Go 语言不是纯粹的面向对象的语言，准确是描述是，Go 语言支持面向对象编程的特性。

* Go 语言中没有传统的面向对象编程语言的 class , 而 Go 语言中的 struct 和 其他编程语言中的。

* class 具有同等地位，也就是说 Go 语言是 基于 struct 来实现 面向对象 的特性的。

* 面向对象编程在 Go 语言中的支持设计得很具有特点，Go 语言放弃了很多面向对象编程的。

  念，如：重载，构造函数，析构函数，隐藏 this 指针等，但是Go 语言可以实现面向对象编程的特性 封装，继承，多态。

* Go 语言仅支持封装，不支持继承和多态 (在 go 语言中使用接口实现)。

* Go 语言面向对象编程的支持是语言类型系统 中天然的组成部分，整个类型系统通过接口串联，灵活度高。

  

#### 属性和方法

1. 属性：例如人的身高，年龄，体重等，这些是他的属性。用go来描述：

   ```go
   type Person struct {
     Length float32
     Age int
     Weight float32
   }
   ```

2. 方法：例如人的吃饭、睡觉、工作等，这些都是他的行为。用go来描述：

   ```go
   type Person struct {
     length float32
     age int
     weight float32
   }
   
   func (P person) Eat(){
     ……
   }
   
   func (P person) Sleep(){
     ……
   }
   
   func (P person) Work(){
     ……
   }
   ```

   当然虽然go没有构造函数，但是是可以实现类似构造函数的函数的：

   ```go
   package main
   
   
   func NewPerSon() Person {
     return Person{}
   }
   
   
   func main(){
     
     //实例化Person
     p := NewPerSon()
     
     //使用属性
     p.Length = 185.4
     P.Age = 22
     P.Weight = 140.55
     
     //调用方法
     p.Eat()
     p.Sleep()
     p.Work()
   }
   ```

   同时go有自己的垃圾回收机制(GC)，所以也不需要析构函数。



### 继承

指子类获得父类的属性和方法， 在go中就是使用结构体嵌套结构体

```go
//父类
type A struct {
  a int
  b int 
}

//子类
type B struct {
  c A  //嵌套A
  d int 
}
```

这个就是go中的继承，即：

```go
testB := B{}
B.a
B.b
B.c
```

如果一个包是别人写的，那么我们想要扩展它怎么办？
在 c++ 及 Java 里面我们会使用继承来扩展，相对较麻烦，所以在 Go 中干脆就取消了继承，在 Go 中扩展系统类型及别人的类型有两种方法：

* 使用组合

* 定义别名

* 使用内嵌

  

#### 使用组合：
tree 包

```go
package tree

import "fmt"
//结构体
type Node struct{
   Value int
   Right, Left *Node
}

//工厂函数(在golang中可以把取地址局部变量给全局变量使用）
func GreatNode(value int) *Node{
   return &Node{Value:value}
}

func (node *Node)Setvalue(value int){
    node.Value = value
}

//遍历树
func (node *Node) Traveser(){
   if node == nil {
      return
  }
   node.Left.Traveser()
   fmt.Println(node.Value)
   node.RightTraveser()
}
```

我们在别人写 tree 包中使用中序遍历，那么我们现在还需要使用后续遍历，而现在 tree 包中没有后序遍历，此时就需要我们自己来实现。

现在我们在 main 中写一个结构体：

```go
type mytreeNode struct{
   node *tree.Node  //node为tree中的Node类型
}
```

实现：

```go
package main

import (
   "awesomeProject/tree"
 "fmt")

type mytreeNode struct{
   node *tree.Node
}

//有接收者的后序遍历
func (myNode *mytreeNode)postOrder() {
   if myNode == nil || myNode.node == nil {
      return
  }
//使用组合，我们构建了mytreeNode的结构体，必须使用mytreeNode{}
   left := mytreeNode{myNode.node.Left}
   right := mytreeNode{myNode.node.Right}

   left.postOrder()
   right.postOrder()
   fmt.Print(myNode.node)
}

//普通后序遍历
func postOrder1(myNode *mytreeNode){
   if myNode == nil || myNode.node == nil{
      return
  }
   left := mytreeNode{myNode.node.Left}
   right := mytreeNode{myNode.node.Right}

   postOrder1(&left)
   postOrder1(&right)
   fmt.Print(myNode.node)

}

func main() {
   var root tree.Node
   root = tree.Node{Value: 0}
   root.Left = &tree.Node{1, nil, nil}
   root.Right = &tree.Node{2, nil, nil}
   root.Left.Left = &tree.Node{3, nil, nil}
   root.Left.Right = &tree.Node{4, nil, nil}
   root.Right.Left = new(tree.Node)
   root.Right.Left = &tree.Node{5, nil, nil}

   fmt.Println()
   //接收者函数
   myRoot := mytreeNode{&root}
   myRoot.postOrder()

   //普通函数
   postOrder1(&mytreeNode{&root})
   //postOrder1(myRoot)
}

//输出：
//341520
//341520
```

#### 使用别名
例如：
现在我们另一个目录下使用别名的方式来编写一个int类型队列 (Queue)。
同样，我们将队列写在包里

```go
//定义别名
type Queue []int
```

```go
package queue

type Queue []int

//加入元素
func (q *Queue) Push(v int){
   *q = append(*q, v)
}

//移出首元素
func (q *Queue) Pop() int {
   head := (*q)[0]
   *q = (*q)[1:]
   return head
}

//判断队列是否为空
func (q *Queue)IsImpty() bool{
   return len((*q)) == 0
}
```

在 main 中调用包：

```go
package main

import (
   "awesomeProject/tree/queue"
 "fmt")

func main() {
   q := queue.Queue{1}
   q.Push(2)
   q.Push(3)
   q.Pop()
   fmt.Println(q)
   fmt.Println(q.IsImpty())
   fmt.Println(q.Pop())
   fmt.Println(q.Pop())
   fmt.Println(q.IsImpty())
}
```

输出：

```go
[2 3]
false
2
3
true
```



#### 内嵌 (emdedding)
内嵌其实也就是语法糖，它能使我们的代码量减少，什么是内嵌呢？看下面：

```go
type mytreeNode struct{
   node *tree.Node
}
```

```go
type mytreeNode struct{
   *tree.Node   //Embedding 内嵌
}
```

很容易看出：内嵌把 node 给省略了
使用内嵌后：

```go
//后序遍历
func (myNode *mytreeNode)postOrder() {
   if myNode == nil || myNode.node == nil {
      return
  }

   left := mytreeNode{myNode.node.Left}
   right := mytreeNode{myNode.node.Right}

   left.postOrder()
   right.postOrder()
   fmt.Print(myNode.node.Value)
}
```

>这里简单总结：
>
>* 定义别名：最简单
>* 使用组合：最常用
>* 使用内嵌：需要省下许多代码



### 封装



#### 封装

- 名字一般使用 CamelCase
- 首字母大写：public
- 首字母小写：private



#### 包
包就是每一个程序中的 `package`

- 每一个目录只有一个包
- main 包包含了可执行入口
- 为机构定义的方法必须放在同一个包内
- 可以是不同的文件

下面是一个包：
对于这个包而言它我们要做可执行程序 main 中使用，就应该将构造体及其方法的首字母写成大写

```go
package tree

import "fmt"
//结构体
type Node struct{
   Value int
   Right, Left *Node
}

//工厂函数(在golang中可以把取地址局部变量给全局变量使用）
func GreatNode(value int) *Node{
   return &Node{Value:value}
}

func (node *Node)Setvalue(value int){
    node.Value = value
}

//遍历树
func (node *Node) Traveser(){
   if node == nil {
      return
  }
   node.Left.Traveser()
   fmt.Println(node.Value)
   node.RightTraveser()
}
```

在可执行的 main 中：
我们引入的包名叫：tree
所以需要在构造方法中都添加 tree

```go
package mian
import ("awesomeProject/tree"
        "fmt"
    )

func  mian(){
    var root tree.Node //创建一个根节点  //root ：= tree.Node{} 
    root = tree.Node{Value:  3}  
    //root.left = &Node{} 没有给值时，则默认值为0 
    //left和right在结构体中都是指针需要加&
    root.Left =  &tree.Node{0,  nil,  nil}
    root.Right =  &tree.Node{5,  nil,  nil}
    root.Left.Right =  &tree.Node{2,  nil,  nil} 
    root.Rigth.Left =  new(tree.Node)  //先new再赋值
    root.Right.Left =  &tree.Node{0,  nil,  nil}

    //调用tree中中序遍历方法
    root.Traveser()
    root.Right.Right = new(tree.Node)

    //调用tree中Setvalue()方法
    root.Right.Right.Stevalue(100)
    fmt.Println(root.Right.Right.Value)
```





在其他编程语言中，有点变量和方法是不想让包外访问的，所以需要对属性和方法进行封装，下面看一下go是如何封装的，在go中属性和方法名首字母以大写开头是可以被包外访问的，小写则不能。

```go
type Person struct {
  Length float32
  Age int
  weight float32  //不希望别人知道我的体重用小写
}

//不想然别人知道我的工作用小写
func (P person) work(){
  ……
}
```





### 多态 

go的多态是基于接口```interface```来实现的，定义一个接口，里面有吃饭、睡觉、行走等

```go
type Tester interface {
  Eat()
  Sleep()
  Wake()
}
```

只要某个结构体实现了这个接口里面的方法，我们就说这个结构体实现了这个接口，例如：我们使用猫(cat)来实现了这个接口的全部方法，那么就可以说猫实现了这个接口，那么这个接口也就是猫，如果是狗(dog)实现了这个接口，那么这个接口就是go；这就是多态。



### 值接收者和指针接收者

值接收者 VS 指针接收者:

* 改变内容必须用指针接收者

* 结构过大也可以考虑指针接收者

* 一致性：如有指针接收者，最好使用指针接收者

  

看下面两个函数：

```go
//值接收者
func (node Node) setvalue(value int){
    node.Value = value
}
```

```go
//指针接收者
func (node *Node) setvalue(value int){
    node.Value = value
}
```

#### go只有值传递

第一个函数是值传递不会真正的修改 node.Value 的值,只是将实参的副本给到形参。
第二个函数是指针传参会真正的改变 node.Value 的值，这里也是副本，但这个副本是真正的地址，也就会真正的修改相关位置的值。

 #### nil 指针也可以调用方法 

```go
func  mian(){
    var root Node //创建一个根节点  //root ：= Node{} 
    root = Node{Value:  3}  
    //root.left = &Node{} 没有给值时，则默认值为0 
    //left和right在结构体中都是指针需要加&
    root.Left =  &Node{0,  nil,  nil}
    root.Right =  &Node{5,  nil,  nil}
    root.Left.Right =  &Node{2,  nil,  nil} 
    root.Rigth.Left =  new(Node)  //先new再赋值
    root.Right.Left =  &Node{0,  nil,  nil}

    root.traverse()
}

//输出:
//0 2 3 0 5
```











