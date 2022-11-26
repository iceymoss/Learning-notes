[toc]

### 文章介绍

本系列文章将快速带您入门go语言，本文将介绍go的内置容器(数据结构)，将介绍数组、slice和map的简单使用。



### 数组(array)

数组：数组是指有序的元素序列。如果将有限个类型相同的变量的集合命名，那么这个名称就是数组名，而组成数组的各个变量称为数组的分量，也称为数组的元素，有时也称为下标变量，而数组中的数据可以使用下标 (索引）来查找到。

其实在编程语言中数组的概念是一样的，下面具体来看看 golang 中数组是如何定义和使用的：

#### 数组的定义

* 方法一：

  ```go
  var arr1 [5]int   //定义一个拥有5个元素的数组,此时元素全为0
  ```

* 方法二：

  ```go
  arr2 := [5]int{0, 1, 2, 3, 4}     //定义并放入值
  ```

* 方法三：

  ```go
  arr3 := [...]int{0, 1, 2, 3, 4} //用[...],此时不需写入数组空间具体是多少
  ```



##### 二维数组的定义及初始化

```go
package main
import "mft"
func main {
//二维数组
var arr4 [4][5]int
fmt.Println(arr4)
}

//输出：[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
```



#### 数组用法

##### 遍历

```go
package main
import "fmt"
arr :=[...]int{0, 1, 2, 3, 4}
for i := range arr{
  fmt.Println(arr[i])
}
```

输出：

```
0
1
2
3
4
```

也可以这样遍历：

```go
//.数组的遍历
for i,v :=  range arr3{   //i为对应元素的索引值，v为元素值
     fmt.Println(i, v)
}
```

```go
0 0
1 1
2 2
3 3
4 4
```

##### 更改值

```go
arr :=[...]int{0, 1, 2, 3, 4}
arr[0] = 100  //将arr[0]改为100
fmt.Println(arr)
arr[0] = 0  //将arr[0]改回0
fmt.Println(arr)
```

输出：

```
[100 1 2 3 4 ]
[0 1 2 3 4 ]
```



#### 数组的值传递
数组的值传递仍然是将整个数组 copy 一份传入函数，不会改变 arr 的值

```go
//数组的值传递
package main
import "fmt"
//定义函数：
func printarr1(arr [5]int){
  arr[0] = 100
  fmt.Println(arr)
}
func main(){
  printarr1(arr)
  fmt.Println(arr)
}

//out: [1 2 3 4]
```



#### 使用指针
使用指针是将数组的相应值的地址拷贝传入函数

```go
//使用指针
package main
import "fmt"
//定义函数：
func printarr2(arr *[5]int) {
  arr[0] = 100
  fmt.Println(arr)
}
func main(){
  printarr2(&arr3)
  fmt.Println(arr3)
}

//out:[100 1 2 3 4]
```





### Slice(切片:不定长数组)

#### 原理

**结构**：

```go
type slice struct { 
    array unsafe.Pointer // 指向底层数组 
    len int // 长度 
    cap int // 容量 
}
```

这里可以slice本质是一个复合类型 其指向底层数组， 有长度和容量，所以对slice的操作实质是对底层数组的操作，即改变slice中的元素，会导致底层数组改变。



##### 扩容

当容量小于1024时，slice的容量使用完后，会直接将容量翻倍:例如：a的cap = 5， len = 5, 此时a的容量使用完了，然后我们再使用append()向a中追加一个数据，这时a就会扩容 cap = cap * 2 = 10， len = 6; 如果a的cap >= 1024时，len = 1024,继续追加数据，此时cap扩容1.25倍。



##### 扩容带来的性能消耗

当slice容量用光后，继续追加数据，go就会开辟一个新的数组(容量为原来的2倍或者1.25倍)，然后将原来的底层数组数据依次放入新开辟的底层数组，并将当前slice指向新的数组，如果老数组没有被其他引用，go的内存管理机制就会将老数组内存回收。可以看出这个扩容的过程我们将老底层数组的数据重新放到新底层数组上，这个过程是需要消耗性能的。并且要知道扩容后的slice就会原来的底层数组没有任何关系了。



##### 传参方式

这里介绍一下slice的函数中的传参方式

* 不改变len和cap字段： 值传递， slice指向底层数组更改slice内容，会改变底层数组

* 改变len和cap字段： 使用指针

  

#### 操作slice

```go
package main

import "fmt"

func main(){
//切片
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s := arr[2:6]
fmt.Println("s = ",s)
fmt.Println("arr[2:6 =",arr[2:6])
fmt.Println("arr[:6] = ",arr[:6])
fmt.Println("arr[2:] = ",arr[2:])
fmt.Println("arr[:] = ",arr[:])
}
```

输出：

```go
s =  [2 3 4 5]
arr[2:6 = [2 3 4 5]
arr[:6] =  [0 1 2 3 4 5]
arr[2:] =  [2 3 4 5 6 7]
arr[:] =  [0 1 2 3 4 5 6 7]
```



##### Slice 值的改变
改变 Slice 中元素的值，会导致底层 arr 值的改变

```go
package main

import "fmt"

//定义函数，改变Slice的值
//切片
func updateSlice(s []int){
   s[0] = 100
}
//主函数
func main(){
  arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
  s1 :=arr[2:]
  s2 := arr[:]
  //改变之前
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println("change after:")
  //函数调用
  updateSlice(s1)
  fmt.Println(s1)
  //最后底层的arr的值也会随之改变
  fmt.Println(arr)
}
```

输出：

```go
[2 3 4 5 6 7]
[0 1 2 3 4 5 6 7]
change after:
[100 3 4 5 6 7]
[0 1 100 3 4 5 6 7]
```



当改变s2：

```go
fmt.Println(s2)
updateSlice(s2)
fmt.Println(s2)
fmt.Print(arr)
```

输出：

```go
[0 1 100 3 4 5 6 7]
[100 1 100 3 4 5 6 7]
[100 1 100 3 4 5 6 7]
```



这里还有一种情况：Slice 的更新

```go
package main

import "fmt"

func main(){

  arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
  s2 := arr[:]
  //Slice中值的更新
  //对Slice的操作，其实是对底层arr的修改

  fmt.Println(s2)
  s2 = s2[:5]
  fmt.Println(s2)
  s2 = s2[2:]
  fmt.Println(s2)
  fmt.Println(arr)
}
```

输出：

```go
[0 1 2 3 4 5 6 7]
[0 1 2 3 4]
[2 3 4]
[0 1 2 3 4 5 6 7]
```



##### slice的拓展

```go
package main
import "fmt"
func main(){
  //Slice的扩展
  arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
  fmt.Println("Extneding Slice:")
  fmt.Println(arr)
  s1 = arr[2:6]
  s2 = s1[3:5]
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println(arr)
}
```

现在肯定对这段代码有疑问吧！
`s2 = s1[3:5]` 会不会报错？
答案：不会！
下面来看看运行结果：

```go
[0 1 2 3 4 5 6 7]
[2 3 4 5]
[5 6]
[0 1 2 3 4 5 6 7]
```

**原因：在 golang 中 Slice 是可是往后扩展的，例如上面**

**arr := […]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 = arr[2:6] = [2 3 4 5]
s2 = s1 [3:5]，此时已经超出 s1 的范围了，但是 Slice 是对底层 arr 的操作，并且可以往后扩展的，所以往底层走就应该是 [5 6]，s2 = s1 [3:5] = [5 6] **



##### 向 slice 添加元素
向 slice 添加元素使用 `append(s, x)`

```go
//向slice添加元素
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//切分
s1 :=arr[2:]
s2 := arr[:]
//打印
fmt.Println(s1)
fmt.Println(s2)
//添加元素
s3 := append(s2, 10)
s4 := append(s3, 11)
s5 := append(s4, 12)
fmt.Println("s3, s4, s5 =",s3, s4, s5)
fmt.Println("arr = ",arr)
```

输出：

```go
[2 3 4 5 6 7]
[0 1 2 3 4 5 6 7]
s3, s4, s5 = [0 1 2 3 4 5 6 7 10] [0 1 2 3 4 5 6 7 10 11] [0 1 2 3 4 5 6 7 10 11 12]
arr =  [0 1 2 3 4 5 6 7]  //arr的容量，使用光了，发生扩容
```



##### 验证扩容

```go
package mian

import "fmt"
//定义函数
func printslice(s []int){
 fmt.Printf(" s=%v, len=%d, cap=%d",s, len(s), cap(s))
}
func main(){
  var s []int
  //s的初始化,生成100个奇数
  for i:=0; i < 20; i++{
     printslice(s)
     fmt.Print("\n")
     s = append(s, 2*i + 1)
  }
}
```

输出：

```go
 s=[], len=0, cap=0  //var s []int 不会初始化 当第一次遇到append时完成初始化
 s=[1], len=1, cap=1
 s=[1 3], len=2, cap=2    //发生扩容 容量翻倍
 s=[1 3 5], len=3, cap=4
 s=[1 3 5 7], len=4, cap=4
 s=[1 3 5 7 9], len=5, cap=8     //发生扩容 容量翻倍
 s=[1 3 5 7 9 11], len=6, cap=8
 s=[1 3 5 7 9 11 13], len=7, cap=8
 s=[1 3 5 7 9 11 13 15], len=8, cap=8
 s=[1 3 5 7 9 11 13 15 17], len=9, cap=16    //发生扩容 容量翻倍
 s=[1 3 5 7 9 11 13 15 17 19], len=10, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21], len=11, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21 23], len=12, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25], len=13, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25 27], len=14, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29], len=15, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31], len=16, cap=16
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33], len=17, cap=32   //发生扩容 容量翻倍
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35], len=18, cap=32
 s=[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37], len=19, cap=32
```



查看长度和容量：

```go
package main

import "fmt"

func printslice(s []int){
 	fmt.Printf(" s=%v, len=%d, cap=%d",s, len(s), cap(s))
}
func main(){
  //graet slice
  s1 := []int{0,1,2}
  s2 := make([]int, 16)
  s3 := make([]int, 10, 32)
  //调用函数
  printslice(s1)
  printslice(s2)
  printslice(s3)
}
```



```GO
 s=[0 1 2], len=3, cap=3 s=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16 s=[0 0 0 0 0 0 0 0 0 0], len=10, cap=32

```





##### capy slice

```go
//copy slice
copy(s2, s1)   //将s1的内容copy到s2上
fmt.Println(s2)
```



##### delete Slice

删除操作

```go
package main
import "fmt"
定义函数
func printslice(s []int){
fmt.Printf(" s=%v, len=%d, cap=%d",s, len(s), cap(s))
}
//主函数
func main(){
  //delete "2"
  s5 := []int{0, 1, 2, 3, 4, 5, 6, 7}
  s5 = append(s5[:2], s5[3:]...)  //使用切片(Slice)
  printslice(s5)
  //delete front
  s5 = s5[1:]
  printslice(s5)
  //delete tail
  s5 = s5[:len(s5)-1]
  printslice(s5)
}
```

输出：

```go
s=[0 1 3 4 5 6 7], len=7, cap=8
s=[1 3 4 5 6 7], len=6, cap=7
s=[1 3 4 5 6], len=5, cap=7
```



### map(键值对)

map：是一个无序，key—value 对，不会出现某个 key 对应不同的 value

##### map 的定义

```go
//定义map
m :=  map[string]string{      //[string]对应key，string对应value
 "name":"yangkuang",
 "work":"students",
 "year": "20",
}
m3 := map[string]int{
  "a": 0,
	"b": 1,
	"c": 2,

//空map的定义
m1 := make(map[string]int)   // empty map1
var m2 map[string]int
fmt.Println(m, m1, m2)    //map2 == nil
```

输出：

```go
map[name:yangkuang work:students year:20] map[] map[]
```



##### 遍历

```go
   for k, v := range m3{
       fmt.Println(k, v)
   }


//输出：
// a 0
// b 1
// c 2
```



##### 值查找

```go
func main(){
//定义map
m :=  map[string]string{
 "name":"yangkuang",
 "work":"students",
 "year": "20",
}
iname, ok := m["name"]   //key—name存在则将key对应的value赋值给iname，并且将ok赋值为true，否则ok为false
fmt.Println(iname,ok)   //输出：yangkuang true
}
```



##### 移除key

内置方法：delete

```go
delete(intMap, 1)
```



##### 清空map

```go
for k, _ := range m {
	delete(m, k)
}

```



##### 释放内存

```go
map = nil
```

这之后坐等垃圾回收器回收就好了。

如果你用 map 做缓存，而每次更新只是部分更新，更新的 key 如果偏差比较大，有可能会有内存逐渐增长而不释放的问题。要注意。

