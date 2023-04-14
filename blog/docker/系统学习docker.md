[toc]



# docker

## docker安装

### Linux安装

[docker官网](https://www.docker.com/products/personal/)

### MacOS安装

[docker官网](https://www.docker.com/products/personal/)

### Windows安装

[docker官网](https://www.docker.com/products/personal/)

## docker容器

### docker命令

> docker

使用docker查看所以命令，命令格式：```docker 一级命令 二级命令```

例如：

```sh
docker ps -a
```

查看所有容器(正在运行，退出)



> docker version

```docker version```查看docker版本相关的信息

```sh
[~] docker version                                                                                                                                                                                14:47:03
Client:
 Cloud integration: v1.0.31
 Version:           20.10.23
 API version:       1.41
 Go version:        go1.18.10
 Git commit:        7155243
 Built:             Thu Jan 19 17:35:19 2023
 OS/Arch:           darwin/arm64
 Context:           desktop-linux
 Experimental:      true

Server: Docker Desktop 4.17.0 (99724)
 Engine:
  Version:          20.10.23
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.18.10
  Git commit:       6051f14
  Built:            Thu Jan 19 17:31:28 2023
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.6.18
  GitCommit:        2456e983eb9e37e47538f59ea18f2043c9a73640
 runc:
  Version:          1.1.4
  GitCommit:        v1.1.4-0-g5fd4c4d
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```



> docker info

```docker info``` 显示容器运行情况等信息

```sh
[~] docker info                                                                                                                                                                                   14:47:54
Client:
 Context:    desktop-linux
 Debug Mode: false
 Plugins:
  buildx: Docker Buildx (Docker Inc., v0.10.3)
  compose: Docker Compose (Docker Inc., v2.15.1)
  dev: Docker Dev Environments (Docker Inc., v0.1.0)
  extension: Manages Docker extensions (Docker Inc., v0.2.18)
  sbom: View the packaged-based Software Bill Of Materials (SBOM) for an image (Anchore Inc., 0.6.0)
  scan: Docker Scan (Docker Inc., v0.25.0)
  scout: Command line tool for Docker Scout (Docker Inc., v0.6.0)

Server:
 Containers: 7
  Running: 2
  Paused: 0
  Stopped: 5
 Images: 4
 Server Version: 20.10.23
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
  userxattr: false
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Cgroup Version: 2
 ……
 ……
 ……
```



> docker image ls

可以简写为```docker images```

查看所有容器镜像

```sh
[~] docker images                                                                                                                                                                                 14:49:55
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
redis        latest    cedd18a74eff   2 weeks ago     111MB
alpine/git   latest    9793ee61fc75   4 months ago    43.4MB
ubuntu       latest    d5ca7a445605   18 months ago   65.6MB
ubuntu       15.10     9b9cb95443b5   6 years ago     137MB
```



> docker pull imagename

去远程拉取镜像，这里的远程指```docker Hub```



> docker images rm imagename

删除镜像

```sh
docker image rm ubuntu
```



> docker container ls 

查看正在运行在容器，也可以使用：```docker ps```

```sh
[~] docker container ls                                                                                                                                                                           15:19:39
CONTAINER ID   IMAGE          COMMAND                  CREATED       STATUS      PORTS                    NAMES
5a6f69f16a45   cedd18a74eff   "docker-entrypoint.s…"   9 days ago    Up 9 days   6379/tcp                 angry_euler
1e3b9cd32745   redis:latest   "docker-entrypoint.s…"   10 days ago   Up 4 days   0.0.0.0:6379->6379/tcp   competent_wilson
```



> docker container run imagename

使用镜像创建容器并运行容器，如果本地没有镜像，则去远程拉取并创建运行



> docker container stop containerID/name

停止容器运行





### Docker镜像和容器

#### docker镜像

docker是一个```read-only```文件

包含文件系统，源码，库函数，依赖和工具等application所需要的文件

可以理解为是一个模版



#### docker容器

docker的容器可以理解为是docker镜像的一个运行态

实质是复制docker镜像并在镜像上层加了```read-write```

一个镜像可以创建多个容器



#### docker镜像的获取

* 通过自己制作
* 拉取镜像（例如在docker hub）





### 创建容器

在docker创建容器，就是指把对应docker进行image运行起来，可以使用命令，该命令指，创建某一个容器并运行， 如果在本地没有找到对应镜像就回去远程拉取镜像，然后使用该镜像创建容器并运行：

```
docker container run imagename
```

例如我们创建一个Nginx容器

```sh
[~] docker container run nginx                                                                                                                                                                    15:24:39
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
927a35006d93: Pull complete
fc3910c70f9c: Pull complete
e11bfbf9fd54: Pull complete
fbb8b547daa2: Pull complete
0f1992aeebd8: Pull complete
f929dacee378: Pull complete
Digest: sha256:0d17b565c37bcbd895e9d92315a05c1c3c9a29f762b011a10c54a66cd53c9b31
Status: Downloaded newer image for nginx:latest
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
```

当我们终端输入```docker container ls  ```就可以看到，对应的容器。

```sh
[~] docker container ls                                                                                                                                                                           15:29:10
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS         PORTS                    NAMES
c814bb0656c2   nginx          "/docker-entrypoint.…"   3 minutes ago   Up 3 minutes   80/tcp                   modest_euler
5a6f69f16a45   cedd18a74eff   "docker-entrypoint.s…"   9 days ago      Up 9 days      6379/tcp                 angry_euler
1e3b9cd32745   redis:latest   "docker-entrypoint.s…"   10 days ago     Up 4 days      0.0.0.0:6379->6379/tcp   competent_wilson
```



这里需要解释一些东西：

* CONTAINER ID

  容器的id

* IMAGE

  创建容器使用的镜像

* COMMAND

​		执行的命令

* CREATED

​	 	创建时间

*  STATUS

  容器状态

* PORTS

  运行端口

* NAMES

  名字是系统随机给我们当前容器的



### 批量操作容器

当有多个容器时，我们可以使用命令：

```
docker container rm id1 id2 id3 id4                                                 
```

例如：

```sh
[~] docker ps -a                                                                                                                                                                                  15:55:07
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS                   PORTS                    NAMES
46f2f834f206   nginx          "/docker-entrypoint.…"   6 minutes ago    Up 6 minutes             80/tcp                   hardcore_blackburn
cbbdae41ada5   nginx          "/docker-entrypoint.…"   6 minutes ago    Up 6 minutes             80/tcp                   nifty_sammet
eb028035b1f2   nginx          "/docker-entrypoint.…"   6 minutes ago    Up 6 minutes             80/tcp                   pensive_bohr
c814bb0656c2   nginx          "/docker-entrypoint.…"   33 minutes ago   Up 33 minutes            80/tcp                   modest_euler
b0a8203fc448   ubuntu:15.10   "/bin/echo 'Hello wo…"   9 days ago       Exited (0) 9 days ago                             quirky_curie
2529798509bc   d5ca7a445605   "bash"                   9 days ago       Exited (0) 9 days ago                             gifted_johnson
06eecc8dce25   ubuntu         "bash"                   9 days ago       Exited (0) 9 days ago                             nifty_wing
9232372963cb   d5ca7a445605   "bash"                   9 days ago       Exited (0) 9 days ago                             dreamy_jones
5a6f69f16a45   cedd18a74eff   "docker-entrypoint.s…"   9 days ago       Up 9 days                6379/tcp                 angry_euler
1e3b9cd32745   redis:latest   "docker-entrypoint.s…"   11 days ago      Up 4 days                0.0.0.0:6379->6379/tcp   competent_wilson
22e1fbf2426c   alpine/git     "git clone https://g…"   11 days ago      Exited (0) 11 days ago                            repo
[~] docker container rm b0a8203fc44 06eecc8dce25 9232372963cb 2529798509bc                                                                                                                        15:58:58
b0a8203fc44
06eecc8dce25
9232372963cb
2529798509bc
```

但是这样所还是不方便，接下来介绍另一种方法

```sh
docker ps -aq
```

进行批量操作，```docker ps -aq```返回所有容器id

```sh
[~] docker ps -aq                                                                                                                                                                                 16:00:31
46f2f834f206
cbbdae41ada5
eb028035b1f2
c814bb0656c2
5a6f69f16a45
1e3b9cd32745
22e1fbf2426c
```

比如我们要停止所有容器

```sh
[~] docker ps                                                                                                                                                                                    16:02:11
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                    NAMES
46f2f834f206   nginx          "/docker-entrypoint.…"   10 minutes ago   Up 10 minutes   80/tcp                   hardcore_blackburn
cbbdae41ada5   nginx          "/docker-entrypoint.…"   10 minutes ago   Up 10 minutes   80/tcp                   nifty_sammet
eb028035b1f2   nginx          "/docker-entrypoint.…"   10 minutes ago   Up 10 minutes   80/tcp                   pensive_bohr
c814bb0656c2   nginx          "/docker-entrypoint.…"   37 minutes ago   Up 37 minutes   80/tcp                   modest_euler
5a6f69f16a45   cedd18a74eff   "docker-entrypoint.s…"   9 days ago       Up 9 days       6379/tcp                 angry_euler
1e3b9cd32745   redis:latest   "docker-entrypoint.s…"   11 days ago      Up 4 days       0.0.0.0:6379->6379/tcp   competent_wilson
[~] docker container stop $(docker ps -aq)                                                                                                                                                        16:02:58
46f2f834f206
cbbdae41ada5
eb028035b1f2
c814bb0656c2
5a6f69f16a45
1e3b9cd32745
22e1fbf2426c
[~] docker ps                                                                                                                                                                                     16:03:35
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

当然也可以将容器运行起来：

```sh
[~] docker ps -aq                                                                                                                                                                                 16:03:56
46f2f834f206
cbbdae41ada5
eb028035b1f2
c814bb0656c2
5a6f69f16a45
1e3b9cd32745
22e1fbf2426c
[~] docker container start $(docker ps -aq)
46f2f834f206
cbbdae41ada5
eb028035b1f2
c814bb0656c2
5a6f69f16a45
1e3b9cd32745
22e1fbf2426c
```





### 容器的attached和detached模式

#### attached

attached模式就是指将容器在前台运行，窗口不能退出，假如我们使用命令```control+c```退出前台，那么对应的容器也就停止运行了

例如: 当我们这样运行容器后，日志都会输出在终端里

```sh
[~] docker container run -p 80:80 nginx     #运行容器，将docker运行端口映射到宿主机器端口，我们可以在浏览器中127.0.0.1进行访问                                                                                                                                                 16:15:53
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2023/04/10 08:16:28 [notice] 1#1: using the "epoll" event method
2023/04/10 08:16:28 [notice] 1#1: nginx/1.21.5
2023/04/10 08:16:28 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2023/04/10 08:16:28 [notice] 1#1: OS: Linux 5.15.49-linuxkit
2023/04/10 08:16:28 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2023/04/10 08:16:28 [notice] 1#1: start worker processes
2023/04/10 08:16:28 [notice] 1#1: start worker process 31
2023/04/10 08:16:28 [notice] 1#1: start worker process 32
2023/04/10 08:16:28 [notice] 1#1: start worker process 33
2023/04/10 08:16:28 [notice] 1#1: start worker process 34
172.17.0.1 - - [10/Apr/2023:08:16:45 +0000] "GET / HTTP/1.1" 200 615 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36" "-"
172.17.0.1 - - [10/Apr/2023:08:16:46 +0000] "GET /favicon.ico HTTP/1.1" 404 555 "http://127.0.0.1/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36" "-"
2023/04/10 08:16:46 [error] 32#32: *1 open() "/usr/share/nginx/html/favicon.ico" failed (2: No such file or directory), client: 172.17.0.1, server: localhost, request: "GET /favicon.ico HTTP/1.1", host: "127.0.0.1", referrer: "http://127.0.0.1/"
```

#### detached模式

就是让容器在后台运行，不会占用终端，不会输出日志，当终端退出后，容器依然会运行

```sh
[~] docker container run -d -p 80:80 nginx                                                                                                                                                        16:21:10
d2dbfc20f9f0e79c53c1c96d1999e2af9a4eaeb58983956629649735c470497b
```

当然我们可以将detached转到attached

```sh
[~] docker container run -d -p 80:80 nginx                                                                                                                                                        16:21:10
d2dbfc20f9f0e79c53c1c96d1999e2af9a4eaeb58983956629649735c470497b
[~] docker attach d2d                                                                                                                                                                             16:21:28
172.17.0.1 - - [10/Apr/2023:08:24:44 +0000] "GET / HTTP/1.1" 304 0 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36" "-"

```





### 容器的交互模式

有时候需要进行容器的交互，例如调试，日志等

```sh
[~] docker ps  #查看运行的容器                                                                                                                                                                                   16:30:21
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                NAMES
23efdbab4810   nginx     "/docker-entrypoint.…"   48 seconds ago   Up 47 seconds   0.0.0.0:80->80/tcp   admiring_bouman
[~] docker container logs 23e    #查看容器日志                                                                                                                                                               16:30:28
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2023/04/10 08:29:40 [notice] 1#1: using the "epoll" event method
2023/04/10 08:29:40 [notice] 1#1: nginx/1.21.5
2023/04/10 08:29:40 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2023/04/10 08:29:40 [notice] 1#1: OS: Linux 5.15.49-linuxkit
2023/04/10 08:29:40 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2023/04/10 08:29:40 [notice] 1#1: start worker processes
2023/04/10 08:29:40 [notice] 1#1: start worker process 31
2023/04/10 08:29:40 [notice] 1#1: start worker process 32
2023/04/10 08:29:40 [notice] 1#1: start worker process 33
2023/04/10 08:29:40 [notice] 1#1: start worker process 34


[~] docker container logs -f 23e  #动态监控日志
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2023/04/10 08:29:40 [notice] 1#1: using the "epoll" event method
2023/04/10 08:29:40 [notice] 1#1: nginx/1.21.5
2023/04/10 08:29:40 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)
2023/04/10 08:29:40 [notice] 1#1: OS: Linux 5.15.49-linuxkit
2023/04/10 08:29:40 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2023/04/10 08:29:40 [notice] 1#1: start worker processes
2023/04/10 08:29:40 [notice] 1#1: start worker process 31
2023/04/10 08:29:40 [notice] 1#1: start worker process 32
2023/04/10 08:29:40 [notice] 1#1: start worker process 33
2023/04/10 08:29:40 [notice] 1#1: start worker process 34
```



> docker container run -it imagename sh

使用该命令当容器运行后，进入容器shell，注意：当我们退出后，整个容器就退出了



> docker exec -it id or name sh

进入正在运行的容器shell

例如：

```sh
[~] docker ps                                                                                                                                                                                     16:41:00
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                NAMES
23efdbab4810   nginx     "/docker-entrypoint.…"   11 minutes ago   Up 11 minutes   0.0.0.0:80->80/tcp   admiring_bouman
[~] docker exec -it 23e sh                                                                                                                                                                        16:41:03
# ls
bin  boot  dev	docker-entrypoint.d  docker-entrypoint.sh  etc	home  lib  media  mnt  opt  proc  root	run  sbin  srv	sys  tmp  usr  var
# pwd
/
#
```

当使用```control+c```退出shell后，容器依然会运行



### 容器和虚拟机

![](https://dockertips.readthedocs.io/en/latest/_images/containers-vs-virtual-machines.jpg)

我们要知道容器不等于虚拟机

- 容器其实是进程Containers are just processes

- 容器中的进程被限制了对CPU内存等资源的访问

- 当进程停止后，容器就退出了

  

### docker container run 背后发生了什么？

前面介绍了容器的创建，只是简单是说明了一下run背后的工作，现在来详细介绍一下，以下面命令为例：

```
$ docker container run -d --publish 80:80 --name webhost nginx
```

- 1. 在本地查找是否有nginx这个image镜像，但是没有发现
- 1. 去远程的image registry查找nginx镜像（默认的registry是Docker Hub)
- 1. 下载最新版本的nginx镜像 （nginx:latest 默认)
- 1. 基于nginx镜像来创建一个新的容器，并且准备运行
- 1. docker engine分配给这个容器一个虚拟IP地址
- 1. 在宿主机上打开80端口并把容器的80端口转发到宿主机上
- 1. 启动容器，运行指定的命令（这里是一个shell脚本去启动nginx）





## docker镜像



### 镜像的获取

* 通过registry去拉取
  * 公有：比如docker hub别人都能拉取
  * 私有：比如企业内部制作的镜像，需要企业自己搭建一个registry
* 基于dockerfile自己构建
* load from `file` (offline) 文件导入

<img src="https://dockertips.readthedocs.io/en/latest/_images/docker-stages.png" style="zoom:80%;" />



### 通过registry去拉取



#### 镜像拉取

> docker image pull imagename

使用上面命令进行镜像的拉取。如果没有指定版本，则默然是最新版本：latest



我们的registry默认是使用docker hub， 如果我们需要某一个镜像，可以直接去[docker hub](https://hub.docker.com/)查找，找到需要的版本，使用registry提供的命令直接拉取即可。

例如：我现在我找MongoDB的镜像，直接去docker hub搜索即可，找到需要的版本(这里我以最新版本）然后运行hub提供的命令：

```sh
docker pull mongo
```

假如我们要拉取指定版本：

```sh
docker pull mongo:4.4.20-rc0
```



#### 镜像删除

> docker image rm imageid

**注意：当需要删除的镜像有运行容器，或者退出的容器状态，他都是不可以删除的，我们需要将使用的容器删除后，才可以删除对应的镜像。**





### 镜像的导入导出

有的时候我们的其他设备没有网络但需要使用镜像或者是你只想把某一个镜像发给别人，可以使用进行的导入/导出。

#### 导出

> docker image save imagename:x.x.x -o outname.image



假如现在要导出ubuntu:15.10

```sh
[~] docker images                                                                                                                                                                                 17:58:55
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
redis        latest    cedd18a74eff   2 weeks ago     111MB
alpine/git   latest    9793ee61fc75   4 months ago    43.4MB
nginx        latest    eeb9db34b331   15 months ago   134MB
mongo        latest    b66a51a1f0e6   16 months ago   657MB
ubuntu       latest    d5ca7a445605   18 months ago   65.6MB
ubuntu       15.10     9b9cb95443b5   6 years ago     137MB
[~] docker image save ubuntu:15.10 -o ubuntu.image   #导出镜像，保存在当前目录                                                                                                                                           17:58:58
[~] ls                                                                                                                                                                                            18:00:52
Desktop         Downloads       Library         Music           Postman         ubuntu.image
Documents       GolandProjects
```



#### 导入

> docker image load -i ./your_filepath/name.image



现在我们来导入Ubuntu:15.10

```sh
[~] docker images   #镜像列表                                                                                                                                                                          18:05:26
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
redis        latest    cedd18a74eff   2 weeks ago     111MB
alpine/git   latest    9793ee61fc75   4 months ago    43.4MB
nginx        latest    eeb9db34b331   15 months ago   134MB
mongo        latest    b66a51a1f0e6   16 months ago   657MB
ubuntu       latest    d5ca7a445605   18 months ago   65.6MB
[~] docker image load -i ./ubuntu.image      #导入镜像                                                                                                                                                   18:05:33
f121afdbbd5d: Loading layer [==================================================>]  142.9MB/142.9MB
4b955941a4d0: Loading layer [==================================================>]  15.87kB/15.87kB
af288f00b8a7: Loading layer [==================================================>]  11.78kB/11.78kB
98d59071f692: Loading layer [==================================================>]  4.608kB/4.608kB
Loaded image: ubuntu:15.10
[~] docker images                                                                                                                                                                                 18:06:57
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
redis        latest    cedd18a74eff   2 weeks ago     111MB
alpine/git   latest    9793ee61fc75   4 months ago    43.4MB
nginx        latest    eeb9db34b331   15 months ago   134MB
mongo        latest    b66a51a1f0e6   16 months ago   657MB
ubuntu       latest    d5ca7a445605   18 months ago   65.6MB
ubuntu       15.10     9b9cb95443b5   6 years ago     137MB   #ubuntu:15.10
[~]
```



### DockerFile构建镜像

#### 镜像构建

下面我们直接来使用dockerfile构建一个python容器

hellp_docker.py

```py
print("hello, docker")
```

DockerFile:

```dockerfile
FROM ubuntu:20.04
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y python3.9 python3-pip python3.9-dev
ADD hello_docker.py /
CMD ["python3", "/hello_docker.py"]
```

两个文件在同一命令下然后运行：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker image build -t hello:1.0 .
Sending build context to Docker daemon  4.096kB
Step 1/4 : FROM ubuntu:20.04
20.04: Pulling from library/ubuntu
7b1a6ab2e44d: Pull complete
Digest: sha256:626ffe58f6e7566e00254b638eb7e0f3b11d4da9675088f4781a50ae288f3322
Status: Downloaded newer image for ubuntu:20.04
 ---> ba6acccedd29
Step 2/4 : RUN apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y python3.9 python3-pip python3.9-dev
 ---> Running in fd6c73a00ee1
 ……
 ……
 ……

# 最后看到Successfully构建成功
Step 4/4 : CMD ["python3", "/hello_docker.py"]
 ---> Running in 33d73376ceb6
Removing intermediate container 33d73376ceb6
 ---> d1d214b5b189
Successfully built d1d214b5b189
Successfully tagged hello:1.0
```

查看镜像：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker image ls
REPOSITORY                      TAG       IMAGE ID       CREATED         SIZE
hello                           1.0       d1d214b5b189   9 minutes ago   253MB
redis                           latest    7614ae9453d1   15 months ago   113MB
```

创建并运行容器：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker container run hello:1.0
hello, docker
```

这样的一个简单镜像就构建完成了。



#### 镜像分享

将镜像构建完成后，可能需要上传到docker hub，如果我们要上传镜像需要注意：镜像名要以```您的docker用户名/镜像名```为完整的镜像名称，才能上传成功。

所以需要重新构建镜像为：```iceymoss/hello:1.0```，这里的构建可以使用原来的镜像直接生成新的指定tag镜像，当然我们也可以直接删除之前的镜像：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker image rm hello:1.0
Untagged: hello:1.0
```

构建```iceymoss/hello:1.0```镜像：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker image build -t iceymoss/hello:1.0 .
Sending build context to Docker daemon  4.096kB
Step 1/4 : FROM ubuntu:20.04
 ---> ba6acccedd29
Step 2/4 : RUN apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y python3.9 python3-pip python3.9-dev
 ---> Using cache
 ---> 4a6342e5abda
Step 3/4 : ADD hello_docker.py /
 ---> Using cache
 ---> 57b6fdef7e55
Step 4/4 : CMD ["python3", "/hello_docker.py"]
 ---> Using cache
 ---> d1d214b5b189
Successfully built d1d214b5b189
Successfully tagged iceymoss/hello:1.0
```

查看镜像：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker images
REPOSITORY                      TAG       IMAGE ID       CREATED          SIZE
iceymoss/hello                  1.0       d1d214b5b189   20 minutes ago   253MB
```

上传镜像：

* 登录dockerhub

  ```sh
  root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker login   #登录dockerhub
  Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
  Username: iceymoss  # 输入用户名
  Password:						# 输入密码
  WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
  Configure a credential helper to remove this warning. See
  https://docs.docker.com/engine/reference/commandline/login/#credentials-store
  
  Login Succeeded  #成功！
  ```

* 上传镜像

  ```sh
  root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker image push iceymoss/hello:1.0  #上传镜像
  The push refers to repository [docker.io/iceymoss/hello]
  17ebfe15a563: Pushed
  bd233375145e: Pushed
  9f54eef41275: Mounted from library/ubuntu
  1.0: digest: sha256:33e432cbe808e746eec1ef85658b40447b00a3f18299d556eb9e5a057d08e5d3 size: 948
  ```

  这样我们的镜像就成功的上传到```docker hub```我们的用户下面了,我们登录:

  进行镜像搜索：```iceymoss/hello```就可以看到上传的镜像和对应的pull命令，我们可以将容器pull来验证。

  ```sh
  docker pull iceymoss/hello:1.0
  ```

  然后创建容器并运行即可。



### 通过commit生成镜像

比如我们想将我们pull的在此基础上进行修改，例如我们运行nginx镜像，然后进入容器shell，找到nginx端口首页```index.html```，将内容修改后，我们要将修改后的内容作为我们自己需要的镜像，使用命令：

```sh
docker container commit 容器id iceymoss/nginx
```

```iceymoss/nginx```根据自己需要进行命名即可。



### 基于scratch构建镜像

scratch是docker提供的一个空镜像，他什么都没有，我们无法直接通过pull拉取他，但是他可以帮助我们构建我们自己的镜像，下面以一段golang程序来基于scratch构建镜像。

在同一目录下：

main.go:

```go
package main

import "fmt"

func main(){
	fmt.Println("hello, iceymoss!")
	fmt.Println("this is a container for docker!")
}
```

我们需要将main.go编译成可执行文件：

```
go build main.go
```

然后在同一目录下生成一个main的可执行文件



接着来编写Dockerfile文件：

```dockerfile
FROM scratch 
ADD main /
CMD ["/main"]
```

* ```FROM scratch``` 表示镜像来源
* ```ADD main /  ```表示在当前命令下添加main可执行文件到镜像```/```目录下
* ```CMD ["/main"]```执行命令

然后使用镜像构建命令：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss& docker image build -t iceymoss/hello_go:1.0 .   #构建镜像
Sending build context to Docker daemon  1.819MB
Step 1/3 : FROM scratch
 --->
Step 2/3 : ADD main /
 ---> 3a2d990df0f1
Step 3/3 : CMD ["/main"]
 ---> Running in 1f0882b7dfcd
Removing intermediate container 1f0882b7dfcd
 ---> 138f83cc9e92
Successfully built 138f83cc9e92
Successfully tagged iceymoss/hello_go:1.0
```

这里我们知道scratch本身是一个空镜像，他没有占学习空间，而对于```iceymoss/hello_go:1.0 ```镜像的空间其实就是main可执行文件的大小。



## Dockerfile指南

在上面的一部分内容我们只是简单介绍了一下镜像的构建过程，但是并没有真正的系统学习Dockerfile的语法和细节，这里我们将着重介绍Dockerfile的使用和镜像构建的细节。



### 基础镜像的选择

我们在构建镜像时，是需要选择一个基础镜像进行构建的，这里我们需要有一些选择的原则和方法。

#### 基本原则

* 官方镜像优于非官方镜像，如果没有官方镜像，那也尽量Dockerfile开源的
* 固定版本，而不是每次都使用latest
* 尽量选择体积小的镜像



例如

```dockerfile
FROM golang:alpine3.17
ADD main /
CMD ["/main"]
```

```golang:alpine3.17```有100.41 MB，但是```golang:latest```版本有301.78 MB



### RUN指令

RUN指令主要是在构建镜像的时候去下进行一些软件安装，文件下载等

```sh
$ apt-get update
$ apt-get install wget
$ wget https://github.com/ipinfo/cli/releases/download/ipinfo-2.0.1/ipinfo_2.0.1_linux_amd64.tar.gz
$ tar zxf ipinfo_2.0.1_linux_amd64.tar.gz
$ mv ipinfo_2.0.1_linux_amd64 /usr/bin/ipinfo
$ rm -rf ipinfo_2.0.1_linux_amd64.tar.gz
```

在Dockerfile构建镜像过程中，如果我们使用的```RUN```的次数越多，那么构建的镜像的层数就越多，造成的结果就是镜像体积大。

```sh
$ docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
ipinfo       latest    97bb429363fb   4 minutes ago   138MB
ubuntu       21.04     478aa0080b60   4 days ago      74.1MB
$ docker image history 97b
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
97bb429363fb   4 minutes ago   RUN /bin/sh -c rm -rf ipinfo_2.0.1_linux_amd…   0B        buildkit.dockerfile.v0
<missing>      4 minutes ago   RUN /bin/sh -c mv ipinfo_2.0.1_linux_amd64 /…   9.36MB    buildkit.dockerfile.v0
<missing>      4 minutes ago   RUN /bin/sh -c tar zxf ipinfo_2.0.1_linux_am…   9.36MB    buildkit.dockerfile.v0
<missing>      4 minutes ago   RUN /bin/sh -c wget https://github.com/ipinf…   4.85MB    buildkit.dockerfile.v0
<missing>      4 minutes ago   RUN /bin/sh -c apt-get install -y wget # bui…   7.58MB    buildkit.dockerfile.v0
<missing>      4 minutes ago   RUN /bin/sh -c apt-get update # buildkit        33MB      buildkit.dockerfile.v0
<missing>      4 days ago      /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B
<missing>      4 days ago      /bin/sh -c mkdir -p /run/systemd && echo 'do…   7B
<missing>      4 days ago      /bin/sh -c [ -z "$(apt-get indextargets)" ]     0B
<missing>      4 days ago      /bin/sh -c set -xe   && echo '#!/bin/sh' > /…   811B
<missing>      4 days ago      /bin/sh -c #(nop) ADD file:d6b6ba642344138dc…   74.1MB
```

我们在构建镜像时，编写Dokerfile时，尽量少用```RUN```最好将所有命令放在一个```RUN```中执行

#### 改进版

```sh
FROM ubuntu:20.04
RUN apt-get update && \
    apt-get install -y wget && \
    wget https://github.com/ipinfo/cli/releases/download/ipinfo-2.0.1/ipinfo_2.0.1_linux_amd64.tar.gz && \
    tar zxf ipinfo_2.0.1_linux_amd64.tar.gz && \
    mv ipinfo_2.0.1_linux_amd64 /usr/bin/ipinfo && \
    rm -rf ipinfo_2.0.1_linux_amd64.tar.gz
```

```sh
$ docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
ipinfo-new   latest    fe551bc26b92   5 seconds ago    124MB
ipinfo       latest    97bb429363fb   16 minutes ago   138MB
ubuntu       21.04     478aa0080b60   4 days ago       74.1MB
$ docker image history fe5
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
fe551bc26b92   16 seconds ago   RUN /bin/sh -c apt-get update &&     apt-get…   49.9MB    buildkit.dockerfile.v0
<missing>      4 days ago       /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B
<missing>      4 days ago       /bin/sh -c mkdir -p /run/systemd && echo 'do…   7B
<missing>      4 days ago       /bin/sh -c [ -z "$(apt-get indextargets)" ]     0B
<missing>      4 days ago       /bin/sh -c set -xe   && echo '#!/bin/sh' > /…   811B
<missing>      4 days ago       /bin/sh -c #(nop) ADD file:d6b6ba642344138dc…   74.1MB
$
```



### 文件的复制和目录操作

向镜像里复制文件的方法有两种```COPY```和```ADD```

#### 复制普通文件

`COPY` 和 `ADD` 都可以把local的一个文件复制到镜像里，如果目标目录不存在，则会自动创建

```dockerfile
FROM python:3.9.5-alpine3.13
COPY hello.py /app/hello.py
```

比如我们将本地的hello.py文件复制到镜像的```/app```目录下，实际上/app这个目录不存在，但他会自动创建

**注意：复制的文件在本地的权限也会被复制到镜像中去**



#### 复制压缩文件

`ADD` 比 COPY高级一点的地方就是，如果复制的是一个gzip等压缩文件时，ADD会帮助我们自动去解压缩文件。

```dockerfile
FROM python:3.9.5-alpine3.13
ADD hello.tar.gz /app/
```

将```hello.tar.gz```复制到镜像/app/下并解压



#### 使用原则

因此在 COPY 和 ADD 指令中选择的时候，可以遵循这样的原则，所有的文件复制均使用 COPY 指令，仅在需要自动解压缩的场合使用 ADD。



#### 目录操作

目录操作使用```WORKDIR```进行命令切换，如果没有给定的目录docker则会在镜像里面进行创建该目录， 使用该命令，镜像也会变得多一层，但是不会占用空间

```dockerfile
FROM golang:alpine3.17
WORKDIR /App/
ADD main .
CMD ["/main"]
```

添加到当前/App/目录下



### 构建参数和环境变量

我们先来看一下环境变量：

```dockerfile
FROM ubuntu:20.04
RUN apt-get update && \
    apt-get install -y wget && \
    wget https://github.com/ipinfo/cli/releases/download/ipinfo-2.0.1/ipinfo_2.0.1_linux_amd64.tar.gz && \
    tar zxf ipinfo_2.0.1_linux_amd64.tar.gz && \
    mv ipinfo_2.0.1_linux_amd64 /usr/bin/ipinfo && \
    rm -rf ipinfo_2.0.1_linux_amd64.tar.gz
```

我们可以看到有```RUN```中有很多版本好，如果要进行镜像升级的时候，岂不是每次都要进行修改，这很麻烦的。

所以我们需要环境变量

#### ENV

```dockerfile
FROM ubuntu:20.04
ENV VERSION=2.0.1
RUN apt-get update && \
    apt-get install -y wget && \
    wget https://github.com/ipinfo/cli/releases/download/ipinfo-${VERSION}/ipinfo_${VERSION}_linux_amd64.tar.gz && \
    tar zxf ipinfo_${VERSION}_linux_amd64.tar.gz && \
    mv ipinfo_${VERSION}_linux_amd64 /usr/bin/ipinfo && \
    rm -rf ipinfo_${VERSION}_linux_amd64.tar.gz
```



#### ARG

```dockerfile
FROM ubuntu:20.04
ARG VERSION=2.0.1
RUN apt-get update && \
    apt-get install -y wget && \
    wget https://github.com/ipinfo/cli/releases/download/ipinfo-${VERSION}/ipinfo_${VERSION}_linux_amd64.tar.gz && \
    tar zxf ipinfo_${VERSION}_linux_amd64.tar.gz && \
    mv ipinfo_${VERSION}_linux_amd64 /usr/bin/ipinfo && \
    rm -rf ipinfo_${VERSION}_linux_amd64.tar.gz
```



`ARG` 和 `ENV` 是经常容易被混淆的两个Dockerfile的语法，都可以用来设置一个“变量”。 但实际上两者有很多的不同。



#### 区别

![](https://dockertips.readthedocs.io/en/latest/_images/docker_environment_build_args.png)



ARG 可以在镜像build的时候动态修改value, 通过 `--build-arg`

```sh
$ docker image build -f ./Dockerfile-arg -t ipinfo-arg-2.0.0 --build-arg VERSION=2.0.0 .   #原版本为2.0.1
$ docker image ls
REPOSITORY         TAG       IMAGE ID       CREATED          SIZE
ipinfo-arg-2.0.0   latest    0d9c964947e2   6 seconds ago    124MB
$ docker container run -it ipinfo-arg-2.0.0
root@b64285579756:/#
root@b64285579756:/# ipinfo version
2.0.0
root@b64285579756:/#
```

ENV 设置的变量可以在Image中保持，并在容器中的环境变量里





### 容器启动命令CMD

```CMD```可以用来设置容器启动时默认会执行的命令

- 容器启动时默认执行的命令
- 如果docker container run启动容器时指定了其它命令，则CMD命令会被忽略
- 如果定义了多个CMD，只有最后一个会被执行。(Dockerfile文件中)

例如：

```dockerfile
FROM scratch 
ADD main /
CMD ["/main"]
```

构建镜像:

```sh
$ docker image build -t iceymoss/hello_go:1.0 .   #构建镜像
```

创建容器时：

```sh
$ docker container run iceymoss/hello_go:1.0
```

docker就会去的镜像shell中执行```/main```



##### 清除退出的命令

> docker system prune -f



##### 清除没有使用的镜像

> docker image prune -a



### 容器启动命令 ENTRYPOINT

NTRYPOINT 也可以设置容器启动时要执行的命令，但是和CMD是有区别的。

- `CMD` 设置的命令，可以在docker container run 时传入其它命令，覆盖掉 `CMD` 的命令，但是 `ENTRYPOINT` 所设置的命令是一定会被执行的。
- `ENTRYPOINT` 和 `CMD` 可以联合使用，`ENTRYPOINT` 设置执行的命令，CMD传递参数



实例：

构建三个镜像

Dockerfile-cmd:

```dockerfile
FROM ubuntu:20.04
CMD ["echo","hello,docker"]
```



Dockerfile-ent:

```dockerfile
FROM ubuntu:20.04
CMD ["echo","hello,docker"]
```



Dockerfile:

```dockerfile
FROM ubuntu:20.04
ENTRYPOINT [ "echo" ]
CMD []
```



将他们分别构建

```sh
$ docker image build -f ./Dockerfile-cmd -t dome-cmd .
```

```sh
$ docker image build -f ./Dockerfile-ent -t dome-ent .
```

```sh
$ docker image build -f ./Dockerfile -t dome-both .  
```

```sh
$ docker images                                                                                                    
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
dome-cmd     latest    f6dc13ce942a   18 months ago    65.6MB
dome-ent     latest    f6dc13ce942a   18 months ago    65.6MB
dome-both    latest    b5cb092c67ea   18 months ago    65.6MB
```

他们的大小都是一样的

```sh
$ docker container run dome-cmd echo "hi,iceymoss"                                                              
hi,iceymoss   #原内容被覆盖了
$ docker container run demo-ent echo "hi,iceymoss"                                                               
hello,docker echo hi,iceymoss    #打印原内容，因为容器一定会执行ENTRYPOINT的命令，所以我们运行容器时写入的命令也会被当做参数传入给ENTRYPOINT

$ docker container run dome-both                                                                           
					#没有输入任何参数，为空
$ docker container run dome-both "吃了吗?"    
吃了吗?   
```

#### Shell格式

```dockerfile
CMD echo "hello docker"
```



```dockerfile
ENTRYPOINT echo "hello docker"
```

#### Exec格式

以可执行命令的方式

```dockerfile
ENTRYPOINT ["echo", "hello docker"]
```



```dockerfile
CMD ["echo", "hello docker"]
```



注意shell脚本的问题

```dockerfile
FROM ubuntu:20.04
ENV NAME=docker
CMD echo "hello $NAME"
```



假如我们要把上面的CMD改成Exec格式，下面这样改是不行的, 大家可以试试。

```dockerfile
FROM ubuntu:20.04
ENV NAME=docker
CMD ["echo", "hello $NAME"]
```



它会打印出 `hello $NAME` , 而不是 `hello docker` ,那么需要怎么写呢？ 我们需要以shell脚本的方式去执行

```dockerfile
FROM ubuntu:20.04
ENV NAME=docker
CMD ["sh", "-c", "echo hello $NAME"]
```



#### 构建一个python服务

Python 程序

```python
from flask import Flask

app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Hello, World!'
```

Dockerfile

```go
FROM python:3.9.5-slim

COPY app.py /src/app.py

RUN pip install flask

WORKDIR /src
ENV FLASK_APP=app.py

EXPOSE 5000

CMD ["flask", "run", "-h", "0.0.0.0"]
```

构建：

```sh
$ docker image build -t flask-demo .
```

运行：

```sh
$ docker run -d -p 5000:5000 flask-demo
```



#### 构建一个goweb服务

由于我对go比较熟，所以使用go也来构建一个镜像吧

在此之前你应该需要安装gin框架

```sh
go get -u github.com/gin-gonic/gin
```

main.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handle方法
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "ice_moss",
		"age":    18,
		"school": "家里蹲大学",
	})
}

