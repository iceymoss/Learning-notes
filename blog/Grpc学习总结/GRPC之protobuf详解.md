[toc]



### 文章介绍

本文主要介绍我学习 protobuf 的理解和总结、主要介绍 protobuf 的基本类型、option 的作用、proto 文件中导入其他 proto 文件、嵌套 message、enum 枚举类型、map 类型、proto 中内置的 timetram 类型及 service {}

### protobuf 介绍
protobuf 全称 Google Protocol Buffers，是 google 开发的的一套用于数据存储，网络通信时用于协议编解码的工具库。protobuf 是一种灵活高效的独立于语言平台的结构化数据表示方法。在通信协议和数据存储等领域中使用比较多。protobuf 对于结构中的每个成员会提供 set 系列函数和 get 系列函数。与 XML 相比，protoBuf 更小更快更简单。你可以用定义 protobuf 的数据结构。用 protobuf 编译器生成特定语言的源代码，如 C++，Java，go，Python 等，proto 文件是以 xxx.proto 命名的。

### 基本类型

- ###### int

  ```protobuf
  int32
  int64
  ```

- float

  ```protobuf
  float
  ```

- string

  ```protobuf
  string
  ```

  这是常用的基本类似，更多类型请参考[官方地址](https://developers.google.com/protocol-buffers/docs/proto3)



### option 的作用

我们先来看一个简单的 proto 文件

```protobuf
syntax = "proto3";  //值proto3的语法

option go_package="/.;proto";  
```

`option`: 指生成的哪一个语言的代码及生成目的文件下

`go_package`: 指生成 go 语言的代码

`"/.;proto"`：指在当前文件下生成，并且包名为 proto



### proto 文件中导入其他 proto 文件

我们在开发的过程中，难免会遇到代码重复使用的情况，这时我们可以将 proto 文件导入

例如：我们要使用 case.prot 的内容，我们就需要导入 case.proto 文件，则使用 ```import```

```protobuf
syntax = "proto3";
option go_package="/.;proto";
import "case.proto";  //
```

然后我们就可以使用 case.proto 中的内容了

- case.proto:

  ```protobuf
  syntax = "proto3";
  option go_package = "/.;proto";
  
  message IsEmpty{}
  
  message Pong{
      string id = 1;
  }
  ```

  

- holle.proto:

  ```protobuf
  syntax = "proto3";
  
  option go_package="/.;proto";
  //引入protobuf的内置类型
  import "google/protobuf/timestamp.proto";
  import "case.proto";
  
  //定义接口
  service Greeter {
      rpc SayHello (HolleRequest) returns (HolleReply);
      rpc Ping(IsEmpty) returns (Pong);     //Ping(IsEmpty) returns (Pong)中的IsEmpty，Pong来自case.proto中
  }
  
  message HolleRequest {
      string name = 1;
      string url = 2;
  }
  
  message HolleReply {
      string id = 1;
  }
  ```

  

### message 及嵌套 message

- ###### message

我们还是直接看代码：

```protobuf
message HolleRequest {
    string name = 1;  //name = 1不是赋值，是指在字段的编号
    string url = 2;
}
```

这就是一个简单的 message 他类似于结构体，message 内部我们可以定义他的字段，这里需要注意的是：

```protobuf
string name = 1;
string url = 2;
```

不是赋值，而是给字段的编号

- ###### 嵌套 message

  我们可以在 message 嵌套一个或者多个 message， 下面我们来看示例：

  ```protobuf
  message HolleRequest {
      string name = 1;
      string url = 2;
  }
  
  message HolleReply {
      string id = 1;
      HelloRequest request = 2;   //HelloReply中嵌套HelloRequest
  }
  ```

  这样我们就实现了 message 的嵌套



### enum 枚举类型

enum 枚举类型是我们在业务中经常需要使用的，例如，性别 (男，女)、用户身份认证 (为认证，认证中，认证失败) 等，这里我们以性别为例：

```protobuf
//枚举类型
enum Gender{
    MALE = 0;
    FE_MALE = 1;
}
```

一个简单的枚举类型就定义完成了，那么可以把他放入我们的 message 中：

```protobuf
message HolleRequest {
    string name = 1;
    string url = 2;
    Gender gender = 3;
}
```



### map 类型

我们只需要 `map<string, string> mp = 4;`

```protobuf
message HolleRequest {
    string name = 1;
    string url = 2;
    Gender gender = 3;
    map<string, string> mp = 4;  //proto map类型
}
```



### timestamp 内置类型

proto 中也有一些内置的类型，例如我们要介绍的 timestamp

我们使用它时，需要导入 “google/protobuf/timestamp.proto” 这个是 protobuf 官方定义的

```protobuf
message HolleRequest {
    string name = 1;
    string url = 2;
    Gender gender = 3;
    map<string, string> m = 4;  //proto map类型
    google.protobuf.Timestamp addTime = 5;  //protobuf的内置类型
}
```



### 接口(service{})
这里我的理解是我们上面介绍的所有类型最后都是为了 service {} 而准备的，service {} 我们可以理解是一个接口，里面有我们定义的各种方法，最后我们在业务中调用的方法也来自于此。

```protobuf
service Greeter {
        //定义了一个SayHello方法，入参：HolleRequest， 出参：HelloReply
    rpc SayHello (HolleRequest) returns (HelloReply); 
}
```

假如我们有一个业务需要：我们需要给用户信息绑定一个 id, 以便于业务在后期快速查找

我们这样定义 proto 文件：

```protobuf
syntax = "proto3";

option go_package="/.;proto";
//引入protobuf的内置类型
import "google/protobuf/timestamp.proto";


//定义接口
service Greeter {
    rpc SayHello (HolleRequest) returns (HelloReply);
}

//枚举类型
enum Gender{
    MALE = 0;
    FE_MALE = 1;
}

message HolleRequest {
    string name = 1;
    string url = 2;
    Gender gender = 3;
    map<string, string> m = 4;  //proto map类型
    google.protobuf.Timestamp addTime = 5;  //protobuf的内置类型
}

message HelloReply {
    string id = 1;
    HolleRequest request = 2;
}
```

这样我们的 proto 文件就完成了，那如何使用它生成对应语言的业务代码呢？

我们使用命令：

```shell
protoc -I . holle.proto --go_out=plugins=grpc:. 
```

```-I```: 表示输入 input 即输入文件

```.```: 表示在当前目录，这里也可以使用绝对路径或者相对路径

```holle.proto```: 就是我们写的 ```proto``` 文件

```--go_out```: 输出，以 go 语言的形式输出

