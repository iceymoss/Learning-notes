[toc]

### 介绍

相信很多朋友都知道全球最大的代码托管平台GitHub，对于开发人员来而言就更不要说了，尤其是刚接触编程的朋友来说，登GitHub这类托管平台，看到就很蒙蔽，不知道如何使用，本篇文章我将简单的介绍git让我们快速入手git。

### 初始化仓库

git仓库的初始化

```sh
git init name
```



### 移除仓库

直接删除该目录下的.git 文件夹：`rm -rf .git`

-r ：递归的删除该目录下的文件夹和文件，及子目录下的文件夹和文件

-f ： 忽略不存在的文件

```sh
rm -rf .git
```



### 拉取远程仓库

```sh
git clone ……/name.git
```

例如我们拉取《java面试指南》：

```sh
git clone git@github.com:Snailclimb/JavaGuide.git
```



### 工作区、暂存区、本地库、远程库的关系

#### 1. 工作区

工作区顾名思义，就是我们正常编码的目录

#### 2. 暂存区

当我们在工作区将工作代码完成后，想要上传git的第一步就是将需要提交的内容先添加到暂存区

```shell
git add 需要添加的内容


```

如果将当前目录所有文件上传：

```sh
git add .
```

添加文件至git仓库暂存区```git add```后不给参数，会默认将所有文件添加到git 仓库暂存区

```sh
git add file_name
```

添加指定文件夹/文件

例如：

```sh
➜  gitstudy git:(master) ✗ tree
.
├── LICENSE
├── README.md
├── gittest
    ├── GIS
    │   ├── GIS.html
    │   ├── index.html
    │   ├── information.html
    │   ├── mycoucer.html
    │   └── pome.html
    ├── README.md
```

添加gittest

```sh
git add gittest
```





#### 3. 本地库

当我们把文件添加到暂存区后就可以提交到本地库了

提交文件至git 当前分支

```sh
git commit -m "这里填写的是你想要填写的说明"
```



#### 4 远程仓库

远程仓库就github、gitee等这里远程托管平台

将本地库上传至远程

两种情况：

* 如果我们的项目是直接从远程仓库clone下来的：

  ```sh
  git push -u origin 上传的仓库名
  ```

* 如果我们是直接在本地创建的git仓库想要上传远程git仓库，就需要：

  ```sh
  git remote add origin git@github.com:iceymoss/StudyGit.git
  ```

  将远程仓库与本地绑定

  然后：

  ```sh
  git push -u origin 上传的仓库分支名
  ```

  

所以，我们将完成后的代码上传远程仓库就是: **工作区-->暂存区-->本地仓库-->远程仓库**

说明：

第一步是用`git add`把文件添加进去，实际上就是把文件修改添加到暂存区；

第二步是用`git commit`提交更改，实际上就是把暂存区的所有内容提交到当前分支。

因为我们创建Git版本库时，Git自动为我们创建了唯一一个`master`分支，所以，现在，`git commit`就是往`master`分支上提交更改。

你可以简单理解为，需要提交的文件修改通通放到暂存区，然后，一次性提交暂存区的所有修改。



### commit 规则

Our Pull Request Title follow [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/) 

* **build**：影响构建系统或外部依赖项的更改（示例范围：gulp、broccoli、npm）
* **ci**：对我们的 CI 配置文件和脚本的更改（示例范围：Travis、Circle、BrowserStack、SauceLabs）
* **docs**：仅更改文档
* **feat**：新功能
* **fix**：错误修复
* **perf**：提高性能的代码更改
* **refactor**：既不修复错误也不添加功能的代码更改
* **style**：不影响代码含义的更改（空格、格式、缺少分号等）
* **test**：添加缺失的测试或纠正现有的测试



### 查看仓库文件状态/修改/记录

查看仓库文件状态, 会显示出仓库文件的状态，例如文件被修改

```sh
git status 
```

使用```git diff```查修改的内容

```shell
git diff file_name
```



查看历史记录

```shell
git log
```





### 回退上一个修改 

回退上一个修改 ，注意：`HEAD`指向的版本就是当前版本

```sh
git reset --hard HEAD^
```



