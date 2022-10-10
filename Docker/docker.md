# 	Docker

## 0.目标

- 理解docker仓库，镜像，容器概念
- 了解相关镜像、容器操作命令
- 理解Dockfile文件
- 理解docker-compose



## 1.目的

> 开发环境的统一，开发环境，线上环境版本差别



## 2.Docker思想

> Docker的思想来自于集装箱，集装箱解决了什么问题？在一艘大船上，可以把货物规整的摆放起来。并且各种各样的货物被集装箱标准化了，集装箱（镜像）和集装箱之间不会互相影响。那么我就不需要专门运送水果的船和专门运送化学品的船了。只要这些货物在集装箱里封装的好好的，那我就可以用一艘大船把他们都运走。

![docker_logo_image](..\$Image\Docker\docker_logo.png)

蓝色的大海：宿主机

鲸鱼：docker

集装箱：镜像



## 3.安装

Mac，window

[下载](https://www.docker.com/products/docker-desktop/)

[安装](http://c.biancheng.net/view/3121.html)

加速器

```shell
 Registry Mirrors:
  https://hub-mirror.c.163.com/
  https://registry.aliyuncs.com/
  https://registry.docker-cn.com/
  https://docker.mirrors.ustc.edu.cn/
```



## 4.docker架构

![docker架构](https://img-blog.csdnimg.cn/img_convert/27d622f93fff98388c0dd262972aeefa.png)

通过上图可以得知，Docker 在运行时分为 **Docker引擎**（服务端守护进程） 和 **客户端工具**，我们日常使用各种 docker 命令，其实就是在使用 客户端工具 与 Docker 引擎 进行交互。

> 客户端工具：提供了一套 RESTful API 接口的命令行工具，客户端工具可以和Docker 引擎运行在同一台宿主机上，也可以访问远端的Docker 引擎。
>
> Docker 引擎：一个物理或者虚拟的机器用于执行 Docker 守护进程和容器。
>
> 宿主机：可以是物理机，也可以是虚拟机，上面跑着一个dockers 服务。



## 5.仓库

> 分为公有和私有。公有的就是Docker Hub，私有的可以自己搭建仓库，可以保存自己的镜像。



## 6.镜像

> 镜像由多个层组成，每层叠加之后，从外部看来就如一个独立的对象。镜像内部是一个精简的操作系统（OS），同时还包含应用运行所必须的文件和依赖包。
>
> Docker 镜像可以看作是一个特殊的文件系统，除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（如匿名卷、环境变量、用户等）。

![/resources/articles/docker/12180844322018196a29c55c8de4a2.png](https://img-blog.csdnimg.cn/img_convert/513626342c94b03f34269750fec04c4a.png)

### 6.1.查看仓库镜像

> docker search 镜像名

```shell
$ docker search mysql
NAME                             DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                            MySQL is a widely used, open-source relation…   12466     [OK]
mariadb                          MariaDB Server is a high performing open sou…   4802      [OK]
mysql/mysql-server               Optimized MySQL Server Docker images. Create…   922                  [OK]
percona                          Percona Server is a fork of the MySQL relati…   575       [OK]
phpmyadmin                       phpMyAdmin - A web interface for MySQL and M…   513       [OK]
mysql/mysql-cluster              Experimental MySQL Cluster Docker images. Cr…   93
centos/mysql-57-centos7          MySQL 5.7 SQL database server                   93
bitnami/mysql                    Bitnami MySQL Docker Image                      69                   [OK]
ubuntu/mysql                     MySQL open source fast, stable, multi-thread…   31
circleci/mysql                   MySQL is a widely used, open-source relation…   25
mysql/mysql-router               MySQL Router provides transparent routing be…   23
google/mysql                     MySQL server for Google Compute Engine          21                   [OK]
vmware/harbor-db                 Mysql container for Harbor                      10
mysqlboy/docker-mydumper         docker-mydumper containerizes MySQL logical …   3
bitnami/mysqld-exporter                                                          3
mysqlboy/mydumper                mydumper for mysql logcial backups              3
ibmcom/mysql-s390x               Docker image for mysql-s390x                    2
newrelic/mysql-plugin            New Relic Plugin for monitoring MySQL databa…   1                    [OK]
cimg/mysql                                                                       0
ibmcom/tidb-ppc64le              TiDB is a distributed NewSQL database compat…   0
mysql/mysql-operator             MySQL Operator for Kubernetes                   0
newrelic/k8s-nri-mysql           New Relic Infrastructure MySQL Integration (…   0
mysqlboy/elasticsearch                                                           0
mysqleatmydata/mysql-eatmydata                                                   0
mirantis/mysql                                                                   0
```

也可以去docker hub搜索

![docker_hub_image](..\$Image\Docker\docker_hub_image.png)

### 6.2.查看镜像的Tag

Tag 就是标签，类似版本号，需要去docker hub查看

https://hub.docker.com/_/mysql?tab=tags

![docker_hub_tag](..\$Image\Docker\docker_hub_tag.png)

### 6.3.拉取镜像

> docker pull [OPTIONS] NAME[:TAG|@DIGEST]

- [OPTIONS]

- NAME 镜像名称，比如mysql

- [:TAG] 默认是 latest，最新的

```shell
$ docker pull mysql
Using default tag: latest
latest: Pulling from library/mysql
```

就是去Docker Hub拉取library/mysql，tag latest的镜像

取消拉取，`CTRL-c`

### 6.4.查看

#### 6.4.1.查看本地镜像列表

> docker image ls
>
> docker images

#### 6.4.2.查看具体镜像信息

> docker inspect 镜像ID

### 6.5.删除

> docker rmi 镜像ID

```shell
$ docker image rmi gin-blog-docker
Untagged: gin-blog-docker:latest
Deleted: sha256:3c3048326c3799886442dddea805e596d195dc71fa29ff51ebca5a22c4d08c6
```

### 6.6.修改本地镜像名和Tag

> docker tag 镜像ID 新的镜像名:新的镜像tag

```shell
$ docker image ls
REPOSITORY                  TAG            IMAGE ID       CREATED        SIZE
golang                      latest         65375c930b21   5 days ago     964MB
nginx                       latest         fa5269854a5e   5 days ago     142MB
$ docker tag fa5269854a5e new_image:new_tag
$ docker image ls
REPOSITORY                  TAG            IMAGE ID       CREATED        SIZE
golang                      latest         65375c930b21   5 days ago     964MB
new_image                   new_tag        fa5269854a5e   5 days ago     142MB
```

### 6.6.空悬镜像

```shell
# 搜索空悬镜像
$ docker images -f "dangling=true"
# 查找空悬镜像ID
$ docker images -f "dangling=true" -q
# 根据空悬镜像ID删除空悬镜像
$ docker rmi $(docker images -f "dangling=true" -q)
```



## 7.容器

容器是镜像运行时的实体

镜像->类，容器->对象

[镜像和容器的区别](https://mp.weixin.qq.com/s?__biz=MzI0MDQ4MTM5NQ==&mid=2247512892&idx=1&sn=3a946cb329c38817da3ff08049c0fe01&chksm=e918d620de6f5f362fdfc7c02ebc3f5a07046dd2695859a3d2086a33451516fd848ba2ca6fbb&token=2144552082&lang=zh_CN#rd)

### 7.1.状态机

![容器状态机](https://img-blog.csdnimg.cn/9c573863e70f4032833a81b646dcd94d.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5oq56aaZ6bK45LmL5rW3,size_20,color_FFFFFF,t_70,g_se,x_16)

### 7.2.创建

> docker create 镜像名:tag

```shell
$ docker create golang:latest
29aacb9e4c137e2d38143b49ceeeef802870507b5099393ebcf9702b16499613
#没有该镜像会自动去拉取
$ docker create mysql:5.6
cd5a36bc58215f3ef0a0bc2953ca41c9f793d4c75bfb45ba8030030fa9d508da
```

### 7.3.启动

> docker start 容器名/容器ID

```shell
$ docker start golang
golang
$ docker start cd5a36bc5821
cd5a36bc5821
```

### 7.4.创建并启动

> docker run 镜像名:tag

```shell
$ docker run -p 80:80 -d --name nginx -v /docker/nginx/default.conf:/etc/nginx/conf.d/default.conf -v /docker/www:/docker/www  --privileged=true nginx
$ docker run -p 8088:80 -d --rm --name test_nginx --privileged=true nginx
$ docker run -p 8088:80 -d --name test_nginx --privileged=true nginx

-t 参数让Docker分配一个伪终端并绑定到容器的标准输入上
-i 参数则让容器的标准输入保持打开。
-c 参数用于给运行的容器分配cpu的shares值
-m 参数用于限制为容器的内存信息，以 B、K、M、G 为单位
-v 参数用于挂载一个volume，可以用多个-v参数同时挂载多个volume
-p 参数用于将容器的端口暴露给宿主机端口 格式：host_port:container_port 或者 host_ip:host_port:container_port
--name 容器名称
--net  容器使用的网络
--rm   容器停止，删除容器
```

- -p 36:3600 前面是主机，后面是容器；将容器的3600端口映射到主机的36端口
- -d 后台运行
- --name 容器命名
- --privileged=true 容器内root拥有真正的权限

- -v /docker/nginx/default.conf:/etc/nginx/conf.d/default.conf 将主机的虚拟主机配置文件 **挂载** 到容器的/etc/nginx/conf.d/default.conf
- -v /docker/www:/docker/www 将主机中当前目录下的www挂载到容器的www目录
- --rm 容器退出时，删除容器

![docker_nginx_port](..\$Image\Docker\docker_nginx_port.png)

### 7.5.查看正在运行的容器列表

> docker container ls

### 7.6.查询所有状态容器列表

> docker ps -a

```shell
$ docker ps -a
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS                       PORTS       NAMES
c6d9c51951c4   nginx     "/docker-entrypoint.…"   About a minute ago   Exited (137) 36 seconds ago              test_nginx
```

### 7.7.停止

> docker stop 容器ID/容器名字

```shell
$ docker stop 4dd797365740
4dd797365740
$ docker stop test_nginx
test_nginx
```

### 7.8.进入容器

容器必须是 Up 状态

> docker exec -it 容器ID bash

```shell
# 查看docker的Liunx发行版本
$ docker exec -it bcca857031b1 /bin/sh
root@bcca857031b1:/# cat /etc/issue
Debian GNU/Linux 11 \n \l
root@bcca857031b1:/# apt update
root@bcca857031b1:/# apt install vim

root@c6d9c51951c4:/# echo '<h1>这是docker构建的nginx，端口为8088</h1>' > /usr/share/nginx/html/index.html
```

### 7.9.删除

普通删除

> docker container rm 容器ID

```shell
$ docker container rm cfb682e8c523
Error response from daemon: You cannot remove a running container cfb682e8c523f5ae82c70e9cfa31525bfcec8931c02550883af4043b95ce68c0. Stop the container before attempting removal or force remove
```

强制删除运行中的容器

> docker container rm -f 容器ID

```shell
$ docker container rm -f cfb682e8c523
cfb682e8c523
```

### 7.10.查看日志

> docker log 镜像ID

```shell
$ docker logs ac0b828420a0
2022-04-26 03:51:43+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.6.51-1debian9 started.
2022-04-26 03:51:43+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
2022-04-26 03:51:43+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.6.51-1debian9 started.
2022-04-26 03:51:43+00:00 [ERROR] [Entrypoint]: Database is uninitialized and password option is not specified
    You need to specify one of the following:
    - MYSQL_ROOT_PASSWORD
    - MYSQL_ALLOW_EMPTY_PASSWORD
    - MYSQL_RANDOM_ROOT_PASSWORD
```

需要提供MySQL的用户密码

### 7.11.commit

镜像构成，在 latest 基础上再构建一层自定义的镜像

```bash
$ docker diff test_nginx
$ docker commit --author "TimidHaunter <timidhaunter@gmail.com>" --message "修改了默认网页" test_nginx nginx:v2
$ docker image ls nginx
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
nginx        v2        509c837d1a33   2 minutes ago   142MB
nginx        latest    fa5269854a5e   5 days ago      142MB

$ docker history nginx:v2
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
509c837d1a33   3 minutes ago   nginx -g daemon off;                            1.12kB    修改了默认网页
fa5269854a5e   5 days ago      /bin/sh -c #(nop)  CMD ["nginx" "-g" "daemon…   0B
<missing>      5 days ago      /bin/sh -c #(nop)  STOPSIGNAL SIGQUIT           0B

# 多了一个nginx:v2镜像
$ docker image ls
REPOSITORY                  TAG            IMAGE ID       CREATED         SIZE
nginx                       v2             509c837d1a33   4 minutes ago   142MB
golang                      latest         65375c930b21   5 days ago      964MB
nginx                       latest         fa5269854a5e   5 days ago      142MB

# 启动NGINX:v2
$ docker run -p 8089:80 -d --name test_nginx_v2 --privileged=true nginx:v2

# 之前的镜像，容器删掉，重新用镜像开启，之前在容器的操作都没有了
$ docker run -p 8088:80 -d --name test_nginx --privileged=true nginx
```



### 7.12.命令

![docker_commands_diagram](..\$Image\Docker\docker_commands_diagram.jpg)



## 8.Dockerfile

自定义镜像

### 8.1.是什么？

docker build 构建自己的docker镜像所要用到的定义文件。里面都是命令行。

> docker build [OPTIONS] PATH | URL | -
>
> docker build [选项] <上下文路径 | URL | ->

### 8.2.有什么用？

可以自定义镜像

### 8.3.怎么用?

```bash
# 基准镜像
FROM centos:7

# 作者信息（可以删除）
MAINTAINER "timidhaunter"

# 工作目录
WORKDIR /usr/local/src/

# 定义环境变量
ENV NG_VERSION nginx-1.21.0

# 安装epel仓库
RUN yum -y install epel-release

# 安装wget
RUN yum -y install wget

# 下载nginx文件并解压
RUN wget http://nginx.org/download/$NG_VERSION.tar.gz && tar xzvf $NG_VERSION.tar.gz

# 安装编译依赖包
RUN yum install -y gcc gcc-c++ glibc make autoconf openssl openssl-devel && yum install -y pcre-devel libxslt-devel gd-devel GeoIP GeoIP-devel GeoIP-data

# 清理仓库
RUN yum clean all

# 创建nginx用户
RUN useradd -M -s /sbin/nologin nginx

# 切换工作目录
WORKDIR /usr/local/src/$NG_VERSION

# 编译安装nginx
RUN ./configure --user=nginx --group=nginx --prefix=/usr/local/nginx --with-file-aio --with-http_ssl_module --with-http_realip_module --with-http_addition_module --with-http_xslt_module --with-http_image_filter_module --with-http_geoip_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --with-http_mp4_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_auth_request_module --with-http_random_index_module --with-http_secure_link_module --with-http_degradation_module --with-http_stub_status_module && make && make install

# 复制测试页面到容器中
ADD index.html /usr/local/nginx/html

# 设置容器中要挂在到宿主机的目录
VOLUME /usr/local/nginx/html

# 设置sbin环境变量
ENV PATH /usr/local/nginx/sbin:$PATH

# 暴露镜像的80端口
EXPOSE 80/tcp
ENTRYPOINT ["nginx"]
CMD ["-g","daemon off;"]
```



## 9.docker-compose

### 9.1.是什么

> Docker-Compose运行目录下的所有文件（docker-compose.yml，extends文件或环境变量文件等）组成一个工程，若无特殊指定工程名即为当前目录名。一个工程当中可包含多个服务，每个服务中定义了容器运行的镜像，参数，依赖。一个服务当中可包括多个容器实例，Docker-Compose并没有解决负载均衡的问题，因此需要借助其它工具实现服务发现及负载均衡。
>
> Docker-Compose的工程配置文件默认为docker-compose.yml，可通过环境变量COMPOSE_FILE或-f参数自定义配置文件，其定义了多个有依赖关系的服务及每个服务运行的容器。

### 9.2.干什么用

> 使用一个Dockerfile模板文件，可以让用户很方便的定义一个单独的应用容器。在工作中，经常会碰到需要多个容器相互配合来完成某项任务的情况。例如要实现一个Web项目，除了Web服务容器本身，往往还需要再加上后端的数据库服务容器，甚至还包括负载均衡容器等。
>
> Compose允许用户通过一个单独的docker-compose.yml模板文件（YAML 格式）来定义一组相关联的应用容器为一个项目（project）。

### 9.3.安装

桌面版自带

```bash
$ docker-compose --version
docker-compose version 1.29.2, build 5becea4c
```

### 9.4.使用

docker-compose.yml文件

version、services、networks三大部分，最关键的是services和networks两个部分。

```bash
version: '2'
services:
  #配置nginx服务
  nginx:
    #设置主机名为nginx
    hostname: nginx
    #使用dockerfile创建镜像。Dockerfile文件在当前目录的nginx目录下，文件名为Dockerfile
    build:
      context: ./nginx
      dockerfile: Dockerfile
    #容器名为nginx
    container_name: nginx
    #暴露端口80和443
    ports:
      - 80:80
      - 443:443
    #加入到lnmp网络中，使用ip172.18.0.0.10
    networks:
      lnmp:
        ipv4_address: 172.18.0.10
    #将当前目录的wwwroot目录挂载到容器的/usr/local/nginx/html
    volumes:
      - ./wwwroot/:/usr/local/nginx/html
  #配置服务mysql
  mysql:
    hostname: mysql
    build:
      context: ./mysql
      dockerfile: Dockerfile
    container_name: mysql
    ports:
      - 3306:3306
    networks:
      lnmp:
        ipv4_address: 172.18.0.20
    #设置/usr/local/mysql为数据卷
    volumes:
      - /usr/local/mysql
  #配置服务php  
  php:
    hostname: php
    build:
      context: ./php
      dockerfile: Dockerfile
    container_name: php
    ports:
      - 9000:9000
    networks:
      lnmp:
        ipv4_address: 172.18.0.30
    #从nginx容器和mysql容器获取数据卷  
    volumes_from:
      - nginx
      - mysql
    #php容器需要在nginx和mysql之后启动  
    depends_on:
      - nginx
      - mysql
    #php和容器nginx，容器mysql连接   
    links:
      - nginx
      - mysql
#配置网络模式和网络名    
networks:
  #设置网络名lnmp
  lnmp:
    #网络模式为bridge桥接莫斯
    driver: bridge
    ipam:
      config:
        #使用的网段为172.18.0.0/16
        - subnet: 172.18.0.0/16
```

- image：指定服务的镜像名称
- build：可以基于镜像，还可以基于一份Dockerfile

  - context：可以是Dockerfile的文件路径，也可以是到链接到git仓库的url
  - dockerfile：构建路径
  - args：添加构建参数，这是只能在构建过程中访问的环境变量
- ports：暴露端口（HOST:CONTAINER）
- container_name：容器名
- environment：
- depends_on：启动顺序

  ```bash
  kibana:
      depends_on:
        - elasticsearch
  ```
- volumes：挂载一个目录或者一个已存在的数据卷容器
- expose：暴露端口，但不映射到宿主机，用于内部访问
- restart：重启策略
- command：可以覆盖容器启动后默认执行的命令

启动所有服务

> docker-compose up

查看项目所有容器

> docker-compose ps



## 10.dnmp

就是用docker-compose搭建的一个PHP开发环境项目

gitee：https://gitee.com/yeszao/dnmp/

GitHub：https://github.com/yeszao/dnmp



## 11.学习资料

[1.docker run 命令详解（新手入门必备）](https://blog.csdn.net/anqixiang/article/details/106545603)
[2.Docker学习笔记：Docker 基础用法和命令帮助](https://www.docker.org.cn/dockerppt/106.html)
[3.dockerfile详解](https://blog.csdn.net/zisefeizhu/article/details/83472190)
[4.Docker学习笔记：Dockerfile](https://www.docker.org.cn/dockerppt/114.html)
[5.使用Dockerfile制作nginx镜像](https://blog.csdn.net/weixin_44455388/article/details/117447640)
[6.Docker快速入门——Docker-Compose](https://www.cnblogs.com/phpk/p/11205467.html)
[7.使用Docker Compose 搭建LNMP](https://blog.csdn.net/qq_44135433/article/details/121351799)
[8.docker的架构及工作原理](https://blog.csdn.net/weixin_44621343/article/details/115895268)
8.书籍
链接：[https://pan.baidu.com/s/1ki9d6xq3cunw7QWkFvb7Pw](https://pan.baidu.com/s/1ki9d6xq3cunw7QWkFvb7Pw)
提取码：x0a4