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

#### detached

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

我们知道运行的容器时，是在image加上一层```read-write```层，我们所需要的数据也会存储在该层里面，但是这些数据随着容器的删除而删除了，例如数据库容器等，这些数据一定不能随着容器的删除而删除，这是非常危险的，所以必须要对docker容器数据进行持久化。

默认情况下，在运行中的容器里创建的文件，被保存在一个可写的容器层：

- 如果容器被删除了，则数据也没有了
- 这个可写的容器层是和特定的容器绑定的，也就是这些数据无法方便的和其它容器共享

Docker主要提供了两种方式做数据的持久化

- Data Volume, 由Docker管理，(/var/lib/docker/volumes/ Linux), 持久化数据的最好方式
- Bind Mount，由用户指定存储的数据具体mount在系统什么位置

![](https://dockertips.readthedocs.io/en/latest/_images/types-of-mounts.png)

### 持久化之data-volume

#### VOLUME的使用

Dockerfile：构建一个计划任务的镜像，定时向容器中```/app/my-cron```写入时间数据

```dockerfile
FROM alpine:latest
RUN apk update
RUN apk --no-cache add curl
ENV SUPERCRONIC_URL=https://github.com/aptible/supercronic/releases/download/v0.1.12/supercronic-linux-amd64 \
    SUPERCRONIC=supercronic-linux-amd64 \
    SUPERCRONIC_SHA1SUM=048b95b48b708983effb2e5c935a1ef8483d9e3e
RUN curl -fsSLO "$SUPERCRONIC_URL" \
    && echo "${SUPERCRONIC_SHA1SUM}  ${SUPERCRONIC}" | sha1sum -c - \
    && chmod +x "$SUPERCRONIC" \
    && mv "$SUPERCRONIC" "/usr/local/bin/${SUPERCRONIC}" \
    && ln -s "/usr/local/bin/${SUPERCRONIC}" /usr/local/bin/supercronic
COPY my-cron /app/my-cron
WORKDIR /app

VOLUME ["/app"]

# RUN cron job
CMD ["/usr/local/bin/supercronic", "/app/my-cron"]
```

**注意：```VOLUME ["/app"]```他会将容器里面的```/app```目录的所有内容持久化到宿主机磁盘**



#### 构建镜像

```
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker-demo/ch01# ls
Dockerfile  my-cron
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker-demo/ch01# docker image build -t my-cron .
```



#### 创建容器-不指定参数v

这里我们要了解两个命令：

> docker volume ls    #展示docker本地持久化文件



> docker volume inspect filename  #获取对应docker本地文件相关信息



此时Docker会自动创建一个随机名字的volume，去存储我们在Dockerfile定义的volume `VOLUME["/app"]`

```sh
$ docker run -d my-cron   #运行容器
9a8fa93f03c42427a498b21ac520660752122e20bcdbf939661646f71d277f8f
$ docker volume ls    #获取持久化的本地文件
DRIVER    VOLUME NAME
local     4ffa38075aa01b717256f53275ed5a2860e4c4f8df61e0a312628e59b8a90ed3
$ docker inspect 4ffa38075aa01b717256f53275ed5a2860e4c4f8df61e0a312628e59b8a90ed3   #获取持久化信息
[
    {
        "CreatedAt": "2023-04-15T16:30:38+08:00",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/4ffa38075aa01b717256f53275ed5a2860e4c4f8df61e0a312628e59b8a90ed3/_data",
        "Name": "4ffa38075aa01b717256f53275ed5a2860e4c4f8df61e0a312628e59b8a90ed3",
        "Options": null,
        "Scope": "local"
    }
]
```

输出的这个Volume的mountpoint可以发现容器创建的文件



#### 创建容器-指定参数v

在创建容器的时候通过 `-v` 参数我们可以手动的指定需要创建Volume的名字，以及对应于容器内的路径，这个路径是可以任意的，不必需要在Dockerfile里通过VOLUME定义

比如我们把上面的Dockerfile里的VOLUME删除

```dockerfile
FROM alpine:latest
RUN apk update
RUN apk --no-cache add curl
ENV SUPERCRONIC_URL=https://github.com/aptible/supercronic/releases/download/v0.1.12/supercronic-linux-amd64 \
    SUPERCRONIC=supercronic-linux-amd64 \
    SUPERCRONIC_SHA1SUM=048b95b48b708983effb2e5c935a1ef8483d9e3e
RUN curl -fsSLO "$SUPERCRONIC_URL" \
    && echo "${SUPERCRONIC_SHA1SUM}  ${SUPERCRONIC}" | sha1sum -c - \
    && chmod +x "$SUPERCRONIC" \
    && mv "$SUPERCRONIC" "/usr/local/bin/${SUPERCRONIC}" \
    && ln -s "/usr/local/bin/${SUPERCRONIC}" /usr/local/bin/supercronic
COPY my-cron /app/my-cron
WORKDIR /app

# RUN cron job
CMD ["/usr/local/bin/supercronic", "/app/my-cron"]
```

然后从新构建容器：

```
root@VM-0-6-ubuntu:/home/ubuntu/iceymoss/docker-demo/ch01# docker image build -t my-cron .
```

重新build镜像，然后创建容器，加-v参数

> docker container run -d -v 持久化到本地文件名称:需要进行持久化的命令或者文件

例如：

> docker container run -d -v cron-data:/app my-cron

如果我们的容器不小心被删除了，我们之前的持久化文件还在，我们新建容器只需要使用该命令，指定相同文件或者目录即可。

```sh
$ docker container run -d -v cron-data:/app my-cron
43c6d0357b0893861092a752c61ab01bdfa62ea766d01d2fcb8b3ecb6c88b3de
$ docker volume ls
DRIVER    VOLUME NAME
local     cron-data
$ docker volume inspect cron-data
[
    {
        "CreatedAt": "2023-04-15T16:30:38+08:00",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/cron-data/_data",
        "Name": "cron-data",
        "Options": null,
        "Scope": "local"
    }
]
$ ls /var/lib/docker/volumes/cron-data/_data
my-cron
$ ls /var/lib/docker/volumes/cron-data/_data
my-cron  test.txt
```

Volume也创建了。



#### 数据清理

强制删除所有容器，系统清理和volume清理

```sh
$ docker rm -f $(docker container ps -aq)
$ docker system prune -f
$ docker volume prune -f  
```



### mysql数据持久化实践

现在我们来实践使用volume对MySQL数据库容器的数据做持久化

#### 拉取mysq

```sh
docker pull mysql:latest
```



#### 运行mysql容器

```sh
[~] docker container run --name demo-mysql -e MYSQL_ROOT_PASSWORD=123456 -d -p 3307:3306 -v mysql-data:/var/lib/mysql mysql:latest
```

#### 说明

* ```--name```：运行的容器名称

* ```-e MYSQL_ROOT_PASSWORD```：表示当前MySQL容器的密码
* ```-p 3307:3306```：有容器端口3306映射到宿主机端口
* ```-v mysql-data:/var/lib/mysql```：将容器中的```/var/lib/mysql```持久化到宿主机```mysql-data```中

这里由于我们本地mysql占用3306，所以我开3307



然后我们之间连接MySQL服务(或者直接进入MySQL容器中)，写入如下内容：建立一个user数据库，新建一个user_info表

```sql
CREATE DATABASE user
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE TABLE user_info(
    id int COMMENT '用户id',
    name VARCHAR(30) COMMENT '姓名',
    age INT COMMENT '年龄',
    gender VARCHAR(10) COMMENT '性别'
)COMMENT '用户基本信息表';
```



然后我们来查看运行持久化文件

```sh
[~] docker volume ls                                                                                                                                                                              
DRIVER    VOLUME NAME
local     mysql-data
[~] docker inspect mysql-data                                                                                                                                                                     
[
    {
        "CreatedAt": "2023-04-15T09:17:49Z",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/mysql-data/_data",
        "Name": "mysql-data",
        "Options": null,
        "Scope": "local"
    }
]
```

可以查看：```/var/lib/docker/volumes/mysql-data/_data```

```sh
$ /var/lib/docker/volumes/mysql-data/_data
auto.cnf    client-cert.pem  ib_buffer_pool  ibdata1  performance_schema  server-cert.pem
ca-key.pem  client-key.pem   ib_logfile0     ibtmp1   private_key.pem     server-key.pem
ca.pem      user             ib_logfile1     mysql    public_key.pem      sys
```

**注意：如果是mac和win系统是直接查看不了```/var/lib/docker/volumes/mysql-data/_data```，原因是docker运行在Linux虚拟机中**

**当我们把MySQL这个容器删除后，然后重新运行一个MySQL容器，持久化的数据会被导入新的容器中**





### 持久化之Bind Mount

前面说过：如果是mac和win系统是直接查看不了```/var/lib/docker/volumes/mysql-data/_data```；现在可以直接使用BindMount将容器中的内容映射到宿主机上的指定目录之下。

依然使用MySQL容器作为例子：

当前目录下什么都没有：

```sh
[mysql-data] ls
[mysql-data] 
[mysql-data] pwd                                                                                             
/Users/iceymoss/iceymoss/dockerlearn/mysql-data
```

然后运行：

```sh
[mysql-data] docker container run --name demo-mysql -e MYSQL_ROOT_PASSWORD=123456 -d -p 3307:3306 -v $(pwd):/var/lib/mysql mysql:latest
```

说明：

* ```-v $(pwd):/var/lib/mysql```：表示将宿主机当前目录下进行持久化映射到mysql容器中的```/var/lib/mysql```



只需要这样我们就完成了容器数据持久化到宿主机的指定目录下。

进入容器shell或者mysql连接工具，创建两个数据库：

```mysql
CREATE DATABASE good
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE user
    DEFAULT CHARACTER SET = 'utf8mb4';
```

启动容器后可以查看当前目录的情况：

```
[mysql-data] ls                                                                                                 
#ib_16384_0.dblwr  #innodb_temp       binlog.000002      ca.pem             good               ibtmp1             mysql.sock         public_key.pem     sys                user
#ib_16384_1.dblwr  auto.cnf           binlog.index       client-cert.pem    ib_buffer_pool     mysql              performance_schema server-cert.pem    undo_001
#innodb_redo       binlog.000001      ca-key.pem         client-key.pem     ibdata1            mysql.ibd          private_key.pem    server-key.pem     undo_002

```

可以看到我们创建的数据库。

和volume一样，如果将当前MySQL的容器删除，然后重新创建一个MySQL容器，只需要将宿主机```/Users/iceymoss/iceymoss/dockerlearn/mysql-data```重新映射到容器中```/var/lib/mysql```就可以持久化数据重新加载到容器中。



### 基于docker搭建开发环境

假如我们在本地没有安装对应编程语言的开发环境，这里以golang例，我们可以将编写好的golang程序使用BindMout上传至有golang开发环境的基础镜像中去，然后进行编译运行。

举个例子：

main.go

```go
package "main"
import "fmt"

func main(){
	fmt.Println("hello,docker")
}
```

然后使用命令：

```sh
[ch01] docker container run -it -v $(pwd):/root golang:alpine3.17
```

进入容器：

```
~ # cd /root
~ # ls
main.go
~ # go run main.go
hello,docker
~ # go build main.go
~ # ls
main     main.go
~ # ./main
hello,docker
```

这样我们的程序就运行和编译了。



### 多个机器之间的容器共享数据

![multi-host-volume](https://dockertips.readthedocs.io/en/latest/_images/volumes-shared-storage.png)

官方参考链接 https://docs.docker.com/storage/volumes/#share-data-among-machines

Docker的volume支持多种driver。默认创建的volume driver都是local

```sh
$ docker volume inspect vscode
[
    {
        "CreatedAt": "2021-06-23T21:33:57Z",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/vscode/_data",
        "Name": "vscode",
        "Options": null,
        "Scope": "local"
    }
]
```



这一节我们看看一个叫sshfs的driver，如何让docker使用不在同一台机器上的文件系统做volume

#### 环境准备

准备三台Linux机器，之间可以通过SSH相互通信。

| hostname     | ip             | ssh username | ssh password |
| ------------ | -------------- | ------------ | ------------ |
| docker-host1 | 192.168.200.10 | vagrant      | vagrant      |
| docker-host2 | 192.168.200.11 | vagrant      | vagrant      |
| docker-host3 | 192.168.200.12 | vagrant      | vagrant      |

#### 安装plugin

在其中两台机器上安装一个plugin `vieux/sshfs`

```sh
[vagrant@docker-host1 ~]$ docker plugin install --grant-all-permissions vieux/sshfs
latest: Pulling from vieux/sshfs
Digest: sha256:1d3c3e42c12138da5ef7873b97f7f32cf99fb6edde75fa4f0bcf9ed277855811
52d435ada6a4: Complete
Installed plugin vieux/sshfs
```



```sh
[vagrant@docker-host2 ~]$ docker plugin install --grant-all-permissions vieux/sshfs
latest: Pulling from vieux/sshfs
Digest: sha256:1d3c3e42c12138da5ef7873b97f7f32cf99fb6edde75fa4f0bcf9ed277855811
52d435ada6a4: Complete
Installed plugin vieux/sshfs
```



#### 创建volume

```sh
[vagrant@docker-host1 ~]$ docker volume create --driver vieux/sshfs \
                          -o sshcmd=vagrant@192.168.200.12:/home/vagrant \
                          -o password=vagrant \
                          sshvolume
```



查看

```sh
[vagrant@docker-host1 ~]$ docker volume ls
DRIVER               VOLUME NAME
vieux/sshfs:latest   sshvolume
[vagrant@docker-host1 ~]$ docker volume inspect sshvolume
[
    {
        "CreatedAt": "0001-01-01T00:00:00Z",
        "Driver": "vieux/sshfs:latest",
        "Labels": {},
        "Mountpoint": "/mnt/volumes/f59e848643f73d73a21b881486d55b33",
        "Name": "sshvolume",
        "Options": {
            "password": "vagrant",
            "sshcmd": "vagrant@192.168.200.12:/home/vagrant"
        },
        "Scope": "local"
    }
]
```

#### 创建容器挂载Volume

创建容器，挂载sshvolume到/app目录，然后进入容器的shell，在/app目录创建一个test.txt文件

```sh
[vagrant@docker-host1 ~]$ docker run -it -v sshvolume:/app busybox sh
Unable to find image 'busybox:latest' locally
latest: Pulling from library/busybox
b71f96345d44: Pull complete
Digest: sha256:930490f97e5b921535c153e0e7110d251134cc4b72bbb8133c6a5065cc68580d
Status: Downloaded newer image for busybox:latest
/ #
/ # ls
app   bin   dev   etc   home  proc  root  sys   tmp   usr   var
/ # cd /app
/app # ls
/app # echo "this is ssh volume"> test.txt
/app # ls
test.txt
/app # more test.txt
this is ssh volume
/app #
/app #
```



这个文件我们可以在docker-host3上看到

```sh
[vagrant@docker-host3 ~]$ pwd
/home/vagrant
[vagrant@docker-host3 ~]$ ls
test.txt
[vagrant@docker-host3 ~]$ more test.txt
this is ssh volume
```



## docker网络

### 介绍

docker网络主要有：

* Brigde
* Host
*  null

```sh
$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
b8ca79f73455   bridge    bridge    local
d90ff26166e3   host      host      local
82d6ae2e173d   none      null      local
```

在docker网络中，我们需要回答五个问题

* 容器为什么能获取ip地址。
* 为什么宿主机可以ping通容器ip。

* 为什么docker容器之可以ping通网络。
* 为什么容器可以ping通外网地址。
* 容器端口转发是怎么回事。



### Brigde网络

docker默认的网络是使用一个doekr0作为网络转发，类似于路由器。其通讯类型是：**bridge** 。



<img src="https://dockertips.readthedocs.io/en/latest/_images/two-container-network.png" style="zoom:50%;" />



上图是我们的一台宿主机和运行的两个容器的网络拓扑图，下面我们准备运行的容器：

```sh
$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED        STATUS        PORTS                               NAMES
033243ce3801   mysql:latest   "docker-entrypoint.s…"   24 hours ago   Up 24 hours   33060/tcp, 0.0.0.0:3307->3306/tcp   server-mysql
a51d40b5eeea   mongo:latest   "docker-entrypoint.s…"   25 hours ago   Up 25 hours   0.0.0.0:27017->27017/tcp            keen_bouman
0910b6c23f19   redis:latest   "docker-entrypoint.s…"   25 hours ago   Up 25 hours   0.0.0.0:6379->6379/tcp              keen_burnell
```

然后我们使用命令查看网络：

```sh
$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
b8ca79f73455   bridge    bridge    local
d90ff26166e3   host      host      local
82d6ae2e173d   none      null      local
```

使用命令：

```sh
$ docker network inspect b8ca79f73455
[
    {
        "Name": "bridge",
        "Id": "b8ca79f73455660003be1225a30a08e58d7575d9a0064b575e876b0294b22545",
        "Created": "2023-04-21T06:37:17.281109791Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "033243ce380110beef00dbc5f61dde95560b9f1633d3bfad70e4da0d4ccfaf42": {
                "Name": "server-mysql",
                "EndpointID": "f6561639185394753fafd19e44253640b51e80285d172b397504c86d440ea98e",
                "MacAddress": "02:42:ac:11:00:04",
                "IPv4Address": "172.17.0.4/16",
                "IPv6Address": ""
            },
            "0910b6c23f19f7c557fe925a1e9ed60912c62cb10e53fd9aad75e7efd1e195a6": {
                "Name": "keen_burnell",
                "EndpointID": "a9db648c673688f09206d06d5eb6f865aff95efbbb65583e6ab539dc14097571",
                "MacAddress": "02:42:ac:11:00:02",
                "IPv4Address": "172.17.0.2/16",
                "IPv6Address": ""
            },
            "a51d40b5eeeaa2c917284d0380dcf0a1e4ae2f6d4f674c2c27fb7bbdc7546614": {
                "Name": "keen_bouman",
                "EndpointID": "0b455515d4e5e73b597bb9e038a278b3a1cd4af9a23e125ce3cacecec374d063",
                "MacAddress": "02:42:ac:11:00:03",
                "IPv4Address": "172.17.0.3/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]
```

由上面返回结果：

```sh
"IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
```

我们看到docker0的ip和网关，**连接在docker0上的容器ip统一在172.17.0.0/16网段下**，这看上去类似于家庭中的网络结构，运行的容器就像是连接在家庭路由器的每一台设备。

同样我们可以看到，三个容器都连接在docker0上：

```sh
 "Containers": {
            "033243ce380110beef00dbc5f61dde95560b9f1633d3bfad70e4da0d4ccfaf42": {
                "Name": "server-mysql",
                "EndpointID": "f6561639185394753fafd19e44253640b51e80285d172b397504c86d440ea98e",
                "MacAddress": "02:42:ac:11:00:04",
                "IPv4Address": "172.17.0.4/16",
                "IPv6Address": ""
            },
            "0910b6c23f19f7c557fe925a1e9ed60912c62cb10e53fd9aad75e7efd1e195a6": {
                "Name": "keen_burnell",
                "EndpointID": "a9db648c673688f09206d06d5eb6f865aff95efbbb65583e6ab539dc14097571",
                "MacAddress": "02:42:ac:11:00:02",
                "IPv4Address": "172.17.0.2/16",
                "IPv6Address": ""
            },
            "a51d40b5eeeaa2c917284d0380dcf0a1e4ae2f6d4f674c2c27fb7bbdc7546614": {
                "Name": "keen_bouman",
                "EndpointID": "0b455515d4e5e73b597bb9e038a278b3a1cd4af9a23e125ce3cacecec374d063",
                "MacAddress": "02:42:ac:11:00:03",
                "IPv4Address": "172.17.0.3/16",
                "IPv6Address": ""
            }
        },
```



#### 容器为什么能获取ip地址

docker默认的网络是使用一个doekr0作为网络转发，类似于路由器。其通讯类型是：**bridge** 。

Brigde网络类型，为每一个容器分配了一个局域网ip，容器之间可以通讯。

#### 为什么宿主机可以ping通容器ip

<img src="https://dockertips.readthedocs.io/en/latest/_images/two-container-network.png" style="zoom:50%;" />

我们看到docker0这个bridge连接到了我们的宿主机网络，就像是docker链接了一个通往外网的路由一样，就类似于家庭路由器连接到了网络运营商本地的路由器。



#### 为什么docker容器之可以ping通网络

这个问题其实很好解释了，容器之间都连接了docker0，相当于两个容器在一个局域网内，肯定是可以连通网络的



#### 为什么容器可以ping通外网地址

当我们进入某一个容器内，例如直接ping```www.baidu.com```这也是能ping通的，这又是为什么呢？直接举个例子你就明白了。

在此之前，你需要明白NAT，也是会网络层的地址转换协议

例：

```sh
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS     NAMES
4f3303c84e53   busybox   "/bin/sh -c 'while t…"   49 minutes ago   Up 49 minutes             box2
03494b034694   busybox   "/bin/sh -c 'while t…"   49 minutes ago   Up 49 minutes             box1
```

box1的ip:```172.17.0.3```， box4的ip:```172.17.0.3```

docker0的IP：```172.17.0.1```， 宿主机IP：```192.168.10.4```

本地路由器ip公网ip:```1.14.120.10```



当我们在容器box1中ping：```www.baidu.com```， 在百度的服务器看来其实是路由器ip:```1.14.120.10```进行的，在外网是无法感知到内网的

其过程是：box1的ip:```172.17.0.3```将数据报发送给docker0IP：```172.17.0.1```然后，docker0查看数据报目的地址然后将自己的ip写入数据报的源地址，然后转发给宿主机，然后宿主机将自己的ip填入数据报中的源地址，最后发给路由器ip公网ip:```1.14.120.10```路由器将自己的ip写入数据报中的源地址，最后发给```www.baidu.com```对应的服务器。

**发数据：172.17.0.3 --> 172.17.0.1 -->  192.168.10.4 -->  1.14.120.10 -->``` www.baidu.com```**

**返回数据：``` www.baidu.com``` -->  1.14.120.10 --> 192.168.10.4  --> 172.17.0.1  --> 172.17.0.3**



#### 容器端口转发是怎么回事

比如我现在有一个web服务器NGINX容器，他运行的NGINX容器的80端口上，现在我们从其他设备怎么进行访问到这个80端口？

##### 端口映射

我们来开这条命令：

```sh
$ docker run -d -p 8080:80 nginx:latest
```

```-p 8080:80```:表示将宿主机上的8080端口映射到容器nginx的80端口，当有请求进来时，请求会先到我们的host:8080端口，然后通过端口转发将宿主机上的8080的请求转发到容器NGINX的80端口上。我们在外网直接对宿主机ip:8080端口就可以请求到NGINXweb服务器。

下面的容器都做了端口转发

```sh
$ docker run -d -p 8080:80 nginx:latest
d9f0c712b30919c8854a7e9f06e5d86772a7dca96badbf62f4bc98093b7550c2

# iceymoss @ iceymossdeMacBook-Pro in ~ [16:01:35]
$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS         PORTS                               NAMES
d9f0c712b309   nginx:latest   "/docker-entrypoint.…"   4 seconds ago   Up 3 seconds   0.0.0.0:8080->80/tcp                competent_buck
033243ce3801   mysql:latest   "docker-entrypoint.s…"   25 hours ago    Up 25 hours    33060/tcp, 0.0.0.0:3307->3306/tcp   server-mysql
a51d40b5eeea   mongo:latest   "docker-entrypoint.s…"   25 hours ago    Up 25 hours    0.0.0.0:27017->27017/tcp            keen_bouman
0910b6c23f19   redis:latest   "docker-entrypoint.s…"   25 hours ago    Up 25 hours    0.0.0.0:6379->6379/tcp              keen_burnell
```

##### Dockerfile

我们来看一下一个goweb服务的Dockerfile

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

仔细看```EXPOSE 8083```表示容器对外暴露的端口，他默认是TCP协议端口，如果我们需要使用UDP协议就需要:

```dockerfile
EXPOSE 8083/udp
```

如果我们不编写这一行，当我们在运行这个容器时只需要：

```sh
docker container run -d -p 8083:8083 gin-demo:1.0
```

他依然能正常对外暴露我们指定的端口，这里的```EXPOSE 8083```更多的是提示使用者，该镜像需要进行端口映射。



### 自定义Brigde

docker网络默认是使用docker0来做网关，但是很多时候我们需要根据业务需要自己定义一下Brigde

```sh
$ docker network create -d bridge bridge_name
```

```sh
$ docker network create -d bridge mybridge  #新建一个bridge
9ad7eca18a46ff733eca849c008530deca34f67d9f895e007c0463e73630a24b

$ docker network ls  #查看docker网络
NETWORK ID     NAME       DRIVER    SCOPE
b8ca79f73455   bridge     bridge    local
d90ff26166e3   host       host      local
9ad7eca18a46   mybridge   bridge    local
82d6ae2e173d   none       null      local

$ docker network inspect 9ad7eca18a46  #查看mybridge
[
    {
        "Name": "mybrigde",
        "Id": "9ad7eca18a46ff733eca849c008530deca34f67d9f895e007c0463e73630a24b",
        "Created": "2023-04-22T08:52:55.674403043Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]
```

我们可以看到配置中的ip，当我们的容器连接到我们创建的bridge上，会为其分配该bridge子网的ip

#### 连接到指定Brigde

```sh
$ docker run -d --rm --name box1 busybox:latest /bin/sh -c "while true; do sleep 3600; done"
```

会默认连接到docker0

需要我们指定到：mybridge

```sh
$ docker run -d --rm --name box2 --network mybridge busybox:latest /bin/sh -c "while true; do sleep 3600; done"
3986f6f4ae91633c7c31f60faa0fa136bd884d8b971f570cbeb33563a9bf155c

$ docker run -d --rm --name box3 --network mybridge busybox:latest /bin/sh -c "while true; do sleep 3600; done"
b91b45d9a8dbe5b19672e5633de4595ed42615c3d0f38dacff4b0bb08103472f

$ docker network ls
NETWORK ID     NAME       DRIVER    SCOPE
b8ca79f73455   bridge     bridge    local
d90ff26166e3   host       host      local
9ad7eca18a46   mybridge   bridge    local
82d6ae2e173d   none       null      local

$ docker network inspect 9ad7eca18a46
[
    {
        "Name": "mybridge",
        "Id": "9ad7eca18a46ff733eca849c008530deca34f67d9f895e007c0463e73630a24b",
        "Created": "2023-04-22T08:52:55.674403043Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "3986f6f4ae91633c7c31f60faa0fa136bd884d8b971f570cbeb33563a9bf155c": {
                "Name": "box2",
                "EndpointID": "b3ca1d035c685173124e33d81039ac090dfce7f1a3862aae26bbe3424fa95f9c",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": ""
            },
            "b91b45d9a8dbe5b19672e5633de4595ed42615c3d0f38dacff4b0bb08103472f": {
                "Name": "box3",
                "EndpointID": "8f76ce7a717f77d884d3b4341aa5bd4a03d829a8133a02f28aca1c81b67cbd21",
                "MacAddress": "02:42:ac:12:00:03",
                "IPv4Address": "172.18.0.3/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]
```

#### 容器连接多个网络

可以将一个容器连接到多个网络Brigde

例如前面的box2已经连接了mybridge，现在我们要讲box2再连接到docker0的bridge网络上

```sh
$ docker network connect bridge box2
```



### host网络

前面我们学习了bridge网络，现在我们来学习一下docker的host网络，host顾名思义，就是容器直接使用宿主ip进行通讯，例如NGINX容器，使用host网络，那么他就会占用宿主机的80端口，mysq则会占用3306端口。

```sh
$ docker run -d --name nginx-web --network host nginx:latest
171e3babf161bbda519e4081d11618bbed8e561e276bafebf693687444054fc1

$ docker ps   #PORTS没有信息，相当于在本地启动了一个NGINX服务
CONTAINER ID   IMAGE            COMMAND                  CREATED             STATUS             PORTS                               NAMES
171e3babf161   nginx:latest     "/docker-entrypoint.…"   5 seconds ago       Up 4 seconds                                           nginx-web                                        box1
```

然后查看一下host相关信息：

```sh
$ docker network ls
NETWORK ID     NAME       DRIVER    SCOPE
b8ca79f73455   bridge     bridge    local
d90ff26166e3   host       host      local
9ad7eca18a46   mybrigde   bridge    local
82d6ae2e173d   none       null      local

# iceymoss @ iceymossdeMacBook-Pro in ~ [17:20:02]
$ docker network inspect d90ff26166e3
[
    {
        "Name": "host",
        "Id": "d90ff26166e3e699bf40b570c7987da53752e3ca207c10403a9f62e6f5296529",
        "Created": "2023-04-21T05:43:37.064860583Z",
        "Scope": "local",
        "Driver": "host",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": []
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "171e3babf161bbda519e4081d11618bbed8e561e276bafebf693687444054fc1": {
                "Name": "nginx-web",
                "EndpointID": "62895ff820b05450de65c356c0122a5fa30626c5f1cd0639638642c6372e5bf0",
                "MacAddress": "",
                "IPv4Address": "",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]

```

#### 性能分析

host网络相比于bridge性能更高，因为host直接使用宿主机的ip，不需要进行docker内部bridge复杂的网络，没有更多的消耗。



### none网络

就是不需要连接任何网络，只是在后台运行一个容器，告诉docker你需要运行容器，网络方面你不需要管理，我自己需要自己配置

```sh
$ docker run -d --name nginx-web --network none nginx:latest
```



### 实战

#### Python Flask + Redis 练习

![flask-redis](https://dockertips.readthedocs.io/en/latest/single-host-network/_static/flask-redis.png)

#### 程序准备

准备一个Python文件，名字为 `app.py` 内容如下：

```python
from flask import Flask
from redis import Redis
import os
import socket

app = Flask(__name__)
redis = Redis(host=os.environ.get('REDIS_HOST', '127.0.0.1'), port=6379)


@app.route('/')
def hello():
    redis.incr('hits')
    return f"Hello Container World! I have been seen {redis.get('hits').decode('utf-8')} times and my hostname is {socket.gethostname()}.\n"
```



准备一个Dockerfile

```dockerfile
FROM python:3.9.5-slim

RUN pip install flask redis && \
    groupadd -r flask && useradd -r -g flask flask && \
    mkdir /src && \
    chown -R flask:flask /src

USER flask

COPY app.py /src/app.py

WORKDIR /src

ENV FLASK_APP=app.py REDIS_HOST=redis

EXPOSE 5000

CMD ["flask", "run", "-h", "0.0.0.0"]
```



#### 镜像准备

构建flask镜像，准备一个redis镜像。

```sh
$ docker image pull redis
$ docker image build -t flask-demo .
$ docker image ls
REPOSITORY   TAG          IMAGE ID       CREATED              SIZE
flask-demo   latest       4778411a24c5   About a minute ago   126MB
python       3.9.5-slim   c71955050276   8 days ago           115MB
redis        latest       08502081bff6   2 weeks ago          105MB
```



#### 创建一个docker bridge

```sh
$ docker network create -d bridge demo-network
8005f4348c44ffe3cdcbbda165beea2b0cb520179d3745b24e8f9e05a3e6456d
$ docker network ls
NETWORK ID     NAME           DRIVER    SCOPE
2a464c0b8ec7   bridge         bridge    local
8005f4348c44   demo-network   bridge    local
80b63f711a37   host           host      local
fae746a75be1   none           null      local
$
```



#### 创建redis container

创建一个叫 `redis-server` 的container，连到 demo-network上

```sh
$ docker container run -d --name redis-server --network demo-network redis
002800c265020310231d689e6fd35bc084a0fa015e8b0a3174aa2c5e29824c0e
$ docker container ls
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS      NAMES
002800c26502   redis     "docker-entrypoint.s…"   4 seconds ago   Up 3 seconds   6379/tcp   redis-server
$
```



#### 创建flask container

```sh
$ docker container run -d --network demo-network --name flask-demo --env REDIS_HOST=redis-server -p 5000:5000 flask-demo
```



打开浏览器访问 [http://127.0.0.1:5000](http://127.0.0.1:5000/)

应该能看到类似下面的内容，每次刷新页面，计数加1

Hello Container World! I have been seen 36 times and my hostname is 925ecb8d111a.



#### 总结

如果把上面的步骤合并到一起，成为一个部署脚本

```sh
# prepare image
docker image pull redis
docker image build -t flask-demo .

# create network
docker network create -d bridge demo-network

# create container
docker container run -d --name redis-server --network demo-network redis
docker container run -d --network demo-network --name flask-demo --env REDIS_HOST=redis-server -p 5000:5000 flask-demo
```





## docker compose

### docker compose安装

Windows和Mac在默认安装了docker desktop以后，docker-compose随之自动安装

```
PS C:\Users\Peng Xiao\docker.tips> docker-compose --version
docker-compose version 1.29.2, build 5becea4c
```



Linux用户需要自行安装

最新版本号可以在这里查询 https://github.com/docker/compose/releases

```
$ sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
$ sudo chmod +x /usr/local/bin/docker-compose
$ docker-compose --version
docker-compose version 1.29.2, build 5becea4c
```



熟悉python的朋友，可以使用pip去安装docker-Compose

```
$ pip install docker-compose
```



### docker-compose结构和文件

#### 基本语法结构

```yaml
version: "3.8"

services: # 容器
  servicename: # 服务名字，这个名字也是内部 bridge网络可以使用的 DNS name
    image: # 镜像的名字
    command: # 可选，如果设置，则会覆盖默认镜像里的 CMD命令
    environment: # 可选，相当于 docker run里的 --env
    volumes: # 可选，相当于docker run里的 -v
    networks: # 可选，相当于 docker run里的 --network
    ports: # 可选，相当于 docker run里的 -p
  servicename2:

volumes: # 可选，相当于 docker volume create

networks: # 可选，相当于 docker network create
```



以 Python Flask + Redis练习：为例子，改造成一个docker-compose文件

```sh
docker image pull redis
docker image build -t flask-demo .

# create network
docker network create -d bridge demo-network

# create container
docker container run -d --name redis-server --network demo-network redis
docker container run -d --network demo-network --name flask-demo --env REDIS_HOST=redis-server -p 5000:5000 flask-demo
```



docker-compose.yml 文件如下

```yaml
version: "3.8"

services:
  flask-demo:
    image: flask-demo:latest
    environment:
      - REDIS_HOST=redis-server
    networks:
      - demo-network
    ports:
      - 8080:5000

  redis-server:
    image: redis:latest
    networks:
     - demo-network

networks:
  demo-network:
```

#### docker-compose 语法版本

向后兼容

https://docs.docker.com/compose/compose-file/



### docker compose命令

#### 启动

> docker-compose up

使用启动命令需要先进入对应的目录中，前提条件一定是本地镜像中对应的镜像了

我们需要进入到存储docker-compose.yml文件的目录下

```sh
docker-compose up   #前台启动，会打印日志，退出容器也就退出
```

```sh
docker-compose up -d   #后台启动
```



#### 拉取和构建

例如本地如果没有本地镜像，他会去DockerHub进行拉取

```sh
$ docker-compose pull #拉取
```



docker-compose为我们提供了一个构建需要的镜像的命令
```sh
$ docker-compose build #构建
```

前提是我们需要先编写好docker-compse.yml文件

```yaml
version: "3.8"

services:
  flask-demo:
    build: ./docker
    image: flask-demo:latest
    environment:
      - REDIS_HOST=redis-server
    networks:
      - demo-network
    ports:
      - 8081:5000

  redis-server:
    image: redis:latest
    networks:
     - demo-network

networks:
  demo-network:
```

目录结构：

```test
├── docker
│   ├── Dockerfile
│   └── app.py
├── docker-compose.yml
```

我们添加了，compose会默认去找./docker命令下的文件，并且需要注意我们没有指定dockerfile文件名称则默认是```Dockerfile```，但是有时候是需要我们指定的
```
build: ./docker
```



假设结构是这样的：

```
├── docker
│   ├── Dockerfile.dev
│   └── app.py
├── docker-compose.yml
```

我们需要这要编写docker-compose.yml文件：

```yaml
version: "3.8"

services:
  flask-demo:
    build:
      context: ./docker
      dockerfile: Dockerfile.dev
    image: flask-demo:latest
    environment:
      - REDIS_HOST=redis-server
    networks:
      - demo-network
    ports:
      - 8081:5000

  redis-server:
    image: redis:latest
    networks:
     - demo-network

networks:
 
```



### docker-compose服务更新

#### 更新

##### 更新源文件

当我们在使用docker-compos进行构建镜像的时候，当源文件更改后，要如何更新镜像

比如我们上面的```flask-demo```，我们直接修改app.py文件内容，然后发现更新和构建，compose提供了命令：

```
docker-compose up -d --build
```

该命令会检查镜像是否需要重新build，即使在容器在运行时也是可以进行构建的然后运行。



##### 更新compose文件

当我们更新compose文件后，例如：

```yaml
version: "3.8"

services:
  flask-demo:
    build:
      context: ./docker
      dockerfile: Dockerfile.dev
    image: flask-demo:latest
    environment:
      - REDIS_HOST=redis-server
    networks:
      - demo-network
    ports:
      - 8081:5000

  redis-server:
    image: redis:latest
    networks:
     - demo-network

  busybox:
    image: busybox:latest
    command: sh -c "while true; do sleep 3600; done"
    networks:
      - demo-network

networks:
  demo-network:
```



我们新增了一个busybox的镜像，并执行命令，我们只需要使用：

```
docker-compose up -d
```

然后compose就会去拉取对镜像还在构建镜像



##### 移除镜像

我们在上面实例的基础上移除```busybox```镜像后，运行命令：

```sh
$ docker-compose up -d  #我们会看到提示
WARN[0000] Found orphan containers ([hellopy-busybox-1]) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up.
[+] Running 2/0
 ✔ Container hellopy-flask-demo-1    Running                                                                                                                                                           0.0s
 ✔ Container hellopy-redis-server-1  Running
```

所以我们需要使用命令：

```sh
$ docker-compose up -d --remove-orphans  
[+] Running 3/1
 ✔ Container hellopy-busybox-1       Removed                                                                                                                                                          
 ✔ Container hellopy-flask-demo-1    Running                                                                                                                                                           
 ✔ Container hellopy-redis-server-1  Running                                                                                                                                                           
```



当然我们可以使用命令重启docker-compose完成内容的更新

```
docker-compose restart
```



### 未完待续