只要上面的命令行窗口还没有被关掉，你就可以顺着往上找啊找啊，找到那个`append GPL`的`commit id`是`1094adb...`，于是就可以指定回到未来的某个版本：

```
git reset --hard commit_id
```

列如：1094a为版本号前几位 

```shell
git reset --hard 1094a
```



如果哪天关掉了电脑，第二天早上就后悔了，想恢复到新版本怎么办？找不到新版本的`commit id`怎么办？

在Git中，总是有后悔药可以吃的。当你用`$ git reset --hard HEAD^`回退到`add distributed`版本时，再想恢复到`append GPL`，就必须找到`append GPL`的commit id。Git提供了一个命令`git reflog`用来记录你的每一次命令：

```sh
$ git reflog
e475afc HEAD@{1}: reset: moving to HEAD^
1094adb (HEAD -> master) HEAD@{2}: commit: append GPL
e475afc HEAD@{3}: commit: add distributed
eaadf4e HEAD@{4}: commit (initial): wrote a readme file
```

就可以找到我们修改的版本号



在你准备提交前，一杯咖啡起了作用，你猛然发现了`stupid boss`可能会让你丢掉这个月的奖金！

既然错误发现得很及时，就可以很容易地纠正它。你可以删掉最后一行，手动把文件恢复到上一个版本的状态。如果用`git status`查看一下：

```sh
$ git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   readme.txt

no changes added to commit (use "git add" and/or "git commit -a")
```

你可以发现，Git会告诉你，`git checkout -- file`可以丢弃工作区的修改：

```
$ git checkout -- readme.txt
```

命令`git checkout -- readme.txt`意思就是，把`readme.txt`文件在工作区的修改全部撤销，这里有两种情况：

一种是`readme.txt`自修改后还没有被放到暂存区，现在，撤销修改就回到和版本库一模一样的状态；

一种是`readme.txt`已经添加到暂存区后，又作了修改，现在，撤销修改就回到添加到暂存区后的状态。

总之，就是让这个文件回到最近一次`git commit`或`git add`时的状态。



现在假定是凌晨3点，你不但写了一些胡话，还`git add`到暂存区了：

```
$ cat readme.txt
Git is a distributed version control system.
Git is free software distributed under the GPL.
Git has a mutable index called stage.
Git tracks changes of files.
My stupid boss still prefers SVN.

$ git add readme.txt
```

庆幸的是，在`commit`之前，你发现了这个问题。用`git status`查看一下，修改只是添加到了暂存区，还没有提交：

```
$ git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

	modified:   readme.txt
```

Git同样告诉我们，用命令`git reset HEAD <file>`可以把暂存区的修改撤销掉（unstage），重新放回工作区：

```
$ git reset HEAD readme.txt
Unstaged changes after reset:
M	readme.txt
```

命令`git rm`用于删除一个文件

```sh
git rm file_name
```



### git分支

#### Git 中的分支是什么？

分支是指向提交的指针。

Git 分支是从特定时间点开始的项目及其更改的快照。

在处理大型项目时，有包含所有代码的主存储库，通常称为main或master。

分支允许您创建原始主要工作项目的新的、独立的版本。您可以创建一个分支来编辑它以进行更改、添加新功能或在尝试修复错误时编写测试。一个新的分支可以让你在不以任何方式影响主代码的情况下做到这一点。

总而言之 - 分支让您可以在不影响核心代码的情况下更改代码库，直到您完全准备好实施这些更改。

这有助于您保持代码库整洁有序。



#### 创建分支

```sh
git branch 分支名
```

例如在开发过程中我创建一个自己的开发分支：

```sh
git branch dev_iceymoss
```



#### 查看所有分支

```sh
git branch
```

#### 切换分支

```sh
git checkout 分支名
```

例如切换至 dev_iceymoss分支

```sh
git checkout dev_iceymoss
```



#### 合并分支

当我们将项目的功能完成后，最后就需要将我们写的分支合并的主分支上去

```sh
git  merge 需要合并的分支
```

例如：

假如我们当前在dev分支上，刚开发完项目，执行了下列命令 将代码提交到远程仓库的dev分支中：

```csharp
git  add .
git  commit -m '备注信息'
git  push -u origin dev
```

想将dev分支合并到master分支，操作如下：

