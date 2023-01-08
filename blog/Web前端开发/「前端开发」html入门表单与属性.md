[toc]



### 表单的创建

#### 表单是什么

表单是用来采集信息，比如我们的登录、注册、发表评论、购买商品等等，都需要使用表单来提交信息然后给到后台服务器。

#### 表单标签```<form></form>```

##### action属性

表示提交表单的地址

##### method属性

提交表单的http方法

```html
 <form action="save.php" method="post"></form>
```



### 基本控件

#### 单行文本框

##### type属性

指该输入款输入的类型

>text：文本
>
>file：文件
>
>image：图片
>
>……

```html
<input type="text">
```

##### value属性

表示输入框已经填好的值

```html
<input type="text" value="123456">
```

##### placeholder属性

表示文本框中提示的内容

```html
<input type="text" placeholder="请输入密码">
```

##### disabled属性

将文本框锁死，意思就是不能向输入框中写入信息

```html
 <input type="text" value="123456" disabled>
```

#### 实例

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="save.php" method="post">
                <p>用户名：<input type="text" value="admin"></p>
                <p>密码：<input type="text" placeholder="请输入密码"></p>
                <p>国籍：<input type="text" value="中国" disabled></p>  
            </form>
        </div>
    </body>
</html>
```



#### 单选框radio

radio在输入框中表现为按钮选择样式

value：为选择当期按钮后，提交的值

name：多个当radio类型的输入框的name的值相同时，对应的输入框值按钮选择形成互斥关系

```html
<input type="radio" name="sex" value="男">男
```

##### 实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="save.php" method="post">
                <p>用户名：<input type="text" value="admin"></p>
                <p>密码：<input type="text" placeholder="请输入密码"></p>
                <p>国籍：<input type="text" value="中国" disabled></p>
              
              	<!-- 注意：这里性别和学历是不同的类别，name值不能相同，不然同样会出现互斥关系 -->
                <p>性别：
                    <input type="radio" name="sex" value="男">男
                    <input type="radio" name="sex" value="女">女
                </p>
                <p>学历：
                    <input type="radio" name="degree" value="小学">小学
                    <input type="radio" name="degree" value="初中">初中
                    <input type="radio" name="degree" value="高中">高中
                    <input type="radio" name="degree" value="专科">专科
                    <input type="radio" name="degree" value="本科">本科
                    <input type="radio" name="degree" value="研究生">研究生
                </p>
                
            </form>
        </div>
    </body>
</html>
```



有一种情况：当我们点击男，或者本科学历等这些文本时，网页不会给我们选择输入框按钮，那么如何将点击范围扩大到文本上呢？答案是：``` <label></label>```标签。

##### ```<label></label>```

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="save.php" method="post">
                <p>用户名：<input type="text" value="admin"></p>
                <p>密码：<input type="text" placeholder="请输入密码"></p>
                <p>国籍：<input type="text" value="中国" disabled></p>
               
                <p>性别：
                    <label>
                        <input type="radio" name="sex" value="男">男
                    </label>
                    <label>
                        <input type="radio" name="sex" value="女">女
                    </label>
                </p>      
            </form>
        </div>
    </body>
</html>
```



#### 复选框checkbox

复选框表示可以选择多个

```html
<input type="checkbox" name="hobby" value="篮球">篮球
```

##### 实例：

````html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="save.php" method="post">
                <p>用户名：<input type="text" value="admin"></p>
                <p>密码：<input type="text" placeholder="请输入密码"></p>
                <p>国籍：<input type="text" value="中国" disabled></p>
                <p>爱好：
                    <label>
                        <input type="checkbox" name="hobby" value="篮球">篮球
                    </label>
                    <label>
                        <input type="checkbox" name="hobby" value="足球">足球
                    </label>
                    <label>
                        <input type="checkbox" name="hobby" value="打游戏">打游戏
                    </label>
                </p>
            </form>
        </div>
    </body>
</html>
````



#### 密码框

当前输入框内容会被隐藏为小圆点

```html
<input type="password" placeholder="请输入密码">
```

##### 实例

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="http://127.0.0.1:8080/user/login" method="post">
                <p>用户名：
                    <input type="text" placeholder="请输入用户名">
                </p>
                <P>密码：
                    <input type="password" placeholder="请输入密码">
                </P>
                <p>验证码：
                    <input type="text" value="4820" disabled>
                </p>
            </form>
            </p>
        </div>
    </body>
</html>
```



