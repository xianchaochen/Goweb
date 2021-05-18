package router

import (
	"bluebell/api"
	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

func Register(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	v1.POST("/register", api.RegisterHandler)
	v1.POST("/login", api.LoginHandler)
	// 全站 2s 放一个token
	v1.Use(middleware.JWTAuthMiddleware(), middleware.RateLimitMiddleware(time.Second*2,1))
	{
		v1.GET("/community", api.CommunityHandler)
		//v1.GET("/community/:id", api.CommunityDetailHandler)
		v1.POST("/vote", api.PostVoteHandler)
		// 局部限流
		v1.Use(middleware.RateLimitMiddleware(time.Second*2,1))
	}

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
