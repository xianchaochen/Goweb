package v1

import (
	"bluebell/api/v1"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) *gin.Engine  {
	r.POST("/register", v1.RegisterHandler)
	r.POST("/login", v1.LoginHandler)


	return r
}