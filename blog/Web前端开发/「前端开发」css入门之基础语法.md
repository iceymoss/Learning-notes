[toc]



### 概况

CSS中文名：层叠式样式表，用来给html添加样式的语言

### 前端三层

| 分类   | 语言      | 功能                               |
| ------ | --------- | ---------------------------------- |
| 结构层 | HTML      | 搭建结构、放置部件、语义描述       |
| 样式层 | CSS       | 美化页面、实现布局                 |
| 行为层 | JavaScrip | 实现交互效果、数据收发、表单验证等 |



### 选择器

css的出现实现了页面结构(html)和样式(css)的分离，他们的关联是使用: **选择器**

例如：

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <style>
          
          /* h1、h3、p是选择器，和html中标签对应 */
            h1 {
                color: brown;
            }

            h3 {
                color: aquamarine;
            }

            p {
                color: blueviolet;
            }
        </style>
    </head>
    <body>
        <div>
            <h1>我是一级标题</h1>
            <h2>我是二级标题</h2>
            <h3>我是三级标题</h3>
            <p>我是一个段落</p>
            <p>我是一个段落</p>
            <p>我是一个段落</p>
            <p>我是一个段落</p>
        </div>
    </body>
</html>

```



#### class选择器

当我们要选择特定的标签来添加样式时，可以通过html标签中class的值和css样式中名称关联起来

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
            .spec-h2-2 {
                color: aquamarine;
            }

            .spec-p-4 {
                color: blueviolet;
            }
        </style>
    </head>
    <body>
        <div>
            <h2>我是一级标题</h2>
            <h2 class="spec-h2-2">我是二级标题</h2>
            <h2>我是三级标题</h2>
            <p>我是一个段落</p>
            <p>我是一个段落</p>
            <p>我是一个段落</p>
            <p class="spec-p-4">我是一个段落</p>
        </div>
    </body>
</html>
```



### css书写位置

#### 内嵌式

就是写在html标签```<style></style>```中。

```html
<style>
   .spec-h2-2 {
       	color: aquamarine;
    }
   .spec-p-4 {
        color: blueviolet;
    }
</style>
```



#### 外链式

将css样式单独写在后缀为```.css```的文件中，然后在html中使用```<link></link>```引入

##### rel属性

表示关系，stylesheet就表示样式表

##### href属性

引入文件的路径

```html
<link rel="stylesheet" href="../css/style.css">
```



#### 导入式

```html
<style>
    @import url(../css/style.css);
 </style>
```



#### 行内式

将样式直接写在html需要添加样式的标签上

```html
<h2 style="color: darkorange; font-size: large; font-style: italic;">我是三级标题</h2>
```



### 基本语法

在后续的css内容中我们都以内嵌式和外链式的方式展示css样式，使用选择器进行关联(绑定)，然后使用key: value键值对，key表示属性，value表示属性值。

```css
选择器 {
      key: value;
      ……;
      ……;
      ……
}
```

实例：

```css
<style>
    .spec-h2-2 {
                color: aquamarine;
                font-size: large;
                font-style: italic
     }
     .spec-p-4 {
                color: blueviolet;
                background-color: rgb(26, 216, 22);
     }
</style>
```



