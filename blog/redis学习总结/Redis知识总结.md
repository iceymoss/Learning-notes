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


## 列表(list)

redis中list可以被设计为栈，队列，阻塞队列等。

list使用的核心是所有命令都是以`L`开头的（不区分大小写)，默认是先进后出。

* `lpush listname item`向列表中添加元素（从list左边添加数据)

  实例：

  ```sh
  127.0.0.1:6379> lpush userlist "iceymoss"
  1
  127.0.0.1:6379> lpush userlist "jierkill"
  2
  127.0.0.1:6379> lpush userlist "srenlag"
  3
  127.0.0.1:6379> lpush userlist "moskel"
  4
  ```

* `lrange listname indexstart indexend`索取list数据

  ```sh
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "moskel",
      "srenlag",
      "jierkill",
      "iceymoss"
  ]
  ```

  数据默认以栈的方式加入list

* `rpush listname item`从list右边添加数据

  ```sh
  127.0.0.1:6379> rpush userlist "iceymoss"
  1
  127.0.0.1:6379> rpush userlist "jierkill"
  2
  127.0.0.1:6379> rpush userlist "srenlag"
  3
  127.0.0.1:6379> rpush userlist "moskel"
  4
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "iceymoss",
      "jierkill",
      "srenlag",
      "moskel"
  ]
  ```

* `lpop 默认弹出左边第一个元素`

  ```sh
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "moskel",
      "srenlag",
      "jierkill",
      "iceymoss"
  ]
  127.0.0.1:6379> lpop userlist
  moskel
  127.0.0.1:6379> lpop userlist
  srenlag
  127.0.0.1:6379> lpop userlist
  jierkill
  127.0.0.1:6379> lpop userlist
  iceymoss
  127.0.0.1:6379> lpop userlist
  (nil)
  ```

  我们将`lpush`和`lpop`结合起来userlist变成了一个栈，即先进后出。

* `rpop默认弹出右边第一个元素`

  ```sh
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "iceymoss",
      "jierkill",
      "srenlag",
      "moskel"
  ]
  127.0.0.1:6379> lpop userlist
  iceymoss
  127.0.0.1:6379> lpop userlist
  jierkill
  127.0.0.1:6379> lpop userlist
  srenlag
  127.0.0.1:6379> lpop userlist
  moskel
  127.0.0.1:6379> lpop userlist
  (nil)
  ```

  我们将`rpush`和`lpop`(`lpush`和`rpop`)结合起来userlist就变成了一个队列，即先进先出。

  

* ` lindex` 通过下标获取值

  ```sh
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "iceymoss2",
      "iceymoss1",
      "iceymoss"
  ]
  127.0.0.1:6379> lindex userlist 0
  iceymoss2
  127.0.0.1:6379> lindex userlist -1
  iceymoss
  127.0.0.1:6379> 
  ```

* `llen `获取keylist的长度

  ```sh
  127.0.0.1:6379> llen userlist
  3
  ```

* `lrem`移除指定的值

  ```sh
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "iceymoss3",
      "iceymoss3",
      "iceymoss3",
      "iceymoss2",
      "iceymoss1",
      "iceymoss"
  ]
  127.0.0.1:6379> lrem userlist 1 "iceymoss" #移除uselist中的一个值: iceymoss
  1
  127.0.0.1:6379> lrem userlist 3 "iceymoss3" #移除userlist中的3个值，iceymoss3
  3
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "iceymoss2",
      "iceymoss1"
  ]
  ```

* `ltrim`通过下标截取指定下标的list操作，截取后只剩下截取后的结果

  ```sh
  ltrim mylist 1 2
  ```

* `rpoplpush userlist userlist1`将先对userlist进行rpop操作获取到数据，然后对userlist1进行lpush操作。

  ```sh
  127.0.0.1:6379> lrange userlist 0 -1
  [
      "iceymoss3",
      "iceymoss2",
      "iceymoss1"
  ]
  127.0.0.1:6379> rpoplpush userlist userlist1
  iceymoss1
  127.0.0.1:6379> lrange userlist1 0 -1
  [
      "iceymoss1"
  ]
  ```

  

*  `lset`更新操作，目标元素要先存在，否则失败

  ```sh
  127.0.0.1:6379> lrange userlist1 0 -1
  [
      "iceymoss1"
  ]
  127.0.0.1:6379> lset userlist1 0 "yangkuang"
  OK
  127.0.0.1:6379> lrange userlist1 0 -1
  [
      "yangkuang"
  ]
  127.0.0.1:6379> 
  ```



* ` linsert` 列表插入操作

  ```sh
  127.0.0.1:6379> linsert userlist1 before "yangkuang" "你好呀" #向uselist1中“yangkuang"前后插入
  2 
  127.0.0.1:6379> lrange userlist1 0 -1
  [
      "你好呀",
      "yangkuang"
  ]
  ```

