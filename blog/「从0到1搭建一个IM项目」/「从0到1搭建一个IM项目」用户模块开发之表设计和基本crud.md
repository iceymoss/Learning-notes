[toc]



### 概况

在上一篇我们完成了项目的初始化以及一些基本的配置，目录如下：

```
HiChat   
    ├── common    //放置公共文件
    │  
    ├── config    //做配置文件
    │  
    ├── dao       //数据库crud
    │  
    ├── global    //放置各种连接池，配置等
    │   		|——global.go
    |
    ├── initialize  //项目初始化文件
    │  			|——db.go
    |				|——logger.go
    |
    ├── middlewear  //放置web中间件
    │ 
    ├── models      //数据库表设计
    │   
    ├── router   		//路由
    │   
    ├── service     //对外api
    │   
    ├── test        //测试文件
    │  
    ├── main.go     //项目入口
    ├── go.mod			//项目依赖管理
    ├── go.sum			//项目依赖管理
```



### 用户表设计

#### 表分析

这里我们先来探讨一下需求：

>唯一标识id（主键)
>
>创建时间
>
>更新时间
>
>删除时间

>用户名
>
>密码
>
>头像
>
>性别
>
>电话号码
>
>邮件
>
>标识
>
>登录设备ip和端口
>
>密码加密盐值
>
>登录时间
>
>离线时间
>
>心跳时间
>
>是否在线
>
>登录设备描述



#### 代码实现

分析完成后，我们就来实现这些字段，在models目录下新建user_basic.go

```go
package models

import (
	"gorm.io/gorm"

	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserBasic struct {
	Model
	Name          string
	PassWord      string
	Avatar        string
	Gender        string `gorm:"column:gender;default:male;type:varchar(6) comment 'male表示男， famale表示女'"` //gorm为数据库字段约束
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`   //valid为条件约束
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string `valid:"ipv4"`
	ClientPort    string
	Salt          string     //盐值
	LoginTime     *time.Time `gorm:"column:login_time"`
	HeartBeatTime *time.Time `gorm:"column:heart_beat_time"`
	LoginOutTime  *time.Time `gorm:"column:login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string //登录设备
}

//UserTableName 指定表的名称
func (table *UserBasic) UserTableName() string {
	return "user_basic"
}
```



#### 用户表生成

在test目录下新建一个mai.go文件

```go
package main

import (
	"HiChat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/hi_chat?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	err = db.AutoMigrate(&models.User_basic{})
	if err != nil {
		panic(err)
	}
}
```

可以看到：

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-28%20%E4%B8%8B%E5%8D%884.08.26.png" style="zoom:50%;" />



### 用户crud

下面来分析一下有哪些功能模块，

1. 用户列表
2. 用户名\密码查询
3. 根据用户名查询用户
4. 根据id查询用户
5. 根据电话查询用户
6. 根据邮件查询用户
7. 新建用户
8. 更新用户
9. 删除用户

在dao目录下新建一个user.go文件，下面我们来一一实现这些功能



#### 用户列表

```go
//GetUserList 获取用户列表， 直接在用户
func GetUserList() ([]*models.UserBasic, error) {
	var list []*models.UserBasic
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("获取用户列表失败")
	}
	return list, nil
}
```



#### 用户名\密码查询

```go
//FindUserByNameAndPwd 昵称和密码查询
func FindUserByNameAndPwd(name string, password string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ? and pass_word=?", name, password).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
  
	//登录识别
	t := strconv.Itoa(int(time.Now().Unix()))
  
  //这个md5后续会讲解， 可以先注释
	temp := common.Md5encoder(t)
  
	if tx := global.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp); tx.RowsAffected == 0 {
		return nil, errors.New("写入identity失败")
	}
	return &user, nil
}
```



#### 根据用户名查询用户

1. 登录时使用

```go
//FindUserByName 根据name查询
func FindUserByName(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("没有查询到")
	}
	return &user, nil
}
```



1. 用户注册时使用

```go

func FindUser(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.RowsAffected == 1 {
		return nil, errors.New("当前用户名已存在")
	}
	return &user, nil
}
```



#### 根据id查询用户

```go
func FindUserID(ID uint) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where(ID).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}
```



#### 根据电话查询用户

```go
func FindUserByPhone(phone string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("phone = ?", phone).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}
```



#### 根据邮件查询用户

```go
func FindUerByEmail(email string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("email = ?", email).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}
```



#### 新建用户

```go
//CreateUser 新建用户
func CreateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Create(&user)
	if tx.RowsAffected == 0 {
		zap.S().Info("新建用户失败")
		return nil, errors.New("新增用户失败")
	}
	return &user, nil
}
```



#### 更新用户

```go
func UpdateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Model(&user).Updates(models.UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Salt:     user.Salt,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("更新用户失败")
		return nil, errors.New("更新用户失败")
	}
	return &user, nil
}
```



#### 删除用户

```go
func DeleteUser(user models.UserBasic) error {
	if tx := global.DB.Delete(&user); tx.RowsAffected == 0 {
		zap.S().Info("删除失败")
		return errors.New("删除用户失败")
	}
	return nil
}
```



#### user.go完整代码

```go
package dao

import (
	"errors"
	"strconv"
	"time"

	"HiChat/common"
	"HiChat/global"
	"HiChat/models"

	"go.uber.org/zap"
)

func GetUserList() ([]*models.UserBasic, error) {
	var list []*models.UserBasic
	if tx := global.DB.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("获取用户列表失败")
	}
	return list, nil
}

//查询用户:根据昵称，根据电话，根据邮件

//FindUserByNameAndPwd 昵称和密码查询
func FindUserByNameAndPwd(name string, password string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ? and pass_word=?", name, password).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	//token加密
	t := strconv.Itoa(int(time.Now().Unix()))
	temp := common.Md5encoder(t)
	if tx := global.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp); tx.RowsAffected == 0 {
		return nil, errors.New("写入identity失败")
	}
	return &user, nil
}

func FindUserByName(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("没有查询到记录")
	}
	return &user, nil
}

func FindUser(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.RowsAffected == 1 {
		return nil, errors.New("当前用户名已存在")
	}
	return &user, nil
}

func FindUserByPhone(phone string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("phone = ?", phone).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

func FindUerByEmail(email string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("email = ?", email).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

func FindUserID(ID uint) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where(ID).First(&user); tx.RowsAffected == 0 {
		return nil, errors.New("未查询到记录")
	}
	return &user, nil
}

//CreateUser 新建用户
func CreateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Create(&user)
	if tx.RowsAffected == 0 {
		zap.S().Info("新建用户失败")
		return nil, errors.New("新增用户失败")
	}
	return &user, nil
}

func UpdateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Model(&user).Updates(models.UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Salt:     user.Salt,
	})
	if tx.RowsAffected == 0 {
		zap.S().Info("更新用户失败")
		return nil, errors.New("更新用户失败")
	}
	return &user, nil
}

func DeleteUser(user models.UserBasic) error {
	if tx := global.DB.Delete(&user); tx.RowsAffected == 0 {
		zap.S().Info("删除失败")
		return errors.New("删除用户失败")
	}
	return nil
}

```



### 总结

现在我们就将用户模块dao层的方法写完了，其实都很简单，就是使用gorm堆数据库的增删改查，有不足的欢迎小伙伴们指正，欢迎小伙伴们在评论区讨论。