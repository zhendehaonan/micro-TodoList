package weblib

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"micro-TodoList/api-gateway/weblib/handlers"
	"micro-TodoList/api-gateway/weblib/middleware"
)

func NewRouter(service ...interface{}) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		//用户服务
		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)
	}
	return r
}
