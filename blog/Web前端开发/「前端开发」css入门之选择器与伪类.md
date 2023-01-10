[toc]

### 标签选择器

就是使用html标签名作为选择器，进行css样式的添加

```css
/* 将选择html页面中所有的h3标签，无论h3标签位置的深浅 */
h3 {
      color: blueviolet;
      background-color: rgb(26, 216, 22);
}
```

标签选择器，覆盖范围大，所以通常用来做标签的初始化。

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            ul {
                /* 将无序列表的前面的小圆点去掉 */
                list-style: none;
            }
            a {
                /* 将连接下方的下划线去掉 */
                text-decoration: none;
            }  
        </style>
    </head>
    <body>
        <div>
            <ul>
                <li>面包</li>
                <li>牛奶</li>
                <li>奶酪</li>
            </ul>
            <a href="http://127.0.0.1:8080/user/list">获取用户列表</a>
           
        </div>
    </body>
</html>
```



### class选择器

#### class类名

class类名,命名规则：不能数字开头，区分大小写。

```html
<h2 class="spec-h2-2">我是二级标题</h2>
```

#### class选择器

使用```.类名```对html标签进行选择，只要标签中出现class类名和选择器名相同，样式就会添加到对应的html标签中。

```css
.spec-h2-2 {
           color: aquamarine;
           font-size: large;
           font-style: italic;
}
```



**注意**：同一个标签可以属于多个类，类名使用空格隔开

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            .para {
                color:aqua;
            }
            .spec {
                font-size: xx-large;
            }
        </style>
    </head>
    <body>
        <div>
            <p class="para spec">我是一个段落</p>     
        </div>
    </body>
</html>
```



#### 原子类

就是将各种属性值分别分类出来，例如：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            .fs16 {
                font-size: 16px;
            }
            .fs18 {
                font-size: 18px;
            }
            .fs20 {
                font-size: 20px;
            }
            .fs22 {
                font-size: 22px;
            }
            .fs24 {
                font-size: 24px;
            }
            .fs26 {
                font-size: 26px;
            }
            .color-red {
                color: red;
            }
            .color-green {
                color: green;
            }
            .color-blue {
                color: blue;
            }
            .color-yellow {
                color: yellow;
            }
            .color-aqua {
                color: aqua;
            }
        </style>
    </head>
    <body>
        <div>
            <ul>
                <li class="fs18 color-yellow">面包</li>
                <li class="fs18 color-aqua">牛奶</li>
                <li class="fs18 color-blue">奶酪</li>
            </ul>
            <p class="fs18">我是一个段落</p>
            <p class="fs20">我是一个段落</p>
            <p class="fs22">我是一个段落</p>
            <p class="fs24 color-green">我是一个段落</p>
            <p class="fs26">我是一个段落</p>   
        </div>
    </body>
</html>
```



### id属性

标签可以有id属性，用来唯一标识标签，命名规则：只能以字母、数字、下划线、短横线租成，不能以数字开头，区分大小写；同一个网页不能有相同id值的标签。

#### id选择器

使用```#前缀```选择指定的标签。

```html
 <p id="par">我是一个段落</p>
```

```css
#par {
     color:aqua;
}
```



### 复合选择器

| 选择器的名称 | 示例       | 示例的意义                                |
| ------------ | ---------- | ----------------------------------------- |
| 后代选择器   | .box .spec | 选择类名为box的标签内部的类名为spec的标签 |
| 交集选择器   | li.spec    | 选择既是li的标签，也属于是spec的标签      |
| 并集选择器   | ul, ol     | 选择所有的ul标签和ol标签                  |

