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



## Redis的三种特殊数据类型

### geospatial 地理位置存储

用于存储地理数据，而redis可以使用我们的业务场景中更方便的去使用该数据类型做，用户的定位、位置计算、网约车车费计算，查询附近的人等很多场景。

#### getadd

> 用于添加地理位置

需要注意一下几点：

* 两级无法直接添加，大量数据可以直接通过相关程序一次性导入
* 当经纬度超出一定范围时，会报超范围错误

实例：

```sh
127.0.0.1:6379> geoadd china:city 116.40 39.90 beijing
(integer) 1
127.0.0.1:6379> geoadd china:city 121.47 31.23 shanghai
(integer) 1
127.0.0.1:6379> geoadd china:city 106.50 29.53 chongqing 114.05 22.52 shenzhen
(integer) 2
127.0.0.1:6379> geoadd china:city 120.16 30.24 hangzhou
(integer) 1
127.0.0.1:6379> geoadd china:city 108.96 34.26 xian
```



#### geopos

> 获取地理位置

实例：

```sh
127.0.0.1:6379> geopos china:city beijing
1) 1) "116.39999896287918091"
   2) "39.90000009167092543"
127.0.0.1:6379> geopos china:city beijing shanghai
1) 1) "116.39999896287918091"
   2) "39.90000009167092543"
2) 1) "121.47000163793563843"
   2) "31.22999903975783553"
```



#### geodist

> 获取给定的两个位置之间的直线距离

- m：米
- km：千米
- mi：英里
- ft：英尺

实例：

```sh
# 查看两个城市之间的距离
127.0.0.1:6379> geodist china:city beijing shanghai
"1067378.7564"
127.0.0.1:6379> geodist china:city beijing shanghai km  #以km为单位
"1067.3788"
```



#### georadius

> 给给定的一个坐标为中心，返回某个半径内的元素的值，业务场景：附近的人，附近车辆等。

实例：

```sh
# 找出110 30为中心 500 km半径内的元素
127.0.0.1:6379> georadius china:city 110 30 500 km
1) "chongqing"
2) "xian

# 显示直线距离
127.0.0.1:6379> georadius china:city 110 30 500 km withdist
1) 1) "chongqing"
   2) "341.9374"
2) 1) "xian"
   2) "483.8340"

127.0.0.1:6379> georadius china:city 110 30 500 km withdist withcoord
1) 1) "chongqing"
   2) "341.9374"
   3) 1) "106.49999767541885376"
      2) "29.52999957900659211"
2) 1) "xian"
   2) "483.8340"
   3) 1) "108.96000176668167114"
      2) "34.25999964418929977"
      
# 指定获得数量
127.0.0.1:6379> georadius china:city 110 30 500 km withdist withcoord count 1
1) 1) "chongqing"
   2) "341.9374"
   3) 1) "106.49999767541885376"
      2) "29.52999957900659211"
```



#### georadiusbymember

> 以地点名称为中心查询，查询指定半径内的的元素

实例：

```sh
127.0.0.1:6379>  GEORADIUSBYMEMBER china:city shanghai 2000 km
[
    "chongqing",
    "xian",
    "shenzhen",
    "hangzhou",
    "shanghai",
    "beijing"
]
```



#### geohash

>geohash 返回一个或多个位置元素的geohash表示，使用该命令将返回11 个字符串的geohash字符串

实例：

```sh
127.0.0.1:6379> geohash china:city beijing shenzhen
[
    "wx4fbxxfke0",
    "ws10578st80"
]
```

在Redis的五大基础数据类型学完后，其实geo的底层就是使用Zset进行封装的，下面我们来使用Zset命令操作地理数据

实例：