func main() {
	//初始化一个gin的server对象
	//Default实例化对象具有日志和返回状态功能
	r := gin.Default()
	//注册路由，并编写处理方法
	r.GET("/ping", Pong)
	//监听端口：默认端口listen and serve on 0.0.0.0:8080
	r.Run(":8083")
}
```



然后直接编译成可执行文件(当然我们可以将源码复制到基础镜像中再编译也可以)

```sh
go build main main.go
```

然后来编写Dockerfile文件

```dockerfile
FROM ubuntu:latest
COPY main /
EXPOSE 8083
CMD ["/main"]
```

然后构建：

```sh
docker image build -t gin-gemo:1.0 .
```

最后如将端口映射到宿主机器端口上：8083

```
docker run -d -p 8083:8083 gin-demo:1.0
```

直接访问：127.0.0.1:8083

但是这里要注意：我实在Linux服务器演示的，如果您的是M1的mac就会出现CPU架构不兼容问题，你需要找到对应架构的基础镜像。



### Dockerfile 技巧——合理使用 .dockerignore

Docker是client-server架构，理论上Client和Server可以不在一台机器上。

在构建docker镜像的时候，需要把所需要的文件由CLI（client）发给Server，这些文件实际上就是build context

举例：

```sh
$ dockerfile-demo more Dockerfile
FROM python:3.9.5-slim

