[toc]



### 文章介绍

在篇内容介绍基于golang的gorm,这里我将简单介绍如何安装，连接数据库(以MySQL为例)，以及基本的curd操作

### 安装

[gorm 官方文档](https://gorm.io/zh_CN/docs/index.html)

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

### 模型定义

模型是标准的 struct，由 Go 的基本数据类型、实现了 [Scanner](https://pkg.go.dev/database/sql/?tab=doc#Scanner) 和 [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) 接口的自定义类型及其指针或别名组成，我们最后的表名就是结构体的和结构体名一致，当然字段也一致的

例如：

```go
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
```

GORM 倾向于约定，而不是配置。默认情况下，GORM 使用 `ID` 作为主键，使用结构体名的 `蛇形复数` 作为表名，字段名的 `蛇形` 作为列名，并使用 `CreatedAt`、`UpdatedAt` 字段追踪创建、更新时间

当然，遵循 GORM 已有的约定，可以减少您的配置和代码量。如果约定不符合您的需求，GORM 允许您自定义配置它们

### 自定义配置

##### 使用 `ID` 作为主键

默认情况下，GORM 会使用 `ID` 作为表的主键。

```go
type User struct {
  ID   string // 默认情况下，名为 `ID` 的字段会作为表的主键
  Name string
}
```

你可以通过标签 `primaryKey` 将其它字段设为主键

```go
// 将 `UUID` 设为主键
type Animal struct {
  ID     int64
  UUID   string `gorm:"primaryKey"`
  Name   string
  Age    int64
}
```

此外，您还可以看看 [复合主键](https://gorm.io/zh_CN/docs/composite_primary_key.html)

### gorm.Model

GORM 定义一个 `gorm.Model` 结构体，其包括字段 `ID`、`CreatedAt`、`UpdatedAt`、`DeletedAt`

```go
// gorm.Model 的定义
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

您可以将它嵌入到您的结构体中，以包含这几个字段

### 连接数据库

这里我们以mysql为例：

```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  //注意：pass为MySQL数据库的管理员密码，dbname为要连接的数据库
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

MySQl 驱动程序提供了 [一些高级配置](https://github.com/go-gorm/mysql) 可以在初始化过程中使用

### 建表

这里我们先连接数据库，为了方便查看SQL语句，我们将记入日志

```go
package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//定义表结构
type Producttest struct {
	gorm.Model
	Name  string
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	//用于输出使用的sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	//打开mysql服务中对应的数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	//AutoMigrate 为给定模式运行自动迁移,建立Product类型的数据表
	err = db.AutoMigrate(&Producttest{})
	if err != nil {
		log.Fatal("建表失败", err)
	}
```

我们看到建表成功：

mysql> show tables;
+---------------------+
| Tables_in_gorm_test |
+---------------------+
| dities              |
| products            |
| producttests        |
| subways             |
| test                |
| test1               |
| user2               |
| user_infos          |
| users               |
+---------------------+
9 rows in set (0.00 sec)

### 快速入门

##### 新增数据

这里以结构体的方式插入数据

```go
	// Create 新增数据
	db.Create(&Producttest{Code: "01", Name: "golang程序设计", Price: 100})
	db.Create(&Producttest{Code: "02", Name: "python入门", Price: 200})
```

然后我们可以看到输出结果:

```sh

2022/06/23 09:31:03 /Users/feng/go/src/GormStart/ch01/main.go:51
[16.193ms] [rows:1] INSERT INTO `producttests` (`created_at`,`updated_at`,`deleted_at`,`name`,`code`,`price`) VALUES ('2022-06-23 09:31:03.35','2022-06-23 09:31:03.35',NULL,'golang程序设计','01',100)

2022/06/23 09:31:03 /Users/feng/go/src/GormStart/ch01/main.go:52
[2.331ms] [rows:1] INSERT INTO `producttests` (`created_at`,`updated_at`,`deleted_at`,`name`,`code`,`price`) VALUES ('2022-06-23 09:31:03.364','2022-06-23 09:31:03.364',NULL,'python入门','02',200)
Process 96312 has exited with status 0
Detaching
dlv dap (96294) exited with code: 0
```

这样我们就将两条数据插入了数据库中



##### 查找数据

```go
var product Producttest  //需要实例化一个表结构
	db.First(&product, 2) // 根据整型主键查找
	fmt.Println(product.Name)
	db.First(&product, "code = ?", "02") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)
```

查询结果：

```shell
[3.323ms] [rows:1] SELECT * FROM `producttests` WHERE `producttests`.`id` = 2 AND `producttests`.`deleted_at` IS NULL ORDER BY `producttests`.`id` LIMIT 1
python入门

[1.002ms] [rows:1] SELECT * FROM `producttests` WHERE code = '02' AND `producttests`.`deleted_at` IS NULL AND `producttests`.`id` = 2 ORDER BY `producttests`.`id` LIMIT 1
{{2 2022-06-23 09:31:03.364 +0800 CST 2022-06-23 09:31:03.364 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}} python入门 02 200}
```

##### 更新数据

```go
db.Model(&Producttest{}).Where("Code = ?", "01").Update("Price", "50")
// UPDATE users SET Price='50', updated_at='2022-06-23 10:00:23.654' WHERE Code='01';
```

输出结果:

```sh
[9.204ms] [rows:0] UPDATE `producttests` SET `price`='50',`updated_at`='2022-06-23 10:05:49.719' WHERE Code = '01' AND `producttests`.`deleted_at` IS NULL
```

##### 删除数据

```go
var product Producttest
db.Delete(&product, 1) //删除主码为1的数据
```



### 新增数据(C)

我们先定义表结构

```go
type User struct {
	Name     string
	Age      uint
	Birthday time.Time
  Addr     string
  Work     string
}
```

连接数据库：

```go
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	//用于输出使用的sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	//打开mysql服务中对应的数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
  
  	//建立表
	_ = db.AutoMigrate(&User{})
```

建表成功：

```sql
[18.454ms] [rows:0] CREATE TABLE `users` (`name` longtext,`age` bigint unsigned,`birthday` datetime(3) NULL,`addr` longtext,`work` longtext)
```

新增数据项：

```go
//新增数据项
	var Age uint = 18
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		_ = db.Create(&User{Name: fmt.Sprintf("小杨%d", i), Age: Age, Birthday: time.Now()})
		Age++
	}
```

我们可以看到:

```sql

2022/06/23 10:35:28 /Users/feng/go/src/GormStart/ch03/main.go:73
[6.255ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`,`addr`,`work`) VALUES ('小杨0',18,'2022-06-23 10:35:28.267','','')

2022/06/23 10:35:30 /Users/feng/go/src/GormStart/ch03/main.go:73
[4.036ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`,`addr`,`work`) VALUES ('小杨1',19,'2022-06-23 10:35:30.275','','')

2022/06/23 10:35:32 /Users/feng/go/src/GormStart/ch03/main.go:73
[5.979ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`,`addr`,`work`) VALUES ('小杨2',20,'2022-06-23 10:35:32.284','','')

2022/06/23 10:35:34 /Users/feng/go/src/GormStart/ch03/main.go:73
[4.307ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`,`addr`,`work`) VALUES ('小杨3',21,'2022-06-23 10:35:34.292','','')

2022/06/23 10:35:36 /Users/feng/go/src/GormStart/ch03/main.go:73
[3.353ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`,`addr`,`work`) VALUES ('小杨4',22,'2022-06-23 10:35:36.296','','')
Process 97602 has exited with status 0
```

###### 单值插入：

当然我们可以使用：db.create()

```go
db.Create(&User{Name: "小李", Age: 20, Birthday: time.Now(), Addr: "北京", Work: "程序员"})
```

即：

```sql
[2.769ms] [rows:1] INSERT INTO `users` (`name`,`age`,`birthday`,`addr`,`work`) VALUES ('小李',20,'2022-06-23 10:42:39.84','北京','程序员')
```

###### 批量插入

现在我们修改一下表结构

```go
type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Addr         string
	Work         string
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
```

将原来的users表删除

然后建立表：

```go
//建立表
	_ = db.AutoMigrate(&User{})
```

```go
//批量插入数据
	var users = []User{{Name: "小杨"}, {Name: "小张"}, {Name: "小李"}, {Name: "小冯"}}
	db.Create(&users)
	for _, u := range users {
		fmt.Println(u.Name)
	}
```

输出结果：

```sql
[2.378ms] [rows:4] INSERT INTO `users` (`name`,`email`,`age`,`addr`,`work`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('小杨',NULL,0,'','',NULL,NULL,NULL,'2022-06-23 10:57:56.088','2022-06-23 10:57:56.088'),('小张',NULL,0,'','',NULL,NULL,NULL,'2022-06-23 10:57:56.088','2022-06-23 10:57:56.088'),('小李',NULL,0,'','',NULL,NULL,NULL,'2022-06-23 10:57:56.088','2022-06-23 10:57:56.088'),('小冯',NULL,0,'','',NULL,NULL,NULL,'2022-06-23 10:57:56.088','2022-06-23 10:57:56.088')
小杨
小张
小李
小冯
```

我们还可以这样做：

```go
db.Model(&User{}).Create(map[string]interface{}{
		"Name": "小刚", "Age": 25, "Addr": "广州",
	})
```



### 查找数据(R)

###### 按排序找找

1. 升序

```go
//检索单个数据，升序
	var user User
	_ = db.First(&user)
	fmt.Println(user)
```

输出：

```sql
[1.218ms] [rows:1] SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
{1 小杨 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
```

2. 降序

```go
//降序
	var user User
	_ = db.Last(&user)
	fmt.Println(user)
```



###### 按照位置查找

```go
//按数据表中的位置
	var user User
	_ = db.Take(&user, 2)
	fmt.Println(user)
```

输出：

```sql
[1.653ms] [rows:1] SELECT * FROM `users` WHERE `users`.`id` = 2 LIMIT 1
{2 小张 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
```

###### 按照主键查找

   ```go
	//通过主键查询
	var user User
	result := db.First(&user, 3)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("数据未找到")
	}
	fmt.Println(user)
   ```

###### 查找表中所有数据

```go
//检索全部对象
	var users []User
	result := db.Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(result.RowsAffected)
```

输出：

```sql
[1.553ms] [rows:5] SELECT * FROM `users`
{1 小杨 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
{2 小张 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
{3 小李 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
{4 小冯 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
{5 小刚 <nil> 25 广州  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
5
```



###### 根据条件查找

1. ***匹配一条数据：First()***

这里db.Where().First()只会匹配一条数据

```go
//根据条件检索
	var users []User
	//匹配一条数据
	db.Where("name= ?", "小杨").First(&users)
	for _, user := range users {
		fmt.Println(user)
	}
```

输出：

```sql
[2.400ms] [rows:1] SELECT * FROM `users` WHERE name= '小杨' ORDER BY `users`.`id` LIMIT 1
{1 小杨 <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 10:57:56.088 +0800 CST}
```

2. ***匹配多条数据：Find()***

```go
	db.Where("Addr = ?", "北京").Find((&users))
	for _, user := range users {
		fmt.Println(user)
	}
```

输出：

```sql
[2.133ms] [rows:2] SELECT * FROM `users` WHERE Addr = '北京'
{6 小熊 <nil> 0 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:30:17.018 +0800 CST 2022-06-23 11:30:17.018 +0800 CST}
{8 小张 <nil> 0 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:33:21.414 +0800 CST 2022-06-23 11:33:21.414 +0800 CST}
```



***注意：当我们不知道sql表中字段名的时候可以直接使用结构体，这样可以直接屏蔽数据表底层逻辑，这样我们就可以不用关心数据表的结构了***

例如：

```go
//此方法可屏蔽底层SQL数据表字段
	var user User
	db.Where(&User{Name: "小杨"}).First(&user)
	fmt.Println(user)
```



3. ***根据条件```IN```检索***

   

```go
//根据条件检索 IN 
//查找age等于18、19、20的数据
	var users []User
	db.Where("age IN ?", []uint{18, 19, 20}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
```

输出：

```sql
[0.917ms] [rows:3] SELECT * FROM `users` WHERE age IN (18,19,20)
{9 小周 <nil> 18 上海  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.759 +0800 CST 2022-06-23 11:49:05.759 +0800 CST}
{10 小周 <nil> 19 上海  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.763 +0800 CST 2022-06-23 11:49:05.763 +0800 CST}
{11 小周 <nil> 20 上海  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.77 +0800 CST 2022-06-23 11:49:05.77 +0800 CST}
```



4. ***根据条件``` AND```检索***

```go
//根据条件检索 AND
//查找地址在北京并且大于等于18岁的人
	var users []User
	db.Where("addr=? AND age>=?", "北京", 18).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
```

输出：

```go
[2.181ms] [rows:3] SELECT * FROM `users` WHERE addr='北京' AND age>=18
{12 小董 <nil> 23 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.773 +0800 CST 2022-06-23 11:49:05.773 +0800 CST}
{13 小周 <nil> 18 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.964 +0800 CST 2022-06-23 11:54:55.964 +0800 CST}
{16 小画 <nil> 23 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.973 +0800 CST 2022-06-23 11:54:55.973 +0800 CST}
```



根据条件```OR``` 检索

```go
var users []User
db.Where("addr=? OR age>=?", "北京", 20).Find(&users)
for _, user := range users {
	fmt.Println(user)
}
```

输出：

```sql
[1.749ms] [rows:8] SELECT * FROM `users` WHERE addr='北京' OR age>=20
{5 小刚 <nil> 25 广州  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
{6 小熊 <nil> 0 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:30:17.018 +0800 CST 2022-06-23 11:30:17.018 +0800 CST}
{8 小张 <nil> 0 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:33:21.414 +0800 CST 2022-06-23 11:33:21.414 +0800 CST}
{11 小周 <nil> 20 上海  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.77 +0800 CST 2022-06-23 11:49:05.77 +0800 CST}
{12 小董 <nil> 23 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.773 +0800 CST 2022-06-23 11:49:05.773 +0800 CST}
{13 小周 <nil> 18 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.964 +0800 CST 2022-06-23 11:54:55.964 +0800 CST}
{15 小杨 <nil> 20 上海  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.971 +0800 CST 2022-06-23 11:54:55.971 +0800 CST}
{16 小画 <nil> 23 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.973 +0800 CST 2022-06-23 11:54:55.973 +0800 CST}
```



***使用struct & map***

```go
//使用struct
	var users []User
	db.Where(&User{Name: "小周", Age: 18, Addr: "北京"}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
```

输出：

```sql
[2.005ms] [rows:1] SELECT * FROM `users` WHERE `users`.`name` = '小周' AND `users`.`age` = 18 AND `users`.`addr` = '北京'
{13 小周 <nil> 18 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.964 +0800 CST 2022-06-23 11:54:55.964 +0800 CST}
```



```go
	//使用map
	var users []User
	db.Where(map[string]interface{}{"name": "小周", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	for _, user := range users {
		fmt.Println(user)
	}
```

输出：

```sql
[4.219ms] [rows:2] SELECT * FROM `users` WHERE `age` = 18 AND `name` = '小周'
{9 小周 <nil> 18 上海  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:49:05.759 +0800 CST 2022-06-23 11:49:05.759 +0800 CST}
{13 小周 <nil> 18 北京  <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 11:54:55.964 +0800 CST 2022-06-23 11:54:55.964 +0800 CST}
```



### 更新数据(U)

###### 保存所有字段

`Save` 会保存所有的字段，即使字段是零值

```go
//通过save方法更新
	var user User
	_ = db.First(&user)
	fmt.Println(user)

	user.Name = "小旷"
	user.Age = 22
	user.ID = 17
	user.Addr = "深圳"
	user.Work = "go开发工程师&gis开发工程师"
	_ = db.Save(&user)
	fmt.Println(user)
```

输出结果：

```sql
[rows:1] INSERT INTO `users` (`name`,`email`,`age`,`addr`,`work`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`id`) VALUES ('小旷',NULL,22,'深圳','go开发工程师&gis开发工程师',NULL,NULL,NULL,'2022-06-23 10:57:56.088','2022-06-23 13:28:10.252',17)
{17 小旷 <nil> 22 深圳 go开发工程师&gis开发工程师 <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 2022-06-23 10:57:56.088 +0800 CST 2022-06-23 13:28:10.252 +0800 CST}
```

###### 更新指定字段

```go
//通过指定字段更新
//将age等于0的更新为20
	var user User
	db.Model(&user).Where("age", 0).Update("age", 20)
```

输出结果:

```sql
[rows:7] UPDATE `users` SET `age`=20,`updated_at`='2022-06-23 13:33:50.548' WHERE `age` = 0
```



```go
db.Model(&user).Where("addr", "北京").Update("work", "go开发工程师")
```

输出结果：

```sql
[rows:5] UPDATE `users` SET `work`='go开发工程师',`updated_at`='2022-06-23 13:38:32.69' WHERE `addr` = '北京'
```



### 删除数据(D)

###### 删除一条记录

删除一条记录时，删除对象需要指定主键，否则会触发批量 Delete

###### 根据主键删除

```go
//根据主键删除
	var user User
	db.Delete(&user, 17)
```

输出：

````
[rows:1] DELETE FROM `users` WHERE `users`.`id` = 17
````



```go
//使用slice
	var user User
	db.Delete(&user, []int{18, 19, 20, 21})
```

输出：

```sql
[rows:0] DELETE FROM `users` WHERE `users`.`id` IN (18,19,20,21)
```

###### 软删除

如果您的模型包含了一个 `gorm.deletedat` 字段（`gorm.Model` 已经包含了该字段)，它将自动获得软删除的能力！

拥有软删除能力的模型调用 `Delete` 时，记录不会从数据库中被真正删除。但 GORM 会将 `DeletedAt` 置为当前时间， 并且你不能再通过普通的查询方法找到该记录。

```go
	// 批量删除
	db.Where("age = ?", 20).Delete(&User{})
```

输出：

```sql
rows:8] DELETE FROM `users` WHERE age = 20
```



```sql
// 在查询时会忽略被软删除的记录
db.Where("age = 20").Find(&user)
fmt.Println(user)
```

输出：查询无果，已经被软删除了

```
{0  <nil> 0   <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
```



### 关联
##### Belongs To
```belongs to``` 会与另一个模型建立了一对一的连接。 这种模型的每一个实例都 “属于” 另一个模型的一个实例。
例如，假如我们的应用包含 ```user``` 和 ```company```，并且每个 ```user``` 能且只能被分配给一个 ```company```。下面的类型就表示这种关系。 注意，在 ```UserTest ```对象中，有一个和 ```CompanyTest``` 一样的``` CompanyTestID```。 默认情况下， ```CompanyTestID``` 被隐含地用来在 ```UserTest``` 和 ```CompanyTest``` 之间创建一个外键关系， 因此必须包含在 ```UserTest``` 结构体中才能填充 ```CompanyTest``` 内部结构体
外键简单解释：一张表中的外键，该表关联的另一张表的主键，例如：```UserTest``` 的外键 ```CompanyTestID``` 就为 ```CompanyTest``` 表的主键

实例：

```go
// `UserTest` 属于 `CompanyTest`，`CompanyTestID` 是外键
type UserTest struct {
  gorm.Model
  Name          string
  CompanyTestID int // CompanyTestID会默认为外键
  CompanyTest   CompanyTest  //这里必须外键名前缀一致
}

type CompanyTest struct {
  ID   int
  Name string
}
```



建表：

```go
func main() {
   // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
  dsn := "root:password@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

  //用于输出使用的sql语句
  newLogger := logger.New(
      log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
  logger.Config{
         SlowThreshold:             time.Second, // 慢 SQL 阈值
  LogLevel:                  logger.Info, // 日志级别
  IgnoreRecordNotFoundError: true, // 忽略ErrRecordNotFound（记录未找到）错误
  Colorful:                  true, // 禁用彩色打印
  },
  )
   //打开mysql服务中对应的数据库
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
      Logger: newLogger,
  })
   if err != nil {
      panic(err)
   }

   err = db.AutoMigrate(&UserTest{})
   if err != nil {
      panic(err)
   }
```

这里需要注意，gorm 会先建立 CompanyTest 表然后建立 UserTest

写入数据：

```go
//分别插入数据，并且自动写入外键值
db.Create(&UserTest{
   Name: "ice_moss",
   CompanyTest: CompanyTest{
      Name: "腾讯",
  },
})
```

在插入数据的时候，也是先对关联表 `CompanyTest` 插入，然后插入 `UserTest`

或者指定 ID

```go
//可指定外键
db.Create(&UserTest{
   Name: "ice_moss5",
  CompanyTest: CompanyTest{
      ID:   3,
  Name: "字节跳动",
  },
})
```



如下图
UserTest:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/Cf8RbuetRt.png%21large.png)



CompanyTest:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/YaulAJyN5n.png%21large.png)



##### 关联查询

###### db.Preload()

```go
//多表关联查询
var User1 []UserTest
db.Preload("CompanyTest").Find(&User1)
for key, value := range User1 {
   fmt.Println(key, value)
}
```

**db.Joins()**

```go
//多表关联查询
var User2 UserTest
db.Joins("CompanyTest").First(&User2)
fmt.Println(User2.Name, User2.CompanyTest.Name)
```



##### has many
```has many ```与另一个模型建立了一对多的连接。 不同于 ```has one```，拥有者可以有零或多个关联模型。
例如，您的应用包含 ```user``` 和 ```credit card``` 模型，且每个 ```user ```可以有多张 ```credit card```。

```go
// User 有多张 CreditCard，UserID 是外键，多个CreditCard可以对应一个User，所以每一个CreditCard都需要有外键指向User
type User struct {
    gorm.Model
    CreditCards []CreditCard
}

type CreditCard struct {
    gorm.Model
    Number string
    UserID uint   //外键
}
```



**重写外键**

```go
type User struct {
   gorm.Model
  CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
   gorm.Model
  Number    string
  UserRefer uint //外键， 每一张卡需要指向唯一用户，所以每一个CreditCard需要使用外键指向User
}
```



现在来插入几条记录

```go
//插入数据，两条卡记录指向同一user
var user User
db.Create(&user)
db.Create(&CreditCard{
   Number:    "12",
  UserRefer: user.ID,  //CreditCard外键为User的主键
})
db.Create(&CreditCard{
   Number:    "34",
  UserRefer: user.ID,  //给外键
})
```



CreditCard：

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/38MGHICZJ9.png%21large.png)



User:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/GvY3HEal0M.png%21large.png)



反向查询：

```go
//使用User做反向查询
var user User
db.Preload("CreditCards").First(&user)
for _, value := range user.CreditCards {
   fmt.Println(value.Number)
}
```

输出：

```sql
[2.587ms] [rows:2] SELECT * FROM `credit_cards` WHERE `credit_cards`.`user_refer` = 1 AND `credit_cards`.`deleted_at` IS NULL

[5.851ms] [rows:1] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
12
34
```





##### Many to Many
```Many to Many``` 会在两个 ```model ```中添加一张连接表。

例如，您的应用包含了 ```user``` 和 ```language```，且一个 user 可以说多种``` language```，多个 user 也可以说一种 ```language```。

```go
// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
   gorm.Model
  Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
   gorm.Model
  Name string
  Users []*User `gorm:"many2many:user_languages;"`
}
```



```go
// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
   gorm.Model
   Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
   gorm.Model
   Name string
}
```



当使用 GORM 的 `AutoMigrate` 为 `User` 创建表时，GORM 会自动创建连接表
建表：

users:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/yAA9XSIqrE.png%21large.png)



languages:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/04iX6MlHjn.png%21large.png)



user_languages:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/rTkpLNX2d5.png%21large.png)



插入记录：

```go
//写入数据
var language []Language
language = append(language, Language{Name: "golang"})
language = append(language, Language{Name: "c++"})
language = append(language, Language{Name: "java"})
db.Create(&User{
   Languages: language,
})
```



输出：一个执行三条 sql 语句，分别对三张表进行插入

```sql
2022/07/29 22:58:29 /Users/feng/go/src/GormStart/ch11/mian.go:52
[7.552ms] [rows:3] INSERT INTO `languages` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ('2022-07-29 22:58:29.731','2022-07-29 22:58:29.731',NULL,'golang'),('2022-07-29 22:58:29.731','2022-07-29 22:58:29.731',NULL,'c++'),('2022-07-29 22:58:29.731','2022-07-29 22:58:29.731',NULL,'java') ON DUPLICATE KEY UPDATE `id`=`id`

2022/07/29 22:58:29 /Users/feng/go/src/GormStart/ch11/mian.go:52
[1.197ms] [rows:3] INSERT INTO `user_languages` (`user_id`,`language_id`) VALUES (1,1),(1,2),(1,3) ON DUPLICATE KEY UPDATE `user_id`=`user_id`

2022/07/29 22:58:29 /Users/feng/go/src/GormStart/ch11/mian.go:52
[16.804ms] [rows:1] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`) VALUES ('2022-07-29 22:58:29.727','2022-07-29 22:58:29.727',NULL)
```



users:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/B3bww30iuv.png%21large.png)



languages:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/kXCP0xrEu4.png%21large.png)



user_languages:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/sbTwgjhin2.png%21large.png)



查询：

```go
var user User
db.Preload("Languages").Find(&user)
for _, value := range user.Languages {
   fmt.Println(value.Name)
}
```

输出：

```go
[0.523ms] [rows:3] SELECT * FROM `user_languages` WHERE `user_languages`.`user_id` = 1

[0.641ms] [rows:3] SELECT * FROM `languages` WHERE `languages`.`id` IN (1,2,3) AND `languages`.`deleted_at` IS NULL

[3.467ms] [rows:1] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
golang
c++
java
```



获取数据的另一种方式：

```go
var user User
db.First(&user)
var languges []Language
_ = db.Model(&user).Association("Languages").Find(&languges)
for _, value := range languges {
   fmt.Println(value.Name)
}
```

输出：

```go
[3.020ms] [rows:1] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1

[1.765ms] [rows:3] SELECT `languages`.`id`,`languages`.`created_at`,`languages`.`updated_at`,`languages`.`deleted_at`,`languages`.`name` FROM `languages` JOIN `user_languages` ON `user_languages`.`language_id` = `languages`.`id` AND `user_languages`.`user_id` = 1 WHERE `languages`.`deleted_at` IS NULL
golang
c++
java
```



下面来介绍一个多对多的实例：

```go
// User 拥有并属于多种 language，`user_languages` 是连接表type User struct {
   Name string
  gorm.Model
  Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
   gorm.Model
  Name  string
  Users []*User `gorm:"many2many:user_languages;"`
}
```

写入记录：

```go
//写入数据
var language []*Language
language = append(language, &Language{Name: "golang"})
language = append(language, &Language{Name: "c++"})
language = append(language, &Language{Name: "java"})
db.Create(&User{
   Languages: language,
})

var user []*User
user = append(user, &User{Name: "ice_moss1"})
user = append(user, &User{Name: "ice_moss2"})
user = append(user, &User{Name: "ice_moss3"})
db.Create(&Language{
   Users: user,
})
```





users:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/Vgc9bhbzSY.png%21large.png)



languages:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/d3nKBbDfCw.png%21large.png)



user_languages:

![](https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/gorm/5xj7ehQdrw.png%21large.png)



从关联表中我们可以很直观的看出：users 中 ID 为 1 的有 languages 中有 ID 为 1、2、3 的与之对应，同样，languages 中 ID 为 4 的有 languages 中有 ID 为 2、3、4 的与之对应

多表关联查询我们这样做:

```go
func GetAllUsers(db *gorm.DB) ([]User, error) {
   var users []User
  err := db.Model(&User{}).Preload("Languages").Find(&users).Error
   return users, err
}

func GetAllLanguages(db *gorm.DB) ([]Language, error) {
   var languages []Language
  err := db.Model(&languages).Preload("Users").Find(&languages).Error
   return languages, err
}
```

调用：

```go
users, err := GetAllUsers(db)
if err != nil {
   panic(err)
}
for _, value := range users {
   fmt.Println(value.Name)
}

languges, err := GetAllLanguages(db)
if err != nil {
   panic(err)
}
for _, value := range languges {
   fmt.Println(value.Name)
}
```

输出：

```go
2022/08/02 10:08:06 /Users/feng/go/src/GormStart/ch12/mian.go:109
[4.635ms] [rows:6] SELECT * FROM `user_languages` WHERE `user_languages`.`user_id` IN (1,2,3,4)

2022/08/02 10:08:06 /Users/feng/go/src/GormStart/ch12/mian.go:109
[2.984ms] [rows:4] SELECT * FROM `languages` WHERE `languages`.`id` IN (1,2,3,4) AND `languages`.`deleted_at` IS NULL

2022/08/02 10:08:06 /Users/feng/go/src/GormStart/ch12/mian.go:109
[54.201ms] [rows:4] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL

ice_moss1
ice_moss2
ice_moss3

2022/08/02 10:08:06 /Users/feng/go/src/GormStart/ch12/mian.go:115
[1.467ms] [rows:6] SELECT * FROM `user_languages` WHERE `user_languages`.`language_id` IN (1,2,3,4)

2022/08/02 10:08:06 /Users/feng/go/src/GormStart/ch12/mian.go:115
[1.769ms] [rows:4] SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3,4) AND `users`.`deleted_at` IS NULL

2022/08/02 10:08:06 /Users/feng/go/src/GormStart/ch12/mian.go:115
[4.931ms] [rows:4] SELECT * FROM `languages` WHERE `languages`.`deleted_at` IS NULL
golang
c++
java
```