>总结：
>
>- list实际上是一个链表，before node after，left，right 都可以插入值
>- 如果key不存在，创建新的链表
>- 如果key存在，新增内容
>- 如果移除了所有值，空链表，也代表不存在！
>- 在两边插入或者改动值，效率最高！中间元素，相对来说效率会低一点



## 集合(set)

set是集合，里面的元素是无序且唯一的

对set操作的核心是使用`S`

```sh
# 使用sadd进行添加元素
127.0.0.1:6379> sadd username "iceymoss"  
1
127.0.0.1:6379> sadd username "iceymos1"
1
127.0.0.1:6379> sadd username "iceymos2"

# 使用smembers查看元素
127.0.0.1:6379> smembers username 
[
    "iceymos1",
    "iceymos2",
    "iceymos3",
    "iceymoss"
]

# sismember判断某元素是否在集合中
127.0.0.1:6379> sismember username "iceymoss2"
0 #返回0表示不在

# scard返回金集合的元素个数
127.0.0.1:6379> scard username
4

# srem 移除集合中指定元素
127.0.0.1:6379> smembers username
[
    "iceymos1",
    "iceymos2",
    "iceymos3",
    "iceymoss"
]
127.0.0.1:6379> srem username "iceymoss"
1
127.0.0.1:6379> smembers username
[
    "iceymos1",
    "iceymos2",
    "iceymos3"
]

# srandmember随机抽取一个集合中的元素
127.0.0.1:6379> srandmember username
iceymos2
127.0.0.1:6379> srandmember username
iceymos2
127.0.0.1:6379> srandmember username
iceymos3

# spop随机移除元素
127.0.0.1:6379> spop username 
iceymos3
127.0.0.1:6379> spop username 
iceymos1
127.0.0.1:6379> smembers username
[
    "iceymos2"
]

# smove将一个指定的值，移动到另一个set集合中
127.0.0.1:6379> smembers username
[
    "iceymos2",
    "dkfj",
    "iceymoss1",
    "iceymoss2",
    "iceymoss4",
    "iceymoss3"
]
127.0.0.1:6379> sadd username1 "yauso"
1
127.0.0.1:6379> smembers usernam1
[]
127.0.0.1:6379> smembers username1
[
    "yauso"
]
127.0.0.1:6379> smove username username1 "dkfj"
1
127.0.0.1:6379> smembers username1
[
    "yauso",
    "dkfj"
]

# 数字集合类：(例如抖音B站：共同关注功能)
- 差集
- 交集
- 并集
127.0.0.1:6379> smembers user1_star
[
    "汤姆老师",
    "英雄联盟解说",
    "法外狂徒",
    "影视飓风",
    "kuangsheng",
    "偶像练习生蔡徐坤"
]
127.0.0.1:6379> smembers user2_star
[
    "lpl赛事",
    "罗翔说刑法",
    "汤姆老师",
    "偶像练习生蔡徐坤"
]

# sinter交集，两个用户的共同关注
127.0.0.1:6379> sinter user1_star user2_star
[
    "汤姆老师",
    "偶像练习生蔡徐坤"
]

# sunion并集
127.0.0.1:6379>  sunion user1_star user2_star
[
    "罗翔说刑法",
    "汤姆老师",
    "法外狂徒",
    "lpl赛事",
    "kuangsheng",
    "英雄联盟解说",
    "偶像练习生蔡徐坤",
    "影视飓风"
]

# sdiff差集 业务场景：抖音推荐对方还关注了(对方关注了，你没关注，即差集)
127.0.0.1:6379> sdiff user1_star user2_star
[
    "影视飓风",
    "kuangsheng",
    "英雄联盟解说",
    "法外狂徒"
]

```

微博，A用户将所有关注的人放在一个set集合中！将它的粉丝也放在一个集合中！共同关注，共同爱好，二度好友，推荐好友（六度分割理论）



## 散列(Hashes)

map集合，key-value相信学过任何一门编程语言的人都这个map