实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            /* 后代选择器 */
            .title .t-h1 {
                color: brown;
            }
            /* 后代选择器 */
            .title h2 {
                color: rgb(236, 11, 176);
            }
            /* 后代选择器 */
            .title .table .tr {
                background-color: aquamarine;
            }
            /* 后代选择器 */
            .title .table .student-info-head {
                color: rgb(160, 24, 228);
                font-style: italic;
            }

            /* 交集选择器 */
            .p2.green {
                font-size: xx-large;
                color: green
            }

            /* 并集选择器 */
            .ui-li, .p1 {
                font-size: x-small;
                color: rgb(173, 29, 239);
            }  

            /* 交集、后代、并集混合选择器 */
            div.box ol li p, ol {
                font-weight: 200;
                font-size: smaller;
                color: deeppink; 
            }
        </style>
    </head>
    <body>
        <div>
            <ul>
                <li class="ui-li">面包</li>
                <li class="ui-li">牛奶</li>
                <li class="ui-li">奶酪</li>
            </ul>
            <div class="title">
                <h1 class="t-h1">我是一级标题</h1>
                <h2 class="t-h2">我是二级标题</h2>
                <h2 class="t-h2">我是二级标题</h2>

                <table class="table" border="1" width="500" align="center">
                    <caption class="student-info-head">学生信息表</caption>
                    <tr class="tr" align="center">
                        <td>学号</td>
                        <td>姓名</td>
                        <td>专业</td>
                        <td>年级</td>
                    </tr>
                    <tr class="tr">
                        <td>001</td>
                        <td>小明</td>
                        <td>计算机科学与技术</td>
                        <td>2023级</td>
                    </tr class="tr">
                    <tr class="tr">
                        <td>004</td>
                        <td>小化</td>
                        <td>软件工程</td>
                        <td>2022级</td>
                    </tr class="tr">
                </table>
                
            </div>
            <p class="p1">我是一个段落</p>
            <p class="p1">我是一个段落</p>
            <p class="p2 green">我是一个段落</p>
            <p class="p2 green">我是一个段落</p>
            <div class="box">
                <ol class="city-list">
                    <li>
                        <p>北京</p>
                    </li>
                    <li>
                        <p>上海</p>
                    </li>
                    <li>
                        <p>广州</p>
                    </li>
                    <li>
                        <p>深圳</p>
                    </li>
                </ol>
            </div> 
        </div>
    </body>
</html>
```



### 伪类

| 伪类      | 意义                                       |
| --------- | ------------------------------------------ |
| a:link    | 没有被访问的超级链接                       |
| a:visited | 已经被访问过的超级链接                     |
| a:hover   | 正在被鼠标悬停的超级链接                   |
| A:activa  | 正被激活的超级链接(点击是，点击后还没松开) |

实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            a:link {
                color: aqua;
            }
            a:visited {
                color: beige;
            }
            a:hover {
                color: brown;
            }
            a:active {
                color: darkorange;
            }
        </style>
    </head>
    <body>
        <div>
            <p>
                <a href="http://wwww.baidu.com">前往百度</a>
            </p>
            <p>
                <a href="http://wwww.jd.com">前往京东商城</a>
            </p>
        </div>
    </body>
</html>
```

### css3新增伪类

#### :empty

标签中没有内容时，显示该该样式

#### :focus 

用户进行聚焦市，会显示该样式

#### :checked+span

被选中的后面的标签显示该样式

#### :root

里面的样式会影响所有内容

#### 实例

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            p {
                width: 150px;
                height: 150px;
                border: 1px solid rgb(50, 204, 231);
            }

            /* 标签中没有内容时，显示该该样式 */
            p:empty {
                background-color: rgb(255, 0, 119); 
            }

            /* 用户进行聚焦市，会显示该样式 */
            div input:focus {
                background-color: blueviolet;
            }

            /* 被选中的后面的标签显示该样式 */
            .box input:checked+span {
                color: rgb(243, 18, 164);
            }

            /* 影响所有内容 */
            :root {
                font-size: 10px;
            }
        </style>
    </head>
    <body>
        <p></p>
        <p></p>
        <p>123</p>
        <p></p>
        <div>
            <input type="text">
            <input type="text" disabled>
            <input type="text">
            <input type="text">
        </div>
        <hr>
        <div class="box">
            <input type="checkbox"><span>文字</span>
            <input type="checkbox"><span>文字</span>
            <input type="checkbox"><span>文字</span>
            <input type="checkbox"><span>文字</span>
        </div> 
    </body>