- 1.首先切换到master分支上

```undefined
git  checkout master
```

- 2.如果是多人开发的话 需要把远程master上的代码pull下来

```cpp
git pull origin master
//如果是自己一个开发就没有必要了，为了保险期间还是pull
```

- 3.然后我们把dev分支的代码合并到master上

```undefined
git merge dev
```

执行合并命令后我们查看一下状态

```csharp
git status

On branch master
Your branch is ahead of 'origin/master' by 4 commits.
  (use "git push" to publish your local commits)
nothing to commit, working tree clean
```

上面的意思就是你有4个commit，需要push到远程master上 



最后我们执行下面提交命令，推送至远程

```sh
git push origin master
```





#### 删除分支

##### 为什么要删除 Git 中的分支？

您已经创建了一个分支来保存要在项目中进行的更改的代码。然后，您将该更改或新功能合并到项目的原始版本中。

这意味着您不再需要保留和使用该分支，因此删除它是一种常见的最佳做法，以免它弄乱您的代码。

##### 如何在 Git 中删除本地分支

本地分支是您本地机器上的分支，不会影响任何远程分支。

在 Git 中删除本地分支的命令是：

```sh
git branch -d  local_branch_name
```

- git branch 是在本地删除分支的命令。

- -d是一个标志，是命令的一个选项，它是--delete. 顾名思义，它表示您要删除某些内容。-local_branch_name是要删除的分支的名称。

  

##### 删除远程分支

远程分支与本地分支是分开的。

它们是托管在远程服务器上的存储库，可以在那里访问。这与本地分支相比，本地分支是本地系统上的存储库。

删除远程分支的命令是：

```sh
git push remote_name -d remote_branch_name
```

- git branch您可以使用该命令删除远程分支，而不是使用用于本地分支的git push命令。
- 然后您指定遥控器的名称，在大多数情况下是origin.
- -d是删除标志，是--delete.
- remote_branch_name 是要删除的远程分支。

例如：

我想删除远程origin/test分支，所以我使用命令：

```sh
git push origin -d test
```



### 配置SSH

如果我们想要将代码上传某个远程仓库，就必须配置SHH，下面我们两步配置SHH

#### 第1步：创建SSH Key

在用户主目录下，看看有没有.ssh目录，如果有，再看看目录下有没有`id_rsa`和`id_rsa.pub`这两个文件，如果已经有了，可直接跳到下一步。如果没有，打开Shell（Windows下打开Git Bash），创建SSH Key：

```sh
$ ssh-keygen -t rsa -C "youremail@example.com"
```

你需要把邮件地址换成你自己的邮件地址，然后一路回车，使用默认值即可，由于这个Key也不是用于军事目的，所以也无需设置密码。

如果一切顺利的话，可以在用户主目录里找到`.ssh`目录，里面有`id_rsa`和`id_rsa.pub`两个文件，这两个就是SSH Key的秘钥对，`id_rsa`是私钥，不能泄露出去，`id_rsa.pub`是公钥，可以放心地告诉任何人。

#### 第2步：登陆GitHub/Gitee

打开“Account settings”，“SSH Keys”页面：

然后，点“Add SSH Key”，填上任意Title，在Key文本框里粘贴`id_rsa.pub`文件的内容

为什么GitHub需要SSH Key呢？因为GitHub需要识别出你推送的提交确实是你推送的，而不是别人冒充的，而Git支持SSH协议，所以，GitHub只要知道了你的公钥，就可以确认只有你自己才能推送。

当然，GitHub允许你添加多个Key。假定你有若干电脑，你一会儿在公司提交，一会儿在家里提交，只要把每台电脑的Key都添加到GitHub，就可以在每台电脑上往GitHub推送了。



### 实战：将本地仓库推送到远程仓库

create a new repository on the command line

```sh
//新建README.md文件
echo "# StudyGit" >> README.md
//初始化仓库
git init 
//添加README.md到暂存区
git add README.md
//提交本地库
git commit -m "feat：first commit"
//创建一个test分支， github默认是master为主分支
git branch -M test
//将远程仓库和当前本地仓库绑定
git remote add origin git@github.com:iceymoss/StudyGit.git
//上传远程仓库
git push -u origin test
```
