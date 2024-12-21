package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gofileshare.server/config"
	"gofileshare.server/internal/api"
)

func main() {
	// Initialize the Gin router 初始化Gin路由
	r := gin.Default()

	// Add shared directories 添加共享目录
	api.AddSharedDir(config.Cfg.SharedDirectory)

	// Set up routes
	r.GET("/", func(c *gin.Context) {
		path := c.Query("path")
		files, err := api.ListFiles(path)
		if err != nil {
			c.HTML(http.StatusNotFound, "error.tmpl", gin.H{
				"Error": "目录不存在",
			})
			return
		}

		breadcrumbs := api.GetBreadcrumbs(path)
		c.HTML(200, "index.tmpl", gin.H{
			"Title":       "共享文件",
			"Files":       files,
			"CurrentPath": path,
			"Breadcrumbs": breadcrumbs,
		})
	})

	r.GET("/files/*filepath", func(c *gin.Context) {
		api.HandleFileRequestGin(c)
	})

	// Load HTML templates 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	log.Printf("Starting server on :%d...", config.Cfg.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", config.Cfg.Server.Port)); err != nil {
		log.Printf("Server failed to start: %v", err)
	}
}
