package router

import "github.com/gin-gonic/gin"

// NewGroup 新建自定义路由组
func NewGroup(g *gin.Engine, path string) *Group {
	return (*Group)(g.Group(path))
}
