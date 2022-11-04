### 文章介绍

本文我们来介绍一下go语言内置的time包，以实例的方式来介绍time包中常用的方法

### time对象

```go
type Time struct {
    wall uint64
    ext  int64
    loc  *Location
}
```

Methods:

```go
String() string
Format(layout string) string
AppendFormat(b []byte, layout string) []byte
nsec() int32
sec() int64
unixSec() int64
addSec(d int64)
setLoc(loc *time.Location)
stripMono()
setMono(m int64)
mono() int64
After(u time.Time) bool
Before(u time.Time) bool
Equal(u time.Time) bool
IsZero() bool
abs() uint64
locabs() (name string, offset int, abs uint64)
Date() (year int, month time.Month, day int)
Year() int
Month() time.Month
Day() int
Weekday() time.Weekday
ISOWeek() (year int, week int)
Clock() (hour int, min int, sec int)
Hour() int
Minute() int
Second() int
Nanosecond() int
YearDay() int
Add(d time.Duration) time.Time
Sub(u time.Time) time.Duration
AddDate(years int, months int, days int) time.Time
date(full bool) (year int, month time.Month, day int, yday int)
UTC() time.Time
Local() time.Time
In(loc *time.Location) time.Time
Location() *time.Location
Zone() (name string, offset int)
Unix() int64
UnixNano() int64
MarshalBinary() ([]byte, error)
UnmarshalBinary(data []byte) error
GobEncode() ([]byte, error)
GobDecode(data []byte) error
MarshalJSON() ([]byte, error)
UnmarshalJSON(data []byte) error
MarshalText() ([]byte, error)
UnmarshalText(data []byte) error
Truncate(d time.Duration) time.Time
Round(d time.Duration) time.Time
```



### time包方法介绍



##### time.Now()获取当前时间

```go
func Now() Time
```

实例：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now()
	fmt.Println(t)
}
```

输出：

```
2022-07-28 22:01:20.596934 +0800 CST m=+0.000081427
```

 ```+0800```指时区-东八区

##### t.Format()格式化

```go
func (t Time) Format(layout string) string
```

time对象提供了Format()方法，可以将时间格式化

实例:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now()
	fmt.Println(t)

	//时间格式化
	ts := t.Format("2006-01-02 15:04:05")  //go诞生时间123456
	fmt.Println(ts)
}
```

输出：

```
2022-07-28 22:03:57.386954 +0800 CST m=+0.000098870
2022-07-28 22:03:57
```



##### time.LoadLocation()获取相应时区时间

```go
func LoadLocation(name string) (*Location, error)
```

实例:

```go
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now()
	fmt.Println(t)

	//时区
	loc, err := time.LoadLocation("America/New_York")  
	if err != nil {
		log.Panicln("转换时区失败")
	}
  fmt.Println(t.In(loc))    //这里需要配合t.In()
}
```

输出：

```
2022-07-28 22:11:34.610774 +0800 CST m=+0.000089252
2022-07-28 10:11:34.610774 -0400 EDT
```



##### 获取年月日/时分秒

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t := time.Now()
	fmt.Println(t)
	
	//年月日
	fmt.Println(t.Year(), t.Month(), t.Day())

	//时分秒
	fmt.Println(t.Hour(), t.Minute(), t.Second())
}
```

输出：

```go
2022-07-28 22:15:06.056301 +0800 CST m=+0.000086151
2022 July 28
22 15 6
```



##### t.Parse()字符串转Time类型

```go
func Parse(layout string, value string) (Time, error)
```

实例：

```go
	//获取当前时间
	t := time.Now()
	fmt.Println(t)
  //将字符串转换为time类型
	layout := "2006-01-02 15:04:05"
	ts := "2022-07-28 23:28:40"
	t1, err := time.Parse(layout, ts)
	if err != nil {
		panic(err)
  }
	fmt.Println(t1)
```

输出：

```
2022-07-28 22:55:00.753234 +0800 CST m=+0.000100299
2022-07-28 23:28:40 +0000 UTC
```



##### 时间运算

```go
func (t Time) Add(d Duration) Time
```

```当前时间t+参数d```

实例：

```go
//获取当前时间
t := time.Now()
fmt.Println(t)

fmt.Println(t.Add(1 * time.Hour)) //当前时间+1小时
fmt.Println(t.AddDate(0, 1, 1))   //当前时间+年月日
```

输出:

```
2022-07-28 23:04:08.400064 +0800 CST m=+0.000082336
2022-07-29 00:04:08.400064 +0800 CST m=+3600.000082336
2022-08-29 23:04:08.400064 +0800 CST
```



```go
func (t Time) Sub(u Time) Duration
```

```当前时间t - 时间u```

实例：

```go
//获取当前时间
t := time.Now()
fmt.Println(t)
//时间差值
newAdd := t.Add(1 * time.Hour)
fmt.Println(t.Sub(newAdd)) //当前时间-参数
```

输出：

```
2022-07-28 23:04:08.400064 +0800 CST m=+0.000082336
-1h0m0s
```



###### 时间比较

```go
t := time.Now()
fmt.Println(t)
//时间比较
newAdd := t.Add(1 * time.Hour)
fmt.Println(t.Equal(newAdd))  // = t
fmt.Println(t.After(newAdd))  // > t
fmt.Println(t.Before(newAdd)) // < t
```

输出：

```
2022-07-28 23:15:39.781098 +0800 CST m=+0.000088697
false
false
true
```



##### t.Unix()时间戳

```go
func (t Time) Unix() int64
```

返回int64,使用时间比较可以直接使用时间戳比较

时间戳实例

```go
t := time.Now()
fmt.Println(t)
//获取时间戳
fmt.Println(t.Unix())     //秒
fmt.Println(t.UnixNano()) //纳秒
```

输出：

```
2022-07-28 23:21:07.827182 +0800 CST m=+0.000126192
1659021667
1659021667827182000
```