RUN pip install flask

WORKDIR /src
ENV FLASK_APP=app.py

COPY app.py /src/app.py

EXPOSE 5000

CMD ["flask", "run", "-h", "0.0.0.0"]
$ dockerfile-demo more app.py
from flask import Flask

app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Hello, world!'
```

构建的时候，第一行输出就是发送build context。11.13MB （这里是Linux环境下的log）

```sh
$ docker image build -t demo .
Sending build context to Docker daemon  11.13MB
Step 1/7 : FROM python:3.9.5-slim
 ---> 609da079b03a
Step 2/7 : RUN pip install flask
 ---> Using cache
 ---> 955ce495635e
Step 3/7 : WORKDIR /src
 ---> Using cache
 ---> 1c2f968e9f9b
Step 4/7 : ENV FLASK_APP=app.py
 ---> Using cache
 ---> dceb15b338cf
Step 5/7 : COPY app.py /src/app.py
 ---> Using cache
 ---> 0d4dfef28b5f
Step 6/7 : EXPOSE 5000
 ---> Using cache
 ---> 203e9865f0d9
Step 7/7 : CMD ["flask", "run", "-h", "0.0.0.0"]
 ---> Using cache
 ---> 35b5efae1293
Successfully built 35b5efae1293
Successfully tagged demo:latest
```



### .dockerignore 文件

相信使用git的同学都知道，.gitignore 文件， 我们使用他对应的语法，可以指定哪些文件进行上传，哪些文件涉及隐私，不需要生上传等。

目录结构如下：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker_go$ ls
Dockerfile  go.mod  go.sum  main  main.go
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker_go$
```