```bash
127.0.0.1:6379> zrange china:city 0 -1  #查看所有数据
[
    "chongqing",
    "xian",
    "shenzhen",
    "hangzhou",
    "shanghai",
    "beijing"
]

127.0.0.1:6379>  zrem china:city xian  #移除西安
1
127.0.0.1:6379> zrange china:city 0 -1
[
    "chongqing",
    "shenzhen",
    "hangzhou",
    "shanghai",
    "beijing"
]
```



### hyperloglogs 基数存储

> Ps: 这里需要先介绍一下基数，将一个系列的容器中，去重后的元素个数。
>
> 例如：{1, 2, 2, 5, 7, 7, 8, 12, 34} 原来总数为：9，基数：7

hyperloglogs是Redis统计基数的算法，有优点是：占用内存固定，2^64不同元素，只需要2KB的内存，从内存的角度来说使用hyperloglogs是很好的。

##### 使用场景

> 统计网站的访问量，每一个设备只能访问计数一次，多次访问都是一次。

传统的方式，set保存用户的id，然后就可以统计set中的元素数量作为标准判断

这个方式如果保存大量的用户id，就会比较麻烦，我们的目的是为了计数，而不是保存用户id

但是需要注意：

* 0.81%错误率！统计任务，可以忽略不计的

实例：

```bash
# 添加一个hyperloglog ，重复则替换
127.0.0.1:6379> pfadd key1 a b c d e f g
1

# 统计一个key1中有多少元素
127.0.0.1:6379>  PFCOUNT key1
7
127.0.0.1:6379>  pfadd key2 a b c d e f g h i j k l m n
1
127.0.0.1:6379> PFCOUNT key2
14

# 合并两个hyperloglog，求并集
127.0.0.1:6379> PFMERGE mykey key1 key2
OK
127.0.0.1:6379> PFCOUNT mykey
14
127.0.0.1:6379> 
```



### bitmaps 位存储

Bitmaps位图，数据结构！都是操作二进制位来进行记录，只有0和1两个状态！常见的使用场景：用户信息，活跃or不活跃，在线or不在线等等。

实例：

使用bitmaps来记录周一到周日的打卡情况：1打卡；0未打卡

>周一(0): 1;周二(1): 0;周三(2): 1;周四(3): 0;周五(4): 1;周六(5): 1;周日(6): 1   ps：你看你什么成分还想周末休息，赶紧加班吧！(狗头)

```bash
# 某员工打卡情况
127.0.0.1:6379> setbit sign 0 1
0
127.0.0.1:6379> setbit sign 1 0
0
127.0.0.1:6379> setbit sign 2 1
0
127.0.0.1:6379> setbit sign 3 0
0
127.0.0.1:6379> setbit sign 4 1
0
127.0.0.1:6379> setbit sign 5 1
0
127.0.0.1:6379> setbit sign 6 1
0

# 老板心情好查看一下某员工的打卡情况
# 查看对应存的值
127.0.0.1:6379> getbit sign 2
1
# 周四忘打卡，拿不到全勤奖金了！
127.0.0.1:6379> getbit sign 3
0

# 统计打开天数
127.0.0.1:6379> bitcount sign
5
```



## 事务

我们在学关系型数据库的时候都学过事务，以及事务的几大特性(ACID)，以及隔离级别等。

Redis的事务本质：**一组命令的集合** 一个事务中的所有命令都会被序列化，在事务执行过程中，会按照命令的顺序执行。

Redis事务的特性：

* 一次性

* 顺序性

* 排他性

  ```
  队列queue： set set set set get set 执行
  ```

**注意：redis单条命令是保证原子性的，但是Redis的事务是不保证原子性的，Redis也是没有隔离级别的。**

当开启事务后，所有的命令都会放入队列中，不会别直接执行的，只有发起执行命令的时候才会执行！```exec```

Redis的事务命令：

* 开始事务：(``multi``)
* 命令入队(……)
* 执行事务(`exec`)



#### 正常事务的开启执行

实例：

