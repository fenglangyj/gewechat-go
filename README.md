# go语言实现的Gewechat api接口

## 介绍

本项目基于 Gewechat，请先确认 Gewechat 已经正常启动，否则无法使用本模块。

## 🚀 快速入门

### 安装Docker（如果你已有Docker请忽略此步骤）

> Centos Docker安装，已安装Docker可跳过

1. 安装gcc相关

    ```
    yum -y install gcc
    yum -y install gcc-c++
    ```

2. 配置镜像

    ```
    yum install -y yum-utils
    yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
    yum makecache fast
    ```

3. 安装docker

    ```
    yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    ```

4. 启动docker

    ```
    systemctl start docker
    //将docker设置成开机自启动
    systemctl enable docker.service
    ```

### 启动 Gewechat 服务

1. 拉取镜像

    ```
     docker pull registry.cn-hangzhou.aliyuncs.com/gewe/gewe:latest
     
     docker tag registry.cn-hangzhou.aliyuncs.com/gewe/gewe gewe
    ```

2. 运行镜像容器（请注意端口2531和2532不要被占用）

    ```
    
    mkdir -p /root/temp
    
    docker run -itd -v /root/temp:/root/temp -p 2531:2531 -p 2532:2532 --privileged=true --name=gewe gewe /usr/sbin/init
    
    ```

3. 将容器设置成开机运行

    ```
    docker update --restart=always gewe
    
    ```

### Gewechat 服务调用接口如下：

1. API服务调用地址 `http://{服务ip}:2531/v2/api/{接口名}`

2. 文件下载地址 `http://{服务ip}:2532/download/{接口返回的文件路径}`

## 注意事项：


## 交流群：

<p align="center">
  <img src="docs/mine.png" width="300px" height="300px" alt=" Logo">
</p>
<p align="center">
 添加好友备注加群交流
</p>

## 版本更新

### 1.0.0

* 正式1.0版本发布