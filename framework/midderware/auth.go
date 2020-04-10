package middleware

import (
	"github.com/gin-gonic/gin"
)

func AppAuth(c *gin.Context) {
	c.Set("app_auth_id", 1)
	c.Next()
}
