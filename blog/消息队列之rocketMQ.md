[TOC]



### 文章介绍

本文来简单介绍一下消息队列 ，这里将什么是MQ, 介绍RocketMQ的安装，RocketMQ的基本概念，消息类型，并使用go做各类消息的收发

### 什么是MQ

#### 1.什么是mq

消息队列是一种“先进先出”的数据结构


![queue1.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/1.png)



#### 2.应用场景

其应用场景主要包含以下3个方面

- 应用解耦


系统的耦合性越高，容错性就越低。以电商应用为例，用户创建订单后，如果耦合调用库存系统、物流系统、支付系统，任何一个子系统出了故障或者因为升级等原因暂时不可用，都会造成下单操作异常，影响用户使用体验。




![解耦1.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/2.png)
使用消息队列解耦合，系统的耦合性就会提高了。比如物流系统发生故障，需要几分钟才能来修复，在这段时间内，物流系统要处理的数据被缓存到消息队列中，用户的下单操作正常完成。当物流系统回复后，补充处理存在消息队列中的订单消息即可，终端系统感知不到物流系统发生过几分钟故障。


![解耦2.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/3.png)

- 流量削峰



![mq-5.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/4.png)
应用系统如果遇到系统请求流量的瞬间猛增，有可能会将系统压垮。有了消息队列可以将大量请求缓存起来，分散到很长一段时间处理，这样可以大大提到系统的稳定性和用户体验。

![mq-6.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/5.png)

一般情况，为了保证系统的稳定性，如果系统负载超过阈值，就会阻止用户请求，这会影响用户体验，而如果使用消息队列将请求缓存起来，等待系统处理完毕后通知用户下单完毕，这样总不能下单体验要好。

处于经济考量目的：

业务系统正常时段的QPS如果是1000，流量最高峰是10000，为了应对流量高峰配置高性能的服务器显然不划算，这时可以使用消息队列对峰值流量削峰

- 数据分发


![mq-1.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/6.png)


通过消息队列可以让数据在多个系统更加之间进行流通。数据的产生方不需要关心谁来使用数据，只需要将数据发送到消息队列，数据使用方直接在消息队列中直接获取数据即可。



#### MQ的优点和缺点


优点：解耦、削峰、数据分发![mq-2.png](https://mxshop-files.oss-cn-hangzhou.aliyuncs.com/28week/mq/7.png)


缺点包含以下几点：

- 系统可用性降低
  系统引入的外部依赖越多，系统稳定性越差。一旦MQ宕机，就会对业务造成影响。
  如何保证MQ的高可用？
- 系统复杂度提高
  MQ的加入大大增加了系统的复杂度，以前系统间是同步的远程调用，现在是通过MQ进行异步调用。
  如何保证消息没有被重复消费？怎么处理消息丢失情况？那么保证消息传递的顺序性？
- 一致性问题
  A系统处理完业务，通过MQ给B、C、D三个系统发消息数据，如果B系统、C系统处理成功，D系统处理失败。
  如何保证消息数据处理的一致性？

### RocketMQ的安装

使用docker安装

[docker安装RocketMQ](https://www.cnblogs.com/franson-2016/p/12714692.html)

### RocketMQ的基本概念

- Producer：消息的发送者；例如：发信人
- Consumer：消息接收者；例如：收信人
- Broker：暂存和传输消息；例如：邮局、中转站
- NameServer：管理Broker；例如：各个邮局的管理机构
- Topic：区分消息的种类；一个发送者可以发送消息给一个或者多个Topic；一个消息的接收者可以订阅一个或者多个Topic消息
- Message Queue：相当于是Topic的分区；用于并行发送和接收消息



### 消息类型

#### go实战

需要拉取

```shell
go get github.com/apache/rocketmq-client-go/v2
go get github.com/apache/rocketmq-client-go/v2/primitive
go get github.com/apache/rocketmq-client-go/v2/producer
```

这里我以实战的角度来介绍rocketMQ的消息类型：

#### 1. 普通消息

只是消息的收发，发送成功后接收者就直接可以收到消息

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	//初始化生产者
	q, err := rocketmq.NewProducer(producer.WithNameServer([]string{"101.1.12.202:9876"}))
	if err != nil {
		panic("生成q生产者失败")
	}
	if err := q.Start(); err != nil {
		panic("启动q生产者失败")
	}

	msg := []byte("您好呀， 我是ice_moss")
	mq := primitive.NewMessage("msg_test_hello", msg)  //msg_test_hello是为Topic

	res, err := q.SendSync(context.Background(), mq)
	if err != nil {
		fmt.Printf("发送失败%s", err)
	}
	fmt.Println("消息发送成功")
	fmt.Println(res.String())

	err = q.Shutdown()
	if err != nil {
		panic("shutdown fail err")
	}
}
```

这里需要注意的是如果我们需要在一个进程中启动多个```rocketmq.NewProducer()```就必须将他的第二个参数配置上：```producer.WithGroupName("sendMsg")```

```go
q, err := rocketmq.NewProducer(producer.WithNameServer([]string{"101.1.12.202:9876"}), producer.WithGroupName("sendMsg"))
```

不然就会报：**生产者组已经被创建**

原因：我们没有不设置```WithGroupName```在调用时，会自动为我们创建一个默认名称的```WithGroupName```，当第二次```rocketmq.NewProducer```仍然是默认名，这时整个```GroupName```就冲突了

好了已经将"普通消息"发送到队列中了，现在我们来接收



#### 2. 消费消息

**注意：两端的Topic必须保持一直**

```go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		//接收者组
		consumer.WithGroupName("msg_test"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"101.1.12.202:9876"})),
	)
	//订阅消息
	err := c.Subscribe("msg_test_hello", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
  
  //程序运行2分钟
	time.Sleep(time.Second * 120)
  
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
```

输出：

```shell
subscribe callback: [Message=[topic=msg_test_hello, body=您好呀， 我是ice_moss, Flag=0, properties=map[CONSUME_START_TIME:1668255347270 MAX_OFFSET:2 MIN_OFFSET:0 UNIQ_K251664A6E000000003cf040100001], TransactionId=], MsgId=0A0251664A6E000000003cf040100001, OffsetMsgId=010EB4CA00002A9F000000000004BC14,QueueId=1, StoreSize=174, QueueOffset=0, SysFlag=0, BornTimestamp=1668254378888, BornHost=112.21.20.248:43010, StoreTimestamp=1668254379066, StoreHost=101.1.12.202:10911, CommitLogOffset=310292, BodyCRC=1573027761, ReconsumeTimes=0, PreparedTransactionOffset=0] 

```



#### 3. 延时消息

延时消息，指我们将我们需要发送的发送消息延迟多少时间后接收方才能收到，其中一个应用场景就是分布式电商系统的**下单——>支付**， 例如：12306官网买车票，当我们购买一张车票后，后台会做车票库存扣减，但是如果我们只下单，不支付这就很要命，该买票的人买不到票，该卖出去的票没有卖出去；其实仔细一点就会发现，12306购买下单后，在规定时间没有完成支付，就会取消相应的订单， 然后做库存归还。

现在来看一下延迟消息怎么发送：

```go
package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

//SendMessage 生成消息，延迟消息
func SendMessage(q rocketmq.Producer) {
	if err := q.Start(); err != nil {
		panic("启动q生产者失败")
	}

	msg := primitive.NewMessage("msg_test_hello", []byte("这是一个延迟消息"))

	//延迟时间级别
	//messageDelayLevel=1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	msg.WithDelayTimeLevel(3)  //10s

	res, err := q.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("发送失败%s", err)
	}
	err = q.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
	fmt.Println(res.String())
}