```bash
127.0.0.1:6379> multi  #开启事务
OK
127.0.0.1:6379> set k1 v1    #命令入队
QUEUED
127.0.0.1:6379> set k2 v2    #命令入队
QUEUED
127.0.0.1:6379> get k2       #命令入队
QUEUED
127.0.0.1:6379> set k3 v3    #命令入队
QUEUED
127.0.0.1:6379> exec         #执行事务
[
    "OK",
    "OK",
    "v2",
    "OK"
]
```



#### 放弃事务(discard)

当我们开启事务后然后不想执行当前事务，使用```discard```命令放弃事务

实例：

```bash
127.0.0.1:6379> multi     #开启事务
OK
127.0.0.1:6379> set k1 v1   
QUEUED
127.0.0.1:6379> set k2 v2
QUEUED
127.0.0.1:6379> get k1
QUEUED
127.0.0.1:6379> discard   #放弃事务
OK
127.0.0.1:6379> set k1 v1
OK
127.0.0.1:6379> get k1
v1
```

#### 两大异常

##### 编译型异常

编译型异常指代码有问题！命令错误！在事务中所有的命令都不会被执行！不会别编译过去！

实例：

```bash
127.0.0.1:6379> multi   #开启事务
OK
127.0.0.1:6379> set k1 v1
QUEUED
127.0.0.1:6379> set k2 v2
QUEUED
127.0.0.1:6379> getset k2   #写入一个错误命令，直接报错
ERR wrong number of arguments for 'getset' command
127.0.0.1:6379> set k4 v4   #继续向事务队列中写入命令
QUEUED
127.0.0.1:6379> exec        #执行事务，此时所有命令都不会被执行
EXECABORT Transaction discarded because of previous errors.
```



##### 运行时异常

运行时异常指，给出命令没有问题，例如1/0，那么在执行命令的时候其他没有问题的命令是可以正常执行的，错误命令抛出异常！

实例：

```sh
127.0.0.1:6379> multi
OK
127.0.0.1:6379> get k1
QUEUED
127.0.0.1:6379> set k2 v2
QUEUED
127.0.0.1:6379> set k3 v3
QUEUED
127.0.0.1:6379> exec
Cannot read properties of null (reading 'type')  #这里会报错，但是可以正常获取其他值
127.0.0.1:6379> get k1    #没有k1
(nil)
127.0.0.1:6379> get k2    #有k2
v2
```



## 监控(watch )

### 锁

#### 悲观锁

悲观锁可以理解为做什么事都以一种悲观的态度，认为事情总会发生，就是指在任何情况下都加锁。



#### 乐观锁

很乐观，认为什么情况下都不会发生，持乐观态度，所以不会上锁，更新的时候去去检查在此期间(指从查询到数据->更新数据之间)操作的数据在数据库有没有被更新过，字段version。

步骤：

* 获取verison
* 更新数据的时候比较version

#### 基于Redis的乐观锁

##### 使用watch开启监视

**使用watch开启监视(获取version)，需要注意的是，在监控期间数据没有发生变动的时候事务才能正常执行成功，如果失败了，也需要使用unwatch解除监控。**

假设你和你女朋友一共有一百块钱money = 100去消费，那么的消费情况刚开始out = 0



**模拟单线程情况下：**

```bash
127.0.0.1:6379> set money  100
OK
127.0.0.1:6379> set out 0
ok

# 监控money对象
127.0.0.1:6379> watch money 
ok

# 期间数据没有发生变动，这时候就正常执行成功
# 开启事务
127.0.0.1:6379> multi
OK
127.0.0.1:6379> decrby money 20
QUEUED
127.0.0.1:6379> incrby out 20
QUEUED

# 执行事务
127.0.0.1:6379> exec
[
    80,
    20
]
```



**模拟多线程情况下：**

线程1执行：

```bash
127.0.0.1:6379> watch money
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379> decrby money 20
QUEUED
127.0.0.1:6379> incrby out 20
QUEUED
127.0.0.1:6379> exec
(nil)
127.0.0.1:6379> get money
```

