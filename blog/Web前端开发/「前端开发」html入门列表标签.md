[toc]

### 概况

我们下面来学习html的三种列表

|      标签       | 标签名称 |
| :-------------: | :------: |
| ```<ul></ul>``` | 无序列表 |
| ```<ol></ol>``` | 有序列表 |
| ```<dl></dl>``` | 定义列表 |

### 无序列表

无序列表指没有刻意的排序, 无序列表使用```<ul></ul>```每一个列表中都有```<li></li>```标签

**注意:```<ul></ul>```和```<li></li>```必须组合(嵌套)使用**, **, ```<li></li>```必须放在```<ul></ul>```或者```<ol></ol>```中使用，```<ul></ul>```标签中只能放```<li>```, 但是```<li>```里面可以放其他标签。**

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
    <ul>
        <li>面包</li>
        <li>牛奶
          <p>注意：要低脂的</p>
      	</li>
        <li>鸡蛋</li>
        <li>奶酪</li>
    </ul>
</body>
</html>
```



### 有序列表

指有刻意排序的列表

**注意： ```<li></li>```必须放在```<ul></ul>```或者```<ol></ol>```中使用，```<ol></ol>```标签中只能放```<li>```, 但是```<li>```里面可以放其他标签。** 



#### ol标签有type属性

| type属性值 | 意义                 |
| ---------- | -------------------- |
| a          | 表示小写编号排序     |
| A          | 表示大写编号排序     |
| i          | 表示小写罗马编号排序 |
| I          | 表示大写罗马编号排序 |
| 1          | 表示按数字排序(默认) |

#### start属性

指排序从多少开始



#### reversed属性

指顺序还是倒序， 在属性栏出添加```reversed```就是变成倒序, 

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
  	<h1>一年级2班期末考试成绩排名</h1>
    <ol type="A", start="5">
        <li>小慕</li>
        <li>小李</li>
        <li>小明</li>
        <li>小刚</li>
    </ol>
</body>
</html>
```



### 定义列表

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
    <dl>
        <dt>北京</dt>
        <dd>国家首都、政治中心、文化中心</dd>
        <dt>上海</dt>
        <dd>国际经济、金融、科技创新中心</dd>
        <dt>深圳</dt>
        <dd>经济特区、国际化都市</dd>
    </dl>  
