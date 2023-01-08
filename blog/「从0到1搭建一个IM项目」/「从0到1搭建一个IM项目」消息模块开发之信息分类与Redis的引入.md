[toc]



### 概况

前面我们已经完成了单聊模式的开发，并且已经完成了测试同时对信息发送流程进行了升级改造引入了udp，在本篇内容中，我们将开始完善聊天功能，将聊天类型划分为单聊和群聊分别来实现他们的信息转发；与此同时对应聊天系统一定是要有将聊天记录进行存储的，由于我们的应用是web服务，不能将其聊天记录存储到浏览器的cook中，这是不安全的，所以最后是将聊天记录存储到服务端，所以选择Redis来缓存聊天记录。

下面目录结构：

```
HiChat   
    ├── common    //放置公共文件
    |      |——md5.go
    |      |——resp.go
    │  
    ├── config    //做配置文件
    │  
    ├── dao//数据库crud
    │     |——user.go
    |     |——relation.go
    |     |——community.go
    |
    ├── global    //放置各种连接池，配置等
    │           |——global.go
    |
    ├── initialize  //项目初始化文件
    │              |——db.go
    |              |——logger.go
    |
    ├── middlewear  //放置web中间件
    |              |——jwt.go
    ├── models      //数据库表设计
    │           |——user_basic.go
    |           |——relation.go
    |           |——message.go
    |           |——community.go
    |
    ├── router           //路由
    |       |——router.go
    │   
    ├── service     //对外api
    |       |——user.go
    |       |——relation.go
    │   
    ├── test        //测试文件
    │  
    ├── main.go     //项目入口
    ├── go.mod            //项目依赖管理
    ├── go.sum            //项目依赖管理
```



### 单聊信息的收发及存储

在上一篇文章中，我们已经完成了消息收发以及udp连接，所以这里我们将重点讲解消息的分类，对单聊消息的处理和存储。

在udp服务端中调用了函数：```dispatch(data []byte)```

```go
func dispatch(data []byte) {
	//解析消息
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		zap.S().Info("消息解析失败", err)
		return
	}

	fmt.Println("解析数据:", msg, "msg.FormId", msg.FormId, "targetId:", msg.TargetId, "type:", msg.Type)

	//判断消息类型
	switch msg.Type {
	case 1: //私聊
		sendMsgAndSave(msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(uint(msg.FormId), uint(msg.TargetId), data)
	}
}
```



#### 消息的私发与存储

```go

//sendMsgTest 发送消息 并存储聊天记录到redis
func sendMsgAndSave(userId int64, msg []byte) {

  
	rwLocker.RLock()              //保证线程安全，上锁
	node, ok := clientMap[userId] //对方是否在线
	rwLocker.RUnlock()						//解锁

	jsonMsg := Message{}
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.FormId))

	
	if ok {
    //如果当前用户在线，将消息转发到当前用户的websocket连接中，然后进行存储
		node.DataQueue <- msg
	}

	//userIdStr和targetIdStr进行拼接唯一key
	var key string
	if userId > jsonMsg.FormId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}

	//创建记录
	res, err := global.RedisDB.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
    return
	}

	//将聊天记录写入redis缓存中
	score := float64(cap(res)) + 1
	ress, e := global.RedisDB.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
    return
	}
	fmt.Println(ress)
}

```





### 消息的群消息收发与存储

群发消息的逻辑：当群成员在当前群聊发送一条消息后，然后服务器将该群所有成员获取到，然后向除发送消息的用户外的所有群成员一一单发消息；简单的说就是：将群消息，给群成员都进行一次单聊信息的发送。

```go
func dispatch(data []byte) {
	//解析消息
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		zap.S().Info("消息解析失败", err)
		return
	}

	fmt.Println("解析数据:", msg, "msg.FormId", msg.FormId, "targetId:", msg.TargetId, "type:", msg.Type)

	//判断消息类型
	switch msg.Type {
	case 1: //私聊
		sendMsgAndSave(msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(uint(msg.FormId), uint(msg.TargetId), data)
	}
}
```



#### 群信息收发及存储

```go
//sendGroupMsg 群发
func sendGroupMsg(formId, target uint, data []byte) (int, error) {
	//群发的逻辑：1获取到群里所有用户，然后向除开自己的每一位用户发送消息
	userIDs, err := FindUsers(target)
	if err != nil {
		return -1, err
	}

	for _, userId := range *userIDs {
		if formId != userId {  //不能给当前发送消息的成员进行转发
      //调用单聊的函数，群聊变成了多次单聊
			sendMsgAndSave(int64(userId), data)
		}
	}
	return 0, nil
}
```

调用的```FindUsers(target)```在community.go中

