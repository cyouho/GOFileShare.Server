package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gofileshare.server/config"
	"gofileshare.server/internal/api"
	"gofileshare.server/internal/routes"
)

func main() {
	// Initialize the Gin router 初始化Gin路由
	r := gin.Default()

	// Add shared directories 添加共享目录
	api.AddSharedDir(config.Cfg.SharedDirectory)

	// Load HTML templates 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 设置路由和模板
	routes.SetupRoutes(r)

	// 启动服务器
	log.Printf("Starting server on :%d...", config.Cfg.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", config.Cfg.Server.Port)); err != nil {
		log.Printf("Server failed to start: %v", err)
	}
}
