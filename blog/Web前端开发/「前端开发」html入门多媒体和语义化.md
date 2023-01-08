[toc]



### 概况

下面我们来介绍一下多媒体和语义化标签的使用



### 音频audio

#### autoplay属性

```autoplay```会自动播放音频，但是有点浏览器可能会不允许这种行为(防止打扰用户)

#### loop属性

```loop``` 循环播放，在属性栏加上即可

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
  
    <!--../file/-主角-1.mp3文件相对路径， controls为播放器控件， 如果不写将不会显示播放控件，如果浏览器不支持播放的话，就会显示：抱歉，您的浏览器不支持音频播放-->
    <audio src="../file/-主角-1.mp3" controls autoplay loop>
        抱歉，您的浏览器不支持音频播放
    </audio>

</body>
</html>
```



### 视频video

#### autoplay属性

```autoplay```会自动播放视频，但是有点浏览器可能会不允许这种行为(防止打扰用户)

#### loop属性

```loop``` 循环播放，在属性栏加上即可

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
    <video src="../file/meidusha.mp4" controls autoplay loop></video>

</body>
</html>
```



### 图片img 

#### alt属性

当图片显示不出来时，将会将alt的内容作为备用，显示在图片处

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
    <img src="../file/meidusha.jpeg" alt="美杜莎女王">

</body>
</html>
```



### 超级链接a

跳转到其他网页或网站， 例如现在有两个文件：a.html 、b.html

#### titel文本悬停

将鼠标放置超链接位置就出现相应的文本

#### target新窗口

```target="blank"```时，点击超链接，浏览器会打开一个新窗口来显示跳转后的网页。

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
    <a href="../b.html" title="是不是很帅！" target="blank">点击跳转</a>
</body>
</html>
```



### 网页锚点

当一个网页很长，需要用向下滚动鼠标时，就可以直接使用锚点到达指定位置。

例如我们在:

```html
<h2 id="meidusha">美杜莎女王3</h2>
```

添加了```id="meidusha"```

当我们在浏览器搜索：

```
http://127.0.0.1/dome/test.html#meidusha
```

网页就会跳转到锚点位置：

```html
<h2 id="meidusha">美杜莎女王3</h2>
```



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
        <h2>美杜莎女王0</h2>
        <p>
            <img src="../file/meidusha.jpeg" alt="美杜莎女王">
        </p>
    </div>
    <div>
        <h2>美杜莎女王1</h2>
        <p>
            <img src="../file/meidusha.jpeg" alt="美杜莎女王">
        </p>
    </div>
    <div>
        <h2>美杜莎女王2</h2>
        <p>
            <img src="../file/meidusha.jpeg" alt="美杜莎女王">
        </p>
    </div>
    <div>
        <h2 id="meidusha">美杜莎女王3</h2>
        <p>
            <img src="../file/meidusha.jpeg" alt="美杜莎女王">
        </p>
    </div>
</body>
</html>
```



### 更多功能链接

#### 下载链接

标签a支持exe、zip、rar等文件的下载

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
    <a href="../file/test.zip" title="是不是很帅！" target="blank">点击下载</a>
</body>
</html>
```



#### 邮件链接

点击链接后，系统会打开你的邮件，并指向示例中的邮件地址

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
    <a href="mailto:yourEmail.com" title="是不是很帅！" target="blank">给小编发邮件</a>
</body>
</html>
```



#### 电话链接

用手机打开，点击链接后，系统会调用手机拨号功能，并会显示示例中的电话号码

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
    <a href="tel:12306" title="是不是很帅！" target="blank">打电话购票</a>
</body>
</html>
```



### 更多标签

在上一篇文章中介绍了div标签，但是整个网页的div标签太多，导致代码可读性不高，所以我们现在来试一试其他标签

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-27%20%E4%B8%8B%E5%8D%8811.17.29.png" style="zoom:50%;" />

下面来看看实例

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

    <!-- 网页头部 -->
    <header>

        <!-- logo -->
        <div class="logo"></div>

        <!-- 导航条 -->
        <nav>导航条</nav>
    </header>

    <!-- 网页核心 -->
    <main>

        <!-- 广告 -->
        <aside>我是广告</aside>

        <!-- 文章 -->
        <article>
            <h1>文章标题</h1>
            <section>部分1</section>
            <section>部分2</section>
            <section>部分3</section>
            <section>部分4</section>
        </article>
    </main>

    <!-- 页脚 -->
    <footer>页脚</footer>

</body>
</html>
```



### span标签

span要结合css来使用，将需要添加样式的内容用```<span>需要的内容</span>```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        span {
            color: rgb(248, 6, 6);
        }
    </style>
</head>
<body>
    <div>
        <p><span>北京</span>在区号是<span>010</span></p>
        <p><span>上海</span>在区号是<span>021</span></p>
    </div>
</body>
</html>
```



### ```<b></b>、<u><\u>、<i><\i>```标签

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-27%20%E4%B8%8B%E5%8D%8811.42.08.png" style="zoom:50%;" />

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
    <u>前端学习</u>
    <m>前端学习</m>
    <i>前端学习</i>
</body>
</html>
```



### strong、em、mark标签

需要表示强调的内容

<img src="https://blogfiles-iceymoss.oss-cn-hangzhou.aliyuncs.com/blogs/%E6%88%AA%E5%B1%8F2022-12-27%20%E4%B8%8B%E5%8D%8811.45.48.png" style="zoom:50%;" />



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
    <strong>加油</strong>
    <em>加油</em>
    <mark>加油</mark>
</body>
</html>
```



### figure、figcaption标签

两个标签混合使用

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
    <p>美杜莎女王是《斗破苍穹》小说中被塑造得非常好的人物</p>
    <p>
        <figure>
            <img src="../file/meidusha.jpeg" alt="美杜莎女王" width="300">
            <figcaption>图一：突破斗宗的美杜莎女王</figcaption>
        </figure>
    </p>
</body>
</html>
```

