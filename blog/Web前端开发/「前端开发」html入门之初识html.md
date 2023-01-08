[toc]



### html骨架

在详细信息html前先看一下html的基本骨架

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>小慕医生</title>
        <meta name="Description", content="小慕医生是专业的团队">
        <meta name="Keywords", content="责任、关爱">

    </head>
    <body>
        <div class="header">

            <!-- 网页logo -->
            <div class="loge">
                <h1>小慕医生</h1>
            </div>
        </div>

        <!-- 轮播图 -->
        <div class="banner"></div>

        <!-- 主要内容 -->
        <div class="content"></div>

        <!-- 页脚 -->
        <div class="footer">

            <!-- 友情链接 -->
            <div class="friend-links"></div>

            <!-- 地址 -->
            <div class="address"></div>
        </div>
        
    </body>
</html>
```



### head标签

head标签是我们用来给html整个文件做配置的，选择字符集、网页标题、搜索引擎关键字、js/css文件引入等

```html
<!DOCTYPE html>
<html lang="en">
    <head>
      
				<!--字符集-->
        <meta charset="UTF-8">
      
     		 <!--网页标题-->
        <title>小慕医生</title>
      
   			 <!--搜索引擎爬取关键字-->
        <meta name="Keywords", content="责任、关爱">

    </head>
  
    <body> </body>
</html>
```



### body标签

body标签用来写网页内容



#### ```<h></h>```标签

```<h></h>```标签指标题标签

```html
<h1>这是一级标题</h1>
<h2>这是二级标题</h2>
<h3>这是三级标题</h3>
<h4>这是四级标题</h4>
<h5>这是五级标题</h4>
<h6>这是六级标题</h6>
```



#### ```<p></p>```标签

这个是段落标签

```html
 <h3>方法一</h3>
     <p>下载压缩包：将要下载在 golang 版本和对应操作系统在 golang 中文网或者在 golang 官方下载到本地。</p>
		 <p>将下载后的压缩包上传服务器</p>
```



#### ```<div></div>``` 标签

div又称为盒子， 主要和css用来搭配使用

````html
<div class="box"> </div>
````

css样式

```css
.box {
         		color: rgb(6, 173, 151);
            width: 450px;
            background: repeat;
            padding: 40px;
            text-align: center;
            margin: auto;
            margin-top: 5%;
            font-family: 'Century Gothic', sans-serif;
        }
```





### 实战

这里就以一个还未加入css的网页首页做实战吧

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>小慕医生</title>
        <meta name="Description", content="小慕医生是专业的团队">
        <meta name="Keywords", content="责任、关爱">

    </head>
    <body>
        <div class="header">

            <!-- 网页logo -->
            <div class="loge">
                <h1>小慕医生</h1>
            </div>

            <!-- 工具栏 -->
            <div class="tool"></div>

            <!-- 导航栏 -->
            <div class="nav"></div>
        </div>

        <!-- 轮播图 -->
        <div class="banner"></div>

        <!-- 主要内容 -->
        <div class="content">

            <!-- 常用链接 -->
            <div clasee="userful-link"></div>

            <!-- 动态和公告 -->
            <div class="news-and-notice">
                <div class="news">
                    <h2>医院动态</h2>
                </div>
                <div class="notice">
                    <h2>医院公告</h2>
                </div>
            </div>

            <!-- 广告图片 -->
            <div class="ad-images"></div>

            <!-- 科室介绍 -->
            <div class="dep-info">
                <h2>科室介绍</h2>
            </div>

            <!-- 专家介绍 -->
            <div class="exp-info">
                <h2>专家介绍</h2>
            </div>
        </div>

        <!-- 页脚 -->
        <div class="footer">

            <!-- 友情链接 -->
            <div class="friend-links"></div>

            <!-- 地址 -->
            <div class="address"></div>
        </div>
        
    </body>
</html>
```









