package main

import (
	"gin-service/framework/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	err := bootstrap.InitMySQL("config.json", "")
	if err != nil {
		panic(err)
	}
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	RegistryRoutes(engine)
	err = engine.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
