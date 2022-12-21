[toc]



### 文章介绍

在该系列文章中我们学习了go的基础语法、数据结构、面向”对象“编程和并发编程，现在我们来系统的学习go语言的标准库的学习，在本文中将介绍到权限，目录，文件的打开和关闭，文件的读\写、进程相关和环境相关

### 权限

权限**perm**，在创建文件时才需要指定，不需要创建新文件时可以将其设定为０

| 权限项   | 文件类型          | 读         | 写         | 执行       | 读             | 写             | 执行           | 读       | 写       | 执行     |
| -------- | :---------------- | ---------- | ---------- | ---------- | -------------- | -------------- | -------------- | -------- | -------- | -------- |
| 字符表示 | （d\|l\|c\|s\|p） | r          | w          | x          | r              | w              | x              | r        | w        | x        |
| 数字表示 |                   | 4          | 2          | 1          | 4              | 2              | 1              | 4        | 2        | 1        |
| 权限分配 |                   | 文件所有者 | 文件所有者 | 文件所有者 | 文件所属组用户 | 文件所属组用户 | 文件所属组用户 | 其他用户 | 其他用户 | 其他用户 |

go语言在syscall包中定义了很多关于文件操作的权限的常量例如(部分)；

```go
const (
	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
```



### 目录

##### os.Create方法创建文件

```go
func Create(name string) (*File, error)
```

实例：

```go
//CreateFile 创建文件
func CreateFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("file:%v\n", file)
	}
}

```

```go
func main() {
	name := "/Users/feng/go/src/StudyGin/OSlearn/hello_test.txt"
	CreateFile(name)
}
```



##### 创建目录

创建的单个目录：

```go
func Mkdir(name string, perm FileMode) error
```

实例：

```go
//CreateDir 创建单个目录
func CreateDir(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
```

创建多级目录:

```go
func MkdirAll(path string, perm FileMode) error
```

实例：

```go
//CreateDirAll 创建多级目录
func CreateDirAll(name string) {
	err := os.MkdirAll(name, os.ModePerm)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
```



##### 删除目录

删除单个空目录或文件:

```go
func Remove(name string) error
```

实例：

```go
//RemoveDir 删除单个空目录或文件
func RemoveDir(name string) {
	err := os.Remove(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
```

强制删除当前目录下所有目录：

```go
func RemoveAll(path string) error
```

实例：

```go
//RemoveDirAll 强制删除所有目录
func RemoveDirAll(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
```



##### os.Getwd()获取当前目录

```go
func Getwd() (dir string, err error)
```

实例：

```go
//GetWd 获取当前目录
func GetWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("dir:%v\n", dir)
	}
}
```



##### os.Chdir()修改当前工作目录

```go
func Chdir(dir string) error
```

实例：

```go
//ChWd 修改当前工作目录
func ChWd() {
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}
```





##### os.TempDir()获取临时文件

```go
func TempDir() string
```

实例：

```go
//TemDir 获取临时文件目录
func TemDir() {
	fmt.Println(os.TempDir())
}
```



##### os.Rename()修改文件名

```go
func Rename(oldpath string, newpath string) error
```

实例：

```go
//RenameFile 修改文件名
func RenameFile(oldpath, Newpath string) {
	err := os.Rename(oldpath, Newpath)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
```



### 文件读操作

对于文件的读/写操作我们，我们需要拥有相关权限，才能对文件进行读/写

这里我们定义一些关于权限的常量：

```go
const (
	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
```

##### 打开文件

默认打开方式：

```go
func Open(name string) (*File, error)
```

```os.Open()```返回一个File对象的指针

实例：

```go
//Open 只读文件，不能写
func Open(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("file:%s\n", file)
	}
}
```



以指定权限打开

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

实例：

```go
//OpenFile 以指定权限打开
func OpenFile(name string, perm os.FileMode) {
	file, err := os.OpenFile(name, O_RDONLY, perm)  //只读模式打开文件
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("file:%v\n", file.Name())
	}
}
```



##### file.Close()关闭文件

```go
func (f *File) Close() error
```

在对文件操作结束后我们需要文件关闭，所以以经常和```defer```一起使用

实例：

```go
func ReadFile(name){
  file, err := os.Open(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
  
  defer file.Close()
  
  //将文件读取到缓冲区
	buf := make([]byte, 100)
	n, err := file.Read(buf)
}
```





##### file.Read()读取文件

```go
func (f *File) Read(b []byte) (n int, err error)
```

file.Read中需要一个byte类型的切片做缓冲区，并将文件数据读入这个缓冲区，然后返回读取到的数据长度和error

实例：

```go
//ReadFile 读取文件
func ReadFile(name string) {
  //打开文件
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	for {
		//将文件读取到缓冲区
		buf := make([]byte, 5)  //设置一个缓冲区,一次读取5个字节
		n, err := file.Read(buf)
		fmt.Printf("buf:%v\n", string(buf))
		fmt.Printf("数字:%d\n", n)
		if err == io.EOF { //表示文件读取完毕
			break
		}
	}
	file.Close()
}
```



