package main

import (
	"gin-service/demo/app/api"
	"gin-service/framework/router"

	"github.com/gin-gonic/gin"
)

// RegistryRoutes 路由注册函数
func RegistryRoutes(r *gin.Engine) {
	grp := router.NewGroup(r, "/ping")
	grp.RegRestAPIFunc(api.NewPing)
}
