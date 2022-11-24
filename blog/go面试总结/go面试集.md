[toc]



### go面试题

#### 1、```=```和```:=```的区别

```=``` 是赋值； ```:=``` 是定义变量



#### 2、```new```和```make```的区别

Go语言中new和make都是用来内存分配的原语。简 单的说，new只分配内存，make用于slice，map，和channel的初始化。

**new** 的使用：

当我们声明一个指针变量：

```go
var v *int

fmt.Println(v) //<nil>

*v = 2 

fmt.Println(*v) 
```

会报错：

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10a3143]
```

这很正常，我们可以看到初始化一个指针变量，其值为nil，nil的值是不能直接赋值的。可以通过 new其返回一个指向新分配的类型为int的指针， 像这样：

```go
var v *int
v = new(int)
*v = 2
fmt.Println(v)  //0xc0000ae008
fmt.Println(*v) //2
```



###### 数组中使用new

```go
var a [5]int
fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc04200a180 [5]int{0, 0, 0, 0, 0}

av := new([5]int)
fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[5]int{0, 0, 0, 0, 0}

(*av)[1] = 8
fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000006028 &[5]int{0, 8, 0, 0, 0}
```



###### 在slice中使用new:

```go
var a *[]int
fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc042004028 (*[]int) (nil)

av := new([]int)
fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[]int(nil)

(*av)[0] = 8

fmt.Printf("av: %p %#v \n", &av, av) //panic: runtime error: index out of range
```



###### 在map中使用new:

```go
	var m map[string]string
	fmt.Printf("m: %p %#v \n", &m, m)	//m: 0xc042068018 map[string]string(nil)
	
	mv := new(map[string]string)
	fmt.Printf("mv: %p %#v \n", &mv, mv)	//mv: 0xc000006028 &map[string]string(nil)
	
	(*mv)["a"] = "a"
	fmt.Printf("mv: %p %#v \n", &mv, mv)	//这里会报错panic: assignment to entry in nil map

```



###### 在channel中使用new:

```go
cv := new(chan string)
fmt.Printf("cv: %p %#v \n", &cv, cv)//cv: 0xc000074018 (*chan string)(0xc000074020)
cv <- "good" //会报 invalid operation: cv <- "good" (send to non-chan type *chan string)

```



>总结：
>
>通过上面示例我们看到数组通过new处理，数组av初始化零值，数组虽然是复合类型，但不是引用类型，其他silce、map、channel类型也属于引用类型，go会给引用类型初始化为nil，nil是不能直接赋值的。并且不能用new分配内存。无法直接赋值。那么用make函数处理会是怎么样呢？



**make**

```go
	arr := make([]int, 5)
	fmt.Printf("arr: %p, v: %v\n", &arr, arr)

	arr[1] = 10
	fmt.Printf("arr: %p, v: %v\n", &arr, arr)

	m := make(map[string]string, 5)
	fmt.Printf("m: %p, v: %v\n", &m, m)

	m["address"] = "北京"
	fmt.Printf("m: %p, v: %v\n", &m, m)

	ch := make(chan string)
	fmt.Printf("ch: %p, v: %v\n", &ch, ch)
	go func(msg string) {
		ch <- msg
	}("hello")

	fmt.Println(<-ch)

	close(ch)
```

输出：

```
arr: 0xc00011a018, v: [0 0 0 0 0]
arr: 0xc00011a018, v: [0 10 0 0 0]
m: 0xc000126020, v: map[]
m: 0xc000126020, v: map[address:北京]
ch: 0xc000126028, v: 0xc000100060
hello
```

make不仅可以开辟一个内存，还能给这个内存的类型初始化其零值。



###### 总结

**make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。 make返回的还是引用类型本身；而new返回的是指向类型的指针。 make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配 任意类型的数据。**



#### 3、请你讲一下Go面向对象是如何实现的？

* go的面向对象是使用```struct```和```interface```两个关键字来实现的

* 封装：对于同一个包，包内对象文件可见；对于不同的包，方法及属性首字母大写对外可见

* 继承：继承使用```struct```嵌套来实现， 例如：

  ```go
  type A struct{}
  
  type B struct{
  		A
  }
  ```

* 多态：多态是运行时特征，Go多态通过```interface```来实现。类型和接口是松耦合的，某 个类型的实例可以赋给它所实现的任意接口类型的变量。
* Go支持多重继承，就是在类型中嵌入所有必要的父类型。



#### 4、如何对二维slice初始化

使用内建方法```make```

要注意的是，对二维切片初始化分配内存后，内部的一维slice是没有分配内存的，因 此要使用二维切片保存数据还需要对一维slice分配内存。 否则，会出现 “panic: runtime error: index out of range [0] with length 0”的错误。

例如:

```go
a := make([][]int, 0, 10)
	for i := 0; i < 10; i++ {
		a[i] = make([]int, 0, 10)
	}
