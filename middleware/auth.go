package middleware

import (
	"bluebell/api"
	"bluebell/pkg/errno"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带请求头三种方式 1. 请求头，请求体，url
		authoriation := c.Request.Header.Get("Authoriation")
		if authoriation == "" {
			c.JSON(http.StatusOK, errno.ErrUserTokenEmpty)
			c.Abort()
			return
		}

		// Authoriation: Bearer xxxx.xx.xxx
		parts := strings.SplitN(authoriation, " ", 2)
		if len(parts) != 2 && parts[0] != "Bearer" {
			c.JSON(http.StatusOK, errno.ErrUserTokenEmpty)
			c.Abort()
			return
		}

		token, err := jwt.Parse(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, errno.ErrUserTokenInvalid)
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(api.ContextUserIDKey, token.UserID) // 避免middlerware api互相调用
		c.Set(api.ContextUsernameKey, token.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