```sh
# hset设置hashe
127.0.0.1:6379> hset user name "iceymoss"
1

# hget获取值
127.0.0.1:6379> hget user1 name
dkfiel

# hmset设置hashe多个属性
127.0.0.1:6379> hmset userinfo name "iceymoss" age 18 gender "男"
OK

# hmget设置hashe多个属性
127.0.0.1:6379> hmget userinfo name age gender
[
    "iceymoss",
    "18",
    "男"
]

# 获取hashe的所有key-value
127.0.0.1:6379> hgetall userinfo
{
    "name": "iceymoss",
    "age": "18",
    "gender": "男"
}

# 删除hashe指定key，对应的value也就删除了
127.0.0.1:6379> hgetall userinfo
{
    "name": "iceymoss",
    "age": "18",
    "gender": "男"
}
127.0.0.1:6379> hdel userinfo gender
1
127.0.0.1:6379> hgetall userinfo
{
    "name": "iceymoss",
    "age": "18"
}
127.0.0.1:6379> 

# hlen查看hashe的key数量
127.0.0.1:6379> hlen userinfo
2

# hexists判断hashe的某个key是否存在
127.0.0.1:6379> hexists userinfo name
1  #返回1表示存在，0不在

# hkeys返回所有key
127.0.0.1:6379> hkeys userinfo
[
    "name",
    "age"
]

# hvals返回所有value
127.0.0.1:6379> hvals userinfo
[
    "iceymoss",
    "18"
]
127.0.0.1:6379> 

# hincrby给hash某个字段指定增量
127.0.0.1:6379> hincrby userinfo age 1
19
127.0.0.1:6379> 
127.0.0.1:6379> hincrby userinfo age 1
20
127.0.0.1:6379> hincrby userinfo age 1
21
127.0.0.1:6379> hincrby userinfo age 1
22
127.0.0.1:6379> hincrby userinfo age 1
23
127.0.0.1:6379> hincrby userinfo age 1
24
127.0.0.1:6379> hgetall userinfo
{
    "name": "iceymoss",
    "age": "24"
}

# hsetnx判断hash某个字段是否存在，不存在则可以设置，存在则不可以设置（应用分布式锁）
127.0.0.1:6379> hgetall userinfo
{
    "name": "iceymoss",
    "age": "24"
}
127.0.0.1:6379> hsetnx userinfo gender "男"
1
127.0.0.1:6379> hsetnx userinfo gender "男"
0
127.0.0.1:6379> hgetall userinfo
{
    "name": "iceymoss",
    "age": "24",
    "gender": "男"
}
```

hash变更的数据user name age，尤其是用户信息之类的，经常变动的信息！hash更适合对象的存储



## 有序集合(Zset)

有序集合

在set的基础上，增加了一个值我们可以理解为权值，对比：sadd myset hello，zadd myzset 1 hello

所有zset的命令都是以z开头的

```sh
# zadd添加元素
127.0.0.1:6379> zadd myzset 1 chinese
1
127.0.0.1:6379> zadd myzset 2 math

# zrange获取有序集合所有元素
127.0.0.1:6379> zrange myzset 0 -1
[
    "chinese",
    "math"
]

# 排序
127.0.0.1:6379> zadd salary 2500 xiaohong
(integer) 1
127.0.0.1:6379> zadd salary 5000 zhangsan
(integer) 1
127.0.0.1:6379> zadd salary 500 xiaoming
(integer) 1

# 升序排序 从小到大
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf
1) "xiaoming"
2) "xiaohong"
3) "zhangsan"

# 升序排序带参数
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf withscores
1) "xiaoming"
2) "500"
3) "xiaohong"
4) "2500"
5) "zhangsan"
6) "5000"

# 从负无穷到2500进行排序并且附带参数
127.0.0.1:6379> ZRANGEBYSCORE salary -inf 2500 withscores
1) "xiaoming"
2) "500"
3) "xiaohong"
4) "2500"

# 降序排序，从大到小
127.0.0.1:6379> ZREVRANGE salary 0 -1
1) "zhangsan"
2) "xiaohong"
127.0.0.1:6379> ZREVRANGE salary 0 -1 withscores
1) "zhangsan"
2) "5000"
3) "xiaohong"
4) "2500"

# zrem移除元素
127.0.0.1:6379> ZRANGE salary 0 -1
1) "xiaoming"
2) "xiaohong"
3) "zhangsan"
127.0.0.1:6379> zrem salary xiaoming
(integer) 1
127.0.0.1:6379> ZRANGE salary 0 -1
1) "xiaohong"
2) "zhangsan"

# 获取有序集合中的个数
127.0.0.1:6379> zcard salary
(integer) 2

# 获取集合不同区间中的个数
127.0.0.1:6379> zcount myzset 1 2
(integer) 2
127.0.0.1:6379> zcount myzset 1 2
(integer) 2
127.0.0.1:6379> zcount myzset 1 3
(integer) 3
127.0.0.1:6379> zcount myzset 1 4
(integer) 3
127.0.0.1:6379> zcount myzset 0 4
(integer) 3
```

业务场景: 薪资排序，排序榜等

普通消息：1，重要消息：2 带权重进行判断

## 未完待续
……

## 参考文献

[B站up主-狂神说Java](https://www.bilibili.com/video/BV1S54y1R7SB/?spm_id_from=333.999.0.0)