```

```go
a := make([][]int, 0, 10)
	var c []int

	for i := 0; i < 10; i++ {
		c = []int{}
		a = append(a, c)
	}
```





#### 5、uint型变量值分别为 1，2，它们相减的结果是多少？

结果：会溢出，如果是32位系统，结果是2^32-1，如果是64位系统，结果2^64-1



#### 6、go有没有在main之前执行的函数

go中的```init()```在main之前执行，其特点如下：

* 初始化不能采用初始化表达式初始化的变量；

*  程序运行前执行注册 

* 实现sync.Once功能 不能被其它函数调用 

* init函数没有入口参数和返回值：

  ```go
  func init(){
  	register…
  }
  ```

*  每个包可以有多个init函数， 每个源文件也可以有多个init函数 。 

* 同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个 执行顺序。 

* 不同包的init函数按照包导入的依赖关系决定执行顺序。



#### 7、init() 函数是什么时候执行的？

在main函数之前执行。

init()函数是go初始化的一部分，由runtime初始化每个导入的包，初始化不是按照从 上到下的导入顺序，而是按照解析的依赖关系，没有依赖的包最先初始化。

每个包首先初始化包作用域的常量和变量（常量优先于变量），然后执行包的 init() 函数。同一个包，甚至是同一个源文件可以有多个 init() 函数。 init() 函数没有入参和返回值，不能被其他函数调用，同一个包内多个 init() 函数的执行 顺序不作保证。

执行顺序：import –> const –> var –> init() –> main() 一个文件可以有多个 init() 函数！



#### 8、如何知道一个对象是分配在栈上还是堆上？

Go和C++不同，Go局部变量会进行逃逸分析。如果变量离开作用域后没有被引用，则优先分配到栈上，否则分配到堆上。那么如何判断是否发生了逃逸呢？

```go
go build -gcflags '-m -m -l' xxx.go
```

关于逃逸的可能情况：变量大小不确定，变量类型不确定，变量分配的内存超过用户 栈最大值，暴露给了外部指针。



#### 9、下面这句代码是什么作用，为什么要定义一个空值？

```go
var _ Codec = (*GobCodec)(nil)
type GobCodec struct{
 conn io.ReadWriteCloser
 buf *bufio.Writer
 dec *gob.Decoder
 enc *gob.Encoder
}

type Codec interface {
 io.Closer
 ReadHeader(*Header) error
 ReadBody(interface{})  error
 Write(*Header, interface{}) error
}
```

将nil转换为**GobCodec类型，然后再转换为Codec接口，如果转换失败，说明** GobCodec没有实现Codec接口的所有方法。



#### 10、什么是 rune 类型

ASCII 码只需要 7 bit 就可以完整地表示，但只能表示英文字母在内的128个字符，为 了表示世界上大部分的文字系统，发明了 Unicode， 它是ASCII的超集，包含世界上 书写系统中存在的所有字符，并为每个代码分配一个标准编号（称为Unicode CodePoint），在 Go 语言中称之为 rune，是 int32 类型的别名。 Go 语言中，字符串的底层表示是 byte (8 bit) 序列，而非 rune (32 bit) 序列。

列如:

```go
	sample := "我爱GO"
	runeSamp := []rune(sample)
	fmt.Println(runeSamp)       //[25105 29233 71 79]

	runeSamp[0] = '你'
	fmt.Println(string(runeSamp)) // "你爱GO"
	fmt.Println(len(runeSamp))    // 4
```



#### 11、如何判断 map 中是否包含某个 key ？

```go
	m := make(map[string]int)
	m["age"] = 18
	if v, ok := m["age"]; !ok {
		fmt.Println("key 不存在")
	} else {
		fmt.Println(v)
	}
```



#### 12、什么是协程（Goroutine）

**协程是用户态轻量级线程，它是线程调度的基本单位**。通常在函数前加上go关键 字就能实现并发。一个Goroutine会以一个很小的栈启动2KB或4KB，当遇到栈空间不足时，栈会自动伸缩， 因此可以轻易实现成千上万个goroutine同时启动。









 







