[toc]

# Redis入门

## 概述

### Redis是什么

Redis（Remote Dictionary Serve），即远程字典服务

### Redis能干嘛？

1、内存存储，持久化，内存中是断电即失，所以说持久化很重要（rdb、aof）

2、效率高，可以用于高速缓存

3、发布订阅系统

4、地图信息分析

5、计时器、计数器（浏览量！ ）

### 特性

1、多样的数据类型

2、持久化

3、集群

4、事物



## 安装(docker安装)

使用命令：

```
docker run -p 6379:6379 -d redis:latest redis-server

docker container update --restart=always 容器名字
```



## 基础知识

redis默认共有16个数据库，0为默认使用的数据库，使用```select```进行切换

```
127.0.0.1:6379> select 1   #切换到数据库1
OK
127.0.0.1:6379> 
```

* `keys *`查看所有的key

* `flushdb`，`flushall`清空当前数据库，清空所有数据库

* `exists key`，判断值是否存在

* `move name 1`，移动key为name到数据库1

* `expire name 20`，设置key为name的值20s后过期

* `ttl name`，查看当前key剩余秒数

  

### redis是单线程

redis是基于内存，性能瓶颈不在cpu，redis的瓶颈是机器的内存和网络带宽

> redis为什么单线程还这么块?

首先我们需要了解：

* 高性能不一定就是多线程
* 多线程(上下文切换)不一定比单线程速率高

*** 注意：redis的所有数据都是放在内存的，不用发生缺页中断不用进程置换，也不用进行额外的IO，单线程不需要进行上下文切换，所以单线程的效率高。***



## 数据类型

* `string`字符串
* `list`列表
* `hashes` 散列
* `set`集合
* `sort set`有序集合
* 范围查询， [bitmaps](http://www.redis.cn/topics/data-types-intro.html#bitmaps)， [hyperloglogs](http://www.redis.cn/topics/data-types-intro.html#hyperloglogs) 和 [地理空间（geospatial）](http://www.redis.cn/commands/geoadd.html) 索引半径查询



## 基础命令

```sh
keys *  #查看所有的key
flushdb   #清空当前数据库，
flushall  #清空所有数据库
exists key #判断值是否存在
move name 1 #移动key为name到数据库1
expire name 10 #设置key为name的值10s后过期
ttl name  #查看当前key剩余秒数
type name #查看当前key为name的值类型
append key1 "hello" #追加
strlen key1 # 查看key为key1的长度
```

## String字符串

```sh
set key1 v1 #设置值
get key1 #获得值
append key1 "hello" #追加，如果key不存在，相当于set key
strlen key1 # 查看key为key1的长度
incr views # 自加1
decr views # 自减1
incrby views 10 # 自加10，步长为10
decrby views 10 # 自减10，步长为10
getrange key1 0 3 # 查看字符串范围 0到3（包括3）
getrange key1 0 -1 # 查看字符串全部范围
setrange key1 3 xxx # 替换指定位置3的字符串
setex #(set with expire) 设置过期时间
setex key3 30 "hello"
setnx #(set if not exist)    不存在再设置
setnx mykey "redis"
mset k1 v1 k2 v2 k3 v3  # 批量设置值
mget k1 k2 k3     # 批量获取值
msetnx k1 v1 k4 v4 # 当不存在时批量设置值 msetnx是一个原子性的操作，要么一起成功，要么一起失败
# 对象
set user:1 {name:zhangsan,age:3}
# 批量设置对象
# 这里的key是一个巧妙的设计：user:{id}:{filed}，如此设计可以极大的提升复用率
mset user:2:name lisi user:2:age 21
OK
mget user:2:name user:2:age
1) "lisi"
2) "21"
getset key1 redis #先获取，再设置，如果不存在返回nil，设置新值，如果存在，替换原来的值

```

数据结构是相同的，string类似的使用场景：value除了是我们的字符串还可以是我们的数字

- 计数器
- 统计多单位的数量



## 未完待续

……

## 参考文献

[B站up主-狂神说Java](https://www.bilibili.com/video/BV1S54y1R7SB/?spm_id_from=333.999.0.0)

