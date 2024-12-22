package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gofileshare.server/internal/api"
)

// SetupRoutes 配置所有路由
func SetupRoutes(r *gin.Engine) {
	// 设置页面路由
	setupPageRoutes(r)
	// 设置文件处理路由
	setupFileRoutes(r)
	// Load HTML templates 加载HTML模板
	r.LoadHTMLGlob("templates/*")
}

// setupPageRoutes 设置页面路由
func setupPageRoutes(r *gin.Engine) {
	// 主页路由
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
}

// setupFileRoutes 设置文件处理路由
func setupFileRoutes(r *gin.Engine) {
	// 文件处理路由
	r.GET("/files/*filepath", func(c *gin.Context) {
		api.HandleFileRequestGin(c)
	})
}