</body>
</html>
```



### 实战

学完三种标签后，我们就开始在实际的网页中，来写列表

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
            <div class="nav">
                <ul>
                    <li>首页</li>
                    <li>医院概况</li>
                    <li>医院动态</li>
                    <li>专家学科</li>
                    <li>服务指南</li>
                    <li>医院文化</li>
                    <li>信息公开</li>
                    <li>交流互动</li>
                </ul>
            </div>
        </div>

        <!-- 轮播图 -->
        <div class="banner">

            <!-- 小圆点 -->
            <ol>
                <li></li>
                <li></li>
                <li></li>
                <li></li>
            </ol>
        </div>


        <!-- 主要内容 -->
        <div class="content">

            <!-- 常用链接 -->
            <div clasee="userful-link">
                <ul>
                    <li>就诊须知</li>
                    <li>就医流程</li>
                    <li>专家介绍</li>
                    <li>医保就医</li>
                    <li>健康体检</li>
                </ul>
            </div>

            <!-- 动态和公告 -->
            <div class="news-and-notice">
                <div class="news">
                    <h2>医院动态</h2>

                    <!-- 新闻 -->
                    <div class="news-content">

                        <!-- 图片新闻 -->
                        <div class="news-hots"></div>
                            
                        <ul>
                            <li>为病理诊断装上“千里眼”！深医与和平县人民医院实现远程病理会诊</li>
                            <li>市医保局调研组莅临深医开展门诊共济保障改革工作调研</li>
                            <li>深医许楠医生获得2022年度“豪韵达人秀”甲状腺手术视频大赛全国总冠</li>
                            <li>深圳市输血行业两项团体标准正式发布！</li>
                            <li>深医首次入选复旦版《2021年度中国医院综合排行榜》综合百强!</li>
                        </ul>
                        

                    </div>
                </div>

                <!-- 医院公告 -->
                <div class="notice">
                    <h2>医院公告</h2>

                    <!-- 有序列表 -->
                    <ol>
                        <li>
                            <dl>
                                <dt>养生堂</dt>
                                <dd>郭主任， 特殊时期糖尿病人的新冠...</dd>
                            </dl>
                        </li>
                        <li>
                            <dl>
                                <dt>养生堂</dt>
                                <dd>郭主任， 特殊时期糖尿病人的新冠...</dd>
                            </dl>
                        </li>
                        <li>
                            <dl>
                                <dt>养生堂</dt>
                                <dd>郭主任， 特殊时期糖尿病人的新冠...</dd>
                            </dl>
                        </li>
                        <li>
                            <dl>
                                <dt>养生堂</dt>
                                <dd>郭主任， 特殊时期糖尿病人的新冠...</dd>
                            </dl>
                        </li>
                        <li>
                            <dl>
                                <dt>养生堂</dt>
                                <dd>郭主任， 特殊时期糖尿病人的新冠...</dd>
                            </dl>
                        </li>
                    </ol>

                </div>
            </div>

            <!-- 广告图片 -->
            <div class="ad-images"></div>

            <!-- 科室介绍 -->
            <div class="dep-info">
                <h2>科室介绍</h2>

                <div class="dep-content">
                    <div>
                        <h3>内科</h3>
                        <ul>
                            <li>呼吸内科</li>
                            <li>消化内科</li>
                            <li>神经内科</li>
                            <li>心血管内科</li>
                            <li>免疫内科</li>
                            <li>内分泌科</li>
                            <li>肾内科</li>
                        </ul>
                    </div>
                    <div>
                        <h3>肿瘤科</h3>
                        <ul>
                            <li>肿瘤内科</li>
                            <li>肿瘤外科</li>
                            <li>肿瘤妇科</li>
                            <li>骨肿瘤科</li>
                            <li>放疗科</li>
                            <li>肿瘤康复科</li>
                            <li>肿瘤综合科</li>
                        </ul>
                    </div>
                    <div>
                        <h3>肿瘤科</h3>
                        <ul>
                            <li>肿瘤内科</li>
                            <li>肿瘤外科</li>
                            <li>肿瘤妇科</li>
                            <li>骨肿瘤科</li>
                            <li>放疗科</li>
                            <li>肿瘤康复科</li>
                            <li>肿瘤综合科</li>
                        </ul>
                    </div>
                    <div>
                        <h3>外科</h3>
                        <ul>
                            <li>普通外科</li>
                            <li>神经外科</li>
                            <li>心胸外科</li>
                            <li>泌尿外科</li>
                            <li>肝胆外科</li>
                            <li>肛肠外科</li>
                            <li>心血管外科</li>
                            <li>烧伤科</li>
                            <li>骨外科</li>
                            <li>乳腺外科</li>
                        </ul>
                    </div>
                    <div>
                        <h3>儿科</h3>
                        <ul>
                            <li>儿科总和</li>
                            <li>小儿内科</li>
                            <li>小儿外科</li>
                            <li>新生儿科</li>
                            <li>儿童营养科</li>
                            <li>消化内科</li>
                        </ul>
                    </div>
                </div>

            </div>

            <!-- 专家介绍 -->
            <div class="exp-info">
                <h2>专家介绍</h2>
                <ul>
                    <li>
                        <div class="picbox"></div>
                        <div class="wordbox">
                            <p>姓名：李琳</p>
                            <p>科室：肿瘤内科</p>
                            <p>职称：主任医师</p>
                            <p>介绍：北京医院肿瘤内科科室主任，党支部书记，副教授，硕士研究生导师，中国老年肿瘤专业委员会肺癌分委会常务委员，北京医学……</p>
                        </div>
                    </li>
                    <li>
                        <div class="picbox"></div>
                        <div class="wordbox">
                            <p>姓名：毛永辉</p>
                            <p>科室：肾脏内科</p>
                            <p>职称：主任医师</p>
                            <p>介绍：北京医院肾内科主任，主任医师，硕士研究生导师。1989年毕业于华西医科大学临床医学院，后在北京医院内科、肾内科工作……</p>
                        </div>
                    </li>
                    <li>
                        <div class="picbox"></div>
                        <div class="wordbox">
                            <p>姓名：黄慈波</p>
                            <p>科室：风湿免疫科</p>
                            <p>职称：主任医师</p>
                            <p>介绍：教授主任医师研究生导师卫生部北京医院风湿免疫科主任 北京大学医学部风湿免疫系副主任；海峡两岸医师交流协会风湿免疫……</p>
                        </div>
                    </li>
                    <li>
                        <div class="picbox"></div>
                        <div class="wordbox">
                            <p>姓名：曹素艳</p>
                            <p>科室：特需医疗部</p>
                            <p>职称：主任医师</p>
                            <p>介绍：北京医院特需医疗部（健康管理中心）/全科医疗部主任，老年与全科医学中心副主任，主任医师，医学硕士。北京大学医学部硕士……</p>
                        </div>
                    </li>
                    <li>
                        <div class="picbox"></div>
                        <div class="wordbox">
                            <p>姓名：陈海波</p>
                            <p>科室：神经内科</p>
                            <p>职称：主任医师</p>
                            <p>介绍：北京医院神经内科主任，主任医师，北京大学医学部神经病学系副主任、教授。中国医师协会神经内科医师分会副会长兼帕金……</p>
                        </div>
                    </li>
                    <li>
                        <div class="picbox"></div>
                        <div class="wordbox">
                            <p>姓名：Jack</p>
                            <p>科室：普通外科</p>
                            <p>职称：主任医师</p>
                            <p>介绍：北京医院普通外科主任，胃肠外科主任，北京大学医学部硕士研究生导师，1985年毕业于白求恩医科大学，从事普外科临床工……</p>
                        </div>
                    </li>
                </ul>

            </div>
        </div>

        <!-- 页脚 -->
        <div class="footer">

            <!-- 友情链接 -->
            <div class="friend-links"></div>
                <h2>友情链接</h2>
                <ul>
                    <li>院长信箱</li>
                    <li>投诉信箱</li>
                    <li>在线调查</li>
                    <li>地理位置</li>
                    <li>网站地图</li>
                    <li>网站帮助</li>
                    <li>隐私声明</li>
                </ul>

            <!-- 地址 -->
            <div class="address">
                <h2>小慕医生</h2>
                <ul>
                    <li>地址：北理工国防大厦111号</li>
                    <li>电话：088-88888888</li>
                    <li>邮箱：kefu@imooc.com</li>
                    <li>邮编：666666</li>
                    <li>网址：http://imooc.com</li>
                    <li>举报热线：088-88888888</li>
                </ul>
            </div>   
        </div>
    </body>
</html>
```