线程2突然进来执行：

```go
127.0.0.1:6379> get money
100
127.0.0.1:6379> set money 200
OK
127.0.0.1:6379> 
```

然后回到线程1执行：

```bash
127.0.0.1:6379> exec
(nil)
```

我们的线程1的事务就执行失败了，此时的money = 200。

那么为什么线程1的事务会执行失败呢？答案：watch是Redis的乐观锁，线程1获取到version后，线程2执行了对数据库的更该，线程1在执行事务时(更新数据)发现前后的version值不想同，进而导致失败。

注意：如果发现事务失败后，需要解除监控，从新监控获取最新值。



## Go-Redis

### Go-Redis是什么？

Go-Redis是支持Redis Server和Redis Cluster的Golang客户端，接下来我们要使用golang来操作Redis。

#### 多种客户端

支持单机Redis Server、Redis Cluster、Redis Sentinel、Redis分片服务器

#### 数据类型

go-redis会根据不同的redis命令处理成指定的数据类型，不必进行繁琐的数据类型转换

#### 功能完善

go-redis支持管道(pipeline)、事务、pub/sub、Lua脚本、mock、分布式锁等功能

#### 安装

开源地址：https://github.com/redis/go-redis

直接使用命令安装，前提是您的设备已经配置了golang的开发环境。

