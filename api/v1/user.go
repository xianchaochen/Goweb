package v1

import (
	"bluebell/common"
	"bluebell/entity"
	"bluebell/pkg/translator"
	"bluebell/repository"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var userService service.IUserService

func init()  {
	if userService == nil {
		userService = service.NewUserService(repository.NewUserRepository(""))
	}
}

func RegisterHandler(c *gin.Context) {
	p := new(entity.ParamRegister)
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
			"msg": common.RemoveTopStruct(errs.Translate(translator.Trans)),
		})
		return
	}

	err := userService.Register(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "注册失败,原因:"+err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg" : "success",
	})
}


func LoginHandler(c *gin.Context) {
	p := new(entity.ParamLogin)
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
			"msg": common.RemoveTopStruct(errs.Translate(translator.Trans)),
		})
		return
	}

	//userRepository := repository.NewUserRepository("")
	//userService:=service.NewUserService(userRepository)
	//err := userService.Register(p)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg" : "注册失败,原因:"+err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"msg" : "success",
	})
}