</html>
```











### 重叠性与权重计算

#### 重叠性

指多个选择器同时作用于一个标签

```html
<p id="para" class="para">我是一个段落</p>
```

```css
#para {
       color: blueviolet;
}
.para {
       font-size:x-large
}
p {
   font-style: italic;
}
```



#### 权重

当重叠性出现冲突时，比如出现不同颜色，此时就应该按照权重来选择：

>id > class > 标签



### 属性选择器

直接看实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            img {
                width: 300px;
            }
            /* 属性匹配 */
            img[alt] {
                border: 6px solid rgb(14, 1, 1);
            }

            /* 精准匹配 */
            img[alt=小黑子] {
                border: 6px solid rgb(101, 23, 235);
            }

            /* 开头匹配 */
            img[alt^=我们] {
                border: 6px solid rgb(185, 189, 226);
            }
            
            /* 以ikun-开头匹配 */
            img[alt|=ikun] {
                border: 6px solid rgb(239, 8, 8);
            }

            /* 结尾匹配 */
            img[alt$=哎呀] {
                border: 6px solid rgb(246, 205, 5);
            }

            /* 任意位置匹配 */
            img[alt*=树枝] {
                border: 6px solid rgb(48, 229, 11);
            }

            /* 以只因空格加独立部分匹配 */
            img[alt~=只因] {
                border: 10px solid rgb(236, 236, 6);
            }
        </style>
    </head>
    <body>
        <div>
            <img src="../../file/ikun.jpeg" alt="小黑子">
            <img src="../../file/ikun.jpeg" alt="小黑子">
            <img src="../../file/ikun.jpeg" alt="小黑子没有树枝">
            <img src="../../file/ikun.jpeg" alt="我们都是荔枝的">
            <img src="../../file/ikun.jpeg" alt="我们都是荔枝的">
            <img src="../../file/ikun.jpeg" alt="你干嘛哎呀">
            <img src="../../file/ikun.jpeg" alt="鸡 只因">
            <img src="../../file/ikun.jpeg" alt="ikun-永远滴神">
            <img src="../../file/ikun.jpeg" alt="ikun-永远滴神">
        </div>
    </body>
</html>
```



### 序号选择器

在一个盒子中，我们可以任意选择我们需要的标签进行添加样式。

直接看实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            /* 选择第一个 */
            .box1 p:first-child {
                color: aqua;
            }
            /* 选择最后一个 */
            .box1 p:last-child {
                color: rgb(246, 14, 14);
            }
            /* 选择第2个 */
            .box1 p:nth-child(2) {
                color: rgb(94, 82, 222);
            }
            /* 选择第3个 */
            .box1 p:nth-child(3) {
                color: rgb(22, 234, 18);
            }
            /* 偶数 */
            .box2 p:nth-child(2n) {
                color: rgb(104, 15, 206);
            }
            /* 奇数 */
            .box2 p:nth-child(2n+1) {
                color: rgb(240, 232, 9);
            }

            /* n的计数从0开始依次叠加+2：从第二个开始 */
            .box3 p:nth-child(3n+2) {
                color: rgb(240, 232, 9);
            }

            /* 选择第四个p标签 */
            .box4 p:nth-of-type(4) {
                color: rgb(32, 230, 160);
            }

            /* 选择第二个h4标签 */
            .box4 h4:nth-of-type(2) {
                color: rgb(25, 38, 224);
            }

            /* 选择倒数第2个p标签 */
            .box5 p:nth-last-child(2) {
                color: rgb(224, 25, 217);
            }
        </style>
    </head>
    <body>
        <div class="box1">
            <p>第一个段落</p>
            <p>第二个段落</p>
            <p>第三个段落</p>
            <p>第四个段落</p>
        </div>
        <hr>
        <div class="box2">
            <p>第一个段落</p>
            <p>第二个段落</p>
            <p>第三个段落</p>
            <p>第四个段落</p>
            <p>第五个段落</p>
            <p>第六个段落</p>
            <p>第七个段落</p>
            <p>第八个段落</p>
            <p>第九个段落</p>
            <p>第十个段落</p>
            <p>第十一个段落</p>
            <p>第十二个段落</p>
        </div>
        <hr>
        <div class="box3">
            <p>第一个段落</p>
            <p>第二个段落</p>
            <p>第三个段落</p>
            <p>第四个段落</p>
            <p>第五个段落</p>
            <p>第六个段落</p>
            <p>第七个段落</p>
            <p>第八个段落</p>
            <p>第九个段落</p>
            <p>第十个段落</p>
            <p>第十一个段落</p>
            <p>第十二个段落</p>
        </div>
        <hr>
        <div class="box4">
            <p>第一号段落</p>
            <p>第二号段落</p>
            <p>第三号段落</p>
            <h4>第一号标题</h4>
            <h4>第二号标题</h4>
            <p>第四号段落</p>
            <p>第五号段落</p>
            <h4>第三号标题</h4>
            <p>第六号段落</p>
        </div>
        <hr>
        <div class="box5">
            <p>第一号段落</p>
            <p>第二号段落</p>
            <p>第三号段落</p>
            <p>第四号段落</p>
    </body>