编写一个.dockerignore文件：

```
/go.mod
/go.sum
/main.go
```

下面再进行构建：

```sh
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker_go$ docker image build -t gin-demo:1.1 .
Sending build context to Docker daemon   10.5MB    #文件大小就变小了，镜像的构建会更快
Step 1/4 : FROM ubuntu:latest
 ---> ba6acccedd29
Step 2/4 : COPY main /
 ---> Using cache
 ---> d236d1c477d1
Step 3/4 : EXPOSE 8083
 ---> Using cache
 ---> a1d844c6e4ca
Step 4/4 : CMD ["/main"]
 ---> Using cache
 ---> f002e06a38b9
Successfully built f002e06a38b9
Successfully tagged gin-demo:1.1
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker_go#
```



### 多阶段构建

当我们本地没有对应编程语言的开发环境，但是有需要使用运行我们写出来程序，或者是本地CPU架构不一致，我们不能直接在本地直接编译成可执行文件去构建镜像，那么我们可以在基础镜像中去配置对应程序的编译和运行环境然后再将源码在镜像中镜像编译，最后运行容器即可。

例外我现在本地没有go的开发环境，然后我们来构建一个go的web服务的镜像。

main.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handle方法
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "ice_moss",
		"age":    18,
		"school": "家里蹲大学",
	})
}

