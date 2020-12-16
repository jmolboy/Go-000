package midware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
用户登录验证中间件
 */
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//userName, _ := c.Get("username")
		c.Next()
		return
	}
}