```go
//FindUsers 获取群成员id
func FindUsers(groupId uint) (*[]uint, error) {
	relation := make([]Relation, 0)
	if tx := global.DB.Where("target_id = ? and type = 2", groupId).Find(&relation); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到成员信息")
	}

	userIDs := make([]uint, 0)
	for _, v := range relation {
		userId := v.OwnerId
		userIDs = append(userIDs, userId)
	}
	return &userIDs, nil
}
```

这样消息分类和聊天记录的存储就完成了





### 聊天记录的获取

对于聊天记录，不管是单聊还是群聊，目前为止都已经将其有标识的存储到缓存Redis中了，对于聊天记录我们只需要完成相应的api，当客户端进行请求时就可以请求到，对应的聊天记录

#### 聊天记录的获取

```go
//RedisMsg 获取缓存里面的聊天记录
func RedisMsg(userIdA int64, userIdB int64, start int64, end int64, isRev bool) []string {
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))

	//userIdStr和targetIdStr进行拼接唯一key
	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}

	var rels []string
	var err error
	if isRev {
		rels, err = global.RedisDB.ZRange(ctx, key, start, end).Result()
	} else {
		rels, err = global.RedisDB.ZRevRange(ctx, key, start, end).Result()
	}
	if err != nil {
		fmt.Println(err) //没有找到
	}
	return rels
}
```



#### 对外api的暴露

在service目录下relation.go中编写函数：

```go
func RedisMsg(c *gin.Context) {
	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	common.RespOKList(c.Writer, "ok", res)
}
```



#### 路由

在router目录下的router.go中：

```go
//聊天记录
	v1.POST("/user/redisMsg", service.RedisMsg).Use(middlewear.JWY())
```



### message完整代码

```go
package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"HiChat/global"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
)

type Message struct {
	Model
	FormId   int64  `json:"userId"`   //信息发送者
	TargetId int64  `json:"targetId"` //信息接收者
	Type     int    //聊天类型：群聊 私聊 广播
	Media    int    //信息类型：文字 图片 音频
	Content  string //消息内容
	Pic      string `json:"url"` //图片相关
	Url      string //文件相关
	Desc     string //文件描述
	Amount   int    //其他数据大小
}

func (m *Message) MsgTableName() string {
	return "message"
}

//Node 构造连接
type Node struct {
	Conn      *websocket.Conn //连接
	Addr      string          //客户端地址
	DataQueue chan []byte     //消息
	GroupSets set.Interface   //好友 / 群
}

//映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwLocker sync.RWMutex

//Chat	需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(w http.ResponseWriter, r *http.Request) {
	//1.  获取参数 并 检验 token 等合法性
	query := r.URL.Query()
	fmt.Println("handle:", query)
	Id := query.Get("userId")
	//token := query.Get("token")

	userId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}

	//升级为socket
	var isvalida = true
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取socket连接,构造消息节点
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//用户关系

	//将userId和Node绑定
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	fmt.Println("uid", userId)

	//发送接收消息
	//发送消息
	go sendProc(node)
	//接收消息
	go recProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Info("写入消息失败", err)
				return
			}
			fmt.Println("数据发送socket成功")
		}

	}
}

func recProc(node *Node) {
	for {
		//获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("读取消息失败", err)
			return
		}

		//dispatch(data)

		brodMsg(data)

		//这里是简单实现的一种方法
		//msg := Message{}
		//err = json.Unmarshal(data, &msg)
		//if err != nil {
		//	zap.S().Info("json解析失败", err)
		//	return
		//}
		//
		//if msg.Type == 1 {
		//	zap.S().Info("这是一条私信:", msg.Content)
		//	tarNode, ok := clientMap[msg.TargetId]
		//	if !ok {
		//		zap.S().Info("不存在对应的node", msg.TargetId)
		//		return
		//	}
		//
		//	tarNode.DataQueue <- data
		//	fmt.Println("发送成功：", string(data))
		//}

	}
}

var upSendChan chan []byte = make(chan []byte, 1024)

func brodMsg(data []byte) {
	upSendChan <- data

}

func init() {
	go UdpSendProc()
	go UpdRecProc()
}

//UdpSendProc 完成upd数据发送
func UdpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//192.168.31.147
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
		Zone: "",
	})
	if err != nil {
		zap.S().Info("拨号udp端口失败", err)
		return
	}

	defer udpConn.Close()

	for {
		select {
		case data := <-upSendChan:
			_, err := udpConn.Write(data)
			if err != nil {
				zap.S().Info("写入udp消息失败", err)
				return
			}
			fmt.Println("数据成功发送到udp服务端:", string(data))
		}
	}

}

//UpdRecProc 完成udp数据的接收
func UpdRecProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		zap.S().Info("监听udp端口失败", err)
		return
	}

	defer udpConn.Close()

	for {
		var buf [1024]byte
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			zap.S().Info("读取udp数据失败", err)
			return
		}

		fmt.Println("udp服务端接收udp数据", buf[0:n])

		//处理发送逻辑
		dispatch(buf[0:n])
	}
}

func dispatch(data []byte) {
	//解析消息
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		zap.S().Info("消息解析失败", err)
		return
	}

	fmt.Println("解析数据:", msg, "msg.FormId", msg.FormId, "targetId:", msg.TargetId, "type:", msg.Type)

	//判断消息类型
	switch msg.Type {
	case 1: //私聊
		sendMsgAndSave(msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(uint(msg.FormId), uint(msg.TargetId), data)
	}
}

//sendGroupMsg 群发
func sendGroupMsg(formId, target uint, data []byte) (int, error) {
	//群发的逻辑：1获取到群里所有用户，然后向除开自己的每一位用户发送消息
	userIDs, err := FindUsers(target)
	if err != nil {
		return -1, err
	}

	for _, userId := range *userIDs {
		if formId != userId {
			sendMsgAndSave(int64(userId), data)
		}
	}
	return 0, nil
}

//sendMs 向用户发送消息
func sendMsg(id int64, msg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[id]
	rwLocker.Unlock()

	if !ok {
		zap.S().Info("userID没有对应的node")
		return
	}

	zap.S().Info("targetID:", id, "node:", node)
	if ok {
		node.DataQueue <- msg
	}
}

//sendMsgTest 发送消息 并存储聊天记录到redis
func sendMsgAndSave(userId int64, msg []byte) {

	rwLocker.RLock()
	node, ok := clientMap[userId] //对方是否在线
	rwLocker.RUnlock()

	jsonMsg := Message{}
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.FormId))

	//如果不在线
	if ok {
		zap.S().Info(userId, "不在线， 没有对应的node,")

		node.DataQueue <- msg
	}

	//拼接记录名称
	var key string
	if userId > jsonMsg.FormId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}

	//创建记录
	res, err := global.RedisDB.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}

	//将聊天记录写入数据库
	score := float64(cap(res)) + 1
	ress, e := global.RedisDB.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(ress)
}

//MarshalBinary 需要重写此方法才能完整的msg转byte[]
func (msg Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(msg)
}

//RedisMsg 获取缓存里面的消息
func RedisMsg(userIdA int64, userIdB int64, start int64, end int64, isRev bool) []string {
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))

	//拼接key
	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}

	var rels []string
	var err error
	if isRev {
		rels, err = global.RedisDB.ZRange(ctx, key, start, end).Result()
	} else {
		rels, err = global.RedisDB.ZRevRange(ctx, key, start, end).Result()
	}
	if err != nil {
		fmt.Println(err) //没有找到
	}
	return rels
}
```





