# 本文参考了：
[go-jaeger-demo](https://github.com/xinliangnote/go-jaeger-demo)

[go-gin-api](https://github.com/xinliangnote/go-gin-api)

## 项目介绍

这是一个 Jaeger 链路追踪的 Demo，里面包括 5 个 Service 端，如图所示：

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_1.png)

5 个 Service 端 Demo 分别是：

#### 听（listen）

- 端口：9901
- 通讯：gRPC

#### 说（speak）

- 端口：9902
- 通讯：gRPC

#### 读（read）

- 端口：9903
- 通讯：gRPC

#### 写（write）

- 端口：9904
- 通讯：gRPC

#### 唱（sing）

- 端口：9905
- 通讯：HTTP

其中服务之间又相互调用：

- Speak 服务，又调用了 Listen 服务 和 Sing 服务。
- Read 服务，又调用了 Listen 服务 和 Sing 服务。
- Write 服务，又调用了 Listen 服务 和 Sing 服务。

咱们要实现就是 API 调用 5 个服务的链路，以及服务与服务之间相互调用的链路。

## 运行

#### 1、部署 jaeger 服务

选择docker安装：

```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one
```

目测启动后，访问地址：http://localhost:16686/

看到下图，表示启动成功。

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_4.png)


#### 2、部署 mysql 服务

选择docker安装：

```
docker run -itd -p 3306:3306 --name mysqltest --restart=always -e MYSQL_ROOT_PASSWORD=123456 mysql
```

然后执行命令：docker exec -it mysqltest /bin/bash

能正确进入，代表安装成功！！

#### 3、启动 Service 服务

```
// 启动 Listen 服务
cd listen && go run main.go

// 启动 Speak 服务
cd speak && go run main.go

// 启动 Read 服务
cd read && go run main.go

// 启动 Write 服务
cd write && go run main.go

// 启动 Sing 服务
cd sing && go run main.go
```

#### 4、访问路由

访问 API 项目：http://127.0.0.1:9905/jaeger_test

## 效果

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_2.png)

![](https://github.com/xinliangnote/Go/blob/master/03-go-gin-api%20%5B文档%5D/images/jaeger_demo_3.png)
