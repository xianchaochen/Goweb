package v1

import (
	"bluebell/api/v1"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) *gin.Engine  {
	//r.POST("/auth", v1.AuthHandler)
	r.POST("/register", v1.RegisterHandler)
	r.POST("/login", v1.LoginHandler)
	r.GET("/ping",middleware.JWTAuthMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":"404",
		})
	})
	return r
}