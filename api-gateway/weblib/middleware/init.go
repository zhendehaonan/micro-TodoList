package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 接收服务实例,并存到gin.Keys中
func InitMiddleware(service []any) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存到gin.Keys中
		context.Keys = make(map[string]any)
		context.Keys["userService"] = service[0]
		context.Keys["taskService"] = service[1]
		context.Next()
	}
}

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
