package v1

import (
	"bluebell/common"
	"bluebell/entity/user"
	"bluebell/pkg/errno"
	"bluebell/pkg/translator"
	"bluebell/repository"
	"bluebell/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	ContextUserIDKey   = "userID"
	ContextUsernameKey = "username"
)

var userService *service.UserService

func init() {
	userService, _ = (service.NewUserService(repository.NewUserRepository(""))).(*service.UserService)
}

func RegisterHandler(c *gin.Context) {
	p := new(user.ParamRegister)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, errno.ErrParam)
			return
		}
		//// validator.ValidationErrors类型错误则进行翻译
		msg, _ := json.Marshal(common.FormatTranslateMsg(errs.Translate(translator.Trans)))
		s := string(msg)

		c.JSON(http.StatusOK, errno.ErrParam.WithMsg(s))
		return
	}

	err := userService.Register(p)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrUserRegisterFailed.WithMsg("注册失败,原因:"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, errno.OK)
}

func LoginHandler(c *gin.Context) {
	p := new(user.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		// 并使用removeTopStruct函数去除字段名中的结构体名称标识
		c.JSON(http.StatusOK, gin.H{
			"msg": common.FormatTranslateMsg(errs.Translate(translator.Trans)),
		})
		return
	}

	aToken, rToken, err := userService.Login(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登陆失败,原因:" + err.Error(),
		})
		return
	}

	data := make(map[string]string)
	data["access_token"] = aToken
	data["refresh_token"] = rToken

	c.JSON(http.StatusOK, errno.OK.WithData(data))
}