[go-redis 支持 2 个最新的 Go 版本，并且需要具有模块](https://github.com/golang/go/wiki/Modules)支持的 Go 版本 。所以一定要初始化一个 Go 模块：

```
go mod init github.com/my/repo
```

然后安装 go-redis/ **v9**：

```
go get github.com/redis/go-redis/v9
```

#### 快速使用

```go
package main

import (
  "fmt"
	"time"
	"context"
 
  "github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
  
	defer rdb.Close()
  
	rdb.Ping(context.Background())

	status := rdb.Set(context.Background(), "money", 1000, time.Second*100)
	if status.Err() != nil {
		panic(status.Err())
	}

	res := rdb.Get(context.Background(), "money")
	if res.Err() != nil {
		panic(res.Err())
	}
	fmt.Println(res.Val())
}
```



我们在前面使用的命令都封装在go-redis中是一个一个的方法，可以参考go-redis:https://github.com/redis/go-redis/blob/master/commands.go#L158

下面将介绍go-redis的一些常用方法：

```go
//go-redis需要接收一个上下文
ctx := context.Background()

	//清空数据库
	rdb.FlushAll(ctx)
	rdb.FlushDB(ctx)

	//判断是否存在key
	rdb.Exists(ctx, "username")

	//设置key
	rdb.Set(ctx, "username", "iceymoss", time.Hour*24*30)
	rdb.Set(ctx, "password", "123456", time.Hour*24*30)

	//获取key
	rdb.Get(ctx, "username").Val()
	rdb.Get(ctx, "password").Val()
	
	//删除key
	rdb.Del(ctx, "username")
	rdb.Del(ctx, "password")

	//获取key类型
	rdb.Type(ctx, "username")
	
	//随机获取一个key
	rdb.RandomKey(ctx)

	//重命名key
	rdb.Rename(ctx, "username", "name")
	rdb.Get(ctx, "name")
	
	//返回key的总数
	rdb.DBSize(ctx)
```

下面将给出几大类型的方法，具体实例就展开了

#### 操作string的方法

- Get   获取值
- Set   设置key
- GetSet  设置一个key的值，并返回这个key的旧值:
- SetNX  如果key不存在，则设置这个key
- MGet  批量获取
- MSet   批量设置key
- Incr,IncrBy  自增1，自增指定长
- Decr,DecrBy  自减，自减指定长
- Del   删除key
- Expire  设置过期时间



#### 操作list的方法

-  LPush  从list左插入

-  LPushX  跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。

- RPop 从list移除第一个元素，并返回

- RPush  从list右插入

- RPushX  跟RPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。

- LPop  从list左边移出第一个元素并返回

- LLen 返回list大小

- LRange  返回列表的一个范围内的数据，也可以返回全部数据

- LRem 删除列表中的数据

  ```go
  // 从列表左边开始，删除100， 如果出现重复元素，仅删除1次，也就是删除第一个
  dels, err := rdb.LRem(ctx,"key",1,100).Result()
  if err != nil {
  	panic(err)
  }
  ```

* LIndex  根据索引坐标，查询列表中的数据

- LInsert  根据指定位置向list插入元素



#### 操作Hasha的方法

**内部采用数组+链表结构，采用链地址法解决哈希冲突。**

- HSet  根据key和field字段设置，field字段的值

  ```go
  // user_1 是hash key，username 是字段名, zhangsan是字段值
  err := rdb.HSet(ctx,"user_1", "username", "zhangsan").Err()
  if err != nil {
  	panic(err)
  }
  ```

- HGet  根据key和field字段获取field的值

  ```go
  // user_1 是hash key，username是字段名
  username, err := rdb.HGet(ctx,"user_1", "username").Result()
  if err != nil {
  	panic(err)
  }
  fmt.Println(username)
  ```

- HGetAll  根据根据key查询所有字段

  ```go
  // 一次性返回key=user_1的所有hash字段和值
  data, err := rdb.HGetAll(ctx,"user_1").Result()
  if err != nil {
  	panic(err)
  }
  
  // data是一个map类型，这里使用使用循环迭代输出
  for field, val := range data {
  	fmt.Println(field,val)
  }
  ```

- HIncrBy  根据key和field字段，累加字段的数值

- HKeys  根据key返回所有字段名

- HLen 根据key，查询hash的字段数量

- HMGet 根据key和多个字段获取多个字段值

- HMSet  根据key和多个字段名和字段值，批量设置hash字段值

- HSetNX  如果key不存在则创建

  ```go
  err := rdb.HSetNX(ctx,"key", "id", 100).Err()
  if err != nil {
  	panic(err)
  }
  ```

- HDel 根据key和字段名，删除hash字段，支持批量删除hash字段

  ```go
  // 删除一个字段id
  rdb.HDel(ctx,"key", "id")
  
  // 删除多个字段
  rdb.HDel(ctx,"key", "id", "username")
  
  ```

- HExists  检测hash字段名是否存在

  ```go
  // 检测id字段是否存在
  err := rdb.HExists(ctx,"key", "id").Err()
  if err != nil {
  	panic(err)
  }
  ```

  

#### 操作set的方法

* SAdd 向集合中添加元素

  ```go
  // 添加100到集合中
  err := rdb.SAdd(ctx,"key",100).Err()
  if err != nil {
  	panic(err)
  }
  
  // 将100,200,300添加到集合中
  rdb.SAdd(ctx,"key",100, 200, 300)
  ```

* SCard 获取集合元素个数
* SIsMember 判断是否在集合中
* SMembers 获取集合中的所有元素
* SRem 删除集合所有元素
* Spop, SpopN 随机返回集合中的元素，并且删除返回的元素



#### 操作Zset的方法

* ZAdd  向有序集合中添加元素

* ZCard  获取有序集合的元素个数

* ZCount  根据权重统计范围

* ZIncrBy  增加元素的分数

* ZRange,ZRevRange  返回集合中某个索引范围的元素，根据分数从小到大排序

* ZRangeByScore  根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。

* ZRevRangeByScore  用法类似ZRangeByScore，区别是元素根据分数从大到小排序。

* ZRangeByScoreWithScores  用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数

* ZRem  删除集合元素

* ZRemRangeByRank  根据索引范围删除元素

* ZRemRangeByScore  根据分数范围删除元素

* ZScore  查询元素对应的分数

* ZRank  根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序



#### 其他操作

##### geospatial

```go
	GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *IntCmd
	GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd
	GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *GeoSearchLocationCmd
	GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
```



##### hyperloglogs

```go
PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
PFCount(ctx context.Context, keys ...string) *IntCmd
PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd
……
……
……
```

​	

##### bitmaps

```
SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
GetBit(ctx context.Context, key string, offset int64) *IntCmd
BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
……
……
……
```



#### go-redis完成事务

假设我们有money = 100, 去消费，刚开开始out = 0

```go
package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	defer rdb.Close()

	rdb.Set(ctx, "money", 100, time.Hour*24*30)
	rdb.Set(ctx, "out", 0, time.Hour*24*30)

	//开启一个事务
	multi := rdb.TxPipeline()

	//消费20元
	dec := rdb.DecrBy(ctx, "money", 20)
	//执行失败取消事务
	if dec.Err() != nil {
		multi.Discard()
		return
	}
	inc := rdb.IncrBy(ctx, "out", 20)
	if inc.Err() != nil {
		multi.Discard()
		return
	}

  //执行事务
	multi.Exec(ctx)

	fmt.Println(rdb.Get(ctx, "money").Val())
	fmt.Println(rdb.Get(ctx, "out").Val())
}
```

输出：

```
80
20
```



当消费超出本金后，需要取消事务：

```go
package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	defer rdb.Close()

	rdb.Set(ctx, "money", 100, time.Hour*24*30)
	rdb.Set(ctx, "out", 0, time.Hour*24*30)

	//开启一个事务
	multi := rdb.TxPipeline()

	wantOut := 120
	money, _ := strconv.Atoi(rdb.Get(ctx, "out").Val())
	if money < wantOut {
		multi.Discard()
		return
	}

	//消费20元
	dec := rdb.DecrBy(ctx, "money", int64(wantOut))
	//执行失败取消事务
	if dec.Err() != nil {
		multi.Discard()
		return
	}
	inc := rdb.IncrBy(ctx, "out", int64(wantOut))
	if inc.Err() != nil {
		multi.Discard()
		return
	}

	multi.Exec(ctx)
}
```

事务会被取消



#### go-redis乐观锁

```go
ctx := context.Background()

	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get(ctx, "key").Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 这里可以处理业务
		v++

		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set(ctx, "key", v, 0)
			return nil
		})
		return err
	}

	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(ctx,fn, "key1", "key2", "key3")
	rdb.Watch(ctx, fn, "key")
```





# Redis进阶



## Redis.conf配置介绍

* #### 单位

  ```
  # Note on units: when memory size is needed, it is possible to specify
  # it in the usual form of 1k 5GB 4M and so forth:
  #
  # 1k => 1000 bytes
  # 1kb => 1024 bytes
  # 1m => 1000000 bytes
  # 1mb => 1024*1024 bytes
  # 1g => 1000000000 bytes
  # 1gb => 1024*1024*1024 bytes
  #
  # units are case insensitive so 1GB 1Gb 1gB are all the same.
  ```

  配置文件对大小写不敏感

  

* #### redis配置文件可以有多个导入（可以配置多个文件)

  ```
  ################################## INCLUDES ###################################
  
  # Include one or more other config files here.  This is useful if you
  # have a standard template that goes to all Redis servers but also need
  # to customize a few per-server settings.  Include files can include
  # other files, so use this wisely.
  #
  # Notice option "include" won't be rewritten by command "CONFIG REWRITE"
  # from admin or Redis Sentinel. Since Redis always uses the last processed
  # line as value of a configuration directive, you'd better put includes
  # at the beginning of this file to avoid overwriting config change at runtime.
  #
  # If instead you are interested in using includes to override configuration
  # options, it is better to use include as the last line.
  #
  # include /path/to/local.conf
  # include /path/to/other.conf
  ```

  

* #### 网络: 是否提供远程访问等

  ```
  ################################## NETWORK #####################################
  
  # By default, if no "bind" configuration directive is specified, Redis listens
  # for connections from all the network interfaces available on the server.
  # It is possible to listen to just one or multiple selected interfaces using
  # the "bind" configuration directive, followed by one or more IP addresses.
  #
  # Examples:
  #
  # bind 192.168.1.100 10.0.0.1
  # bind 127.0.0.1 ::1
  ```

  可以使用命令修改ip:

  ```sh
  bind 127.0.0.1 #绑定的ip
  protected-mode yes #受保护模式
  port 6379 #端口设置
  ```



* #### 通用配置

  ```
  ################################# GENERAL #####################################
  
  # By default Redis does not run as a daemon. Use 'yes' if you need it.
  # Note that Redis will write a pid file in /var/run/redis.pid when daemonized.
  ……
  ……
  ……
  ```

  相应的配置作用如下：

  ```sh
  daemonize yes # 以守护进程的方式运行，默认no，需要开启yes（允许后台运行）
  pidfile /var/run/redis_6379.pid #如果以后台的方式运行，我们需要指定一个pid进程文件
  loglevel notice #日志的级别
  logfile "" #日志的文件位置名
  databases 16 #默认的16个数据库，数据库的数量
  always-show-logo yes #是否显示logo
  ```

  

* #### 快照

  持久化，在规定的时间内，执行了多少次操作，则会持久化到文件==.rbd==，==.aof==

  redis是基于内存数据库，如果没有持久化，那么数据断电即失。

  ```sh
  save 900 1   #900s内，至少有一个key进行了修改，就进行持久化操作
  save 300 10
  save 60 10000
  
  stop-writes-on-bgsave-error yes  #持久化失败后，redis是否还要进行工作
  
  dbfilename dump.rdb #是否压缩rdb文件，需要消耗cpu资源
  
  rdbchecksum yes #保存reb文件的时候，进行检查修复rdb文件错误
  
  dir ./ #rdb文件保存位置
  
  
  ```

  

* #### 主从复制

  

* #### 安全

  密码，redis默认是没有密码的

  ```
  # requirepass foobared
  ```

  如果需要使用密码可直接在配置文件中进行修改：

  ```
  requirepass 123456
  ```

  ```sh
  # 使用命令设置密码
  config get requirepass #获取redis的密码
  config set requirepass "123456" #设置redis密码
  auth 123456 #使用密码登录，校验密码
  ```

* #### 客户端

  ```sh
  maxclients 10000 #设置能连接上redis的最大客户端的数量
  maxmemory <bytes> #redis配置最大的内存容量
  maxmemory-policy noeviction #内存到达上限之后的处理策略#移除一些过期的key#报错 
      redis.conf中的默认的过期策略是 volatile-lru
      maxmemory-policy 六种方式
      1、volatile-lru：只对设置了过期时间的key进行LRU（默认值） 
      2、allkeys-lru ： 删除lru算法的key   
      3、volatile-random：随机删除即将过期key   
      4、allkeys-random：随机删除   
      5、volatile-ttl ： 删除即将过期的   
      6、noeviction ： 永不过期，返回错误
  ```

  

* #### aof配置（持久化配置）

  ```sh
  appendonly no #默认不开启aof模式，默认是使用rdb方式持久化，在大部分所有的情况下，rdb完全够用
  appendfilename "appendonly.aof" #持久化的文件名字 
  appendfsync always # 每次修改都会sync，消耗性能！
  appendfsync everysec # 每秒执行一次sync，可能会丢失这1s的数据！
  appendfsync no # 不执行sync同步，这个时候操作系统自己同步数据，速度最快！
  ```

  





## 未完待续

……

## 参考文献

[B站up主-狂神说Java](https://www.bilibili.com/video/BV1S54y1R7SB/?spm_id_from=333.999.0.0)

[码神之路](https://www.mszlu.com/go/go-redis/02/02.html#_1-set)