func main() {
	//初始化一个gin的server对象
	//Default实例化对象具有日志和返回状态功能
	r := gin.Default()
	//注册路由，并编写处理方法
	r.GET("/ping", Pong)
	//监听端口：默认端口listen and serve on 0.0.0.0:8080
	r.Run(":8083")
}
```

Dockerfile:

```dockerfile
FROM golang:alpine3.17
COPY main.go /src/
COPY go.mod /src/
COPY go.sum /src/

ENV GO111MODULE=on

WORKDIR /src
RUN go mod download && \
    go build main.go

EXPOSE 8083
CMD ["/src/main"]
```

然后构建：

```sh
docker image build -t gin-demo:1.0 .
```

构建成功后：

```sh
[web] docker images                                                                                           
REPOSITORY        TAG       IMAGE ID       CREATED          SIZE
gin-demo          1.0       0addd4b55b80   4 minutes ago    519MB    #一个简单的web应用居然这么大！我不能接受
alpine/git        latest    9793ee61fc75   4 months ago     43.4MB
nginx             latest    eeb9db34b331   15 months ago    134MB
ubuntu            latest    d5ca7a445605   18 months ago    65.6MB
```

原因是基础镜像golang:alpine3.17太大了

我们可以创建容器：

```sh
docker container run -d -p 8083:8083 gin-demo:1.0
```



#### 多阶段构建

我们可以先将man.go在golang:alpine3.17编译成可执行文件，然后再将可执行文件放入ubuntu:latest镜像中

```dockerfile
FROM golang:alpine3.17 AS builder
COPY main.go /src/
COPY go.mod /src/
COPY go.sum /src/   
 