##### file.ReadAt从指定位置开始读取

```go
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
```

实例：

```go
//FileReadAt 从指定位置开始读取
func FileReadAt(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}

	buf := make([]byte, 10)       // 设置一个缓冲区，一次读10个字节
	n, err := file.ReadAt(buf, 5) //从第五个开始读取
	fmt.Printf("buf:%v\n", string(buf))
	fmt.Printf("n:%s\n", n)

	file.Close()
}
```



##### file.ReadDir() 读取目录

```go
func (f *File) ReadDir(n int) ([]DirEntry, error)
```

返回当前目录下的所有目录及文件放入```[]DirEntry```中

实例：

```go
//ReadDir 获取目录
func ReadDir(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	dir, err := file.ReadDir(-1)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	for key, value := range dir {
		fmt.Printf("dir:  key:%v, value: %v\n", key, value)
	}
}
```



##### 设置偏移量

```go
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
```

实例：

```go
func Seek(name string) {
	file, err := os.Open(name)   //打开文件光标默认为文件开头
	if err != nil {
		fmt.Println("err:", err)
	}
	buf := make([]byte, 5)
	file.Seek(3, 0) // 从索引值为3处开始读
	n, err := file.Read(buf)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("info:%s\n", string(buf))
	fmt.Printf("n:%s\n", n)
	file.Close()
}
```





##### file.Stat()获取文件信息

```go
func (f *File) Stat() (FileInfo, error)
```

实例：

```go
//StatFile 获取文件信息
func StatFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("f是否是一个文件: %v\n", fInfo.IsDir())
	fmt.Printf("f文件的修改时间: %v\n", fInfo.ModTime().String())
	fmt.Printf("f文件的名称: %v\n", fInfo.Name())
	fmt.Printf("f文件的大小: %v\n", fInfo.Size())
	fmt.Printf("f文件的权限: %v\n", fInfo.Mode().String())
}

```



### 文件写操作



##### file.Write()数据写入

```go
func (f *File) Write(b []byte) (n int, err error)
```

实例:

```go
//WritFile 写入数据
func WritFile(name string) {
	file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	for i := 0; i < 10; i++ {
		file.Write([]byte(fmt.Sprintf("hello golang 我是%d\n  ", i)))
	}
	file.Close()
}
```

运行结果:

```tex
hello golang 我是0
hello golang 我是1
hello golang 我是2
hello golang 我是3
hello golang 我是4
hello golang 我是5
hello golang 我是6
hello golang 我是7
hello golang 我是8
hello golang 我是9
```



##### file.WriteString()写入字符串

```go
func (f *File) WriteString(s string) (n int, err error）
```

实例:

```go
//WriteStringFile 写入字符串
func WriteStringFile(name string) {
	file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
	if err != nil {
		fmt.Printf("err:%v\n")
	}
	file.WriteString("您好 golang")
}
```





##### file.WriteAt()写入指定位置

```go
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
```

实例：

```go
//WriteFileAt 写入指定位置
func WriteFileAt(name string) {
	file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
	if err != nil {
		fmt.Printf("err:%v\n")
	}
	file.WriteAt([]byte("学习使我快乐"), 10) // 从索引值为10的地方开始写入并覆盖原来当前位置的数据
}
```



实例:使用缓冲区

```go
//WriteFile 使用缓冲区
func WriteFile(name string) {
	file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
	if err != nil {
		fmt.Printf("err:%v\n")
	}
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	writefile := bufio.NewWriter(file)
	for i := 0; i < 50; i++ {
		writefile.WriteString(fmt.Sprintf("您好，我是第%d个帅哥  \n", i))
	}
	//Flush将缓存的文件真正写入到文件中
	writefile.Flush()
}
```



##### 逐行读取

os包没有给我们提供逐行读取的方法，这需要我们自己实现:

```go
//ReadLine 逐行读
func ReadLine(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	//将文件读到缓冲区
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Printf("行数据:%v\n", line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
```



### 进程相关

```go
func Exit(code int) // 让当前程序以给出的状态码（code）退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。

func Getuid() int // 获取调用者的用户id

func Geteuid() int // 获取调用者的有效用户id

func Getgid() int // 获取调用者的组id

func Getegid() int // 获取调用者的有效组id

func Getgroups() ([]int, error) // 获取调用者所在的所有组的组id

func Getpid() int // 获取调用者所在进程的进程id

func Getppid() int // 获取调用者所在进程的父进程的进程id

```



### 环境相关

```go
func Hostname() (name string, err error) // 获取主机名

func Getenv(key string) string // 获取某个环境变量

func Setenv(key, value string) error // 设置一个环境变量,失败返回错误，经测试当前设置的环境变量只在 当前进程有效（当前进程衍生的所以的go程都可以拿到，子go程与父go程的环境变量可以互相获取）；进程退出消失

func Clearenv() // 删除当前程序已有的所有环境变量。不会影响当前电脑系统的环境变量，这些环境变量都是对当前go程序而言的

```