package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gofileshare.server/internal/api"
)

// SetupRoutes 配置所有路由
func SetupRoutes(r *gin.Engine) {
	// 添加根路径重定向
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web/")
	})

	// 设置页面路由
	setupWebRoutes(r)
	// 设置文件处理路由
	setupFileRoutes(r)
}

// setupWebRoutes 设置页面路由
func setupWebRoutes(r *gin.Engine) {
	web := r.Group("/web")
	{
		web.GET("/", func(c *gin.Context) {
			path := c.Query("path")
			files, err := api.ListFiles(path)
			if err != nil {
				c.HTML(http.StatusNotFound, "error.html", gin.H{
					"Error": "目录不存在",
				})
				return
			}

			breadcrumbs := api.GetBreadcrumbs(path)
			c.HTML(200, "index.html", gin.H{
				"Title":       "共享文件",
				"Files":       files,
				"CurrentPath": path,
				"Breadcrumbs": breadcrumbs,
			})
		})

	}
}

// setupFileRoutes 设置文件处理路由
func setupFileRoutes(r *gin.Engine) {
	// 文件处理路由
	r.GET("/files/*filepath", func(c *gin.Context) {
		api.HandleFileRequestGin(c)
	})
}
