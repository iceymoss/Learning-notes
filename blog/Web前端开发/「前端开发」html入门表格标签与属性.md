[toc]



### 表格标签

#### table

table用来定义表格

```html
 <table></table>
```

#### tr

定义表格的行

```html
<table>
   <tr>
     <td>第一行</td>
   </tr>
   <tr>
     <td>第二行</td>
   </tr>
   <tr>
     <td>第三行</td>
   </tr>
</table>
```

#### td

定义表格单元格

```html
<tr>
    <td>学号</td>
    <td>姓名</td>
    <td>专业</td>
    <td>年级</td>
</tr>
```



#### th

定义表头单元格



#### table属性

##### border属性

给表格加上边框的，值就是边框的厚度



##### width属性

指表格的宽度



##### cellpadding属性

定义内容和单元格之间的距离



##### align属性

该属性可以在table标签也可以在tr标签

表格放置网页位置，有：center、left、right三个值



##### bgcolor属性

背景颜色，该属性可以在table标签也可以在tr标签



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
            <table border="1" width="500" align="center">
                <caption>学生信息表</caption>
                <tr bgcolor="red" align="center">
                    <td>学号</td>
                    <td>姓名</td>
                    <td>专业</td>
                    <td>年级</td>
                </tr>
                <tr>
                    <td>001</td>
                    <td>小明</td>
                    <td>计算机科学与技术</td>
                    <td>2023级</td>
                </tr>
                <tr>
                    <td>004</td>
                    <td>小化</td>
                    <td>软件工程</td>
                    <td>2022级</td>
                </tr>
            </table>
        </div>
    </body>
</html>

```



####  表格标题标签

```html
<caption></caption>
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
            <table border="1" width="500" cellpadding="20">
                <caption>学生信息表</caption>
                <tr>
                    <td>学号</td>
                    <td>姓名</td>
                    <td>专业</td>
                    <td>年级</td>
                </tr>
                <tr>
                    <td>001</td>
                    <td>小明</td>
                    <td>计算机科学与技术</td>
                    <td>2023级</td>
                </tr>
                <tr>
                    <td>004</td>
                    <td>小化</td>
                    <td>软件工程</td>
                    <td>2022级</td>
                </tr>
            </table>
        </div>
    </body>
</html>
```



#### 合并单元格

来介绍一下html如何实现单元格的合并

下面是一个4列3行的表格

|      |      |      |      |
| ---- | ---- | ---- | ---- |
|      |      |      |      |
|      |      |      |      |

比如我们要将表格的第1列第1行和第1列第2行进行合并；将第3列第3行和第4列第3行进行合并，将其内容合并后，其实就是将第一行变成4列，将第2行变成3列(第一列合并到第1行第1列)，第3行变成3列(第3行第3列和第4列合并)。



##### 合并行rowspan

使用rowspan属性来选择合并的行数

##### 合并列colspan

使用colspan属性来合并的列数

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
            <table border="1" width="500" align="center">
                
                <tr>
                    <td rowspan="2">哈哈</td>
                    <td>哈哈</td>
                    <td>哈哈</td>
                    <td>哈哈</td>
                </tr>
                <tr>
                    <td>哈哈</td>
                    <td>哈哈</td>
                    <td>哈哈</td>
                </tr>
                <tr>
                    <td>哈哈</td>
                    <td>哈哈</td>
                    <td colspan="2">哈哈</td>
                </tr>
            </table>
        </div>
    </body>
</html>
```