</html>
```



### 元素关系选择器

直接看实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            
            /* 只会选择儿子标签 */
            .box>p {
                color: aqua;
            }

            /* 相邻兄弟选择器 */
            .box p+div {
                color: rgb(0, 110, 255);
            }

            /* 通用兄弟(一定是同层级)选择器:表示p后面的div标签 */
            .box2 p~div {
                color: rgb(234, 36, 241);
            }
        </style>
    </head>
    <body>
        <div class="box">
            <p>儿子段落</p>
            <p>儿子段落</p>
            <div>
                <p>孙子段落</p>
                <p>孙子段落</p>
            </div>
            <p>儿子段落</p>
        </div>

        <div class="box2">
            <p>儿子段落</p>
            <p>儿子段落</p>
            <div>
                <p>div1孙子段落</p>
                <p>div1孙子段落</p>
            </div>
            <p>儿子段落</p>
            <div>
                <p>div2孙子段落</p>
                <p>div2孙子段落</p>
            </div>
        </div>
    </body>
</html>
```



### 伪元素

#### ::before

创建一个伪元素，将其成为匹配选中的元素的第一个子元素，必须设置content属性表示其中内容

```css
a::before {
	content: '⭐️'
}
```



#### ::after

创建一个伪元素，将其成为匹配选中的元素的最后一个子元素，必须设置content属性表示其中内容

```css
a::after {
     content: '⭐️';
}
```



实例：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
            .box1 a::before {
                content: '⭐️'
            }
            .box2 a::after {
                content: '⭐️';
            }

        </style>
    </head>
    <body>
        <div class="box1">
            <a href="http://www.baidu.com">这是前往百度的连接</a>
        </div>
        <div class="box2">
            <a href="http://www.baidu.com">这是前往百度的连接</a>
        </div>
    </body>
</html>
```



#### ::selection

选中文字时显示的样式



#### ::first-letter

第一个文字添加样式



#### ::first-line

第一行文字添加样式



实例：

````html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
             .box3, .box4 {
                width: 400px;
                height: 300px;
                border: 2px solid rgb(130, 36, 213);
            }

            /* 选中文字时的颜色 */
            .box3::selection {
                /* 背景颜色 */
                background-color: rgb(23, 216, 139); 
                /* 字体颜色 */
                color: rgb(165, 42, 122);
            }

            /* 第一个文字添加样式 */
            .box4::first-letter {
                font-size: xx-large;
            }

             /* 第一行文字添加样式 */
            .box4::first-line {
                text-decoration: underline red;
            }
        </style>
    </head>
    <body>
        <div class="box1">
            <a href="http://www.baidu.com">这是前往百度的连接</a>
        </div>
        <div class="box2">
            <a href="http://www.baidu.com">这是前往百度的连接</a>
        </div>
        <div class="box3">
            吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗
        </div>
        <div class="box4">
            吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗 吃了吗
        </div>
    </body>
</html>
````