ENV GO111MODULE=on

WORKDIR /src
RUN go mod download && \
    go build main.go

FROM ubuntu:20.04
COPY --from=builder /src/main /src/main

EXPOSE 8083
CMD ["/src/main"]
```

然后构建：

```
docker image build -t gin-demo:2.0 .
```

可以看到：

```sh
[web] docker images                                                                                                                                                                                                                 21:42:53
REPOSITORY        TAG       IMAGE ID       CREATED          SIZE
gin-demo          2.0       c49a0acbcf4c   2 minutes ago    75.7MB  #可以看到2.0变得很小了
gin-demo          1.0       0addd4b55b80   19 minutes ago   519MB
ubuntu            latest    d5ca7a445605   18 months ago    65.6MB
```

这样我的的web应用就得到了瘦身，总体上效果还可以的，这就是多阶段构建。



### 尽量使用非root用户

> 这里在操作需要在Linux系统下

#### Root的危险性

docker的root权限一直是其遭受诟病的地方，docker的root权限有那么危险么？我们举个例子。

假如我们有一个用户，叫demo，它本身不具有sudo的权限，所以就有很多文件无法进行读写操作，比如/root目录它是无法查看的。

```sh
[demo@docker-host ~]$ sudo ls /root
[sudo] password for demo:
demo is not in the sudoers file.  This incident will be reported.
[demo@docker-host ~]$
```



但是这个用户有执行docker的权限，也就是它在docker这个group里。

```sh
[demo@docker-host ~]$ groups
demo docker
[demo@docker-host ~]$ docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED      SIZE
busybox      latest    a9d583973f65   2 days ago   1.23MB
[demo@docker-host ~]$
```



这时，我们就可以通过Docker做很多越权的事情了，比如，我们可以把这个无法查看的/root目录映射到docker container里，你就可以自由进行查看了。

```sh
[demo@docker-host vagrant]$ docker run -it -v /root/:/root/tmp busybox sh
/ # cd /root/tmp
~/tmp # ls
anaconda-ks.cfg  original-ks.cfg
~/tmp # ls -l
total 16
-rw-------    1 root     root          5570 Apr 30  2020 anaconda-ks.cfg
-rw-------    1 root     root          5300 Apr 30  2020 original-ks.cfg
~/tmp #
```



更甚至我们可以给我们自己加sudo权限。我们现在没有sudo权限

```sh
[demo@docker-host ~]$ sudo vim /etc/sudoers
[sudo] password for demo:
demo is not in the sudoers file.  This incident will be reported.
[demo@docker-host ~]$
```



但是我可以给自己添加。

```sh
[demo@docker-host ~]$ docker run -it -v /etc/sudoers:/root/sudoers busybox sh
/ # echo "demo    ALL=(ALL)       ALL" >> /root/sudoers
/ # more /root/sudoers | grep demo
demo    ALL=(ALL)       ALL
```



然后退出container，bingo，我们有sudo权限了。

```sh
[demo@docker-host ~]$ sudo more /etc/sudoers | grep demo
demo    ALL=(ALL)       ALL
[demo@docker-host ~]$
```

#### 如何使用非root用户

我们准备两个Dockerfile，第一个Dockerfile如下，其中app.py文件源码请参考 [一起构建一个 Python Flask 镜像](https://dockertips.readthedocs.io/en/latest/dockerfile-guide/python-flask.html#python-flask) ：

```dockerfile
FROM python:3.9.5-slim