#### 下拉框

下拉菜单

```html
<select>
       <option value="Alipay">支付宝</option>
       ……
  		 ……
</select>
```



##### 实例

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="http://127.0.0.1:8080/user/pay" method="post">
                <p>支付方式：
                    <select>
                        <option value="Alipay">支付宝</option>
                        <option value="Wachatpay">微信</option>
                        <option value="UnionPay">银联</option>
                    </select>
                </p>
            </form>
            </p>
        </div>
    </body>
</html>
```



#### 多行文本框

多行文本框可以用来做留言、评论等

###### cols属性

指例数(宽度)

###### rows属性

指行数(高度)

```html
<textarea cols="70" rows="30"></textarea>
```



##### 实例

````html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="http://127.0.0.1:8080/user/login" method="post"> 
                <p>留言：
                    <textarea cols="70" rows="30"></textarea>
                </p>
            </form>
            </p>
        </div>
    </body>
</html>
````



#### 三种按钮

| type属性值 | 按钮种类                                         |
| ---------- | ------------------------------------------------ |
| button     | 按钮，可以写成```<button></button>```            |
| submit     | 提交按钮,将form标签中的输入框的内容进行向url提交 |
| reset      | 重置按钮，将form标签中的输入框的内容进行重置     |

##### 实例

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="http://127.0.0.1:8080/user/login" method="post">
                <p>用户名：
                    <input type="text" placeholder="请输入用户名">
                </p>
                <P>密码：
                    <input type="password" placeholder="请输入密码">
                </P>
                <p>验证码：
                    <input type="text" value="4820" disabled>
                </p>
                <p>普通按钮：
                    <input type="button" value="每日打卡">
                </p>

                <p>提交按钮：
                    <input type="submit" value="点击提交">
                </p>

                <p>重置按钮：
                    <input type="reset" value="点击重置">
                </p>
            </form>
            </p>
        </div>
    </body>
</html>
```



#### 更多输入框

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="http://127.0.0.1:8080/user/test" method="post">
                <p>颜色选择控件：
                    <input type="color" >
                </p>
                <p>日期选择控件：
                    <input type="date">
                </p>
                <p>时间选择控件：
                    <input type="time">
                </p>

                <!-- 点击提交时会验证该输入是否符合邮件命名规则 -->
                <P>电子邮件选择控件：
                    <input type="email">
                </P>

                <!-- 点击提交时会验证此输入框是否有输入 -->
                <p>必填项：
                    <input type="text" required>
                </p>
                <p>数字范围：
                    <input type="number" min="0" max="120">
                </p>
                <p>拖拽条：
                    <input type="range" min="10" max="100">
                </p>
                <p>搜索框：
                    <input type="search">
                </p>

                <!-- 点击提交时会验证此输入的url是否合法 -->
                <p>网址：
                    <input type="url">
                </p>
                <p>提交按钮：
                    <input type="submit" value="点击提交">
                </p>
            </form>
            </p>
        </div>
    </body>
</html>
```



#### 智能感应```<datalist></datalist>```

该标签和```<input>```绑定起来使用

##### 实例

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
    </head>
    <body>
        <div>
            <form action="http://127.0.0.1:8080/user/login" method="post">
                <p>省份：
                    <input type="text" list="province-list">
                    <datalist id="province-list">
                        <option value="山西"></option>
                        <option value="山东"></option>
                        <option value="河南"></option>
                        <option value="河北"></option>
                        <option value="湖南"></option>
                        <option value="湖北"></option>
                        <option value="广西"></option>
                        <option value="广东"></option>
                    </datalist>
                </p>
            </form>
            </p>
        </div>
    </body>
</html>
```







