package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	sharedDirs   []string
	sharedDirsMu sync.RWMutex
)

// AddSharedDir adds a directory to the list of shared directories 添加共享目录
func AddSharedDir(dir string) {
	sharedDirsMu.Lock()
	defer sharedDirsMu.Unlock()

	sharedDirs = append(sharedDirs, dir)
}

// FileInfo 结构体用于存储文件或目录信息
type FileInfo struct {
	Name  string
	IsDir bool
	Path  string
}

// ListFiles 返回指定路径下的文件和文件夹列表
func ListFiles(requestPath string) ([]FileInfo, error) {
	sharedDirsMu.RLock()
	defer sharedDirsMu.RUnlock()

	// 遍历共享目录，查找匹配的路径
	for _, baseDir := range sharedDirs {
		fullPath := filepath.Join(baseDir, requestPath)

		// 检查请求的路径是否在允许的共享目录内
		if !strings.HasPrefix(fullPath, baseDir) {
			continue
		}

		// 检查路径是否存在
		if _, err := os.Stat(fullPath); err != nil {
			continue
		}

		entries, err := os.ReadDir(fullPath)
		if err != nil {
			continue
		}

		var files []FileInfo
		for _, entry := range entries {
			files = append(files, FileInfo{
				Name:  entry.Name(),
				IsDir: entry.IsDir(),
				Path:  filepath.Join(requestPath, entry.Name()),
			})
		}
		return files, nil
	}

	return nil, fmt.Errorf("path not found")
}

// GetBreadcrumbs 生成面包屑导航
func GetBreadcrumbs(path string) []map[string]string {
	if path == "" {
		return []map[string]string{{
			"name": "Home",
			"path": "",
		}}
	}

	// 将路径转换为系统无关的格式（统一使用 '/'）
	path = filepath.ToSlash(path)
	parts := strings.Split(strings.Trim(path, "/"), "/")
	breadcrumbs := make([]map[string]string, len(parts)+1)

	// 添加 Home
	breadcrumbs[0] = map[string]string{
		"name": "Home",
		"path": "",
	}

	// 构建每一级的路径
	currentPath := ""
	for i, part := range parts {
		if i == 0 {
			currentPath = part
		} else {
			currentPath = currentPath + "/" + part
		}

		// 确保 URL 中的路径是正确编码的
		breadcrumbs[i+1] = map[string]string{
			"name": part,
			"path": currentPath,
		}
	}

	return breadcrumbs
}

// HandleFileRequestGin serves files requested via Gin 处理文件请求
func HandleFileRequestGin(c *gin.Context) {
	filename := c.Param("filename")

	sharedDirsMu.RLock()
	defer sharedDirsMu.RUnlock()

	for _, dir := range sharedDirs {
		filePath := filepath.Join(dir, filename)
		if _, err := os.Stat(filePath); err == nil {
			c.File(filePath)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
}