RUN pip install flask

COPY app.py /src/app.py

WORKDIR /src
ENV FLASK_APP=app.py

EXPOSE 5000

CMD ["flask", "run", "-h", "0.0.0.0"]
```



假设构建的镜像名字为 `flask-demo`

第二个Dockerfile，使用非root用户来构建这个镜像，名字叫 `flask-no-root` Dockerfile如下：

- 通过groupadd和useradd创建一个flask的组和用户
- 通过USER指定后面的命令要以flask这个用户的身份运行

```dockerfile
FROM python:3.9.5-slim

RUN pip install flask && \
    groupadd -r flask && useradd -r -g flask flask && \
    mkdir /src && \
    chown -R flask:flask /src

USER flask

COPY app.py /src/app.py

WORKDIR /src
ENV FLASK_APP=app.py

EXPOSE 5000

CMD ["flask", "run", "-h", "0.0.0.0"]
```



```sh
$ docker image ls
REPOSITORY      TAG          IMAGE ID       CREATED          SIZE
flask-no-root   latest       80996843356e   41 minutes ago   126MB
flask-demo      latest       2696c68b51ce   49 minutes ago   125MB
python          3.9.5-slim   609da079b03a   2 weeks ago      115MB
```



分别使用这两个镜像创建两个容器

```sh
$ docker run -d --name flask-root flask-demo
b31588bae216951e7981ce14290d74d377eef477f71e1506b17ee505d7994774
$ docker run -d --name flask-no-root flask-no-root
83aaa4a116608ec98afff2a142392119b7efe53617db213e8c7276ab0ae0aaa0
$ docker container ps
CONTAINER ID   IMAGE           COMMAND                  CREATED          STATUS          PORTS      NAMES
83aaa4a11660   flask-no-root   "flask run -h 0.0.0.0"   4 seconds ago    Up 3 seconds    5000/tcp   flask-no-root
b31588bae216   flask-demo      "flask run -h 0.0.0.0"   16 seconds ago   Up 15 seconds   5000/tcp   flask-root
```



### 导航

[Dockerfile](https://docs.docker.com/reference/)

[github/docker](https://github.com/docker-library/official-images)



## docker的存储

