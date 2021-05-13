package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/logger"
)

func SetUp() *gin.Engine  {
	r:=gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok!")
	})

	return r
}