### 测试

#### 消息的发送与接收

消息发送与接收不管是单聊还是群聊测试方法都和《「从0到1搭建一个IM项目」信息模块开发之消息体的设计》中的测试是一样的，这里就不做过多介绍了。

#### 获取聊天记录

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2023-01-08%20%E4%B8%8B%E5%8D%887.48.05.png" style="zoom:40%;" />







### 总结

本篇中主要讲解了聊天分类，聊天记录存储和获取，IM系统的核心还是消息的收发过程，这个过程你可能需要多看代码，先把这个思路整理出来，我想这才是最重要的。到这里，我们的项目大部分功能就完成了，整个api的开发就完成，所有api:

```go
package router

import (
	"HiChat/middlewear"
	"HiChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("v1")
  
	//用户模块
	user := v1.Group("user")
	{
		user.GET("/list", middlewear.JWY(), service.List)
		user.POST("/login_pw", service.LoginByNameAndPassWord)
		user.POST("/new", middlewear.JWY(), service.NewUser)
		user.DELETE("/delete", middlewear.JWY(), service.DeleteUser)
		user.POST("/updata", middlewear.JWY(), service.UpdataUser)
		user.GET("/SendUserMsg", middlewear.JWY(), service.SendUserMsg)
	}

	//图片、语音模块
	upload := v1.Group("upload").Use(middlewear.JWY())
	{
		upload.POST("/image", service.Image)
	}

	//好友关系
	relation := v1.Group("relation").Use(middlewear.JWY())
	{
		relation.POST("/list", service.FriendList)
		relation.POST("/add", service.AddFriendByName)
		relation.POST("/new_group", service.NewGroup)
		relation.POST("/group_list", service.GroupList)
		relation.POST("/join_group", service.JoinGroup)
	}

	//聊天记录
	v1.POST("/user/redisMsg", service.RedisMsg).Use(middlewear.JWY())

	return router
}
```

在后续文章中，我们将项目参数使用viper进行配置化，然后加入前端最后进行部署。