func main() {
	//初始化生产者
	q, err := rocketmq.NewProducer(producer.WithNameServer([]string{"101.1.12.202:9876"}))
	if err != nil {
		panic("生成q生产者失败")
	}
	SendMessage(q)
}
```

10秒后：

```shell
subscribe callback: [Message=[topic=msg_test_hello, body=这是一个延迟消息, Flag=0, properties=map[CONSUME_START_TIME:1668256662984 DELAY:3 MAX_OFFSET:5 MIN_OFFSET:0 REAREAL_TOPIC:msg_test_hello UNIQ_KEY:0A0251664BE9000000003d12f2e00001]………
```





#### 4.事务消息

##### 什么是事务

事务是指是程序中一系列严密的逻辑操作，而且所有操作必须全部成功完成，否则在每个操作中所作的所有更改都会被撤消。可以通俗理解为：就是把多件事情当做一件事情来处理，好比大家同在一条船上，要活一起活，要完一起完 

**事物的四个特性（ACID）**

 **● 原子性**（Atomicity）**：**操作这些指令时，要么全部执行成功，要么全部不执行。只要其中一个指令执行失败，所有的指令都执行失败，数据进行回滚，回到执行指令前的数据状态。

>  **eg：**拿转账来说，假设用户A和用户B两者的钱加起来一共是20000，那么不管A和B之间如何转账，转几次账，事务结束后两个用户的钱相加起来应该还得是20000，这就是事务的一致性。 

 **● 一致性**（Consistency）**：**事务的执行使数据从一个状态转换为另一个状态，但是对于整个数据的完整性保持稳定。

 **● 隔离性**（Isolation）**：**隔离性是当多个用户并发访问数据库时，比如操作同一张表时，数据库为每一个用户开启的事务，不能被其他事务的操作所干扰，多个并发事务之间要相互隔离，可以使用锁机制来实现隔离，其实就是将并发场景下对数据操作的部分对并发请求进行串行化。

 **● 持久性**（Durability）**：**当事务正确完成后，它对于数据的改变是永久性的。

>  **eg：** 例如我们在使用JDBC操作数据库时，在提交事务方法后，提示用户事务操作完成，当我们程序执行完成直到看到提示后，就可以认定事务以及正确提交，即使这时候数据库出现了问题，也必须要将我们的事务完全执行完成，否则就会造成我们看到提示事务处理完毕，但是数据库因为故障而没有执行事务的重大错误。



##### MQ的事务消息

这里的事务消息实现接口：

```go
type TransactionListener interface {
	//  When send transactional prepare(half) message succeed, this method will be invoked to execute local transaction.
	ExecuteLocalTransaction(*Message) LocalTransactionState

	// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
	// method will be invoked to get local transaction status.
	CheckLocalTransaction(*MessageExt) LocalTransactionState
}
```

我们的业务代码需要放在```ExecuteLocalTransaction(*Message) LocalTransactionState```方法中执行，对应返回相应的状态

```go
const (
	CommitMessageState LocalTransactionState = iota + 1   //返回状态：事务执行成功发现消息
	RollbackMessageState          												//返回状态：进行事务回查
	UnknowState																						//仍然会回查
)
```

我们回查机制业务需要在```CheckLocalTransaction(*MessageExt) LocalTransactionState```方法中完成

下面我们来实现该接口(创建订单场景下)：

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"google.golang.org/grpc/codes"
)

//Order 模拟订单
type Order struct {
	OrderSrvID string
	UserID     int32
	GoodsID    int32
	TotalPrice float32
	Post       string
	Address    string
	Mobile     string
}

//OrderLister 接口实现者，事务可以将一下配置\信息写入该结构体中
type OrderLister struct {
	Code codes.Code      //返回状态码
	Ctx  context.Context //上下文数据
	ID   int32           //订单id
}

//ExecuteLocalTransaction  When send transactional prepare(half) message succeed, this method will be invoked to execute local transaction.
func (o *OrderLister) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {

	//执行本地业务逻辑
	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second * 3)

	orderInfo := Order{}
	err := json.Unmarshal(msg.Body, &orderInfo)
	if err != nil {
		o.Code = codes.Unavailable
		log.Fatal("解析失败:", err)

		//调用回查逻辑
		return primitive.RollbackMessageState
	}

	fmt.Println("订单信息：", orderInfo)

	fmt.Println("本地逻辑执行成功")

	//CommitMessageState 提交信息至mq
	//CommitMessageState/RollbackMessageState都不会回查
	return primitive.CommitMessageState
}

//CheckLocalTransaction When no response to prepare(half) message. broker will send check message to check the transaction status, and this method will be invoked to get local transaction status.
func (o *OrderLister) CheckLocalTransaction(*primitive.MessageExt) primitive.LocalTransactionState {
	//回查
	fmt.Println("事务未通过，开始回查")
	return primitive.RollbackMessageState
}

func (o *Order) CreateOrder(q rocketmq.TransactionProducer) {
	order, err := json.Marshal(o)
	if err != nil {
		panic("marshal fail")
	}

	msg := primitive.NewMessage("msg_test_order", order)
	res, err := q.SendMessageInTransaction(context.Background(), msg)
	if err != nil {
		fmt.Printf("发送失败%s", err)
	} else {
		fmt.Println("发送成功", res.String())
	}

	time.Sleep(time.Hour)

	err = q.Shutdown()
	if err != nil {
		panic("shutdown fail err")
	}
}

func main() {

	//初始化事务对象
	orderLister := &OrderLister{}
	q, err := rocketmq.NewTransactionProducer(orderLister,
		producer.WithNameServer([]string{"1.14.180.202:9876"}), producer.WithGroupName("msg_test"))
	if err != nil {
		panic("生成q生产者失败")
	}

	if err = q.Start(); err != nil {
		panic("启动q生产者失败")
	}
	orderInfo := &Order{
		OrderSrvID: "343435",
		UserID:     21,
		GoodsID:    214,
		TotalPrice: 150.5,
		Post:       "请尽快发货",
		Address:    "无锡市",
		Mobile:     "18389202834",
	}

	orderInfo.CreateOrder(q)
}
```

执行输出：

```shell
开始执行本地逻辑
订单信息： {343435 21 214 150.5 请尽快发货 无锡市 18389202834}
本地逻辑执行成功
发送成功 SendResult [sendStatus=0, msgIds=0A0266DB4E24000000003da94f100001, offsetMsgId=010EB4CA00002A9F000000000004C28F, queueOffset=364, messageQueue=MessageQueue [tomsg_test_order, brokerName=broker-a, queueId=1]]
```

接收者接收到：

```shell
subscribe callback: [Message=[topic=msg_test_order, body={"OrderSrvID":"343435","UserID":21,"GoodsID":214,"TotalPrice":150.5,"Post":"请尽快发货","Address":"无锡市","Mob389202834"}, Flag=0, properties=map[CONSUME_START_TIME:1668266524665 MAX_OFFSET:1 MIN_OFFSET:0 PGROUP:msg_test REAL_QID:1 REAL_TOPIC:msg_test_order TRAN_MSG:true UNIQ_KEY:0A0266DB4E24000000003da94f100001], TransactionId=0A0266DB4E24000000003da94f100001], MsgId=0A0266DB4E24000000003da94f100001…………
```















