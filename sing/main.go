package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"sing/app/config"
	"sing/app/route"
	"sing/app/util/jaeger_service"
)

func main() {
	gin.SetMode(config.AppMode)

	_, closer, err := jaeger_service.NewJaegerTracer(config.AppName, config.JaegerHostPort)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	defer closer.Close()

	engine := gin.New()

	// 设置路由
	route.SetupRouter(engine)

	log.Println("server listen port" + config.AppPort)

	// 启动服务
	if err := engine.Run(config.AppPort); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
