package router

import (
	"bluebell/api"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) *gin.Engine  {
	v1 := r.Group("/api/v1")
	v1.POST("/register", api.RegisterHandler)
	v1.POST("/login", api.LoginHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", api.CommunityHandler)
		//v1.GET("/community/:id", api.CommunityDetailHandler)
		v1.POST("/vote",api.PostVoteHandler)

	}
	//
	//v1.GET("/ping",middleware.JWTAuthMiddleware(), func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, "pong")
	//})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":"404",
		})
	})
	return r
}