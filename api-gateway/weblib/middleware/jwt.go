package middleware

import (
	"github.com/gin-gonic/gin"
	"micro-TodoList/api-gateway/pkg/e"
	"micro-TodoList/api-gateway/pkg/utils"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt { //验证token是否过期
				code = e.ErrorAuthCheckTokenTimeOut
			}
		}
		if code != e.Success {
			c.JSON(e.Error, gin.H{
				"status": code,
				"msg":    "鉴权失败",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
