package route

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"sing/app/model/listen"
	"sing/app/model/read"
	"sing/app/model/speak"
	"sing/app/model/write"
	"sing/app/route/middleware/jaeger"
	"sing/app/util"
	"sing/app/util/grpc_client"
)

func SetupRouter(engine *gin.Engine) {

	engine.Use(jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/sing", func(c *gin.Context) {

		fmt.Println("####################### sing #####################")
		utilGin := util.Gin{Ctx: c}
		utilGin.Response(1, "sing", nil)
	})

	// 测试链路追踪
	engine.GET("/jaeger_test", func(c *gin.Context) {

		fmt.Println("=======jaeger_testjaeger_testjaeger_testjaeger_test=======")
		// 调用 gRPC 服务
		conn := grpc_client.CreateServiceListenConn(c)
		grpcListenClient := listen.NewListenClient(conn)
		resListen, _ := grpcListenClient.ListenData(context.Background(), &listen.Request{Name: "listen"})

		fmt.Println("=======listenlistenlistenlistenlisten=======")

		// 调用 gRPC 服务
		conn = grpc_client.CreateServiceSpeakConn(c)
		grpcSpeakClient := speak.NewSpeakClient(conn)
		resSpeak, _ := grpcSpeakClient.SpeakData(context.Background(), &speak.Request{Name: "speak"})

		fmt.Println("=======speakspeakspeakspeakspeak=======")

		// 调用 gRPC 服务
		conn = grpc_client.CreateServiceReadConn(c)
		grpcReadClient := read.NewReadClient(conn)
		resRead, _ := grpcReadClient.ReadData(context.Background(), &read.Request{Name: "read"})

		// 调用 gRPC 服务
		conn = grpc_client.CreateServiceWriteConn(c)
		grpcWriteClient := write.NewWriteClient(conn)
		resWrite, _ := grpcWriteClient.WriteData(context.Background(), &write.Request{Name: "write"})

		defer conn.Close()

		// 调用 HTTP 服务
		resHttpGet := ""
		_, err := util.HttpGet("http://127.0.0.1:9905/sing", c)
		if err == nil {
			resHttpGet = "[HttpGetOk]"
		}

		// 业务处理...

		msg := resListen.Message + "-" +
			resSpeak.Message + "-" +
			resRead.Message + "-" +
			resWrite.Message + "-" +
			resHttpGet + "=================="

		utilGin := util.Gin{Ctx: c}
		utilGin.Response(1, msg, nil)
	})
}
